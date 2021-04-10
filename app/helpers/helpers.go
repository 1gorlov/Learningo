package helpers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func JSONResponse(w http.ResponseWriter, v interface{}, code int) {
	header := w.Header()
	header.Set("Content-Type", "application/json; charset=utf-8")
	header.Set("X-Content-Type-Options", "nosniff")
	header.Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)

	var r Response

	if code != 200 {
		r.Status = "error"
	} else {
		r.Status = "success"
	}
	r.Data = v
	json.NewEncoder(w).Encode(r)
}
