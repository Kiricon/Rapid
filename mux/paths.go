package mux

import (
	"fmt"
	"net/http"
	"strings"
)

type Path struct {
	path     string
	method   string
	subPaths map[string]Path
}

// Paths - All paths recrusively registered
var Paths map[string]path

// PathHandlers - All path handler's keyed by their path string
var PathHandlers map[string]map[string]func(connection.Connection)

type fileServerPath struct {
	path string
	dir  string
}

var fileServerPaths []fileServerPath

// AddStaticPath - Add a static file server path
func AddStaticPath(path string, dir string) {

	if fileServerPaths == nil {
		fileServerPaths = []fileServerPath{}
	}

	fileServerPaths = append(fileServerPaths, fileServerPath{path, dir})
}

func StaticServerHandler(sPath fileServerPath) func(connection.Connection) {
	return func(c connection.Connection) {
		dirPath := strings.Replace(c.R.URL.Path, sPath.path, "", 1)
		dirPath = sPath.dir + "/" + dirPath
		fmt.Println(dirPath)
		http.ServeFile(c.W, c.R, dirPath)
	}
}

// AddPath - Add route to the map of paths
func AddPath(pathString string, handler func(connection.Connection), method string) {

	pathArr := formatPath(pathString)

	if Paths == nil {
		Paths = make(map[string]path)
		PathHandlers = make(map[string]map[string]func(connection.Connection))
	}

	insertPath(Paths, pathArr, 0, method)
	if PathHandlers[pathString] == nil {
		PathHandlers[pathString] = make(map[string]func(connection.Connection))
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
	singlePath = strings.ToLower(singlePath)

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

func FindStaticServer(path string) fileServerPath {

	for i := 0; i < len(fileServerPaths); i++ {
		sPath := fileServerPaths[i].path
		if path[0:len(sPath)] == sPath {
			return fileServerPaths[i]
		}
	}

	return fileServerPath{"", ""}
}
