package rapid

import (
	"github.com/Kiricon/Rapid/connection"
	"github.com/Kiricon/Rapid/mux"
)

type Server struct {
	Paths        map[string]mux.Path
	PathHandlers map[string]map[string]func(connection.Connection)
}

func App() Server {
	return Server{}
}

// Get - Create http GET rest endpoint
func (s *Server) Get(path string, handler routeHandler) {
	createRoute(path, handler, "GET")
}

// Post - Create http POST rest endpoint
func (s *Server) Post(path string, handler routeHandler) {
	createRoute(path, handler, "POST")
}

// Put - Create http PUT rest endpoint
func (s *Server) Put(path string, handler routeHandler) {
	createRoute(path, handler, "PUT")
}

// Delete - Create http DELETE rest endpoint
func (s *Server) Delete(path string, handler routeHandler) {
	createRoute(path, handler, "DELETE")
}

// Route - Create a route for your webserver
func (s *Server) Route(path string, handler routeHandler) {
	createRoute(path, handler, "ALL")
}
