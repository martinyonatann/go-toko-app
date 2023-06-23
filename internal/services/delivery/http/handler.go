package http

import (
	"github.com/martinyonatann/go-toko-app/internal/usecase/product_usecase"
	"github.com/martinyonatann/go-toko-app/internal/usecase/user_usecase"
	"github.com/rs/zerolog"
)

type (
	Handler struct {
		log zerolog.Logger
		UseCases
	}

	UseCases struct {
		UserUC    user_usecase.UserUC
		ProductUC product_usecase.ProductUC
	}

	responseBody struct {
		StatusCode int         `json:"rc"`
		Message    string      `json:"message"`
		Error      string      `json:"error,omitempty"`
		Data       interface{} `json:"data"`
	}
)

func New(log zerolog.Logger, uc UseCases) *Handler {
	return &Handler{log, uc}
}
