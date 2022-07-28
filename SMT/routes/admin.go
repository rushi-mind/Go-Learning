package routes

import (
	"SMT/controllers/api/v1/admin"
	"SMT/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(router *gin.Engine) {

	adminRoutesGroup := router.Group("/api/v1")
	adminRoutesGroup.Use(middlewares.AuthMiddleware)

	// students
	adminRoutesGroup.POST("/student", admin.AddStudent)
	adminRoutesGroup.PUT("/student/:rollNo", admin.UpdateStudent)
	adminRoutesGroup.GET("/students", admin.GetAllStudents)
	adminRoutesGroup.GET("/student/:id", admin.GetStudent)
	adminRoutesGroup.DELETE("/student/:rollNo", admin.DeleteStudent)

	// faculties
	adminRoutesGroup.POST("/teacher", admin.AddTeacher)
	adminRoutesGroup.PUT("/teacher/:empId", admin.UpdateTeacher)
	adminRoutesGroup.GET("/teachers", admin.GetAllTeachers)
	adminRoutesGroup.GET("/teacher/:empId", admin.GetTeacher)
	adminRoutesGroup.DELETE("/teacher/:empId", admin.DeleteTeacher)

	// departments
	adminRoutesGroup.POST("/department", admin.AddDepartment)
	adminRoutesGroup.PUT("/department/:id", admin.UpdateDepartment)
	adminRoutesGroup.DELETE("/department/:id", admin.DeleteDepartment)

	// auth
	router.POST("/auth", admin.Auth)
	adminRoutesGroup.PUT("/change-password", admin.ChangePassword)

}
