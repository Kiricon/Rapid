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

// Send - Return string back to user
func (c *Connection) Send(message string) {
	fmt.Fprintf(c.W, message)
}

// View - Render HTML view without templating
func (c *Connection) View(path string) {
	c.Render(path, nil)
}

// Render - Render HTML view with templating
func (c *Connection) Render(path string, object interface{}) {
	t, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
	}
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(c.W, object)
}

// Redirect - Redirect a request to another rest end point
func (c *Connection) Redirect(path string) {
	http.Redirect(c.W, c.R, path, 301)
}

// StaticFolder - Specify application public/static folder
func StaticFolder(path string, dir string) {
	http.Handle("/"+path+"/", http.StripPrefix("/"+path+"/", http.FileServer(http.Dir(dir))))
}

// Listen - Start webserver on specified port
func Listen(port int) {
	portString := strconv.Itoa(port)
	http.ListenAndServe(":"+portString, nil)
}
