package main

import (
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"

	"github.com/amsokol/kendo-ui-echo-samples/data"
	"github.com/amsokol/kendo-ui-echo-samples/api/product"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"github.com/tylerb/graceful"
	"log"
	"os"
	"time"
)

const (
	DEFAULT_PORT = "8080"
	DEFAULT_HOST = "localhost"
)

func main() {
	var port string
	if port = os.Getenv("VCAP_APP_PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	var host string
	if host = os.Getenv("VCAP_APP_HOST"); len(host) == 0 {
		host = DEFAULT_HOST
	}

	// prepare folder for DB with sample data
	dir := "./.db"
	os.RemoveAll(dir)
	defer os.RemoveAll(dir)

	// init sample data
	if err := data.InitDb(dir); err != nil {
		log.Panic(err)
	}
	defer data.Db.Close()

	// init sample data
	if err := data.InitData(); err != nil {
		log.Panic(err)
	}

	// Echo instance
	e := echo.New()

	// Static test web pages
	e.Static("/", "website")
	e.Index("website/index.html")

	// Customization
	e.SetDebug(true)

	// Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	// Routes
	routes(e)

	s := e.Server(host + ":" + port)

	// HTTP2 is currently enabled by default in echo.New(). To override TLS handshake errors
	// you will need to override the TLSConfig for the server so it does not attempt to validate
	// the connection using TLS as required by HTTP2
	s.TLSConfig = nil

	// Start server
	log.Printf("Starting app on %+v:%+v\n", host, port)
	graceful.ListenAndServe(s, 5*time.Second)
}

func routes(e *echo.Echo) {
	e.Get("/products", product.GetProducts)
}
