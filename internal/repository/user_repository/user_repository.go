package user_repository

import (
	"context"
	"database/sql"

	"github.com/martinyonatann/go-toko-app/internal/repository/contract"
	"github.com/martinyonatann/go-toko-app/internal/repository/postgresql_query"
)

func New(db *sql.DB) contract.UserRepository {
	return &Repo{db: db}
}

type Repo struct {
	db *sql.DB
}

func (x *Repo) CreateUser(
	ctx context.Context,
	request contract.CreateUserRequest) (
	response contract.CreateUserResponse,
	err error) {
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

func (x *Repo) GetUserById(ctx context.Context, request contract.GetUserByIDRequest) (response contract.GetUserByIDResponse, err error) {
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

func (x *Repo) GetUserByEmail(
	ctx context.Context,
	request contract.GetUserByEmailRequest) (
	response contract.GetUserByEmailResponse,
	err error) {
	query := postgresql_query.GetUserByEmail

	if err = x.db.QueryRowContext(ctx, query, request.Email).Scan(
		&response.ID,
		&response.FullName,
		&response.Email,
		&response.Password,
		&response.CreatedAt); err != nil {
		return response, err
	}

	return response, err
}

func (x *Repo) ListUsers(
	ctx context.Context,
	req contract.ListUsersRequest) (
	resp contract.ListUsersResponse,
	err error) {
	query := postgresql_query.ListUsers

	a, err := x.db.Query(query)
	if err != nil {
		return resp, err
	}

	defer a.Close()

	for a.Next() {
		user := contract.User{}
		if err := a.Scan(&user.ID, &user.FullName, &user.Email, &user.CreatedAt); err != nil {
			return resp, err
		}
		resp = append(resp, user)
	}

	return resp, err
}
