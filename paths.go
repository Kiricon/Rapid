package rapid

import (
	"fmt"
	"net/http"
	"strings"
)

type path struct {
	path     string
	method   string
	subPaths map[string]path
}

type fileServerPath struct {
	path string
	dir  string
}

// AddStaticPath - Add a static file server path
func (s *Server) addStaticPath(path string, dir string) {

	if s.fileServerPaths == nil {
		s.fileServerPaths = []fileServerPath{}
	}

	s.fileServerPaths = append(s.fileServerPaths, fileServerPath{path, dir})
}

func (s *Server) staticServerHandler(sPath fileServerPath) func(Connection) {
	return func(c Connection) {
		dirPath := strings.Replace(c.R.URL.Path, sPath.path, "", 1)
		dirPath = sPath.dir + "/" + dirPath
		http.ServeFile(c.W, c.R, dirPath)
	}
}

// AddPath - Add route to the map of paths
func (s *Server) addPath(pathString string, handler func(Connection), method string) {

	pathArr := formatPath(pathString)
	fmt.Println(pathArr)
	if s.paths == nil {
		s.paths = make(map[string]path)
		s.pathHandlers = make(map[string]map[string]func(Connection))
	}

	s.insertPath(s.paths, pathArr, 0, method)
	if s.pathHandlers[pathString] == nil {
		s.pathHandlers[pathString] = make(map[string]func(Connection))
	}
	s.pathHandlers[pathString][method] = handler

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
func (s *Server) insertPath(paths map[string]path, pathArr []string, index int, method string) {

	singlePath := pathArr[index]
	singlePath = checkPathParams(singlePath)
	singlePath = strings.ToLower(singlePath)

	if _, ok := paths[singlePath]; ok && index+1 < len(pathArr) {
		s.insertPath(paths[singlePath].subPaths, pathArr, index+1, method)
	} else {
		emptySlice := make(map[string]path)
		paths[singlePath] = path{strings.Join(pathArr, ""), method, emptySlice}

		if index+1 < len(pathArr) {
			s.insertPath(paths[singlePath].subPaths, pathArr, index+1, method)
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
func (s *Server) findCorrectPath(path string, method string) string {
	path = strings.ToLower(path)
	pathArr := formatPath(path)
	currentPath := s.paths
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

func (s *Server) findStaticServer(path string) fileServerPath {

	for i := 0; i < len(s.fileServerPaths); i++ {
		sPath := s.fileServerPaths[i].path
		if path[0:len(sPath)] == sPath {
			return s.fileServerPaths[i]
		}
	}

	return fileServerPath{"", ""}
}
