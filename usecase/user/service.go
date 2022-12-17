package user

import (
	"context"
	"errors"

	"github.com/martinyonatann/go-invoice/infrastructure/repository"
	"github.com/martinyonatann/go-invoice/infrastructure/repository/contract"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func New(r repository.Repository) *Feat {
	return &Feat{userRepository: r}
}

type Feat struct {
	userRepository repository.Repository
}

var ErrUserIdNotFound = errors.New("user_id not found")

func (x *Feat) CreateUser(ctx context.Context,
	request CreateUserRequest,
) (response CreateUserResponse, err error) {
	newPassword, err := generatePassword(request.Password)
	if err != nil {
		return response, errors.New("failed generate password")
	}

	createUserPayload := contract.CreateUserRequest{
		FullName: request.FullName,
		Email:    request.Email,
		Password: newPassword,
	}

	userData, err := x.userRepository.CreateUser(ctx, createUserPayload)
	if err != nil {
		return response, err
	}

	logrus.New().Info("userData", userData)

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
	// validation
	{
		if request.UserID == 0 {
			return response, ErrUserIdNotFound
		}
	}

	userData, err := x.userRepository.GetUserById(ctx, contract.GetUserRequest{
		UserID: request.UserID,
	})
	if err != nil {
		return response, err
	}

	response = GetUserResponse{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Password:  userData.Password,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return response, err
}

func generatePassword(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
