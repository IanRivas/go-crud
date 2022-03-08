package handler

import "net/http"

// RoutePerson ruta para hacer nuestro crud de memory
func RoutePerson(mux *http.ServeMux, storage Storage) {
	h := newPerson(storage)

	mux.HandleFunc("/v1/persons/create", h.create)
	mux.HandleFunc("/v1/persons/update", h.update)
	mux.HandleFunc("/v1/persons/delete", h.delete)
	mux.HandleFunc("/v1/persons/get-all", h.getAll)
}
