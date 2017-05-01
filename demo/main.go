package main

import (
	"rapid"
)

func main() {

	rapid.Route("/", func(c rapid.Connection) string {
		return "How are you doing?"
	})

	rapid.Route("/hello", func(c rapid.Connection) string {
		return "Testing"
	})

	rapid.Start(3000)
}
