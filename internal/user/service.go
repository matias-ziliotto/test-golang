package user

import (
	"context"

	"github.com/matias-ziliotto/test-golang/internal/domain"
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
