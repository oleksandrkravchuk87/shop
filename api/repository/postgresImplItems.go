package repository

import (
	"github.com/gobuffalo/pop"
	"github.com/shop/api/models"
)

type ItemImpl struct {
	db *pop.Connection
}

func NewItemImpl() *ItemImpl {
	return &ItemImpl{db: models.DB}
}

func (i *ItemImpl) FindOne(id string) (*models.Category, error) {
	category := models.Category{}
	err := i.db.Find(&category, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (i *ItemImpl) List(offset, limit int) (models.Categories, error) {
	categories := models.Categories{}
	err := i.db.Paginate(offset, limit).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
