package repository

import (
	"github.com/gobuffalo/pop"
	"github.com/shop/api/models"
)

type CategoryImpl struct {
	db *pop.Connection
}

func NewCategoryImpl() *CategoryImpl {
	return &CategoryImpl{db: models.DB}
}

func (c *CategoryImpl) FindOne(id string) (*models.Category, error) {
	category := models.Category{}
	err := c.db.Find(&category, id)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (c *CategoryImpl) List(offset, limit int) (models.Categories, error) {
	categories := models.Categories{}
	err := c.db.Paginate(offset, limit).All(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
