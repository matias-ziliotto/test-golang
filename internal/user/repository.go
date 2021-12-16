package user

import (
	"context"
	"database/sql"

	"github.com/matias-ziliotto/test-golang/internal/domain"
)

type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.User, error) {
	query := "SELECT * FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		u := domain.User{}
		err = rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.DocumentTypeId, &u.DocumentNumber)
		if err == nil {
			users = append(users, u)
		}
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}
