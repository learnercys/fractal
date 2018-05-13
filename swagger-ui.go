package main

import (
	"github.com/goadesign/goa"
)

// SwaggerUIController implements the swagger-ui resource.
type SwaggerUIController struct {
	*goa.Controller
}

// NewSwaggerUIController creates a swagger-ui controller.
func NewSwaggerUIController(service *goa.Service) *SwaggerUIController {
	return &SwaggerUIController{Controller: service.NewController("SwaggerUIController")}
}
