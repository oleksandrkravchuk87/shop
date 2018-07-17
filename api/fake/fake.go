package fake

import (
	"time"

	"github.com/gobuffalo/uuid"
	"github.com/icrowley/fake"
	"github.com/shop/api/models"
)

func GenerateFakeItems() {
	for i := 0; i < 100; i++ {
		item := models.Item{ID: uuid.UUID{}, CreatedAt: time.Now(), Alias: fake.ProductName(), Title: fake.Product(),
			Desc: fake.Brand(), Pictures: fake.IPv4(), Price: fake.LatitudeSeconds(), Count: fake.WeekdayNum(),
			CategoryID: uuid.UUID{}}
		models.DB.Save(&item)
	}
}
