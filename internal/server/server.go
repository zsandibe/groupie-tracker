package server

import (
	"net/http"
	"zsandibe/internal/delivery"
)

func Server(handler delivery.Handler) *http.Server {
	mux := http.NewServeMux()
	handler.Routes(mux)
	return &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}
}
