package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, rawData any, success bool, message string, statusCode int) {

	response := map[string]interface{}{
		"data":    rawData,
		"success": success,
		"message": message,
	}

	data, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, _ = w.Write(data)
}