package ServerAPI

import (
	"encoding/json"
	"net/http"
)

func (s *Server) getPeople(writer http.ResponseWriter, request *http.Request) {
	// checking for valid key
	key, err := checkKey(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	// checking if user exists
	value, err := s.cache.Get(key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(value)

}
