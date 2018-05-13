package main

import (
	"github.com/goadesign/goa"
)

// DocsController implements the docs resource.
type DocsController struct {
	*goa.Controller
}

// NewDocsController creates a docs controller.
func NewDocsController(service *goa.Service) *DocsController {
	return &DocsController{Controller: service.NewController("DocsController")}
}
