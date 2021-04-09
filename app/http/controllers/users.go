package controllers

import (
	"encoding/json"
	"fmt"
	"learningo/app/helpers"
	"learningo/app/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UsersController struct{}

func AllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := models.All()
	if err != nil {
		helpers.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := models.Find(ps.ByName("id"))
	if err != nil {
		helpers.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func NewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	nu, err := models.Create(user)
	if err != nil {
		helpers.JSONError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(nu)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
