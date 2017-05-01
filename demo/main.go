package main

import (
	. "rapid"
)

func main() {

	Route("/", func(c Connection) string {
		return c.View("index.html")
	})

	Route("/hello", func(c Connection) string {
		return "Testing"
	})

	PublicFolder("public")

	StartServer(3000)
}
