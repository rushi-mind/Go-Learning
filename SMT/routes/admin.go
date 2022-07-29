package routes

import (
	adminV1Controller "SMT/controllers/api/v1/admin"
	"SMT/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(router *gin.Engine) {

	adminRoutesGroup := router.Group("/api/v1")
	adminRoutesGroup.Use(middlewares.AuthMiddleware)

	// students
	adminRoutesGroup.POST("/student", adminV1Controller.AddStudent)
	adminRoutesGroup.PUT("/student/:rollNo", adminV1Controller.UpdateStudent)
	adminRoutesGroup.GET("/students", adminV1Controller.GetAllStudents)
	adminRoutesGroup.GET("/student/:id", adminV1Controller.GetStudent)
	adminRoutesGroup.DELETE("/student/:rollNo", adminV1Controller.DeleteStudent)

	// faculties
	adminRoutesGroup.POST("/teacher", adminV1Controller.AddTeacher)
	adminRoutesGroup.PUT("/teacher/:empId", adminV1Controller.UpdateTeacher)
	adminRoutesGroup.GET("/teachers", adminV1Controller.GetAllTeachers)
	adminRoutesGroup.GET("/teacher/:empId", adminV1Controller.GetTeacher)
	adminRoutesGroup.DELETE("/teacher/:empId", adminV1Controller.DeleteTeacher)

	// departments
	adminRoutesGroup.POST("/department", adminV1Controller.AddDepartment)
	adminRoutesGroup.PUT("/department/:id", adminV1Controller.UpdateDepartment)
	adminRoutesGroup.DELETE("/department/:id", adminV1Controller.DeleteDepartment)

	// auth
	router.POST("/auth", adminV1Controller.Auth)
	adminRoutesGroup.PUT("/change-password", adminV1Controller.ChangePassword)

}
