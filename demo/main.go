package main

import (
	. "rapid"
)

func main() {

	Route("/", func(c Connection) string {
		return c.View("index.html")
	})

	Route("/test", func(c Connection) string {

		obj := struct {
			Name string
		}{"Dominic"}
		return c.Render("test.html", obj)
	})

	Route("/hello", func(c Connection) string {
		return "Testing"
	})

	PublicFolder("public")

	StartServer(3000)
}
