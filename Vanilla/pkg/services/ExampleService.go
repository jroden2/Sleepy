package services

import (
	"net/http"
)

type exampleService struct {
	// This is where we would store and dependancies we want to call, i.e repositories.
}

func NewExampleService(/*we would pass dependencies here as part of the constructor*/) ExampleService {
	return &exampleService{
		/*link the constructor variable to the struct*/
	}
}

type ExampleService interface {
	ExposedFunction(w http.ResponseWriter)
}

func (s *exampleService) ExposedFunction(w http.ResponseWriter) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World - Vanilla\n"))
}