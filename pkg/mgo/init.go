package mgo

import (
	"bot/config"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMgo(cfg config.Config) *mongo.Client {
	ctx := context.TODO()
	mgoCon := options.Client().ApplyURI(cfg.Mongo.Url)
	mgoClient, err := mongo.Connect(ctx, mgoCon)
	if err != nil {
		fmt.Println(err)
	}
	return mgoClient
}
