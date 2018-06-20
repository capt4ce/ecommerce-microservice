package main

import (
	"log"
	"net/http"

	"github.com/capt4ce/ecommerce-microservice/products/common"
	"github.com/capt4ce/ecommerce-microservice/products/routers"
)

// Entry point for the program
func main() {

	// Calls startup logic
	common.StartUp()
	// Get the mux router object
	router := routers.InitRoutes()

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: router,
	}
	log.Println("Listening [products]...")
	server.ListenAndServe()
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// 	"github.com/sirupsen/logrus"
// )

// var log = logrus.New()

// func init() {
// 	log.Formatter = new(logrus.JSONFormatter)
// 	log.Formatter = new(logrus.TextFormatter) // default
// 	log.Level = logrus.DebugLevel
// }

// func main() {
// 	fmt.Println("Main function :")
// 	e := echo.New()
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())
// 	// Routes
// 	e.GET("/go-docker", goWithDocker)
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// func goWithDocker(c echo.Context) error {
// 	return c.JSON(http.StatusOK, "Go with Docker Container")
// }
