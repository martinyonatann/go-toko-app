package http

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-invoice/internal/feature/user"
	"github.com/rs/zerolog"
)

type UserHandler struct {
	userUC user.UseCase
	log    zerolog.Logger
}

func New(s user.UseCase, log zerolog.Logger) *UserHandler {
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
		createUserPayload := user.CreateUserRequest{}

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

// func (x *UserHandler) GetUserByID() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 		if err != nil {
// 			x.log.Err(err).Msg("[handlerUser]GetUserByID_ParseInt")

// 			return err
// 		}

// 		getUserPayload := user.GetUserRequest{UserID: id}

// 		userData, err := x.userUC.GetUser(c.Request().Context(), getUserPayload)
// 		if err != nil {
// 			logger.Err(err).Msg("[handlerUser][getUserById]")

// 			return err
// 		}
// 		return c.JSON(http.StatusOK, responseBody{
// 			StatusCode: http.StatusOK,
// 			Message:    http.StatusText(http.StatusOK),
// 			Data:       userData,
// 		})
// 	}
// }

// func (x *UserHandler) ListUsers() echo.HandlerFunc {
// 	return func(c echo.Context) (err error) {
// 		defer func() {
// 			x.log.Err(err).Interface("size", c.Response().Size).Interface("status", c.Response().Status).Msg("handler.user.ListUsers")
// 		}()
// 		payload := user.ListUsersRequest{}

// 		dataUsers, err := x.userUC.ListUsers(c.Request().Context(), payload)
// 		if err != nil {
// 			x.log.Err(err).Msg("[ListUsers][ListUsers]")

// 			return c.JSON(500, responseBody{StatusCode: 500, Message: http.StatusText(500), Error: err.Error()})
// 		}

// 		return c.JSON(http.StatusOK, responseBody{StatusCode: 200, Message: http.StatusText(200), Data: dataUsers})
// 	}
// }

func (x *UserHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// tokenAuth := c.Request().Header.Get("Authorization")
		// if !strings.HasPrefix(tokenAuth, "Bearer") {
		// 	return c.JSON(http.StatusUnauthorized, responseBody{StatusCode: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized), Error: "invalid token"})
		// }

		// claims, err := utils.VerifyToken(strings.TrimPrefix(tokenAuth, "Bearer "))
		// if err != nil {
		// 	return c.JSON(http.StatusUnauthorized, responseBody{StatusCode: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized), Error: "token not found"})
		// }
		// return c.JSON(http.StatusOK, responseBody{StatusCode: 200, Data: claims, Message: http.StatusText(200)})
		payload := user.LoginRequest{}

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

/*
func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenAuth := c.Request().Header.Get("Authorization")
		if !strings.HasPrefix(tokenAuth, "Bearer") {
			return c.JSON(http.StatusUnauthorized, responseBody{StatusCode: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized), Error: "invalid token"})
		}

		claims, err := utils.VerifyToken(strings.TrimPrefix(tokenAuth, "Bearer "))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, responseBody{StatusCode: http.StatusUnauthorized, Message: http.StatusText(http.StatusUnauthorized), Error: "token not found"})
		}
		return c.JSON(http.StatusOK, responseBody{StatusCode: 200, Data: claims, Message: http.StatusText(200)})
	}
}
*/
