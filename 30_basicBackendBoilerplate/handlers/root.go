package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Error   bool   `json:"error"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	resp := Response{
		Error:   false,
		Data:    nil,
		Message: "api server is live",
	}
	json.NewEncoder(w).Encode(resp)
}
