package repository

import (
	"errors"

	"github.com/gobuffalo/pop"
	"github.com/shop/api/models"
)

type OrderImpl struct {
	db *pop.Connection
}

func NewOrderImpl() *OrderImpl {
	return &OrderImpl{db: models.DB}
}

func (o *OrderImpl) CreateOrder(order *models.Order) (int, error) {
	verrs, err := o.db.ValidateAndCreate(order)
	if err != nil {
		return 0, err
	}
	if verrs.Count() > 0 {
		return 0, errors.New("validation error: " + verrs.Error())
	}
	return order.ID, nil
}

func (o *OrderImpl) FindOne(id string) (*models.Order, error) {
	order := models.Order{}
	if err := o.db.Find(&order, id); err != nil {
		return nil, err
	}
	return &order, nil
}

func (o *OrderImpl) Update(order *models.Order) (int, error) {
	verrs, err := o.db.ValidateAndUpdate(order)
	if err != nil {
		return 0, err
	}
	if verrs.Count() > 0 {
		return 0, errors.New("validation error: " + verrs.Error())
	}
	return order.ID, nil
}

func (o *OrderImpl) CreateOrdered(ordered *models.Ordered) (int, error) {
	verrs, err := o.db.ValidateAndCreate(ordered)
	if err != nil {
		return 0, err
	}
	if verrs.Count() > 0 {
		return 0, errors.New("validation error: " + verrs.Error())
	}
	return ordered.ID, nil
}
