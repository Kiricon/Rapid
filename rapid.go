package rapid

import "net/http"

type Connection struct {
	R *http.Request
	W http.ResponseWriter
}

type routeHandler func(Connection)

func Route(path string, handler routeHandler) {
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(Connection{r, w})
	})
}

func Start(port string) {
	http.ListenAndServe(":"+port, nil)
}
