package main

import (
	"bot/service/domain/user/delivery/http"
	"bot/service/domain/user/usecase"

	messageHandler "bot/service/domain/message/delivery/http"
	messageUse "bot/service/domain/message/usecase"
	"bot/service/repository"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.TODO()
	mgoCon := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	mgoClient, err := mongo.Connect(ctx, mgoCon)
	if err != nil {
		fmt.Println(err)
	}
	repo := repository.NewIRepo(mgoClient)
	userUsecase := usecase.NewUserUseCase(repo)
	userHandler := http.NewUserHandler(userUsecase)

	messageUsecase := messageUse.NewMessageUsecase(repo)
	messageHandler := messageHandler.NewMessageHandler(messageUsecase)

	r := gin.Default()
	group := r.Group("/v1")
	userHandler.NewUserRoute(group)
	messageHandler.NewMessageRoute(group)
	r.Run(":3070")
}
