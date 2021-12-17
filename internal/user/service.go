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
	Get(ctx context.Context, userId int) (domain.User, error)
	Store(ctx context.Context, firstName string, lastName string, documentTypeId int, documentNumber int) (domain.User, error)
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

func (s *service) Get(ctx context.Context, userId int) (domain.User, error) {
	return s.repository.Get(ctx, userId)
}

func (s *service) Store(ctx context.Context, firstName string, lastName string, documentTypeId int, documentNumber int) (domain.User, error) {
	user := domain.User{
		FirstName:      firstName,
		LastName:       lastName,
		DocumentTypeId: documentTypeId,
		DocumentNumber: documentNumber,
	}

	id, err := s.repository.Save(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	user.Id = id

	return user, nil
}
