package user

import (
	"context"
	"errors"
	"strings"

	user_repository "github.com/martinyonatann/go-invoice/internal/repository"
	"github.com/martinyonatann/go-invoice/internal/utils"
	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type Feat struct {
	userRepository user_repository.UserRepository
	log            zerolog.Logger
}

func New(r user_repository.UserRepository, log zerolog.Logger) *Feat {
	return &Feat{userRepository: r, log: log}
}

var ErrUserIdNotFound = errors.New("user_id not found")

func (x *Feat) CreateUser(ctx context.Context,
	request CreateUserRequest,
) (response CreateUserResponse, err error) {
	defer func() {
		x.log.Err(err).Interface("userData", response).Msg("[FeatureUser][CreateUser]")
	}()

	newPassword, err := generatePassword(request.Password)
	if err != nil {
		x.log.Err(err).Msg("[FeatureUser][CreateUser]generatePassword")
		return response, err
	}

	createUserPayload := user_repository.CreateUserRequest{
		FullName: request.FullName,
		Email:    request.Email,
		Password: string(newPassword),
	}

	userData, err := x.userRepository.CreateUser(ctx, createUserPayload)
	if err != nil {
		// if strings(err.Error())
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			return response, errors.New("email already used")
		}

		x.log.Err(err).Msg("[FeatureUser][CreateUser]CreateUser")

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

func (x *Feat) GetUser(
	ctx context.Context,
	request GetUserRequest,
) (response GetUserResponse, err error) {
	defer func() {
		logger.Err(err).Interface("userData", response).Msg("GetUser_UseCase")
	}()

	// validation
	{
		if request.UserID == 0 {
			return response, ErrUserIdNotFound
		}
	}

	userData, err := x.userRepository.GetUserById(ctx, user_repository.GetUserByIDRequest{
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

func (x *Feat) ListUsers(ctx context.Context, request ListUsersRequest) (resp ListUsersResponse, err error) {
	dataUsers, err := x.userRepository.ListUsers(ctx, user_repository.ListUsersRequest{})
	if err != nil {
		x.log.Err(err).Msg("[ListUsers]ListUsers")
		return resp, err
	}

	for _, v := range dataUsers {
		resp = append(resp, GetUserResponse{
			ID:        v.ID,
			FullName:  v.FullName,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		})
	}

	return resp, err
}

func generatePassword(raw string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (x *Feat) Login(ctx context.Context, request LoginRequest) (response LoginResponse, err error) {
	defer func() {
		x.log.Err(err).Interface("response", response).Msg("[FeatureUser][Login]")
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

	userData, err := x.userRepository.GetUserByEmail(ctx, user_repository.GetUserByEmailRequest{
		Email: request.Email,
	})
	if err != nil {
		x.log.Err(err).Msg("[FeatureUser][Login] GetUserByEmail")
		return response, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(request.Password)); err != nil {
		x.log.Err(err).Msg("[FeatureUser][Login]CompareHashAndPassword")
		return response, err
	}

	token, err := utils.CreateToken(request.Email, utils.UserInfo{ID: userData.ID, FullName: userData.FullName, Email: userData.Email})
	if err != nil {
		x.log.Err(err).Msg("[FeatureUser][Login]CreateToken")
		return response, err
	}

	return LoginResponse{ID: userData.ID, FullName: userData.FullName, Email: userData.Email, CreatedAt: userData.CreatedAt, Token: token}, err
}
