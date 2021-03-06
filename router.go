package main

import (
	"github.com/gorilla/mux"
	h "github.com/hiscox-mock-server/handlers"
)

func RegisterRoutes(mux *mux.Router) {
	mux.Handle(
		"/api/test",
		&h.DummyHandler{FilePath: "./schemas/test.json"},
	).Methods("GET")

	// Error Routes
	mux.Handle(
		"/api/error/401",
		&h.DummyErrorHandler{Status: 401},
	)

	mux.Handle(
		"/api/error/404",
		&h.DummyErrorHandler{Status: 404},
	)

	mux.Handle(
		"/api/error/422",
		&h.DummyErrorHandler{Status: 422},
	)

	mux.Handle(
		"/api/error/500",
		&h.DummyErrorHandler{Status: 500},
	)
}
