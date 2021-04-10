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
	u, err := models.All()
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.JSONResponse(w, u, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, err := models.Find(ps.ByName("id"))
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.JSONResponse(w, u, http.StatusOK)
}

func NewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	nu, err := models.Create(u)
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, nu, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Delete User Endpoint Hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "Update User Endpoint Hit")
}
