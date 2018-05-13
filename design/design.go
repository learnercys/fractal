package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("files", func() {
	Title("The files example")
	Description("An example to serve static assets")
})

var _ = Resource("swagger-ui", func() {
	Files("/ui/*filepath", "public/swagger-ui/")
})

var _ = Resource("swagger", func() {
	Origin("*", func() { // CORS policy that applies to all actions and file servers
		Methods("GET") // of "public" resource
	})

	Files("docs/swagger.json", "swagger/swagger.json")
})
