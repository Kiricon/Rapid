package main

import (
	r "github.com/Kiricon/Rapid"
)

func main() {

	r.Get("/", func(c r.Connection) {
		c.View("index.html")
	})

	r.Route("/test", func(c r.Connection) {

		obj := struct {
			Name string
		}{"Dominic"}
		c.Render("test.html", obj)
	})

	r.Put("/Put", func(c r.Connection) {
		c.Send("Put page")
	})

	r.Get("/Put", func(c r.Connection) {
		c.Send("Put page - Accessed from GET")
	})

	r.Get("/hello/:FirstName/:LastName", func(c r.Connection) {
		c.Render("test2.html", c.Params)
	})

	r.Route("/hello/foo/bar", func(c r.Connection) {
		c.Render("test2.html", map[string]string{"FirstName": "Dominic", "LastName": "Balance"})
	})

	r.Route("/hello", func(c r.Connection) {
		c.Send("Testing")
	})

	r.Route("/blah", func(c r.Connection) {
		c.Redirect("/hello")
	})

	r.StaticFolder("/", "public")

	r.Listen(3000)
}
