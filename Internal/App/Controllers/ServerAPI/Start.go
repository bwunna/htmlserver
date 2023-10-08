package ServerAPI

import (
	"SimpleServer/Internal/App/Repositories/Cache_"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) Start(defaultExpiration time.Duration, cleanUpInterval time.Duration, endlessLifeTimeAvailability bool) {
	// starting server
	s.cache = Cache_.NewCache(defaultExpiration, cleanUpInterval, endlessLifeTimeAvailability)
	http.HandleFunc("/people", s.peopleHandler)
	fmt.Println("http server is working ")
	err := http.ListenAndServe(s.address, s.handler)
	if err != nil {
		panic("unable to connect to ServerAPI")
	}
}
