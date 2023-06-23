package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/martinyonatann/go-toko-app/internal/services/app"
)

func main() {
	if os.Getenv("APP_ENV") == "local" {
		godotenv.Load()
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
