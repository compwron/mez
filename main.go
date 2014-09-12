package main

import (
	// "github.com/bmizerany/pat"
	// "log"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Instructions())
	http.HandleFunc("/game", Game())
	http.HandleFunc("/game/koan", CreateKoan())
	http.HandleFunc("/game/guess", GuessRule())
	http.HandleFunc("/game/generate", StartGameWithUnknownRule())

	port := ":3000"
	fmt.Println("Starting server on port", port)
	http.ListenAndServe(port, nil)
}
