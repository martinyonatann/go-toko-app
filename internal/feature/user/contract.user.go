package user

import (
	"context"
	"time"
)

type UseCase interface {
	GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error)
	CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error)
	ListUsers(ctx context.Context, request ListUsersRequest) (ListUsersResponse, error)
}

type GetUserRequest struct {
	UserID int64 `json:"user_id"`
}

type GetUserResponse struct {
	ID        int64     `json:"user_id"`
	FullName  string    `json:"fullname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type CreateUserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID        int64     `json:"user_id,omitempty"`
	FullName  string    `json:"fullname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ListUsersRequest struct{}

type ListUsersResponse []GetUserResponse
