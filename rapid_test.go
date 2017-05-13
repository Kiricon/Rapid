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

	app.StaticFolder("static", "public")

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

	time.Sleep(3 * time.Second)

	app.Shutdown()
}
