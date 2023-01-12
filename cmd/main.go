package main

import (
	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/rs/zerolog/log"

	"github.com/martinyonatann/go-invoice/infrastructure/database"
	"github.com/martinyonatann/go-invoice/internal/server"
)

func main() {
	db := database.DBConn()
	defer db.Close()

	s := server.NewServer(database.DBConn())
	if err := s.Run(); err != nil {
		log.Err(err).Msg("[main]Run")
	}
}
