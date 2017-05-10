package main

import (
	r "github.com/Kiricon/Rapid"
	"github.com/Kiricon/Rapid/connection"
)

func main() {

	r.Get("/", func(c connection.Connection) {
		c.View("index.html")
	})

	r.Route("/test", func(c connection.Connection) {

		obj := struct {
			Name string
		}{"Dominic"}
		c.Render("test.html", obj)
	})

	r.Put("/Put", func(c connection.Connection) {
		c.Send("Put page")
	})

	r.Get("/Put", func(c connection.Connection) {
		c.Send("Put page - Accessed from GET")
	})

	r.Get("/hello/:FirstName/:LastName", func(c connection.Connection) {
		c.Render("test2.html", c.Params)
	})

	r.Route("/hello/foo/bar", func(c connection.Connection) {
		c.Render("test2.html", map[string]string{"FirstName": "Dominic", "LastName": "Balance"})
	})

	r.Route("/hello", func(c connection.Connection) {
		c.Send("Testing")
	})

	r.Route("/blah", func(c connection.Connection) {
		c.Redirect("/hello")
	})

	r.StaticFolder("/", "public")

	r.Listen(3000)
}
