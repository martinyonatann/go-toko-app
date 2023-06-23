package http

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-toko-app/internal/usecase/user_usecase"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	userUC user_usecase.UserUC
	log    zerolog.Logger
}

func New(s user_usecase.UserUC, log zerolog.Logger) *UserHandler {
	return &UserHandler{s, log}
}

type responseBody struct {
	StatusCode int         `json:"rc"`
	Message    string      `json:"message"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data"`
}

func (x *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		createUserPayload := user_usecase.CreateUserRequest{}

		err := json.NewDecoder(c.Request().Body).Decode(&createUserPayload)
		if err != nil {
			x.log.Err(err).Msg("[handlerUser]createUser_decode")

			return err
		}

		userData, err := x.userUC.CreateUser(c.Request().Context(), createUserPayload)
		if err != nil {
			x.log.Err(err).Msg("[handlerUser][CreateUser]")

			return c.JSON(500, responseBody{StatusCode: 500, Message: http.StatusText(500), Error: err.Error()})

		}

		return c.JSON(http.StatusOK, responseBody{http.StatusOK, http.StatusText(http.StatusOK), "", userData})
	}
}

func (x *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		payload := user_usecase.LoginRequest{}

		err = json.NewDecoder(c.Request().Body).Decode(&payload)
		if err != nil {
			x.log.Err(err).Msg("[Login]createUser_decode")

			return err
		}

		userData, err := x.userUC.Login(c.Request().Context(), payload)
		if err != nil {
			x.log.Err(err).Msg("[Login]Login")

			return c.JSON(500, responseBody{StatusCode: 500, Message: http.StatusText(500), Error: err.Error()})
		}
		return c.JSON(http.StatusOK, responseBody{http.StatusOK, http.StatusText(http.StatusOK), "", userData})
	}
}
