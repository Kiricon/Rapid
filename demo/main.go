package main

import (
	. "rapid"
)

func main() {

	Route("/", func(c Connection) string {
		return View(c.W, "index.html")
	})

	Route("/hello", func(c Connection) string {
		return "Testing"
	})

	StartServer(3000)
}
