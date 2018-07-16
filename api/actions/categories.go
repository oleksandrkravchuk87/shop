package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/shop/api/models"
	"strconv"
	"fmt"
)

// CategoriesList default implementation.
func CategoriesList(c buffalo.Context) error {
	categories := models.Categories{}
	off := c.Request().FormValue("offset")
	lim := c.Request().FormValue("limit")
	o, err := strconv.Atoi(off)
	if err != nil{
		fmt.Println(err) // TODO
	}
	l,err :=strconv.Atoi(lim)
	if err != nil{
		fmt.Println(err) // TODO
	}
	err = models.DB.Paginate(o, l).All(categories)
	if err != nil{
		fmt.Println(err) // TODO
	}
	return c.Render(200, r.JSON(categories))


}

// CategoriesIndex default implementation.
func CategoriesIndex(c buffalo.Context) error {
	category := models.Category{}
	id := c.Param("id")
	err := models.DB.Find(&category, id)
	if err != nil{
		fmt.Println(err) // TODO
	}
	return c.Render(200, r.JSON(category))
}
