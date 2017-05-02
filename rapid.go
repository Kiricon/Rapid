package rapid

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

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

		if resp != "" {
			fmt.Fprintf(w, resp)
		}
	})
}

func (c *Connection) View(path string) string {
	return c.Render(path, nil)
}

func (c *Connection) Render(path string, object interface{}) string {
	t, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
	}
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(c.W, object)
	return ""
}

func PublicFolder(path string) {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(path))))
}

// StartServer - Start webserver on specified port
func StartServer(port int) {
	portString := strconv.Itoa(port)
	http.ListenAndServe(":"+portString, nil)
}
