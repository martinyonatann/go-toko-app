package user_repository

import (
	"context"
	"database/sql"

	"github.com/martinyonatann/go-invoice/internal/repository/postgresql_query"
)

func New(db *sql.DB) UserRepository {
	return &Repo{db: db}
}

type Repo struct {
	db *sql.DB
}

func (x *Repo) CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error) {
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

func (x *Repo) GetUserById(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error) {
	query := postgresql_query.GetUserById

	if err = x.db.QueryRowContext(ctx, query, request.UserID).Scan(
		&response.ID,
		&response.FullName,
		&response.Email,
		&response.CreatedAt); err != nil {
		return response, err
	}

	return response, err
}
