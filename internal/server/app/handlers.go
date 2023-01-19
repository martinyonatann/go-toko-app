package app

import (
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/martinyonatann/go-invoice/internal/feature/user"
	user_repository "github.com/martinyonatann/go-invoice/internal/repository"
	"github.com/martinyonatann/go-invoice/internal/server/delivery/http"
	"github.com/rs/zerolog/log"
)

func (x *Server) InitHandlers(ctx context.Context) error {
	if err := x.initDB(ctx); err != nil {
		log.Err(err).Msg("app.DBConn")
	}

	// Init repositores
	userRepository := user_repository.New(x.DB.DB)

	// Init UseCases
	userService := user.New(userRepository, x.log)

	// Init Handlers
	userHandlers := http.New(userService, x.log)

	x.echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXRequestID},
	}))

	x.echo.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:         1 << 10,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))

	x.echo.Use(middleware.RequestID())

	x.echo.Use(middleware.BodyLimit("2M"))

	v1 := x.echo.Group("/v1")

	health := v1.Group("/health")

	http.InitRoute(http.RoutePayload{
		Version: v1,
		Users:   userHandlers,
	})

	health.GET("", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "OK"})
	})

	return nil
}
