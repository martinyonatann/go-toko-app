package api

import (
	"github.com/labstack/echo/v4"
)

func InitRoute(g *echo.Group, h UsersHandlers) {
	g.POST("/register", h.Register())
	g.GET("/:id", h.GetUserByID())
}
