package handlers

import (
	"encoding/json"
	"net/http"
	"os"
)

func RegisterHealthCheck() {
	http.HandleFunc("/health", handleHealthCheck)
	_ = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func handleHealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"status": "OK",
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(data)
}
