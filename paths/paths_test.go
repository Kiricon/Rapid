package paths

import (
	"testing"

	"github.com/Kiricon/Rapid"
)

func TestAddPath(t *testing.T) {
	AddPath("/Hello/World", func(c rapid.Connection) {})
	if len(paths) != 1 {
		t.Error("Path not added")
	}

	if _, ok := paths["/"].subPaths["Hello/"].subPaths["World"]; !ok {
		t.Error("Sub paths not building correctly")
	}

	AddPath("/Hello/:World", func(c rapid.Connection) {})

	if _, ok := paths["/"].subPaths["Hello/"].subPaths["*"]; !ok {
		t.Error("Sub paths not working with wild cards")
	}

	AddPath("/Foo/:Bar/", func(c rapid.Connection) {})

	if _, ok := paths["/"].subPaths["Foo/"].subPaths["*/"]; !ok {
		t.Error("Sub paths not working with wild cards")
	}
}
