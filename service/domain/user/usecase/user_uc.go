package usecase

import (
	"bot/service/models"
	"bot/service/repository"
	"context"
)

type IUserUsecase interface {
	AddUser(ctx context.Context, user models.User) error
	UpdateUser(ctx context.Context, user models.User) error
	GetUser(ctx context.Context, req models.User) (*models.User, error)
}

type UserUsecase struct {
	repo repository.IRepo
}

func NewUserUseCase(repo repository.IRepo) IUserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}
func (u *UserUsecase) AddUser(ctx context.Context, user models.User) error {
	_, err := u.repo.NewUserRepo().AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) UpdateUser(ctx context.Context, user models.User) error {
	err := u.repo.NewUserRepo().UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) GetUser(ctx context.Context, req models.User) (*models.User, error) {
	user, err := u.repo.NewUserRepo().GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return user, nil
}
