package contract

import (
	"context"
	"time"
)

type ItemRepository interface {
	GetItemByID(ctx context.Context, req GetItemByIDRequest) (resp GetItemByIDResponse, err error)
	ListItems(ctx context.Context, req ListItemsRequest) (resp ListItemsResponse, err error)
}

type Item struct {
	ID          int64
	ItemName    string
	Description string
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

type GetItemByIDRequest struct {
	ID int64
}

type GetItemByIDResponse Item

type ListItemsRequest struct{}

type ListItemsResponse []Item
