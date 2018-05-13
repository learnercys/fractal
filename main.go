//go:generate goagen bootstrap -d github.com/learnercys/fractal/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/learnercys/fractal/app"
)

func main() {
	// Create service
	service := goa.New("files")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "swagger" controller
	c := NewSwaggerController(service)
	app.MountSwaggerController(service, c)
	// Mount "swagger-ui" controller
	c2 := NewSwaggerUIController(service)
	app.MountSwaggerUIController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}