package rapid

import "net/http"
import "strconv"
import "fmt"

// Connection - struct for handling http request and write
type Connection struct {
	R *http.Request
	W http.ResponseWriter
}

type routeHandler func(Connection) string

// Route - Create a route for your webserver
func Route(path string, handler routeHandler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		resp := handler(Connection{r, w})
		fmt.Fprintf(w, resp)
	})
}

// Start - Start webserver on specified port
func Start(port int) {
	portString := strconv.Itoa(port)
	http.ListenAndServe(":"+portString, nil)
}
