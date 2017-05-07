package paths

import (
	"testing"

	"fmt"

	"github.com/Kiricon/Rapid"
)

func TestAddPath(t *testing.T) {
	AddPath("/Hello/World", func(c rapid.Connection) {})
	if len(paths) != 1 {
		t.Error("Path not added")
	}
	AddPath("/Hello/:World", func(c rapid.Connection) {})
	AddPath("/Foo/:Bar/", func(c rapid.Connection) {})

	AddPath("/Foo/:Bar/Baz/:test", func(c rapid.Connection) {})
	AddPath("/Foo/:Bar/Hello/:test", func(c rapid.Connection) {})

	fmt.Println("/Hello/World :" + findCorrectPath("/Hello/World"))
	fmt.Println("/Hello/Foo :" + findCorrectPath("/Hello/Foo"))
	fmt.Println("/Foo/Baz :" + findCorrectPath("/Foo/Baz"))
	fmt.Println("/Foo/one/Baz/two :" + findCorrectPath("/Foo/one/Baz/two"))
	fmt.Println("/Foo/one/Hello/two :" + findCorrectPath("/Foo/one/Hello/two"))
	fmt.Println("/America : " + findCorrectPath("/America"))
}
