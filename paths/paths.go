package paths

import (
	"strings"

	"fmt"

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

	pathArr := formatPath(pathString)

	if paths == nil {
		paths = make(map[string]path)
		pathHandlers = make(map[string]func(rapid.Connection))
	}

	insertPath(paths, pathArr, 0)
	pathHandlers[pathString] = handler

}

func formatPath(pathString string) []string {
	pathArr := strings.Split(pathString, "/")
	//loop := true

	for i := 0; i < len(pathArr); i++ {
		if i+1 != len(pathArr) {
			pathArr[i] += "/"
		}
	}

	return pathArr
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

// Find the correct path to asscociate with a request url
func findCorrectPath(path string) string {
	pathArr := formatPath(path)
	currentPath := paths
	lastMatch := "/"

	for i := 0; i < len(pathArr); i++ {
		dir := pathArr[i]

		if _, ok := currentPath[dir]; ok {
			lastMatch = currentPath[dir].path
			currentPath = currentPath[dir].subPaths
		} else if _, ok := currentPath["*"]; ok {
			lastMatch = currentPath["*"].path
			currentPath = currentPath["*"].subPaths
		} else if _, ok := currentPath["*/"]; ok {
			lastMatch = currentPath["*/"].path
			currentPath = currentPath["*/"].subPaths
		} else {
			fmt.Println("404")
			return "404 No Match"
		}
	}
	fmt.Println(lastMatch)
	return lastMatch

}
