package main

// THIS IS A GENERATED FILE. DO NOT MODIFY
// main.tmpl

import (
	qcontrollers "github.com/beaconsoftwarellc/quimby/controllers"
	"github.com/beaconsoftwarellc/quimby/example/config"
	qhttp "github.com/beaconsoftwarellc/quimby/http"
	"github.com/beaconsoftwarellc/quimby/example/controllers"
    "github.com/beaconsoftwarellc/gadget/log"
	"fmt"
)

//go:generate codegen definition.yaml

func main() {
	log.NewGlobal("ExampleGateway", log.FunctionFromEnv())
	// Constants are defined on http
	// see: https://golang.org/pkg/net/http/#
	specification := config.New()
	rootController := &qcontrollers.HealthCheckController{}
	server := qhttp.CreateRESTServer(fmt.Sprintf(":%d", specification.Port), rootController)
	server.Router.AddController(&qcontrollers.HealthCheckController{})
	server.Router.AddController(controllers.NewDocController(specification))

	server.Router.AddController(controllers.NewEchoController(specification))
	server.Router.AddController(controllers.NewResourceController(specification))
	server.Router.AddController(controllers.NewWidgetsController(specification))
	server.Router.AddController(controllers.NewWidgetController(specification))
	
	log.Infof("Server starting ... http://localhost:%d/", specification.Port)
	log.Error(server.ListenAndServe())
}
