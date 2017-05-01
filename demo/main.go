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

	StartServer(3000)
}
