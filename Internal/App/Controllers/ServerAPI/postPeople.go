package ServerAPI

import (
	"SimpleServer/Internal/App/Models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (s *Server) postPeople(writer http.ResponseWriter, request *http.Request) {
	var person Models.Person
	// checking for successful decoding person from json
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	// adding the user
	err = s.cache.Set(person.Name, &person, time.Minute*2)
	if err != nil {
		// checking for errors while adding the user
		if _, err = fmt.Fprintln(writer, err.Error()); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// if user exists, update info about him
		if err = s.cache.Update(person.Name, person.Age, person.Sex); err != nil {
			http.Error(writer, err.Error(), http.StatusConflict)
			return
		} else {
			// checking for errors while printing the message
			if _, err := fmt.Fprintln(writer, "Information was updated"); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}

}
