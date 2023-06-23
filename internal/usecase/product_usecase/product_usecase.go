package product_usecase

import (
	"context"
	"time"

	"github.com/martinyonatann/go-toko-app/internal/repository/product_repository"
	"github.com/rs/zerolog"
)

type (
	ProductUC interface {
		GetProductByID(context.Context, GetProductByIDRequest) (GetProductByIDResponse, error)
		ListProduct(context.Context, ListProductRequest) (ListProduct, error)
		InsertProduct(context.Context, InsertProductRequest) (ListProduct, error)
	}

	Product struct {
		ProductID   int64     `json:"id,omitempty"`
		ProductName string    `json:"name,omitempty"`
		Description string    `json:"description,omitempty"`
		Price       string    `json:"price,omitempty"`
		CategoryID  int64     `json:"category_id,omitempty"`
		CreatedAt   time.Time `json:"-"`
		UpdatedAt   time.Time `json:"-"`
		DeletedAt   time.Time `json:"-"`
	}

	ListProductRequest struct{}

	ListProduct []Product

	GetProductByIDRequest struct {
		ProductID int64
	}

	GetProductByIDResponse Product

	InsertProductRequest Product

	InsertProductResponse struct {
		Status string
	}

	UseCase struct {
		productRepository product_repository.ProductRepository
		log               zerolog.Logger
	}
)

func New(r product_repository.ProductRepository, log zerolog.Logger) ProductUC {
	return &UseCase{productRepository: r, log: log}
}

func (x *UseCase) GetProductByID(ctx context.Context, req GetProductByIDRequest) (res GetProductByIDResponse, err error) {
	defer func() {
		x.log.Err(err).Interface("res", res).Interface("req", req).Msg("[product_usecase]][GetProductByID]")
	}()

	product, err := x.productRepository.GetProductByID(ctx, product_repository.GetProductByIDRequest{ID: req.ProductID})
	if err != nil {
		x.log.Err(err).Msg("[product_usecase][GetProductByID][productRepository.GetProductByID]")

		return res, err
	}

	res.ProductID = product.ProductID
	res.ProductName = product.ProductName
	res.Description = product.Description
	res.Price = product.Price

	return res, err
}

func (x *UseCase) ListProduct(ctx context.Context, req ListProductRequest) (res ListProduct, err error) {
	return res, err
}

func (x *UseCase) InsertProduct(ctx context.Context, req InsertProductRequest) (res ListProduct, err error) {
	return res, err
}
