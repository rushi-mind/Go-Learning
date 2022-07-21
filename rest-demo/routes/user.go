package routes

import "rest-demo/controllers"

func InitUserRoutes() {
	Router.HandleFunc("/api/users", controllers.GetUsers).Methods("GET")
	Router.HandleFunc("/api/users/{id}", controllers.GetUser).Methods("GEt")
	Router.HandleFunc("/api/users", controllers.AddUser).Methods("POST")
	Router.HandleFunc("/api/users/{id}", controllers.UpdateUser).Methods("PUT")
	Router.HandleFunc("/api/users/{id}", controllers.DeleteUser).Methods("DELETE")
}
