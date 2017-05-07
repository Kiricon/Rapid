package rapid

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	Get("/", func(c Connection) {
		fmt.Println("GET:/")
		c.View("./demo/index.html")
	})

	Put("/put", func(c Connection) {
		fmt.Println("PUT:/put")
		c.Send("Hello")
	})

	Post("/post", func(c Connection) {
		fmt.Println("POST:/post")
		c.Send("Hello")
	})

	Delete("/delete", func(c Connection) {
		fmt.Println("DELETE:/Delete")
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

	if err := server.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
