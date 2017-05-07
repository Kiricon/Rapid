package rapid

import (
	"strings"
)

type path struct {
	path     string
	method   string
	subPaths map[string]path
}

// Paths - All paths recrusively registered
var Paths map[string]path

// PathHandlers - All path handler's keyed by their path string
var PathHandlers map[string]map[string]func(Connection)

// AddPath - Add route to the map of paths
func AddPath(pathString string, handler func(Connection), method string) {

	pathString = strings.ToLower(pathString)
	pathArr := formatPath(pathString)

	if Paths == nil {
		Paths = make(map[string]path)
		PathHandlers = make(map[string]map[string]func(Connection))
	}

	insertPath(Paths, pathArr, 0, method)
	if PathHandlers[pathString] == nil {
		PathHandlers[pathString] = make(map[string]func(Connection))
	}
	PathHandlers[pathString][method] = handler

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
func insertPath(paths map[string]path, pathArr []string, index int, method string) {

	singlePath := pathArr[index]
	singlePath = checkPathParams(singlePath)

	if _, ok := paths[singlePath]; ok && index+1 < len(pathArr) {
		insertPath(paths[singlePath].subPaths, pathArr, index+1, method)
	} else {
		emptySlice := make(map[string]path)
		paths[singlePath] = path{strings.Join(pathArr, ""), method, emptySlice}

		if index+1 < len(pathArr) {
			insertPath(paths[singlePath].subPaths, pathArr, index+1, method)
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

// FindCorrectPath - Find the correct path to asscociate with a request url
func FindCorrectPath(path string, method string) string {
	path = strings.ToLower(path)
	pathArr := formatPath(path)
	currentPath := Paths
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
			return "404"
		}
	}

	return lastMatch
}
