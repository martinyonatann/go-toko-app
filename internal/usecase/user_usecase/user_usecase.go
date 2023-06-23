package user_usecase

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/martinyonatann/go-toko-app/internal/repository/user_repository"
	"github.com/martinyonatann/go-toko-app/internal/utils"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUC interface {
		GetUser(ctx context.Context, request GetUserRequest) (response GetUserResponse, err error)
		CreateUser(ctx context.Context, request CreateUserRequest) (response CreateUserResponse, err error)
		ListUsers(ctx context.Context, request ListUsersRequest) (ListUsersResponse, error)
		Login(ctx context.Context, request LoginRequest) (LoginResponse, error)
	}

	GetUserRequest struct {
		UserID int64 `json:"user_id"`
	}

	GetUserResponse struct {
		ID        int64     `json:"user_id"`
		FullName  string    `json:"fullname,omitempty"`
		Email     string    `json:"email,omitempty"`
		Password  string    `json:"password,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	}

	CreateUserRequest struct {
		FullName string `json:"fullname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	CreateUserResponse struct {
		ID        int64     `json:"user_id,omitempty"`
		FullName  string    `json:"fullname,omitempty"`
		Email     string    `json:"email,omitempty"`
		Password  string    `json:"password,omitempty"`
		CreatedAt time.Time `json:"created_at,omitempty"`
	}

	ListUsersRequest struct{}

	ListUsersResponse []GetUserResponse

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		ID        int64     `json:"id"`
		FullName  string    `json:"fullname"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		Token     string    `json:"token"`
	}

	UseCase struct {
		userRepository user_repository.UserRepository
		log            zerolog.Logger
	}
)

func New(r user_repository.UserRepository, log zerolog.Logger) UserUC {
	return &UseCase{userRepository: r, log: log}
}

var ErrUserIdNotFound = errors.New("user_id not found")

func (uc *UseCase) CreateUser(
	ctx context.Context,
	request CreateUserRequest,
) (response CreateUserResponse, err error) {
	defer func() {
		uc.log.Err(err).Interface("res", response).Interface("req", request).Msg("[user_usecase]][CreateUser]")
	}()

	newPassword, err := utils.GeneratePassword(request.Password)
	if err != nil {
		uc.log.Err(err).Msg("[user_usecase]CreateUser_generatePassword")
		return response, err
	}

	createUserPayload := user_repository.CreateUserRequest{
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(newPassword),
	}

	userData, err := uc.userRepository.CreateUser(ctx, createUserPayload)
	if err != nil {
		// if strings(err.Error())
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return response, errors.New("email already used")
		}

		uc.log.Err(err).Msg("[user_usecase]CreateUser_CreateUser")

		return response, err
	}

	response = CreateUserResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}

	return response, err
}

func (uc *UseCase) GetUser(
	ctx context.Context,
	request GetUserRequest,
) (response GetUserResponse, err error) {
	defer func() {
		uc.log.Err(err).Interface("UserID", request.UserID).Interface("resp", response).Msg("[user_usecase][GetUser]")
	}()

	// validation
	{
		if request.UserID == 0 {
			return response, ErrUserIdNotFound
		}
	}

	userData, err := uc.userRepository.GetUserById(ctx, user_repository.GetUserByIDRequest{
		UserID: request.UserID,
	})
	if err != nil {
		return response, err
	}

	response = GetUserResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}

	return response, err
}

func (uc *UseCase) ListUsers(
	ctx context.Context,
	/*request*/ request ListUsersRequest) (
	/*response*/ response ListUsersResponse, err error) {
	dataUsers, err := uc.userRepository.ListUsers(ctx, user_repository.ListUsersRequest{})
	if err != nil {
		uc.log.Err(err).Msg("[user_usecase][ListUsers]ListUsers")

		return response, err
	}

	for _, v := range dataUsers {
		response = append(response, GetUserResponse{
			ID:        v.ID,
			FullName:  v.FullName,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		})
	}

	return response, err
}

func (uc *UseCase) Login(
	ctx context.Context,
	request LoginRequest) (
	response LoginResponse, err error) {
	defer func() {
		uc.log.Err(err).Interface("response", response).Msg("[user_usecase][Login]")
	}()

	// validation
	{
		if request.Password == "" {
			return response, errors.New("password can't be null")
		}

		if request.Email == "" {
			return response, errors.New("email can't be null")
		}
	}

	userData, err := uc.userRepository.GetUserByEmail(ctx, user_repository.GetUserByEmailRequest{
		Email: request.Email,
	})
	if err != nil {
		uc.log.Err(err).Msg("[user_usecase][Login] GetUserByEmail")
		return response, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password)); err != nil {
		uc.log.Err(err).Msg("[user_usecase][Login]CompareHashAndPassword")
		return response, err
	}

	token, err := utils.CreateToken(request.Email, utils.UserInfo{ID: userData.ID, FullName: userData.FullName, Email: userData.Email})
	if err != nil {
		uc.log.Err(err).Msg("[user_usecase][Login]CreateToken")
		return response, err
	}

	response = LoginResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		Token:     token,
	}

	return response, err
}
