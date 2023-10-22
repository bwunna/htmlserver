package Controller

import (
	"SimpleServer/Internal/App/Providers/Provider"
	"SimpleServer/Internal/App/Repositories/Cache"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	cache   *Cache.Cache
	handler http.Handler
	address string
}

func NewServer(address string, handler http.Handler) *Server {
	// new server
	return &Server{address: address, handler: handler}
}

func (s *Server) AskForPromotion(writer http.ResponseWriter, request *http.Request) {
	key, err := checkKey(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	err = s.cache.AskForPromotion(key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	_, err = fmt.Fprintln(writer, "Congratulations! You've been promoted")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
	}
}

func (s *Server) getPeople(writer http.ResponseWriter, request *http.Request) {
	// checking for valid key
	key, err := checkKey(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
		return
	}
	// checking if user exists
	userData, err := s.cache.Get(key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(writer).Encode(userData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	salaryData, err := s.cache.GetSalaryData(key)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(salaryData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}

}

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

func checkKey(request *http.Request) (string, error) {
	// checking if key from json is valid
	var keyInJson map[string]interface{}

	err := json.NewDecoder(request.Body).Decode(&keyInJson)
	if err != nil {
		return "", err
	}
	if key, ok := keyInJson["key"].(string); !ok {
		return "", errors.New("key is not valid")
	} else {
		return key, nil
	}

}

func (s *Server) Start(defaultExpiration time.Duration, cleanUpInterval time.Duration, endlessLifeTimeAvailability bool, db *Provider.DataBase, promotionInterval time.Duration) error {
	// starting server
	s.cache = Cache.NewCache(defaultExpiration, cleanUpInterval, endlessLifeTimeAvailability, db, promotionInterval)
	http.HandleFunc("/people", s.peopleHandler)
	fmt.Println("http server is working ")
	err := http.ListenAndServe(s.address, s.handler)
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) postPeople(writer http.ResponseWriter, request *http.Request) {
	// checking for successful decoding person from json
	//fmt.Println(request.Body)
	decoder := json.NewDecoder(request.Body)
	answerFromCache, err := s.cache.ParseJson(decoder)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusConflict)
	}
	// adding the user
	err = s.cache.Set(answerFromCache, time.Minute*2)
	//fmt.Println(request.Body)
	if err != nil {
		// checking for errors while adding the user
		if _, err = fmt.Fprintln(writer, err.Error()); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		// if user exists, update info about him
		if err = s.cache.Update(answerFromCache); err != nil {
			http.Error(writer, err.Error(), http.StatusConflict)
			return
		} else {
			// checking for errors while printing the message
			if _, err = fmt.Fprintln(writer, "Information was updated"); err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}
		}

	}

}

func (s *Server) peopleHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		s.getPeople(writer, request)
	case http.MethodPost:
		s.postPeople(writer, request)
	case http.MethodDelete:
		s.deletePeople(writer, request)
	case http.MethodPatch:
		s.AskForPromotion(writer, request)
	default:
		http.Error(writer, "Undefined method", http.StatusMethodNotAllowed)
	}
}
