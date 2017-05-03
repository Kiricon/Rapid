package main

import (
	. "rapid"
)

func main() {

	Get("/", func(c Connection) string {
		return c.View("index.html")
	})

	Route("/test", func(c Connection) string {

		obj := struct {
			Name string
		}{"Dominic"}
		return c.Render("test.html", obj)
	})

	Route("/hello/:FirstName/:LastName", func(c Connection) string {
		return c.Render("test2.html", c.Params)
	})

	Route("/hello", func(c Connection) string {
		return "Testing"
	})

	PublicFolder("public")

	StartServer(3000)
}
