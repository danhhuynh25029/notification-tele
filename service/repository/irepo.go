package repository

import (
	"bot/service/repository/mgo"

	"go.mongodb.org/mongo-driver/mongo"
)

type IRepo struct {
	client *mongo.Client
}

func NewIRepo(client *mongo.Client) IRepo {
	return IRepo{
		client: client,
	}
}

func (i *IRepo) NewUserRepo() mgo.IUserRepository {
	return mgo.NewUserRepository(i.client)
}
func (i *IRepo) NewMessageRepo() mgo.IMessageRepo {
	return mgo.NewMessageRepo(i.client)
}
