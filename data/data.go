package data

import (
	"github.com/HouzuoGuo/tiedot/db"
)

var Db *db.DB

func InitDb(dir string) (err error) {
	// (Create if not exist) open a database
	if Db, err = db.OpenDB(dir); err != nil {
		return
	}
	return
}

func InitData() (err error) {
	if err = initProducts(); err != nil {
		return
	}
	return
}

func initProducts() (err error) {
	if err = Db.Create("Products"); err != nil {
		return
	}
	products := Db.Use("Products")
	products.Insert(map[string]interface{}{"Name": "iPhone 6s"})
	products.Insert(map[string]interface{}{"Name": "iPad 2"})
	products.Insert(map[string]interface{}{"Name": "Huawey P8"})
	products.Insert(map[string]interface{}{"Name": "Nokia 51"})
	products.Insert(map[string]interface{}{"Name": "Samsung 6"})
	return
}
