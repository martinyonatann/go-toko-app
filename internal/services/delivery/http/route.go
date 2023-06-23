package http

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-toko-app/internal/services/delivery"
)

type RoutePayload struct {
	Version *echo.Group
	Users   delivery.UsersHandlers
}

func InitRoute(request RoutePayload) {
	request.Version.POST("/user/register", request.Users.Register())
	request.Version.POST("/user/auth", request.Users.Login())
}
