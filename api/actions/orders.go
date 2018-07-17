package actions

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/shop/api/models"
)

// OrdersCreate default implementation.
func OrdersCreate(c buffalo.Context) error {
	o := models.Order{Status: "created"}
	verrs, err := models.DB.ValidateAndCreate(&o)
	if verrs.Count() > 0 || err != nil {
		log.Println(verrs, err)
		return c.Error(500, errors.New("could not create new order"))
	}
	return c.Render(200, r.JSON(&o.ID))
}

// OrdersUpdate default implementation.
func OrdersUpdate(c buffalo.Context) error {
	order := models.Order{}
	id := c.Param("orderID")
	err := models.DB.Find(&order, id)
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
	err = models.DB.Transaction(func(tx *pop.Connection) error {
		verrs, err := tx.ValidateAndUpdate(&order)
		if verrs.Count() > 0 || err != nil {
			return errors.New("could not update order")
		}

		verrs, err = tx.ValidateAndCreate(&ordered)
		if verrs.Count() > 0 || err != nil {
			return errors.New("could not create new ordered")
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return c.Error(500, err)
	}
	return c.Render(200, r.JSON(&ordered.ID))
}
