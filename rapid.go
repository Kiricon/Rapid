package rapid

func App() Server {
	server := Server{}
	server.handler.server = &server
	return server
}
