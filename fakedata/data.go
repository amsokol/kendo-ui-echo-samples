package fakedata

import (
	"errors"
	"github.com/amsokol/kendo-ui-echo-samples/api/data"
)

func AddData() (err error) {
	if err = addProducts(); err != nil {
		return
	}
	return
}

func addProducts() (err error) {
	if err = data.Db.Create("Products"); err != nil {
		return
	}
	products := data.Db.Use("Products")
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
