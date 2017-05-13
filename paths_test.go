package rapid

import (
	"testing"

	"fmt"
)

func TestAddPath(t *testing.T) {
	AddPath("/Hello/World", func(c Connection) {}, "GET")
	if len(Paths) != 1 {
		t.Error("Path not added")
	}
	AddPath("/Hello/:World", func(c Connection) {}, "GET")
	AddPath("/Foo/:Bar/", func(c Connection) {}, "GET")

	AddPath("/Foo/:Bar/Baz/:test", func(c Connection) {}, "GET")
	AddPath("/Foo/:Bar/Hello/:test", func(c Connection) {}, "GET")

	fmt.Println("/Hello/World :" + FindCorrectPath("/Hello/World", "GET"))
	fmt.Println("/Hello/Foo :" + FindCorrectPath("/Hello/Foo", "GET"))
	fmt.Println("/Foo/Baz :" + FindCorrectPath("/Foo/Baz", "GET"))
	fmt.Println("/Foo/one/Baz/two :" + FindCorrectPath("/Foo/one/Baz/two", "GET"))
	fmt.Println("/Foo/one/Hello/two :" + FindCorrectPath("/Foo/one/Hello/two", "GET"))
	fmt.Println("/America : " + FindCorrectPath("/America", "GET"))
}
