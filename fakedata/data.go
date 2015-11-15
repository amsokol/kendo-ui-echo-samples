package fakedata

import (
	"errors"
	"github.com/HouzuoGuo/tiedot/db"
)

func AddData(data *db.DB) (err error) {
	if err = addProducts(data); err != nil {
		return
	}
	return
}

func addProducts(data *db.DB) (err error) {
	if err = data.Create("Products"); err != nil {
		return
	}
	products := data.Use("Products")
	if products == nil {
		return errors.New("can't find 'Products' collection")
	}
	insert := func(data map[string]interface{}) {
		if err != nil {
			return
		}
		_, err = products.Insert(data)
	}
	insert(map[string]interface{}{"Name": "iPhone 6s"})
	insert(map[string]interface{}{"Name": "iPad 2"})
	insert(map[string]interface{}{"Name": "Huawey P8"})
	insert(map[string]interface{}{"Name": "Nokia 51"})
	insert(map[string]interface{}{"Name": "Samsung 6"})
	return
}
