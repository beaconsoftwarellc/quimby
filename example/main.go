package main

import (
	"fmt"

	qcontrollers "github.com/Kasita-Inc/quimby/controllers"
	"github.com/Kasita-Inc/quimby/example/controllers"
	"github.com/Kasita-Inc/quimby/http"
)

func main() {
	rootController := &qcontrollers.HealthCheckController{}
	address := "localhost:8080"
	server := http.CreateRESTServer(address, rootController)
	server.Router.AddController(&qcontrollers.HealthCheckController{})
	server.Router.AddController(&controllers.ResourceController{})
	server.Router.AddController(&controllers.EchoController{})

	// API Controllers
	server.Router.AddController(&controllers.APIController{})
	storage := controllers.NewWidgetStorage()
	server.Router.AddController(controllers.NewWidgetController(storage))
	server.Router.AddController(controllers.NewWidgetsController(storage))
	fmt.Printf("Serving Example API on '%s'\n", address)
	fmt.Println("Registered routes are:")
	for _, s := range server.Router.RegisteredRoutes {
		fmt.Printf("\thttp://%s/%s\n", address, s)
	}
	fmt.Println("Listening...")
	server.ListenAndServe()
}
