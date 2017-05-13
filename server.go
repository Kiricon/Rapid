package rapid

import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

type routeHandler func(Connection)
type rapidHandler struct {
	server *Server
}

func (h rapidHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn := Connection{W: w, R: r, server: h.server}
	correctPath := h.server.findCorrectPath(r.URL.Path, r.Method)
	if correctPath != "404" {
		params := getParams(correctPath, r.URL.Path)

		if _, ok := h.server.pathHandlers[correctPath][r.Method]; ok {
			h.server.pathHandlers[correctPath][r.Method](Connection{R: r, W: w, Params: params, server: h.server})
		} else if _, ok := h.server.pathHandlers[correctPath]["ALL"]; ok {
			h.server.pathHandlers[correctPath]["ALL"](Connection{R: r, W: w, Params: params, server: h.server})
		} else {
			fileServer := h.server.findStaticServer(r.URL.Path)
			if fileServer.path != "" {
				h.server.staticServerHandler(fileServer)(Connection{R: r, W: w})
			} else {
				conn.NotFound()
			}
		}
	} else {
		fileServer := h.server.findStaticServer(r.URL.Path)
		if fileServer.path != "" {
			h.server.staticServerHandler(fileServer)(Connection{R: r, W: w})
		} else {
			conn.NotFound()
		}
	}
}

// Server - Individual instance of a rapid server
type Server struct {
	paths           map[string]path
	pathHandlers    map[string]map[string]func(Connection)
	srv             *http.Server
	wg              sync.WaitGroup
	handler         rapidHandler
	fileServerPaths []fileServerPath
	notFoundMessage string
}

// Get - Create http GET rest endpoint
func (s *Server) Get(path string, handler routeHandler) {
	s.createRoute(path, handler, "GET")
}

// Post - Create http POST rest endpoint
func (s *Server) Post(path string, handler routeHandler) {
	s.createRoute(path, handler, "POST")
}

// Put - Create http PUT rest endpoint
func (s *Server) Put(path string, handler routeHandler) {
	s.createRoute(path, handler, "PUT")
}

// Delete - Create http DELETE rest endpoint
func (s *Server) Delete(path string, handler routeHandler) {
	s.createRoute(path, handler, "DELETE")
}

// Route - Create a route for your webserver
func (s *Server) Route(path string, handler routeHandler) {
	s.createRoute(path, handler, "ALL")
}

func (s *Server) createRoute(path string, handler routeHandler, method string) {
	s.addPath(path, handler, method)
}

// StaticFolder - Specify application public/static folder
func (s *Server) StaticFolder(path string, dir string) {
	s.addStaticPath(path, dir)
}

// Listen - Start webserver on specified port
// Returns the instance of the http server currently runing.
// You can use this instance to shutdown the server if need be.
func (s *Server) Listen(port int) {
	s.ListenAndWait(port, true)
}

// ListenAndWait - Gives user option of waiting for server or not
func (s *Server) ListenAndWait(port int, wait bool) {
	portString := strconv.Itoa(port)
	s.handler.server = s

	s.srv = &http.Server{Addr: ":" + portString, Handler: s.handler}

	go func() {
		if err := s.srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	s.wg.Add(1)
	if wait {
		s.wg.Wait()
	}
}

// Shutdown - Gracefully shut down the server and unblock
// the server thread.
func (s *Server) Shutdown() {
	s.wg.Done()
	if err := s.srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
