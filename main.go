package main

import (
	"github.com/bmizerany/pat"
	"log"
	"mez/api"
	"net/http"
)

func main() {
	mux := pat.New()
	http.Handle("/", mux)

	mux.Get("/", http.HandlerFunc(api.Instructions))

	mux.Post("/game", http.HandlerFunc(api.CreateGame))
	mux.Get("/game", http.HandlerFunc(api.ViewGame))

	mux.Post("/game/koan", http.HandlerFunc(api.CreateKoan))
	mux.Post("/game/guess", http.HandlerFunc(api.GuessRule))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
