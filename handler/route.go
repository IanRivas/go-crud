package handler

import "net/http"

// RoutePerson ruta para hacer nuestro crud de memory
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", h.create)
}