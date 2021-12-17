package user

import (
	"context"
	"errors"

	"github.com/matias-ziliotto/test-golang/internal/domain"
)

// Errors
var (
	ErrNotFound          = errors.New("user not found")
	ErrBuyerAlreadyExist = errors.New("user already exist")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.User, error)
}

type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{repository: repo}
}

func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	return s.repository.GetAll(ctx)
}
