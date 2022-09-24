package mgo

import (
	"bot/pkg/constant"
	"bot/service/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type IMessageRepo interface {
	AddMessage(ctx context.Context, message models.Message) error
	UpdateMessage(ctx context.Context, message models.Message) error
	DeleteMessage(ctx context.Context, message models.Message) error
	GetAllMessage(ctx context.Context) ([]models.Message, error)
}

type MessageRepo struct {
	mgo *mongo.Client
}

func NewMessageRepo(mgo *mongo.Client) IMessageRepo {
	return &MessageRepo{
		mgo: mgo,
	}
}

func (m *MessageRepo) getCollection() *mongo.Collection {
	return m.mgo.Database(constant.DatabaseTele).Collection(constant.CollectionMessage)
}
func (m *MessageRepo) AddMessage(ctx context.Context, message models.Message) error {
	_, err := m.getCollection().InsertOne(ctx, &message)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageRepo) UpdateMessage(ctx context.Context, message models.Message) error {

	model := m.buildUpdateModel(message)
	id, _ := primitive.ObjectIDFromHex(message.ID)

	_, err := m.getCollection().UpdateOne(ctx, id, model)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageRepo) DeleteMessage(ctx context.Context, message models.Message) error {
	id, _ := primitive.ObjectIDFromHex(message.ID)
	_, err := m.getCollection().DeleteOne(ctx, bson.M{
		"_id": id,
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageRepo) GetAllMessage(ctx context.Context) ([]models.Message, error) {
	var messages []models.Message
	cursor, err := m.getCollection().Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &messages)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
func (m *MessageRepo) buildUpdateModel(message models.Message) bson.M {
	model := bson.M{}
	model["updated_at"] = time.Now().Unix()
	model["content"] = message.Content
	model["send_time"] = message.SendTime
	return model
}
