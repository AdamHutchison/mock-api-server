package main

import (
	"github.com/gorilla/mux"
	h "github.com/hiscox-mock-server/handlers"
)

func RegisterRoutes(mux *mux.Router) {
	mux.Handle(
		"/api/test",
		&h.DummyHandler{FilePath: "./schemas/test.json"},
	)

	mux.Handle(
		"/api/error/404",
		&h.DummyErrorHandler{Status: 404},
	)
}
