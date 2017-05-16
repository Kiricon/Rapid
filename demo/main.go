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
		if c.JSON != nil {
			fmt.Println(c.JSON)
		}
		c.SendJSON(c.JSON)
	})

	app.StaticFolder("/", "public")

	app.Listen(3000)

}
