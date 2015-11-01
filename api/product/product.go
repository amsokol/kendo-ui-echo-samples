package product

import (
	"encoding/json"
	"github.com/amsokol/kendo-ui-echo-samples/api/kendoui"
	"github.com/amsokol/kendo-ui-echo-samples/api/logger"
	"github.com/amsokol/kendo-ui-echo-samples/data"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type Product struct {
	Id   int    `json:"Id,omitempty"`
	Name string `json:"Name,omitempty"`
}

func GetProducts(c *echo.Context) (err error) {
	i := kendoui.Input(c.Request())
	logger.Json(i)

	callback := "asd" //c.Query("callback")

	products := make([]Product, 0)

	col := data.Db.Use("Products")
	if col != nil {
		col.ForEachDoc(func(id int, docContent []byte) (willMoveOn bool) {
			var p Product
			if err := json.Unmarshal(docContent, &p); err == nil {
				p.Id = id
				products = append(products, p)
			} else {
				log.Println("failed to unmarshal Product", id, "is", string(docContent))
			}
			return true
		})
	}
	return c.JSONP(http.StatusOK, callback, products)
}
