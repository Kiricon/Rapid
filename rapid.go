package rapid

// App - Create an instance of a rapid server
func App() Server {
	server := Server{}
	server.handler.server = &server
	return server
}
