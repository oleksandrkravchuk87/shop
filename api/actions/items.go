package actions

import (
	"errors"
	"log"
	"strconv"

	"github.com/gobuffalo/buffalo"
	"github.com/shop/api/models"
)

// ItemsList default implementation.
func ItemsList(c buffalo.Context) error {
	var err error
	var o, l int
	items := models.Items{}
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
	err = models.DB.Paginate(o, l).All(&items)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not get items"))
	}
	return c.Render(200, r.JSON(items))
}

// ItemsIndex default implementation.
func ItemsIndex(c buffalo.Context) error {
	item := models.Item{}
	id := c.Param("id")
	err := models.DB.Find(&item, id)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not get item"))
	}
	return c.Render(200, r.JSON(item))
}
