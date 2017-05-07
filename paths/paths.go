package paths

import (
	"strings"

	"github.com/Kiricon/Rapid"
)

type path struct {
	path     string
	subPaths map[string]path
}

var paths map[string]path
var pathHandlers map[string]func(rapid.Connection)

// AddPath - Add route to the map of paths
func AddPath(pathString string, handler func(rapid.Connection)) {
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
	pathHandlers[pathString] = handler

}

// Insert a path in to the global paths map
func insertPath(paths map[string]path, pathArr []string, index int) {

	singlePath := pathArr[index]
	singlePath = checkPathParams(singlePath)

	if _, ok := paths[singlePath]; ok && index+1 < len(pathArr) {
		insertPath(paths[singlePath].subPaths, pathArr, index+1)
	} else {
		emptySlice := make(map[string]path)
		paths[singlePath] = path{strings.Join(pathArr, ""), emptySlice}

		if index+1 < len(pathArr) {
			insertPath(paths[singlePath].subPaths, pathArr, index+1)
		}

	}

}

// Check if the current path directory is a url parameter
// If it is a param then replace it with  * wild card
func checkPathParams(singlePath string) string {
	if len(singlePath) > 1 {
		if singlePath[0] == ':' {
			if singlePath[len(singlePath)-1] == '/' {
				singlePath = "*/"
			} else {
				singlePath = "*"
			}
		}
	}

	return singlePath
}
