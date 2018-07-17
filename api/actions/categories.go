package actions

import (
	"errors"
	"log"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/shop/api/models"
)

// CategoriesList default implementation.
func CategoriesList(c buffalo.Context) error {
	var err error
	var o, l int
	categories := models.Categories{}
	off := c.Request().FormValue("offset")
	lim := c.Request().FormValue("limit")
	if off != "" && lim != "" {
		o, err = strconv.Atoi(off)
		if err != nil {
			log.Println(err)
			return c.Error(500, errors.New("could not handle offset parameter"))
		}
		l, err = strconv.Atoi(lim)
		if err != nil {
			log.Println(err)
			return c.Error(500, errors.New("could not handle limit parameter"))
		}
	}
	err = models.DB.Paginate(o, l).All(&categories)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not get categories"))
	}
	return c.Render(200, r.JSON(categories))

}

// CategoriesIndex default implementation.
func CategoriesIndex(c buffalo.Context) error {
	category := models.Category{}
	id := c.Param("id")
	err := models.DB.Find(&category, id)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not get categorie"))
	}
	return c.Render(200, r.JSON(category))
}
