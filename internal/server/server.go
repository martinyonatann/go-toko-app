package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jmoiron/sqlx"
	prom "github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-invoice/pkg/metric"
	"github.com/rs/zerolog"
)

// Server struct
type Server struct {
	echo *echo.Echo
	db   *sqlx.DB
	log  zerolog.Logger
}

func NewServer(db *sqlx.DB) *Server {
	return &Server{echo: echo.New(), db: db, log: zerolog.New(os.Stdout)}
}

func (s *Server) Run() error {
	server := &http.Server{
		Addr:         ":5000",
		ReadTimeout:  5,
		WriteTimeout: 5,
	}

	_, err := metric.NewPrometheusService()
	if err != nil {
		s.log.Fatal()
	}

	p := prom.NewPrometheus("echo", nil)

	p.Use(s.echo)

	go func() {
		if err := s.echo.Start(server.Addr); err != nil {
			s.log.Err(err).Msg("[Run]Start")
		}
	}()

	go func() {
		if err := http.ListenAndServe(":5555", http.DefaultServeMux); err != nil {
			s.log.Err(err).Msg("[Run]ListenAndServe")
		}
	}()

	if err := s.InitHandlers(s.echo); err != nil {
		return err
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGALRM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	s.log.Info().Msg("Server Exited Properly")
	return s.echo.Server.Shutdown(ctx)

}
