package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	logging "log"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	"github.com/urfave/negroni"

	"github.com/martinyonatann/go-invoice/api/handler"
	"github.com/martinyonatann/go-invoice/api/middleware"
	"github.com/martinyonatann/go-invoice/infrastructure/database"
	"github.com/martinyonatann/go-invoice/infrastructure/repository/user_repository"
	"github.com/martinyonatann/go-invoice/pkg/metric"
	"github.com/martinyonatann/go-invoice/usecase/user"
)

func main() {
	db := database.DBConn()
	defer db.Close()

	userRepository := user_repository.NewUserRepository(db.DB)
	userService := user.New(userRepository)

	metricService, err := metric.NewPrometheusService()
	if err != nil {
		logging.Fatal(err)
	}

	r := mux.NewRouter()

	//handlers
	n := negroni.New(
		negroni.HandlerFunc(middleware.Cors),
		negroni.HandlerFunc(middleware.Metrics(metricService)),
		negroni.NewLogger(),
	)

	handler.MakeUserHandlers(r, *n, userService)

	http.Handle("/", r)
	http.Handle("/metrics", promhttp.Handler())
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthcheck success"))
	})

	logger := logging.New(os.Stderr, "logger: ", logging.Lshortfile)

	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Addr:         ":" + strconv.Itoa(5000),
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}
	err = srv.ListenAndServe()
	if err != nil {
		logging.Fatal(err)
	}

	log.Err(err).Msg("Server Listening on port" + strconv.Itoa(5000))
}
