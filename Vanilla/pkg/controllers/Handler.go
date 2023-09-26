package controllers

import (
	"github.com/jroden2/Sleepy/Vanilla/pkg/services"
	"net/http"
)

type handler struct {
	es services.ExampleService
}

func NewHandler(es *services.ExampleService) Handler {
	if es == nil {
		tmp := services.NewExampleService()
		es = &tmp
	}

	return &handler{
		es: *es,
	}
}

type Handler interface {
	ExampleHandler(w http.ResponseWriter, r *http.Request)
}

func (h *handler) ExampleHandler(w http.ResponseWriter, r *http.Request) {
	h.es.ExposedFunction(w)
}
