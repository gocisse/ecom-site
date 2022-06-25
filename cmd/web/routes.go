package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()


	fs := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))


	mux.Get("/terminal", app.VirtualTerminal)
	mux.Post("/payment-succeeded", app.Succeeded)
	mux.Get("/buy-once", app.BuyOnce)
	return mux
}
