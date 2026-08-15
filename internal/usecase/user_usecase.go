package usecase

import (
	"context"

	"webproject/internal/domain/entity"
)

// UserRepository is the port the usecase depends on. It is implemented by
// internal/infrastructure/postgres — the usecase layer never imports
// pgx or postgres directly, only this interface.
type UserRepository interface {
	Create(ctx context.Context, u *entity.User) error
	GetByID(ctx context.Context, id int64) (*entity.User, error)
}

type UserUsecase struct {
	repo UserRepository
}

func NewUserUsecase(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *entity.User) error {
	return uc.repo.Create(ctx, u)
}

func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	return uc.repo.GetByID(ctx, id)
}
