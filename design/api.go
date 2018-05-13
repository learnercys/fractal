package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("fractal", func() {
	Title("The virtual wine cellar")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
	Version("1")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("swagger-ui", func() {
	Files("/ui/*filepath", "public/swagger-ui/")
})

var _ = Resource("docs", func() {
	Origin("*", func() { // CORS policy that applies to all actions and file servers
		Methods("GET") // of "public" resource
	})

	Files("docs/swagger.json", "swagger/swagger.json")
})

var _ = Resource("bottle", func() { // Resources group related API endpoints
	BasePath("/bottles")          // together. They map to REST resources for REST
	DefaultMedia(BottleMediaType) // services.

	Action("show", func() { // Actions define a single API endpoint together
		Description("Get bottle by id") // with its path, parameters (both path
		Routing(GET("/:bottleID"))      // parameters and querystring values) and payload
		Params(func() {                 // (shape of the request body).
			Param("bottleID", Integer, "Bottle ID")
		})
		Response(OK)       // Responses define the shape and status code
		Response(NotFound) // of HTTP responses.
	})
})
