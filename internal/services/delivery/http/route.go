package http

import (
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-toko-app/internal/services/delivery"
)

type (
	RoutePayload struct {
		Handlers delivery.Handlers
		Version  *echo.Group
	}
)

func InitRoute(request RoutePayload) {
	request.Version.POST("/user/register", request.Handlers.Register())
	request.Version.POST("/user/auth", request.Handlers.Login())

	request.Version.POST("/product/:id", request.Handlers.GetProductByID())
}
