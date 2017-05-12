package rapid

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/Kiricon/Rapid/connection"
	"github.com/Kiricon/Rapid/mux"
)

var wg sync.WaitGroup
var srv *http.Server

type routeHandler func(connection.Connection)

type Server struct {
	Paths        map[string]mux.Path
	PathHandlers map[string]map[string]func(connection.Connection)
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
func StaticFolder(path string, dir string) {
	mux.AddStaticPath(path, dir)
}

// Listen - Start webserver on specified port
// Returns the instance of the http server currently runing.
// You can use this instance to shutdown the server if need be.
func Listen(port int) {
	ListenAndWait(port, true)
}

// ListenAndWait - Gives user option of waiting for server or not
func ListenAndWait(port int, wait bool) {
	portString := strconv.Itoa(port)
	srv = &http.Server{Addr: ":" + portString, Handler: mux.RapidHandler{}}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			// cannot panic, because this probably is an intentional close
			log.Printf("Httpserver: ListenAndServe() error: %s", err)
		}
	}()

	wg.Add(1)
	if wait {
		wg.Wait()
	}
}

// ShutdownServer - Gracefully shut down the server and unblock
// the server thread.
func ShutdownServer() {
	wg.Done()
	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
