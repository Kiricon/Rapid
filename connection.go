package rapid

import (
	"fmt"
	"html/template"
	"net/http"
)

// Connection - struct for handling http request and write
type Connection struct {
	R      *http.Request
	W      http.ResponseWriter
	Params map[string]string
	server *Server
}

// Send - Return plain text string back to http request
func (c *Connection) Send(message string) {
	fmt.Fprintf(c.W, message)
}

// View - Render HTML view without templating
func (c *Connection) View(path string) {
	c.Render(path, nil)
}

// Render - Render HTML view with templating
// Templating uses standard library templating
func (c *Connection) Render(path string, object interface{}) {
	t, _ := template.ParseFiles(path)
	c.W.Header().Set("Content-Type", "text/html; charset=utf-8")
	t.Execute(c.W, object)
}

// Redirect - Redirect a request to another rest end point
func (c *Connection) Redirect(path string) {
	http.Redirect(c.W, c.R, path, 301)
}

// NotFound - Return 404 message to user
func (c *Connection) NotFound() {
	if c.server.notFoundPage != "" {
		http.ServeFile(c.W, c.R, c.server.notFoundPage)
	} else {
		fmt.Fprintf(c.W, c.server.notFoundMessage)
	}
}
