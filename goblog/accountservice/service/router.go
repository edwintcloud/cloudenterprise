package service

import "github.com/gorilla/mux"

// NewRouter is a function that returns a pointer to a mux.Router we an use as a handler.
func NewRouter() *mux.Router {

	// Create a new instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	// Iterate over the routes we declared in routes.go and attach them to the router instance
	for _, route := range routes {

		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
