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

func Get(path string, handler routeHandler) {
	createRoute(path, handler, "GET")
}

// Route - Create a route for your webserver
func Route(path string, handler routeHandler) {
	createRoute(path, handler, "")
}

func createRoute(path string, handler routeHandler, method string) {
	paramLocations, path := getParamLocations(path)
	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {

		if method == "" || r.Method == method {
			requestPath := strings.Split(r.URL.Path, "/")
			params := map[string]string{}

			for i := 0; i < len(requestPath); i++ {
				if val, ok := paramLocations[i]; ok {
					params[val] = requestPath[i]
				}
			}

			resp := handler(Connection{r, w, params})

			if resp != "" {
				fmt.Fprintf(w, resp)
			}
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

func getParamLocations(path string) (map[int]string, string) {

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
	return params, newPath
}
