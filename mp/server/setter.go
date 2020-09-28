package server

import "net/http"

type Setter func(server *Server) error

func WithRequest(request *http.Request) Setter {
	return func(server *Server) error {
		server.Request = request
		return nil
	}
}

func WithWriter(writer http.ResponseWriter) Setter {
	return func(server *Server) error {
		server.Writer = writer
		return nil
	}
}
