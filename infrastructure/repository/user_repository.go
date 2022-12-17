package repository

import (
	"context"
	"database/sql"

	"github.com/martinyonatann/go-invoice/infrastructure/repository/contract"
	"github.com/martinyonatann/go-invoice/infrastructure/repository/postgresql_query"
)

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (x *Repository) CreateUser(ctx context.Context, request contract.CreateUserRequest) (response contract.CreateUserResponse, err error) {
	query := postgresql_query.UserInsert

	if err = x.db.QueryRowContext(
		ctx,
		query,
		request.FullName,
		request.Email,
		request.Password).Scan(
		&response.ID,
		&response.FullName,
		&response.Email,
		&response.Password,
		&response.CreatedAt); err != nil {
		return response, err
	}

	return response, err
}

func (x *Repository) GetUserById(ctx context.Context, request contract.GetUserRequest) (response contract.GetUserResponse, err error) {
	query := postgresql_query.GetUserById

	if err = x.db.QueryRowContext(ctx, query, request.UserID).Scan(response); err != nil {
		return response, err
	}

	return response, err
}
