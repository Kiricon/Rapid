package rapid

import "net/http"

// Connection - struct for handling http request and write
type Connection struct {
	R *http.Request
	W http.ResponseWriter
}

type routeHandler func(Connection)

// Route - Create a route for your webserver
func Route(path string, handler routeHandler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(Connection{r, w})
	})
}

// Start - Start webserver on specified port
func Start(port string) {
	http.ListenAndServe(":"+port, nil)
}
