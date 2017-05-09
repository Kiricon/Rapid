package rapid

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var srv *http.Server

// StaticFolder - Specify application public/static folder
func StaticFolder(path string, dir string) {
	AddStaticPath(path, dir)
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
	srv = &http.Server{Addr: ":" + portString}
	createPaths()
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

func createPaths() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		correctPath := FindCorrectPath(r.URL.Path, r.Method)
		if correctPath != "404" {
			params := getParams(correctPath, r.URL.Path)

			if _, ok := PathHandlers[correctPath][r.Method]; ok {
				PathHandlers[correctPath][r.Method](Connection{R: r, W: w, Params: params})
			} else if _, ok := PathHandlers[correctPath]["ALL"]; ok {
				PathHandlers[correctPath]["ALL"](Connection{R: r, W: w, Params: params})
			} else {
				fileServer := FindStaticServer(r.URL.Path)
				if fileServer.path != "" {
					StaticServerHandler(fileServer)(Connection{R: r, W: w})
				} else {
					fmt.Fprintf(w, "404 Not Found - 1")
				}
			}
		} else {
			fileServer := FindStaticServer(r.URL.Path)
			if fileServer.path != "" {
				StaticServerHandler(fileServer)(Connection{R: r, W: w})
			} else {
				fmt.Fprintf(w, "404 Not Found - 2")
			}
		}
	})
}

func getParams(path string, rPath string) map[string]string {
	paramLocations := getParamLocations(path)
	requestPath := strings.Split(rPath, "/")
	params := map[string]string{}

	if len(paramLocations) > 0 {
		for i := 0; i < len(requestPath); i++ {
			if val, ok := paramLocations[i]; ok {
				params[val] = requestPath[i]
			}
		}
	}

	return params
}

// ShutdownServer - Gracefully shut down the server and unblock
// the server thread.
func ShutdownServer() {
	wg.Done()
	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
