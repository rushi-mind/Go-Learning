package routes

import "github.com/gin-gonic/gin"

var Router = gin.New()

func InitRoutes() {
	InitAdminTeacherRoutes(Router)
	InitStudentRoutes(Router)
	InitOtherRoutes(Router)
}
