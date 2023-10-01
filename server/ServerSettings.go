package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var people []Person = make([]Person, 0, 3)

func SetServer() {

	http.HandleFunc("/people", peopleHandler)
	fmt.Println("http server is working ")
	err := http.ListenAndServe("localhost:8080", nil)
	if err == nil {
		log.Fatal(err)
	}

}

func peopleHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getPeople(writer, request)
	case http.MethodPost:
		postPeople(writer, request)
	default:
		http.Error(writer, "Undefined method", http.StatusMethodNotAllowed)
	}
}
func getPeople(writer http.ResponseWriter, request *http.Request) {
	err := json.NewEncoder(writer).Encode(people)
	if err != nil {
		return
	}

}

func postPeople(writer http.ResponseWriter, request *http.Request) {
	var person Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	for _, value := range people {
		if value.Name == person.Name {
			http.Error(writer, "error: this person is already in the list", http.StatusConflict)
			return
		}
	}
	people = append(people, person)

}
