package rapid

import (
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
	http.Handle("/"+path+"/", http.StripPrefix("/"+path+"/", http.FileServer(http.Dir(dir))))
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
		correctPath := FindCorrectPath(r.URL.Path)
		if correctPath != "404" {
			paramLocations := getParamLocations(r.URL.Path)
			requestPath := strings.Split(r.URL.Path, "/")
			params := map[string]string{}

			for i := 0; i < len(requestPath); i++ {
				if val, ok := paramLocations[i]; ok {
					params[val] = requestPath[i]
				}
			}
			PathHandlers[correctPath](Connection{R: r, W: w, Params: params})
		}
	})
}

// ShutdownServer - Gracefully shut down the server and unblock
// the server thread.
func ShutdownServer() {
	wg.Done()
	if err := srv.Shutdown(nil); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}
}
