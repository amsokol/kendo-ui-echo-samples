package product

import (
	"github.com/amsokol/kendo-ui-echo-samples/Godeps/_workspace/src/github.com/labstack/echo"
	"github.com/amsokol/kendo-ui-echo-samples/kendoui"
	"log"
	"net/http"
)

type Product struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func GetProducts(c *echo.Context) error {
	i := kendoui.Input(c.Request())
	log.Println(i)

	callback := "asd" //c.Query("callback")
	return c.JSONP(http.StatusOK, callback, &Product{Id: "1234567890", Name: "qwertyuiop!!!"})
}
