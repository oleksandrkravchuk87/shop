package populatedata

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"
	"github.com/shop/api/models"
)

const (
	csvReaderComma = ';'
)

var count int

func store(tableName string, in []byte) error {
	f, ok := storeRegistry[tableName]
	if !ok {
		return errors.Errorf("data to model mapping error on %v", tableName)
	}

	c, err := f(in)
	if err != nil {
		log.Println(err)
		return err
	}

	count += c

	return nil
}

// Import
func Import(dir string) error {

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = csvReaderComma
		return r
	})

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Println(err)
		return errors.Wrapf(err, "couldn't read directory '%v'", dir)
	}

	var csvData []byte

	for _, f := range files {
		fileName := strings.ToLower(f.Name())

		if !f.Mode().IsRegular() || !strings.HasSuffix(fileName, ".csv") {
			continue
		}

		tableName := strings.TrimSuffix(fileName, ".csv")

		if csvData, err = ioutil.ReadFile(filepath.Join(dir, f.Name())); err != nil {
			log.Println(err)
			break
		}

		if err = store(tableName, csvData); err != nil {
			log.Println(err)
			break
		}
	}

	if count == 0 {
		return errors.New("no data was imported")
	}

	return err
}

var storeItems = func(in []byte) (count int, err error) {
	out := make([]models.Item, 0)
	if err = gocsv.UnmarshalBytes(in, &out); err != nil {
		log.Println(err)
		return
	}

	for i := range out {
		if err = models.DB.Create(&out[i]); err != nil {
			log.Println(err)
			return
		}
		count++
	}

	return
}

var storeCategories = func(in []byte) (count int, err error) {
	out := make([]models.Category, 0)
	if err = gocsv.UnmarshalBytes(in, &out); err != nil {
		log.Println(err)
		return
	}

	for i := range out {
		if err = models.DB.Create(&out[i]); err != nil {
			log.Println(err)
			return
		}
		count++
	}

	return
}

type storeFunc func(in []byte) (count int, err error)

var storeRegistry = map[string]storeFunc{
	"categories": storeCategories,
	"items":      storeItems,
}
