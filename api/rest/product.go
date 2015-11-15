package rest

import (
	"github.com/amsokol/kendo-ui-echo-samples/api/domain"
	"github.com/amsokol/kendo-ui-echo-samples/api/kendoui"
	"github.com/amsokol/kendo-ui-echo-samples/api/logger"
	"github.com/labstack/echo"
	"net/http"
)

var ps domain.ProductStore

func SetProductStore(s domain.ProductStore) {
	ps = s
}

func GetProducts(c *echo.Context) error {
	i := kendoui.Input(c.Request())
	logger.Json(i)

	callback := c.Query("callback")
	if len(callback) == 0 {
		callback = "test"
	}

	products, err := ps.GetAll()
	if err == nil {
		err = c.JSONP(http.StatusOK, callback, products)
	}

	return err
}
