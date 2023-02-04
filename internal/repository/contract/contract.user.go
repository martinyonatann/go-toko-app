package contract

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

type User struct {
	ID        int64     `json:"id,omitempty"`
	FullName  string    `json:"fullname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type GetUserByIDRequest struct {
	UserID int64 `json:"user_id"`
}

type GetUserByIDResponse User

type GetUserByEmailRequest struct {
	Email string `json:"email"`
}

type GetUserByEmailResponse User

type CreateUserRequest struct {
	FullName string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse User

type ListUsersRequest struct{}

type ListUsersResponse []User
