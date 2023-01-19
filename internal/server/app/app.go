package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	prom "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-invoice/internal/utils"
	"github.com/martinyonatann/go-invoice/pkg/metric"
	"github.com/rs/zerolog"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
)

// Server struct
type Server struct {
	DB   *sqlx.DB
	echo *echo.Echo
	log  zerolog.Logger
}

func Run(ctx context.Context) error {
	require := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
		"DB_USER",
		"DB_PASS",
		"APP_PORT",
	}

	if err := utils.Require(require...); err != nil {
		<-time.After(time.Second * 5)
		panic(err)
	}

	app := new(Server)
	app.echo = echo.New()
	app.log = zerolog.New(os.Stdout)

	defer utils.RecoverWith(func(v interface{}) {
		_, file, line, _ := runtime.Caller(3)
		app.log.Error().Str("file", file).Int("line", line).Msgf("panic: %v", v)
		<-time.After(time.Second * 5)
		panic(v)
	})

	_, err := metric.NewPrometheusService()
	if err != nil {
		app.log.Fatal()
	}

	p := prom.NewPrometheus("echo", nil)

	p.Use(app.echo)

	server := &http.Server{
		Addr:           ":" + os.Getenv("APP_PORT"),
		ReadTimeout:    time.Second * 5,
		WriteTimeout:   time.Second * 5,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		app.log.Info().Msg("Server is listening on PORT " + server.Addr)
		if err := app.echo.StartServer(server); err != nil {
			app.log.Fatal().Err(err).Msg("Error starting Server")
		}
	}()

	/*
		// for TCP
		go func() {
			app.log.Info().Msg("Server is listening on PORT :5555")
			if err := http.ListenAndServe(":5555", http.DefaultServeMux); err != nil {
				app.log.Fatal().Err(err).Msg("Error PPROF ListenAndServe")
			}
		}()
	*/

	if err := app.InitHandlers(ctx); err != nil {
		app.log.Err(err).Msg("app.InitHandlers")
		<-time.After(time.Second * 5)
		panic(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	app.log.Info().Msg("Server Exited Properly")
	return app.echo.Server.Shutdown(ctx)
}
