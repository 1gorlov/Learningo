package main

import (
	"learningo/app/models"
	"learningo/database"
	"learningo/router"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	db := database.Connect()
	db.AutoMigrate(&models.User{})
	router.InitRouter(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
