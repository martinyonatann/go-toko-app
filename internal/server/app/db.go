package app

import (
	"context"
	"fmt"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func (x *Server) initDB(ctx context.Context) (err error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	x.DB, err = sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		x.log.Err(err).Interface("dataSourceName", dataSourceName).Msg("initDB.Connect")
		panic(err)
	}

	x.DB.SetMaxOpenConns(maxOpenConns)
	x.DB.SetConnMaxLifetime(connMaxLifetime * time.Second)
	x.DB.SetMaxIdleConns(maxIdleConns)
	x.DB.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	return x.DB.PingContext(ctx)
}
