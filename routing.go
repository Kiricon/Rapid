package rapid

import (
	"strings"
)

type routeHandler func(Connection)

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
	AddPath(path, handler, method)
}

func getParamLocations(path string) map[int]string {

	routePath := strings.Split(path, "/")

	params := map[int]string{}
	for i := (len(routePath) - 1); i >= 0; i-- {
		dir := routePath[i]
		if dir != "" && dir[0] == ':' {
			params[i] = dir[1:len(routePath[i])]
			routePath = append(routePath[:i], routePath[i+1:]...)
		}
	}
	newPath := strings.Join(routePath, "/")
	if len(params) > 0 {
		newPath += "/"
	}
	return params
}
