package middleware

import (
	"fmt"
	"learningo/app/helpers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Cors(pass httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		header := w.Header()
		header.Set("Content-Type", "application/json")
		header.Set("Access-Control-Allow-Origin", "*")
		pass(w, r, p)
	}
}

func JwtAuth(pass httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		a := r.Header.Get("Authorization")
		fmt.Println("Authorization: ", a)
		if a == "Auth" {
			pass(w, r, p)
		} else {
			helpers.JSONResponse(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}
	}
}
