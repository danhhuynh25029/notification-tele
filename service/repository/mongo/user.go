package mongo

import (
	"bot/pkg/constant"
	"bot/service/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	AddUser(ctx context.Context, user models.User) (string, error)
	UpdateUser()
	DeleteUser()
	GetUser(ctx context.Context, req models.User) (*models.User, error)
}

type UserRepository struct {
	mgo *mongo.Client
}

func NewUserRepository(mgo *mongo.Client) IUserRepository {
	return &UserRepository{
		mgo: mgo,
	}
}
func (u *UserRepository) getCollection() *mongo.Collection {
	return u.mgo.Database(constant.DatabaseUser).Collection(constant.CollectionUser)
}
func (u *UserRepository) AddUser(ctx context.Context, user models.User) (string, error) {
	result, err := u.getCollection().InsertOne(ctx, &user)
	if err != nil {
		return "", err
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}
func (u *UserRepository) UpdateUser() {

}
func (u *UserRepository) DeleteUser() {

}

func (u *UserRepository) GetUser(ctx context.Context, req models.User) (*models.User, error) {
	user := models.User{}
	condition := bson.M{}
	if req.Username != "" {
		condition["username"] = req.Username
	}
	if req.Password != "" {
		condition["password"] = req.Password
	}
	err := u.getCollection().FindOne(ctx, condition).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
