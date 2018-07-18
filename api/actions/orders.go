package actions

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/shop/api/models"
)

// OrdersCreate default implementation.
func OrdersCreate(c buffalo.Context) error {
	o := models.Order{Status: "created"}

	newOrdwerID, err := ordersRepo.CreateOrder(&o)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not create new order"))
	}
	return c.Render(200, r.JSON(newOrdwerID))
}

// OrdersUpdate default implementation.
func OrdersUpdate(c buffalo.Context) error {

	id := c.Param("orderID")
	order, err := ordersRepo.FindOne(id)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not find order"))
	}
	decoder := json.NewDecoder(c.Request().Body)
	var item models.Item
	err = decoder.Decode(&item)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not decode an item from request"))
	}
	order.Sum += item.Price * item.Count
	order.Status = "updated"
	ordered := models.Ordered{
		OrderID: order.ID,
		ItemID:  item.ID,
		ItemCnt: item.Count,
		ItemSum: item.Price * item.Count,
	}
	_, err = ordersRepo.Update(order)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not update order"))
	}
	newOrderedID, err := ordersRepo.CreateOrdered(&ordered)
	if err != nil {
		log.Println(err)
		return c.Error(500, errors.New("could not create new order"))
	}
	return c.Render(200, r.JSON(newOrderedID))
}
