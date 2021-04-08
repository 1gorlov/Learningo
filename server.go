package main

import (
	"learningo/database"
	"learningo/models"
	"learningo/router"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	router.InitRouter(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
