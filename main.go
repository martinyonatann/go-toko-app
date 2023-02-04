package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"
	"github.com/martinyonatann/go-toko-app/internal/services/app"
)

func main() {
	/*	fx.New(
			fx.Options(
				fx.Invoke(app.Run),
			),
		).Run()
	*/
	if os.Getenv("ENVIRONMENT") == "local" {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}
	}

	if err := app.Run(context.Background()); err != nil {
		panic(err)
	}
}
