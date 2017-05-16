package main

import (
	"fmt"

	rapid "github.com/Kiricon/Rapid"
)

func main() {

	app := rapid.App()

	app.Get("/", func(c rapid.Connection) {
		c.Render("index.html", nil)
	})

	app.Post("/post", func(c rapid.Connection) {
		if c.Json != nil {
			fmt.Println(c.Json)
		}
		c.SendJSON(c.Json)
	})

	app.StaticFolder("/", "public")

	app.Listen(3000)

}
