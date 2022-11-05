package rabbit

import (
	repo "bot/service/repository"
)

type RabbitProducer interface {
	SendMessage() error
}
type rabbitProducer struct {
	repo repo.IRepo
}

func NewProducer(repo repo.IRepo) (RabbitProducer, error) {
	return &rabbitProducer{
		repo: repo,
	}, nil
}

func (r *rabbitProducer) SendMessage() error {
	return nil
}
