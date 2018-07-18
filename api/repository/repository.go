package repository

import (
	"github.com/shop/api/models"
)

type CategoryRepo interface {
	FindOne(id string) (*models.Category, error)
	List(offset, limit int) (models.Categories, error)
}

type ItemRepo interface {
	FindOne(id string) (*models.Category, error)
	List(offset, limit int) (models.Categories, error)
}

type OrderRepo interface {
	CreateOrder(order *models.Order) (int, error)
	FindOne(id string) (*models.Order, error)
	Update(order *models.Order) (int, error)
	CreateOrdered(ordered *models.Ordered) (int, error)
}
