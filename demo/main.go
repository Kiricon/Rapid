package main

import "github.com/Kiricon/Rapid/templating"
import "fmt"

func main() {

	/*
		app := rapid.App()

		app.Get("/", func(c rapid.Connection) {
			c.View("index.html")
		})

		app.Post("/post", func(c rapid.Connection) {
			if c.Json != nil {
				fmt.Println(c.Json)
			}
			c.SendJSON(c.Json)
		})

		app.StaticFolder("/", "public")

		app.Listen(3000)
	*/

	fmt.Println(templating.Template("temp.html"))
}
