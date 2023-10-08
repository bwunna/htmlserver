package ServerAPI

import (
	"net/http"
)

func New(address string, handler http.Handler) *Server {
	// new server
	newServer := &(Server{address: address, handler: handler})
	return newServer
}
