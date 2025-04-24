package handlers

import (
	"encoding/json" // Serve per codificare la risposta in JSON
	"net/http"      // Gestione HTTP (router, response writer, ecc.)

	"github.com/nicconuti/cognitive-api/utils" // Importa la funzione che genera la griglia
)

// MemoryTestResponse definisce la struttura della risposta JSON
// che verrà inviata al client quando richiama l'API di memoria visiva
type MemoryTestResponse struct {
	Grid    [][]string `json:"grid"`    // griglia 2D di simboli emoji
	Timeout int        `json:"timeout"` // tempo massimo per visualizzare la griglia, in secondi
}

// MemoryTestHandler è l'handler HTTP associato all'endpoint /test/visual-memory
// Quando un client lo chiama, genera una griglia casuale e restituisce una risposta JSON
func MemoryTestHandler(w http.ResponseWriter, r *http.Request) {
	// Genera una griglia 3x3 di simboli casuali
	grid := utils.GenerateGrid(3)

	// Crea la struttura di risposta con griglia e timeout
	response := MemoryTestResponse{
		Grid:    grid,
		Timeout: 5, // 5 secondi per memorizzare la griglia
	}

	// Imposta l'header della risposta in modo che il client capisca che è JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica la struttura in JSON e la scrive nella risposta
	json.NewEncoder(w).Encode(response)
}
