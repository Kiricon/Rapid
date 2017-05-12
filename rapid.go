package rapid

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Kiricon/Rapid/mux"
)

type routeHandler func(connection.Connection)

type Server struct {
	paths        map[string]mux.Path
	pathHandlers map[string]map[string]func(connection.Connection)
	srv          *http.Server
	wg           sync.WaitGroup
}

func App() Server {
	return Server{}
}

// Get - Create http GET rest endpoint
func (s *Server) Get(path string, handler routeHandler) {
	createRoute(path, handler, "GET")
}

// Post - Create http POST rest endpoint
func (s *Server) Post(path string, handler routeHandler) {
	createRoute(path, handler, "POST")
}

// Put - Create http PUT rest endpoint
func (s *Server) Put(path string, handler routeHandler) {
	createRoute(path, handler, "PUT")
}

// Delete - Create http DELETE rest endpoint
func (s *Server) Delete(path string, handler routeHandler) {
	createRoute(path, handler, "DELETE")
}

// Route - Create a route for your webserver
func (s *Server) Route(path string, handler routeHandler) {
	createRoute(path, handler, "ALL")
}

func (s *Server) createRoute(path string, handler routeHandler, method string) {
	mux.AddPath(path, handler, method)
}

// StaticFolder - Specify application public/static folder
func (s *Server) StaticFolder(path string, dir string) {
	mux.AddStaticPath(path, dir)
}

// Listen - Start webserver on specified port
// Returns the instance of the http server currently runing.
// You can use this instance to shutdown the server if need be.
func (s *Server) Listen(port int) {
	ListenAndWait(port, true)
}

// ListenAndWait - Gives user option of waiting for server or not
func (s *Server) ListenAndWait(port int, wait bool) {
	portString := strconv.Itoa(port)
	s.srv = &http.Server{Addr: ":" + portString, Handler: mux.RapidHandler{}}

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

// ShutdownServer - Gracefully shut down the server and unblock
// the server thread.
func (s *Server) Shutdown() {
	wg.Done()
	if err := s.srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
