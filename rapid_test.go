package rapid

import (
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	Get("/", func(c Connection) {
		c.View("./demo/index.html")
	})

	Put("/put", func(c Connection) {
		c.Send("Hello")
	})

	Post("/post", func(c Connection) {
		c.Send("Hello")
	})

	Delete("/Delete", func(c Connection) {
		c.Send("Hello")
	})

	Route("/hello/:FirstName/:LastName", func(c Connection) {
		c.Render("./demo/test2.html", c.Params)
	})

	Route("/hello", func(c Connection) {
		c.Send("Testing")
	})

	Route("/blah", func(c Connection) {
		c.Redirect("/hello")
	})

	StaticFolder("static", "public")

	server := Listen(3000)

	_, gerr := http.NewRequest("GET", "/", nil)
	if gerr != nil {
		t.Fatal(gerr)
	}

	_, puterr := http.NewRequest("PUT", "/put", nil)
	if puterr != nil {
		t.Fatal(puterr)
	}

	_, posterr := http.NewRequest("POST", "/post", nil)
	if posterr != nil {
		t.Fatal(posterr)
	}

	_, deleteerr := http.NewRequest("DELETE", "/Delete", nil)
	if deleteerr != nil {
		t.Fatal(deleteerr)
	}

	time.Sleep(10 * time.Second)

	if err := server.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
