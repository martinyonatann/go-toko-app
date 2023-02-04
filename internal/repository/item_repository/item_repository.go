package item_repository

import (
	"context"
	"database/sql"

	"github.com/martinyonatann/go-toko-app/internal/repository/contract"
	"github.com/martinyonatann/go-toko-app/internal/repository/postgresql_query"
)

func New(db *sql.DB) contract.ItemRepository {
	return &Repo{db: db}
}

type Repo struct {
	db *sql.DB
}

func (x *Repo) GetItemByID(
	ctx context.Context,
	req contract.GetItemByIDRequest) (
	resp contract.GetItemByIDResponse,
	err error) {
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

func (x *Repo) ListItems(ctx context.Context, req contract.ListItemsRequest) (resp contract.ListItemsResponse, err error) {
	query := postgresql_query.ListItems

	a, err := x.db.Query(query)
	if err != nil {
		return resp, err
	}

	defer a.Close()

	for a.Next() {
		item := contract.Item{}
		if err := a.Scan(&item.ID, &item.ItemName, &item.Description, &item.Price, &item.CreatedAt); err != nil {
			return resp, err
		}
		resp = append(resp, item)
	}

	return resp, err
}
