package routes

import "github.com/gorilla/mux"

var Router = mux.NewRouter()

func Init() {
	InitAttendaceRoutes()
	InitUserRoutes()
}
