package ServerAPI

import (
	"SimpleServer/Internal/App/Repositories/Cache_"
	"net/http"
)

type Server struct {
	cache   *Cache_.Cache
	handler http.Handler
	address string
}
