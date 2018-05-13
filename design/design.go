package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Title("The virtual wine cellar")    // other global properties. There should be one
	Description("A simple goa service") // and exactly one API definition appearing in
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

var _ = Resource("bottle", func() { // Resources group related API endpoints
	BasePath("/bottles")      // together. They map to REST resources for REST
	DefaultMedia(BottleMedia) // services.

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

// BottleMedia defines the media type used to render bottles.
var BottleMedia = MediaType("application/vnd.goa.example.bottle+json", func() {
	Description("A bottle of wine")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("id", Integer, "Unique bottle ID")
		Attribute("href", String, "API href for making requests on the bottle")
		Attribute("name", String, "Name of wine")
		Required("id", "href", "name")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("id")   // Media types may have multiple views and must
		Attribute("href") // have a "default" view.
		Attribute("name")
	})
})
