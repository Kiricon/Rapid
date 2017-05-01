package main

import (
	"rapid"
)

func main() {
	rapid.Route("/", func(c rapid.Connection) string {
		return "Hello World"
	})

	rapid.Start(3000)
}
