package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-invoice/internal/feature/user"
	"github.com/martinyonatann/go-invoice/internal/server/api"
	logger "github.com/rs/zerolog/log"
)

type UserHandler struct {
	userUC user.UseCase
}

func New(s user.UseCase) api.UsersHandlers {
	return &UserHandler{userUC: s}
}

type responseBody struct {
	StatusCode int         `json:"rc"`
	Message    string      `json:"message,omitempty"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (x *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {

		createUserPayload := user.CreateUserRequest{}

		err := json.NewDecoder(c.Request().Body).Decode(&createUserPayload)
		if err != nil {
			logger.Err(err).Msg("[handlerUser]createUser_decode")

			return err
		}

		userData, err := x.userUC.CreateUser(c.Request().Context(), createUserPayload)
		if err != nil {
			logger.Err(err).Msg("[handlerUser][CreateUser]")

			return err
		}

		return c.JSON(http.StatusOK, responseBody{
			StatusCode: http.StatusOK,
			Message:    http.StatusText(http.StatusOK),
			Data:       userData,
		})
	}
}
func (x *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			logger.Err(err).Msg("[handlerUser]GetUserByID_ParseInt")

			return err
		}

		getUserPayload := user.GetUserRequest{UserID: id}

		userData, err := x.userUC.GetUser(c.Request().Context(), getUserPayload)
		if err != nil {
			logger.Err(err).Msg("[handlerUser][getUserById]")

			return err
		}
		return c.JSON(http.StatusOK, responseBody{
			StatusCode: http.StatusOK,
			Message:    http.StatusText(http.StatusOK),
			Data:       userData,
		})
	}
}
