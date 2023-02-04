package app

import (
	"context"
	"os"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/martinyonatann/go-toko-app/internal/repository/user_repository"
	"github.com/martinyonatann/go-toko-app/internal/services/delivery/http"
	"github.com/martinyonatann/go-toko-app/internal/usecase/user_usecase"
	"github.com/martinyonatann/go-toko-app/internal/utils"
	"github.com/rs/zerolog/log"
)

func (x *Server) InitHandlers(ctx context.Context) error {
	if err := x.initDB(ctx); err != nil {
		log.Err(err).Msg("app.DBConn")
	}

	// Init repositores
	userRepository := user_repository.New(x.DB.DB)

	// Init UseCases
	userService := user_usecase.New(userRepository, x.log)

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

	// Configure middleware with the custom claims type
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(utils.MyJWTClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(401, struct {
				StatusCode int    `json:"rc"`
				Message    string `json:"message"`
				Error      string `json:"error,omitempty"`
			}{
				401, "unauthorization", "token invalid or expired",
			},
			)
		},
	}

	health.Use(echojwt.WithConfig(config))

	health.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*utils.MyJWTClaims)
		name := claims.FullName
		return c.String(200, "Welcome "+name+"!")
	})

	return nil
}
