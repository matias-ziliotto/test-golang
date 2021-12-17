package user

import (
	"context"
	"database/sql"

	"github.com/matias-ziliotto/test-golang/internal/domain"
)

// Repository encapsulates the storage of a user.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	Get(ctx context.Context, id int) (domain.User, error)
	Exists(ctx context.Context, cardNumberID string) bool
	Save(ctx context.Context, b domain.User) (int, error)
	Update(ctx context.Context, b domain.User) error
	Delete(ctx context.Context, id int) error
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

func (r *repository) Get(ctx context.Context, id int) (domain.User, error) {
	query := "SELECT * FROM users WHERE id = ?;"
	row := r.db.QueryRow(query, id)
	u := domain.User{}
	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.DocumentTypeId, &u.DocumentNumber)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (r *repository) Exists(ctx context.Context, documentNumber string) bool {
	query := "SELECT document_number FROM users WHERE document_number=?;"
	row := r.db.QueryRow(query, documentNumber)
	err := row.Scan(&documentNumber)
	return err == nil
}

func (r *repository) Save(ctx context.Context, u domain.User) (int, error) {
	query := "INSERT INTO users(first_name,last_name, document_type_id, document_number) VALUES (?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(&u.FirstName, &u.LastName, &u.DocumentTypeId, &u.DocumentNumber)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, u domain.User) error {
	query := "UPDATE users SET first_name=?, last_name=?  WHERE id=?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(&u.FirstName, &u.LastName, &u.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return ErrNotFound
	}

	return nil
}
