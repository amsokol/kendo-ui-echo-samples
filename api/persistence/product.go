package data

import (
	"encoding/json"
	"github.com/HouzuoGuo/tiedot/db"
	"github.com/amsokol/kendo-ui-echo-samples/api/domain"
	"log"
)

type productStore struct {
	data *db.DB
}

func (p *productStore) GetAll() (products []domain.Product, err error) {
	products = make([]domain.Product, 0)
	col := p.data.Use("Products")
	if col != nil {
		col.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
			var product domain.Product
			err = json.Unmarshal(docContent, &product)
			if err == nil {
				product.Id = id
				products = append(products, product)
				willMoveOn = true
			} else {
				log.Println("failed to unmarshal Product", id, "is", string(docContent))
			}
			return
		})
	}
	return
}

func GetProductStore(data *db.DB) domain.ProductStore {
	return &productStore{data: data}
}
