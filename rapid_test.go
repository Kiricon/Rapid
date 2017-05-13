package rapid

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	app := App()
	app.Get("/", func(c Connection) {
		fmt.Println("GET:/")
		c.View("./demo/index.html")
	})

	app.Put("/put", func(c Connection) {
		fmt.Println("PUT:/put")
		c.Send("Hello")
	})

	app.Post("/post", func(c Connection) {
		fmt.Println("POST:/post")
		c.Send("Hello")
	})

	app.Delete("/delete", func(c Connection) {
		fmt.Println("DELETE:/Delete")
		c.Send("Hello")
	})

	app.Route("/hello/:FirstName/:LastName", func(c Connection) {
		c.Render("./demo/test2.html", c.Params)
	})

	app.Route("/hello", func(c Connection) {
		c.Send("Testing")
	})

	app.Route("/blah", func(c Connection) {
		c.Redirect("/hello")
	})

	app.StaticFolder("/static/", "./demo/public")

	app.ListenAndWait(3000, false)

	_, gerr := http.Get("http://localhost:3000/")
	if gerr != nil {
		t.Fatal(gerr)
	}

	_, template := http.Get("http://localhost:3000/hello/Dominic/Valenciana")
	if template != nil {
		t.Fatal(template)
	}
	_, redirect := http.Get("http://localhost:3000/blah")
	if redirect != nil {
		t.Fatal(redirect)
	}
	http.Get("http://localhost:3000/")

	_, puterr := http.NewRequest("PUT", "http://localhost:3000/put", nil)
	if puterr != nil {
		t.Fatal(puterr)
	}

	_, posterr := http.Post("http://localhost:3000/post", "", nil)
	if posterr != nil {
		t.Fatal(posterr)
	}

	_, deleteerr := http.NewRequest("DELETE", "http://localhost:3000/delete", nil)
	if deleteerr != nil {
		t.Fatal(deleteerr)
	}

	_, staticerr := http.NewRequest("GET", "http://localhost:3000/static/css/app.css", nil)
	if staticerr != nil {
		t.Fatal(staticerr)
	}

	_, wrongerr := http.NewRequest("GET", "http://localhost:3000/wrong/wrong", nil)
	if wrongerr != nil {
		t.Fatal(wrongerr)
	}

	time.Sleep(3 * time.Second)

	app.Shutdown()
}

func TestFailures(t *testing.T) {
	app := App()

	app.Get("/", func(c Connection) {
		c.Send("Hello World")
	})

	app.Get("/", func(c Connection) {
		c.Send("Hello World 2")
	})

	app.Listen(3000)
}
