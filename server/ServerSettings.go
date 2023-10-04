package server

import (
	"SimpleServer/Structures"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var cache *Structures.Cache = Structures.NewCache(time.Minute*2, time.Minute, false)

func SetServer() {

	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/people/update", peopleUpdateHandler)
	fmt.Println("http server is working ")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic("unable to connect to server")
	}

}
func peopleUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		updatePeople(writer, request)
	}
}

func peopleHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		getPeople(writer, request)
	case http.MethodPost:
		postPeople(writer, request)
	case http.MethodDelete:
		deletePeople(writer, request)
	default:
		http.Error(writer, "Undefined method", http.StatusMethodNotAllowed)
	}
}
func updatePeople(writer http.ResponseWriter, request *http.Request) {
	var person Structures.Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = cache.Update(person.Name, person.Age, person.Sex)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusMethodNotAllowed)
	}
}
func deletePeople(writer http.ResponseWriter, request *http.Request) {
	var key string
	err := json.NewDecoder(request.Body).Decode(&key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = cache.Delete(key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	_, _ = fmt.Fprint(writer, "user with key ", key, " was deleted")

}
func getPeople(writer http.ResponseWriter, request *http.Request) {
	var key string
	err := json.NewDecoder(request.Body).Decode(&key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	value, ok := cache.Get(key)
	if !ok {
		http.Error(writer, "couldn't find the value", http.StatusConflict)
		return
	}
	err = json.NewEncoder(writer).Encode(value)

}

func postPeople(writer http.ResponseWriter, request *http.Request) {
	var person Structures.Person
	err := json.NewDecoder(request.Body).Decode(&person)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = cache.Set(person.Name, person, time.Minute)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusMethodNotAllowed)
	}

}
