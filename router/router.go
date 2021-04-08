package router

import (
	"learningo/app/http/controllers"

	"github.com/julienschmidt/httprouter"
)

func registerRoutes(r *httprouter.Router) {
	r.GET("/users", controllers.AllUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.POST("/users", controllers.NewUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
	r.PUT("/users/:id", controllers.UpdateUser)
}

func InitRouter(r *httprouter.Router) {
	registerRoutes(r)
}
