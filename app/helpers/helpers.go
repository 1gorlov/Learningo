package helpers

import (
	"encoding/json"
	"net/http"
)

func JSONError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	err := make(map[string]string)
	err["error"] = message
	json.NewEncoder(w).Encode(err)
}
