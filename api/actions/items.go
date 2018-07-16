package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/shop/api/models"
	"fmt"
	"strconv"
)

// ItemsList default implementation.
func ItemsList(c buffalo.Context) error {
	items := models.Items{}
	off := c.Request().FormValue("offset")
	lim := c.Request().FormValue("limit")
	o, err :=strconv.Atoi(off)
	if err != nil{
		fmt.Println(err) // TODO
	}
	l,err :=strconv.Atoi(lim)
	if err != nil{
		fmt.Println(err) // TODO
	}
	err = models.DB.Paginate(o, l).All(items)
	if err != nil{
		fmt.Println(err) // TODO
	}
	return c.Render(200, r.JSON(items))
}

// ItemsIndex default implementation.
func ItemsIndex(c buffalo.Context) error {
	item := models.Item{}
	id := c.Param("id")
	err := models.DB.Find(&item, id)
	if err != nil{
		fmt.Println(err) // TODO
	}
		return c.Render(200, r.JSON(item))
}
