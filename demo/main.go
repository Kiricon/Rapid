package main

import (
	"fmt"
	"rapid"
)

func main() {
	rapid.Route("/", func(c rapid.Connection) {
		fmt.Fprintf(c.W, "Hello World")
	})

	rapid.Start("3000")
}
