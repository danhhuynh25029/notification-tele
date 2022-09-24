package usecase

import (
	"bot/service/models"
	"bot/service/repository"
	"context"
)

type IMessageUsecase interface {
	AddMessage(ctx context.Context, message models.Message) error
	UpdateMesasge(ctx context.Context, message models.Message) error
	DeleteMessage(ctx context.Context, message models.Message) error
	GetAllMessage(ctx context.Context) ([]models.Message, error)
}

type MessageUsecase struct {
	repo repository.IRepo
}

func NewMessageUsecase(repo repository.IRepo) IMessageUsecase {
	return &MessageUsecase{
		repo: repo,
	}
}

func (m *MessageUsecase) AddMessage(ctx context.Context, message models.Message) error {
	err := m.repo.NewMessageRepo().AddMessage(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageUsecase) UpdateMesasge(ctx context.Context, message models.Message) error {
	err := m.repo.NewMessageRepo().UpdateMessage(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageUsecase) DeleteMessage(ctx context.Context, message models.Message) error {
	err := m.repo.NewMessageRepo().DeleteMessage(ctx, message)
	if err != nil {
		return err
	}
	return nil
}

func (m *MessageUsecase) GetAllMessage(ctx context.Context) ([]models.Message, error) {
	messages, err := m.repo.NewMessageRepo().GetAllMessage(ctx)
	if err != nil {
		return nil, err
	}
	return messages, nil
}
