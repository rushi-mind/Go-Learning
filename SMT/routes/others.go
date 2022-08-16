package routes

import (
	assignmentV1Controller "SMT/controllers/api/v1/assignments"
	"SMT/middlewares"

	"github.com/gin-gonic/gin"
)

func InitOtherRoutes(router *gin.Engine) {
	otherRoutes := router.Group("/api/v1")
	otherRoutes.Use(middlewares.AuthAnyMiddleware)
	otherRoutes.GET("/assignments", assignmentV1Controller.GetAssignments)
}
