package routes

import "gorm-demo/controllers"

func InitProductRoutes() {
	Router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")
	Router.HandleFunc("/api/products", controllers.AddProduct).Methods("POST")
	Router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	Router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
	Router.HandleFunc("/api/products/{id}", controllers.GetProduct).Methods("GET")
}
