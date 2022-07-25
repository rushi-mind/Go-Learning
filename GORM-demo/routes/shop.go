package routes

import "gorm-demo/controllers"

func InitShopRoutes() {
	Router.HandleFunc("/api/shops", controllers.GetShops).Methods("GET")
	Router.HandleFunc("/api/shops", controllers.AddProduct).Methods("POST")
	Router.HandleFunc("/api/shops/{id}", controllers.UpdateProduct).Methods("PUT")
	Router.HandleFunc("/api/shops/{id}", controllers.DeleteProduct).Methods("DELETE")
}
