package rapid

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// Connection - struct for handling http request and write
type Connection struct {
	R      *http.Request
	W      http.ResponseWriter
	Params map[string]string
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
