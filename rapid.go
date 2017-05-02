package rapid

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// Connection - struct for handling http request and write
type Connection struct {
	R      *http.Request
	W      http.ResponseWriter
	Params map[string]string
}

type routeHandler func(Connection) string

// Route - Create a route for your webserver
func Route(path string, handler routeHandler) {
	params, path := getParams(path)
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		resp := handler(Connection{r, w, params})

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

func getParams(path string, r *http.Request) (map[string]string, string) {

	routePath := strings.Split(path, "/")
	requestPath := strings.Split(r.URL.Path, "/")

	params := map[string]string{}
	for i := len(routePath) - 1; i >= 0; i-- {
		dir := routePath[i]
		if dir[0] == ':' {
			params[dir[1:len(routePath[i])-1]] = requestPath[i]
			routePath = append(routePath[:i], routePath[i+1:]...)
		}
	}

	return params, strings.Join(routePath, "/")
}
