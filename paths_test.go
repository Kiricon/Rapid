package rapid

import (
	"testing"

	"fmt"
)

func TestAddPath(t *testing.T) {
	AddPath("/Hello/World", func(c Connection) {})
	if len(paths) != 1 {
		t.Error("Path not added")
	}
	AddPath("/Hello/:World", func(c Connection) {})
	AddPath("/Foo/:Bar/", func(c Connection) {})

	AddPath("/Foo/:Bar/Baz/:test", func(c Connection) {})
	AddPath("/Foo/:Bar/Hello/:test", func(c Connection) {})

	fmt.Println("/Hello/World :" + findCorrectPath("/Hello/World"))
	fmt.Println("/Hello/Foo :" + findCorrectPath("/Hello/Foo"))
	fmt.Println("/Foo/Baz :" + findCorrectPath("/Foo/Baz"))
	fmt.Println("/Foo/one/Baz/two :" + findCorrectPath("/Foo/one/Baz/two"))
	fmt.Println("/Foo/one/Hello/two :" + findCorrectPath("/Foo/one/Hello/two"))
	fmt.Println("/America : " + findCorrectPath("/America"))
}
