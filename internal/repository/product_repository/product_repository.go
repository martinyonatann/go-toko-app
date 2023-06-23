package product_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinyonatann/go-toko-app/internal/repository/postgresql_query"
)

type (
	ProductRepository interface {
		GetProductByID(ctx context.Context, req GetProductByIDRequest) (resp GetProductByIDResponse, err error)
		ListProduct(ctx context.Context, req ListProductsRequest) (resp ListProductsResponse, err error)
	}

	Product struct {
		ProductID   int64
		ProductName string
		Description string
		Price       string
		CategoryID  int64
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   time.Time
	}

	GetProductByIDRequest struct {
		ID int64
	}

	GetProductByIDResponse Product

	ListProductsRequest struct{}

	ListProductsResponse []Product
)

func New(db *sql.DB) ProductRepository {
	return &Repo{db: db}
}

type Repo struct {
	db *sql.DB
}

func (x *Repo) GetProductByID(ctx context.Context, req GetProductByIDRequest) (resp GetProductByIDResponse, err error) {
	query := postgresql_query.GetProductByID

	if err = x.db.QueryRowContext(ctx, query, req.ID).Scan(
		&resp.ProductID,
		&resp.ProductName,
		&resp.Description,
		&resp.Price,
		&resp.CategoryID); err != nil {
		return resp, err
	}
	return resp, err
}

func (x *Repo) ListProduct(ctx context.Context, req ListProductsRequest) (resp ListProductsResponse, err error) {
	query := postgresql_query.ListProduct

	a, err := x.db.Query(query)
	if err != nil {
		return resp, err
	}

	defer a.Close()

	for a.Next() {
		Product := Product{}
		if err := a.Scan(&Product.ProductID, &Product.ProductName, &Product.Description, &Product.Price, &Product.CreatedAt); err != nil {
			return resp, err
		}
		resp = append(resp, Product)
	}

	return resp, err
}
