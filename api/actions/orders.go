package actions

import (
	"github.com/gobuffalo/buffalo"

	"fmt"

	"encoding/json"

	"github.com/shop/api/models"
)

// OrdersCreate default implementation.
func OrdersCreate(c buffalo.Context) error {
	decoder := json.NewDecoder(c.Request().Body)

	var o models.Order
	err := decoder.Decode(&o)

	if err != nil {
		panic(err)
	}
	fmt.Println(o)
	err = models.DB.Create(&o)
	if err != nil {
		panic(err)
	}
	return c.Render(200, r.JSON(&o.ID))
}

// OrdersUpdate default implementation.
func OrdersUpdate(c buffalo.Context) error {
	order := models.Order{}
	id := c.Param("orderID")
	err := models.DB.Find(&order, id)

	decoder := json.NewDecoder(c.Request().Body)

	var i models.Item
	err = decoder.Decode(&i)
	if err != nil {
		panic(err)
	}
	fmt.Println(i)

	order.Sum += i.Price * i.Count

	return c.Render(200, r.JSON("orders/update.html"))
}
