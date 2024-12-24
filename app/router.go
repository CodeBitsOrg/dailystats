package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Router(handler *Handler) *mux.Router {
	router := mux.NewRouter()

	router.Methods("GET").Path("/stats").Handler(http.HandlerFunc(handler.Get))

	return router
}
