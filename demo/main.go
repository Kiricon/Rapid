package main

import (
	. "rapid"
)

type Temp struct {
	Name string
}

func main() {

	Route("/", func(c Connection) string {
		return c.View("index.html")
	})

	Route("/test", func(c Connection) string {

		obj := Temp{"Dominic"}
		return c.Render("test.html", obj)
	})

	Route("/hello", func(c Connection) string {
		return "Testing"
	})

	PublicFolder("public")

	StartServer(3000)
}
