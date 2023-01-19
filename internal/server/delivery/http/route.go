package http

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-invoice/internal/server/delivery"
)

type RoutePayload struct {
	Version *echo.Group
	Users   delivery.UsersHandlers
}

func InitRoute(request RoutePayload) {
	request.Version.POST("/user/register", request.Users.Register())
	request.Version.GET("/user/:id", request.Users.GetUserByID())
	request.Version.GET("/user", request.Users.ListUsers())
}
