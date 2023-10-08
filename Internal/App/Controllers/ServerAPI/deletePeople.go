package ServerAPI

import (
	"fmt"
	"net/http"
)

func (s *Server) deletePeople(writer http.ResponseWriter, request *http.Request) {
	// checking for valid key
	key, err := checkKey(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	err = s.cache.Delete(key)
	// checking for successful deleting
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	_, err = fmt.Fprint(writer, "user with key ", key, " was deleted")
	// checking for errors while printing message
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}
