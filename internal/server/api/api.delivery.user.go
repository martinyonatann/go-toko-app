package api

import "github.com/labstack/echo/v4"

type UsersHandlers interface {
	Register() echo.HandlerFunc
	GetUserByID() echo.HandlerFunc
}
