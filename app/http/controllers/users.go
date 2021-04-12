package controllers

import (
	"encoding/json"
	"learningo/app/helpers"
	"learningo/app/models"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var um models.User
	u, err := um.All()
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	helpers.JSONResponse(w, u, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var um models.User

	u, err := um.Find(ps.ByName("id"))
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

	nu, err := u.Create()
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, nu, http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var u models.User

	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	nu, err := u.Update(ps.ByName("id"), u)
	if err != nil {
		helpers.JSONResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	helpers.JSONResponse(w, nu, http.StatusOK)
}

func DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var u models.User
	u.Find(ps.ByName("id"))
	u.Delete()
	helpers.JSONResponse(w, u, http.StatusOK)
}
