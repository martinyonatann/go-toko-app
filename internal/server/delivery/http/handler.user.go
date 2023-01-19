package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-invoice/internal/feature/user"
	api "github.com/martinyonatann/go-invoice/internal/server/delivery"
	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
)

type UserHandler struct {
	userUC user.UseCase
	log    zerolog.Logger
}

func New(s user.UseCase, log zerolog.Logger) api.UsersHandlers {
	return &UserHandler{s, log}
}

type responseBody struct {
	StatusCode int         `json:"rc"`
	Message    string      `json:"message"`
	Error      string      `json:"error,omitempty"`
	Data       interface{} `json:"data"`
}

type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
	Message    map[string]string
	Meta, Data interface{}
}

type CaptureError struct {
	Err         error  `json:"error,omitempty"`
	UserMsg     string `json:"message,omitempty"`
	InternalMsg string `json:"internal_message,omitempty"`
	MoreInfo    string `json:"more_info,omitempty"`
}

func (x *UserHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		createUserPayload := user.CreateUserRequest{}

		err := json.NewDecoder(c.Request().Body).Decode(&createUserPayload)
		if err != nil {
			x.log.Err(err).Msg("[handlerUser]createUser_decode")

			return err
		}

		userData, err := x.userUC.CreateUser(c.Request().Context(), createUserPayload)
		if err != nil {
			x.log.Err(err).Msg("[handlerUser][CreateUser]")

			return c.JSON(500, x.NewResponse(http.StatusInternalServerError, nil, createUserPayload, err, "handler.user.register"))

		}

		return c.JSON(http.StatusOK, x.NewResponse(http.StatusOK, userData, createUserPayload, err, "handler.user.register"))
	}
}
func (x *UserHandler) GetUserByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			x.log.Err(err).Msg("[handlerUser]GetUserByID_ParseInt")

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

func (x *UserHandler) ListUsers() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		defer func() {
			x.log.Err(err).Interface("size", c.Response().Size).Interface("status", c.Response().Status).Msg("handler.user.ListUsers")
		}()
		payload := user.ListUsersRequest{}

		dataUsers, err := x.userUC.ListUsers(c.Request().Context(), payload)
		if err != nil {
			x.log.Err(err).Msg("[ListUsers][ListUsers]")

			return c.JSON(500, x.NewResponse(http.StatusInternalServerError, nil, payload, err, "handler.user.ListUsers"))
		}

		return c.JSON(http.StatusOK, Response{StatusCode: 200, Data: dataUsers})
		// return c.JSON(http.StatusOK, x.NewResponse(http.StatusOK, dataUsers, payload, err, "handler.user.register"))
	}
}

func (x *UserHandler) NewResponse(statusCode int, data, payload interface{}, err error, msg string) (resp responseBody) {
	defer func() {
		x.log.Err(err).Interface("request", payload).Interface("response", resp).Msg(msg)
	}()

	return responseBody{StatusCode: statusCode, Message: http.StatusText(statusCode), Error: err.Error(), Data: data}
}
