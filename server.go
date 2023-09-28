package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func isChannelLive(channelID string, apiKey string) bool {
	// Sua lógica real para verificar o status do canal aqui

	// Chame a função isChannelLive que você já possui
	// Neste exemplo, a lógica é simplificada para retornar true
	// se o canal estiver online, e false se estiver offline
	return true // Substitua pela sua lógica real
}

func checkYouTubeChannelStatus(w http.ResponseWriter, r *http.Request) {
	// Defina sua lógica real para obter o apiKey e channelID
	apiKey := "AIzaSyCrr7c6iPUHc-3HrFcowu0j46GwK5Ww3c8"
	channelID := "@Alternativerock96"

	// Chame a função isChannelLive que você já possui
	isLive := isChannelLive(channelID, apiKey)

	// Retorne a resposta em formato JSON
	response := struct {
		IsLive bool `json:"isLive"`
	}{
		IsLive: isLive,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/checkChannel", checkYouTubeChannelStatus)
	port := ":8080"
	fmt.Printf("Servidor Go está ouvindo na porta %s\n", port)
	http.ListenAndServe(port, r)
}
