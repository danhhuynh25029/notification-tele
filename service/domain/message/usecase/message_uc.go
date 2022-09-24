package usecase

import (
	"bot/service/models"
	"bot/service/repository/mgo"
	"context"
)

type IMessageUsecase interface {
	AddMessage(ctx context.Context, message models.Message) error
	UpdateMesasge(ctx context.Context, message models.Message) error
	DeleteMessage(ctx context.Context, message models.Message) error
	GetAllMessage(ctx context.Context) ([]models.Message, error)
}

type MessageUsecase struct {
	repo mgo.IMessageRepo
}

func NewMessageUsecase(repo mgo.IMessageRepo) IMessageUsecase {
	return &MessageUsecase{
		repo: repo,
	}
}

func (m *MessageUsecase) AddMessage(ctx context.Context, message models.Message) error {
	return nil
}

func (m *MessageUsecase) UpdateMesasge(ctx context.Context, message models.Message) error {
	return nil
}

func (m *MessageUsecase) DeleteMessage(ctx context.Context, message models.Message) error {
	return nil
}

func (m *MessageUsecase) GetAllMessage(ctx context.Context) ([]models.Message, error) {
	return nil, nil
}
