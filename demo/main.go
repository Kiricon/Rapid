package main

import (
	"rapid"
)

func main() {

	rapid.Route("/", func(c rapid.Connection) string {
		return rapid.View(c.W, "index.html")
	})

	rapid.Route("/hello", func(c rapid.Connection) string {
		return "Testing"
	})

	rapid.Start(3000)
}
