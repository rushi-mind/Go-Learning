package routes

import (
	adminV1Controller "SMT/controllers/api/v1/admin"
	"SMT/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAdminRoutes(router *gin.Engine) {

	adminRoutes := router.Group("/api/v1")
	adminRoutes.Use(middlewares.AuthMiddleware)

	// students
	adminRoutes.POST("/student", adminV1Controller.AddStudent)
	adminRoutes.PUT("/student/:id", adminV1Controller.UpdateStudent)
	adminRoutes.GET("/students", adminV1Controller.GetAllStudents)
	adminRoutes.GET("/student/:id", adminV1Controller.GetStudent)
	adminRoutes.DELETE("/student/:id", adminV1Controller.DeleteStudent)

	// faculties
	adminRoutes.POST("/teacher", adminV1Controller.AddTeacher)
	adminRoutes.PUT("/teacher/:empId", adminV1Controller.UpdateTeacher)
	adminRoutes.GET("/teachers", adminV1Controller.GetAllTeachers)
	adminRoutes.GET("/teacher/:empId", adminV1Controller.GetTeacher)
	adminRoutes.DELETE("/teacher/:empId", adminV1Controller.DeleteTeacher)

	// departments
	adminRoutes.POST("/department", adminV1Controller.AddDepartment)
	adminRoutes.PUT("/department/:deptId", adminV1Controller.UpdateDepartment)
	adminRoutes.DELETE("/department/:deptId", adminV1Controller.DeleteDepartment)

	// attendance
	adminRoutes.POST("/attendance", adminV1Controller.AddAttendance)

	// auth
	router.POST("/admin/auth", adminV1Controller.Auth)
	adminRoutes.PUT("/change-password", adminV1Controller.ChangePassword)

}
