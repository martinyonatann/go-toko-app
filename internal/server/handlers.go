package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/martinyonatann/go-invoice/internal/feature/user"
	user_repository "github.com/martinyonatann/go-invoice/internal/repository"
	"github.com/martinyonatann/go-invoice/internal/server/api"
	"github.com/martinyonatann/go-invoice/internal/server/handler"
)

func (x *Server) InitHandlers(e *echo.Echo) error {
	// Init repositores
	userRepository := user_repository.New(x.db.DB)

	// Init UseCases
	userService := user.New(userRepository)

	// Init Handlers
	userHandlers := handler.New(userService)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))

	e.Use(middleware.RequestID())

	e.Use(middleware.BodyLimit("2M"))

	v1 := e.Group("/v1")

	userGroup := v1.Group("/user")
	health := v1.Group("/health")

	api.InitRoute(userGroup, userHandlers)

	health.GET("", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
