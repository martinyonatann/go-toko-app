package user_usecase

import (
	"context"
	"errors"
	"strings"

	repoContract "github.com/martinyonatann/go-toko-app/internal/repository/contract"
	"github.com/martinyonatann/go-toko-app/internal/usecase/contract"
	"github.com/martinyonatann/go-toko-app/internal/utils"
	"github.com/rs/zerolog"
	"golang.org/x/crypto/bcrypt"
)

type UseCase struct {
	userRepository repoContract.UserRepository
	log            zerolog.Logger
}

func New(r repoContract.UserRepository, log zerolog.Logger) *UseCase {
	return &UseCase{userRepository: r, log: log}
}

var ErrUserIdNotFound = errors.New("user_id not found")

func (uc *UseCase) CreateUser(
	ctx context.Context,
	request contract.CreateUserRequest,
) (response contract.CreateUserResponse, err error) {
	defer func() {
		uc.log.Err(err).Interface("res", response).Interface("req", request).Msg("[user_usecase]][CreateUser]")
	}()

	newPassword, err := utils.GeneratePassword(request.Password)
	if err != nil {
		uc.log.Err(err).Msg("[user_usecase]CreateUser_generatePassword")
		return response, err
	}

	createUserPayload := repoContract.CreateUserRequest{
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

	response = contract.CreateUserResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}

	return response, err
}

func (uc *UseCase) GetUser(
	ctx context.Context,
	request contract.GetUserRequest,
) (response contract.GetUserResponse, err error) {
	defer func() {
		uc.log.Err(err).Interface("UserID", request.UserID).Interface("resp", response).Msg("[user_usecase][GetUser]")
	}()

	// validation
	{
		if request.UserID == 0 {
			return response, ErrUserIdNotFound
		}
	}

	userData, err := uc.userRepository.GetUserById(ctx, repoContract.GetUserByIDRequest{
		UserID: request.UserID,
	})
	if err != nil {
		return response, err
	}

	response = contract.GetUserResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
	}

	return response, err
}

func (uc *UseCase) ListUsers(
	ctx context.Context,
	/*request*/ request contract.ListUsersRequest) (
	/*response*/ response contract.ListUsersResponse, err error) {
	dataUsers, err := uc.userRepository.ListUsers(ctx, repoContract.ListUsersRequest{})
	if err != nil {
		uc.log.Err(err).Msg("[user_usecase][ListUsers]ListUsers")

		return response, err
	}

	for _, v := range dataUsers {
		response = append(response, contract.GetUserResponse{
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
	request contract.LoginRequest) (
	response contract.LoginResponse, err error) {
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

	userData, err := uc.userRepository.GetUserByEmail(ctx, repoContract.GetUserByEmailRequest{
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

	response = contract.LoginResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		CreatedAt: userData.CreatedAt,
		Token:     token,
	}

	return response, err
}
