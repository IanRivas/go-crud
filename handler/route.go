package handler

import (
	"net/http"

	"github.com/IanRivas/go-crud/middleware.go"
)

// RoutePerson ruta para hacer nuestro crud de memory
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", middleware.Log(middleware.Authentication(h.create)))
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", middleware.Log(h.delete))
	mux.HandleFunc("/v1/persons/get-all", middleware.Log(h.getAll))
}
