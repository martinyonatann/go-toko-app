package user_repository

import (
	"context"
	"time"
)

type UserRepository interface {
	GetUserById(ctx context.Context, request GetUserByIDRequest) (GetUserByIDResponse, error)
	GetUserByEmail(ctx context.Context, request GetUserByEmailRequest) (GetUserByEmailResponse, error)
	CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error)
	ListUsers(ctx context.Context, request ListUsersRequest) (ListUsersResponse, error)

	// SearchUsers(query string) ([]*models.User, error)

	// UpdateUser(e *models.User) error
	// DeleteUser(id int64) error
}

type GetUserByIDRequest struct {
	UserID int64 `json:"user_id"`
}

type GetUserByIDResponse struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

type GetUserByEmailResponse struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID        int64     `json:"id,omitempty"`
	FullName  string    `json:"fullname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type ListUsersRequest struct{}

type ListUsersResponse []GetUserByIDResponse