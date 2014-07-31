package main

import (
	"github.com/bmizerany/pat"
	"log"
	"net/http"
)

func main() {
	mux := pat.New()
	http.Handle("/", mux)

	mux.Get("/", http.HandlerFunc(Instructions))

	mux.Post("/game", http.HandlerFunc(CreateGame))
	mux.Get("/game", http.HandlerFunc(ViewGame))

	mux.Post("/game/koan", http.HandlerFunc(CreateKoan))
	mux.Post("/game/guess", http.HandlerFunc(GuessRule))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}
