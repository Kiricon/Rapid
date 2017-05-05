package rapid

import "testing"

func TestAddPath(t *testing.T) {
	AddPath("/Hello/World")
	if len(paths) != 1 {
		t.Error("Path not added")
	}

	if _, ok := paths["/"].subPaths["Hello/"].subPaths["World"]; !ok {
		t.Error("Sub paths not building correctly")
	}

	AddPath("/Hello/:World")

	if _, ok := paths["/"].subPaths["Hello/"].subPaths["*"]; !ok {
		t.Error("Sub paths not working with wild cards")
	}
}
