package rapid

import (
	"github.com/Kiricon/Rapid/connection"
	"github.com/Kiricon/Rapid/mux"
)

type routeHandler func(connection.Connection)

// Get - Create http GET rest endpoint
func Get(path string, handler routeHandler) {
	createRoute(path, handler, "GET")
}

// Post - Create http POST rest endpoint
func Post(path string, handler routeHandler) {
	createRoute(path, handler, "POST")
}

// Put - Create http PUT rest endpoint
func Put(path string, handler routeHandler) {
	createRoute(path, handler, "PUT")
}

// Delete - Create http DELETE rest endpoint
func Delete(path string, handler routeHandler) {
	createRoute(path, handler, "DELETE")
}

// Route - Create a route for your webserver
func Route(path string, handler routeHandler) {
	createRoute(path, handler, "ALL")
}

func createRoute(path string, handler routeHandler, method string) {
	mux.AddPath(path, handler, method)
}
