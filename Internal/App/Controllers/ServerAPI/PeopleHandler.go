package ServerAPI

import "net/http"

func (s *Server) peopleHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		s.getPeople(writer, request)
	case http.MethodPost:
		s.postPeople(writer, request)
	case http.MethodDelete:
		s.deletePeople(writer, request)
	default:
		http.Error(writer, "Undefined method", http.StatusMethodNotAllowed)
	}
}
