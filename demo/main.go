package main

import rapid "github.com/Kiricon/Rapid"

func main() {

	app := rapid.App()

	app.Get("/", func(c rapid.Connection) {
		c.View("index.html")
	})

	app.Route("/test", func(c rapid.Connection) {

		obj := struct {
			Name string
		}{"Dominic"}
		c.Render("test.html", obj)
	})

	app.Put("/Put", func(c rapid.Connection) {
		c.Send("Put page")
	})

	app.Get("/Put", func(c rapid.Connection) {
		c.Send("Put page - Accessed from GET")
	})

	app.Get("/hello/:FirstName/:LastName", func(c rapid.Connection) {
		c.Render("test2.html", c.Params)
	})

	app.Route("/hello/foo/bar", func(c rapid.Connection) {
		c.Render("test2.html", map[string]string{"FirstName": "Dominic", "LastName": "Balance"})
	})

	app.Route("/hello", func(c rapid.Connection) {
		c.Send("Testing")
	})

	app.Route("/blah", func(c rapid.Connection) {
		c.Redirect("/hello")
	})

	app.Get("/NotFound", func(c rapid.Connection) {
		c.NotFound()
	})

	//app.StaticFolder("/", "public")

	app.ListenAndWait(3000, false)

	app2 := rapid.App()

	app2.Get("/", func(c rapid.Connection) {
		c.Send("Server Two")
	})

	app2.Listen(3001)
}
