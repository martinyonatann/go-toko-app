package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/martinyonatann/go-invoice/internal/server/app"
)

func main() {
	if os.Getenv("ENVIRONMENT") == "local" {
		godotenv.Load()
	}

	app.Run(context.Background())
}
