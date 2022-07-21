package routes

import (
	"rest-demo/controllers"
)

func InitAttendaceRoutes() {
	Router.HandleFunc("/api/attendance", controllers.GetAttendance).Methods("GET")
}
