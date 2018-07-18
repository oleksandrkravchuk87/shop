package grifts

import (
	"log"

	"github.com/markbates/grift/grift"
	"github.com/shop/api/populatedata"
)

const (
	dir = "./csvs"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		if err := grift.Run("db:seed:categories", c); err != nil {
			return err
		}
		if err := grift.Run("db:seed:items", c); err != nil {
			return err
		}
		return nil
	})
	grift.Add("seed:categories", func(c *grift.Context) error {
		err := populatedata.GenerateFakeCategoryes()
		if err != nil {
			log.Printf("filed to generate categories: %v\n", err)
			return err
		}
		return nil
	})
	grift.Add("seed:items", func(c *grift.Context) error {
		err := populatedata.GenerateFakeItems()
		if err != nil {
			log.Printf("filed to generate items: %v\n", err)
			return err
		}
		return nil
	})

	grift.Desc("import", "Imports data from csv to database")
	grift.Add("import", func(c *grift.Context) error {

		err := populatedata.Import(dir)
		if err != nil {
			log.Printf("filed to import items: %v\n", err)
			return err
		}
		return nil
	})

})
