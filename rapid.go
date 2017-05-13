package rapid

// App - Create an instance of a rapid server
func App() Server {
	return Server{notFoundMessage: "404 Not Found!"}
}
