package main

import (
	//for extracting service credentials from VCAP_SERVICES
	//"github.com/cloudfoundry-community/go-cfenv"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/amsokol/kendo-ui-echo-samples/api/persistence"
	"github.com/amsokol/kendo-ui-echo-samples/api/rest"
	"github.com/amsokol/kendo-ui-echo-samples/fakedata"
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
	dir := "./.fakedata-db"
	os.RemoveAll(dir)
	defer os.RemoveAll(dir)

	var tiedot *db.DB
	var err error

	// init sample data
	if tiedot, err = db.OpenDB(dir); err != nil {
		log.Panic(err)
	}
	defer tiedot.Close()

	// init sample data
	if err = fakedata.AddData(tiedot); err != nil {
		log.Panic(err)
	}

	// init persistence stores
	rest.PR = data.GetProductReader(tiedot)

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
	e.Get("/products", rest.GetProducts)
}
