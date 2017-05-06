package rapid

import "testing"

func TestServer(t *testing.T) {
	Get("/", func(c Connection) {
		c.View("./demo/index.html")
	})

	Put("/put", func(c Connection) {
		c.Send("Hello")
	})

	Post("/post", func(c Connection) {
		c.Send("Hello")
	})

	Delete("/Delete", func(c Connection) {
		c.Send("Hello")
	})

	Route("/hello/:FirstName/:LastName", func(c Connection) {
		c.Render("./demo/test2.html", c.Params)
	})

	Route("/hello", func(c Connection) {
		c.Send("Testing")
	})

	Route("/blah", func(c Connection) {
		c.Redirect("/hello")
	})

	StaticFolder("static", "public")
}
