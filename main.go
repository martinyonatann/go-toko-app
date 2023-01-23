package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/martinyonatann/go-invoice/internal/server/app"
)

func main() {
	if os.Getenv("ENVIRONMENT") == "local" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
