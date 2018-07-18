package actions

import (
	"errors"
	"log"
	"strconv"

	"github.com/gobuffalo/buffalo"
)

// CategoriesList default implementation.
func CategoriesList(c buffalo.Context) error {
	var err error
	var o, l int

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

	categories, err := categoriesRepo.List(o, l)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not get categories"))
	}
	return c.Render(200, r.JSON(categories))
}

// CategoriesIndex default implementation.
func CategoriesIndex(c buffalo.Context) error {
	id := c.Param("id")

	category, err := categoriesRepo.FindOne(id)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not get category"))
	}
	return c.Render(200, r.JSON(category))
}
