package item_repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/martinyonatann/go-toko-app/internal/repository/postgresql_query"
)

type (
	ItemRepository interface {
		GetItemByID(ctx context.Context, req GetItemByIDRequest) (resp GetItemByIDResponse, err error)
		ListItems(ctx context.Context, req ListItemsRequest) (resp ListItemsResponse, err error)
	}

	Item struct {
		ID          int64
		ItemName    string
		Description string
		Price       float64
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   time.Time
	}

	GetItemByIDRequest struct {
		ID int64
	}

	GetItemByIDResponse Item

	ListItemsRequest struct{}

	ListItemsResponse []Item
)

func New(db *sql.DB) ItemRepository {
	return &Repo{db: db}
}

type Repo struct {
	db *sql.DB
}

func (x *Repo) GetItemByID(ctx context.Context, req GetItemByIDRequest) (resp GetItemByIDResponse, err error) {
	query := postgresql_query.GetItemByID

	if err = x.db.QueryRowContext(ctx, query, req.ID).Scan(
		&resp.ID,
		&resp.ItemName,
		&resp.Description,
		&resp.Price,
		&resp.CreatedAt); err != nil {
		return resp, err
	}
	return resp, err
}

func (x *Repo) ListItems(ctx context.Context, req ListItemsRequest) (resp ListItemsResponse, err error) {
	query := postgresql_query.ListItems

	a, err := x.db.Query(query)
	if err != nil {
		return resp, err
	}

	defer a.Close()

	for a.Next() {
		item := Item{}
		if err := a.Scan(&item.ID, &item.ItemName, &item.Description, &item.Price, &item.CreatedAt); err != nil {
			return resp, err
		}
		resp = append(resp, item)
	}

	return resp, err
}
