package main

import (
	r "rapid"
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

	r.Route("/hello/:FirstName/:LastName", func(c r.Connection) {
		c.Render("test2.html", c.Params)
	})

	r.Route("/hello", func(c r.Connection) {
		c.Send("Testing")
	})

	r.StaticFolder("static", "public")

	r.StartServer(3000)

	/*
		r.AddPath("/hello/world/")
		r.AddPath("/hello/world/hotdog")
		r.AddPath("/test/:Name/")
	*/
}
