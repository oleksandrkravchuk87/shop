package populatedata

import (
	"log"
	"math/rand"
	"time"

	"github.com/gobuffalo/uuid"
	"github.com/icrowley/fake"
	"github.com/shop/api/models"
)

const (
	categoriesUpper = 10
	categoriesLower = 15

	itemsMin = 50
	itemsMax = 150
)

var CategoryUpperIDs []uuid.UUID
var CategoryLowerIDs []uuid.UUID

func GenerateFakeCategoryes() error {
	var err error
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < categoriesUpper; i++ {
		category := models.Category{
			Descr: fake.Brand(),
			Alias: fake.ProductName(),
			Title: fake.Product(),
			Logo:  fake.IPv4()}

		err = models.DB.Save(&category)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	err = models.DB.RawQuery("SELECT id FROM categories").All(&CategoryUpperIDs)
	if err != nil {
		log.Println(err)
		return err
	}

	for i := 0; i < categoriesLower; i++ {

		randInd := r1.Intn(len(CategoryUpperIDs))
		category := models.Category{
			Alias:    fake.ProductName(),
			Title:    fake.Product(),
			Descr:    fake.Brand(),
			Logo:     fake.IPv4(),
			ParentID: CategoryUpperIDs[randInd]}
		err = models.DB.Save(&category)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	log.Println("categirise were generated")
	return nil
}

func GenerateFakeItems() error {
	var err error

	err = models.DB.RawQuery("SELECT id FROM categories WHERE parent_id=?", uuid.UUID{}).All(&CategoryLowerIDs)
	if err != nil {
		log.Println(err)
		return err
	}
	for i := 0; i < len(CategoryLowerIDs); i++ {
		amount := random(itemsMin, itemsMax)
		for j := 0; j < amount; j++ {

			item := models.Item{
				Alias:      fake.ProductName(),
				Title:      fake.Product(),
				Descr:      fake.Brand(),
				Pictures:   fake.IPv4(),
				Price:      fake.LatitudeSeconds(),
				Count:      fake.WeekdayNum(),
				CategoryID: CategoryLowerIDs[i]}

			err = models.DB.Save(&item)
			if err != nil {
				log.Println(err)
				return err
			}
		}
	}
	log.Println("items were generated")
	return nil
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
