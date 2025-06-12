package main

import (
	"log"
	"net/http"

	"github.com/nicconuti/cognitive-api/handlers"
)

// corsMiddleware è un middleware che permette al frontend di accedere al backend da un'origine diversa.
// Aggiunge gli header necessari per CORS.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permette tutte le origini (solo per sviluppo)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		// Permette i metodi HTTP usati da React (GET, POST, ecc.)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		// Permette alcuni header usati nei fetch moderni
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Se è una richiesta preflight (OPTIONS), termina subito la richiesta
		if r.Method == "OPTIONS" {
			return
		}

		// Passa la richiesta al prossimo handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Crea un multiplexer (router base)
	mux := http.NewServeMux()

	// Collega l'endpoint al relativo handler
	mux.HandleFunc("/api/memory", handlers.MemoryTestHandler)
	mux.HandleFunc("/api/test", handlers.TestHandler)
	mux.HandleFunc("/api/test/submit", handlers.SubmitHandler)
	mux.HandleFunc("/api/test/results", handlers.ResultsHandler)
	mux.HandleFunc("/api/stroop", handlers.StroopHandler)
	mux.HandleFunc("/api/math", handlers.ArithmeticHandler)
	mux.HandleFunc("/api/pool", handlers.PoolHandler)

	// Wrappa il mux con il middleware CORS
	handlerWithCORS := corsMiddleware(mux)

	// Avvia il server
	log.Println("Server in ascolto su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handlerWithCORS))
}
