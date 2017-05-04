package rapid

import (
	"fmt"
	"strings"
)

type path struct {
	path     string
	subPaths map[string]path
}

var paths map[string]path

func AddPath(pathString string) {
	pathArr := strings.Split(pathString, "/")
	//loop := true

	for i := 0; i < len(pathArr); i++ {
		if i+1 != len(pathArr) {
			pathArr[i] += "/"
		}
	}

	if paths == nil {
		paths = make(map[string]path)
	}

	insertPath(paths, pathArr, 0)
	fmt.Println(paths)

}

func insertPath(paths map[string]path, pathArr []string, index int) {

	if _, ok := paths[pathArr[index]]; ok && index+1 < len(pathArr) {
		insertPath(paths, pathArr, index+1)
	} else {
		emptySlice := make(map[string]path)
		paths[pathArr[index]] = path{pathArr[index], emptySlice}

		if index+1 < len(pathArr) {
			insertPath(paths[pathArr[index]].subPaths, pathArr, index+1)
		}

	}

}
