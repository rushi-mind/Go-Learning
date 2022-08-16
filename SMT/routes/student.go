package routes

import (
	authV1Controllers "SMT/controllers/api/v1/auth"
	studentV1Controllers "SMT/controllers/api/v1/student"
	"SMT/middlewares"

	"github.com/gin-gonic/gin"
)

func InitStudentRoutes(router *gin.Engine) {
	studentRoutes := router.Group("/api/v1")
	studentRoutes.Use(middlewares.AuthMiddlewareStudent)

	// Profile
	studentRoutes.GET("/profile", studentV1Controllers.GetProfile)
	studentRoutes.PUT("/change-password-student", studentV1Controllers.ChangePassword)

	// Auth
	router.POST("/student/auth", authV1Controllers.AuthStudent)
}
