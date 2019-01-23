package main

// THIS IS A GENERATED FILE. DO NOT MODIFY

import (
	"fmt"

	"gitlab.com/beacon-software/gadget/log"
	qcontrollers "gitlab.com/beacon-software/quimby/controllers"
	"gitlab.com/beacon-software/quimby/example/config"
	"gitlab.com/beacon-software/quimby/example/controllers"
	qhttp "gitlab.com/beacon-software/quimby/http"
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

	server.Router.AddController(controllers.NewWidgetsController(specification))
	server.Router.AddController(controllers.NewWidgetController(specification))

	log.Infof("Server starting ... http://localhost:%d/", specification.Port)
	log.Error(server.ListenAndServe())
}
