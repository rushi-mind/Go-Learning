package routes

import (
	adminV1Controller "SMT/controllers/api/v1/admin"
	assignmentV1Controller "SMT/controllers/api/v1/assignments"
	attendanceV1Controller "SMT/controllers/api/v1/attendance"
	authV1Controllers "SMT/controllers/api/v1/auth"
	departmentV1Controller "SMT/controllers/api/v1/department"
	studentV1Controllers "SMT/controllers/api/v1/student"
	teacherV1Controller "SMT/controllers/api/v1/teacher"
	"SMT/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAdminTeacherRoutes(router *gin.Engine) {

	adminRoutes := router.Group("/api/v1")
	teacherRoutes := router.Group("/api/v1")
	adminRoutes.Use(middlewares.AuthMiddlewareAdmin)
	teacherRoutes.Use(middlewares.AuthMiddlewareTeacher)

	router.MaxMultipartMemory = 8 << 20

	// students
	adminRoutes.POST("/student", studentV1Controllers.AddStudent)
	adminRoutes.PUT("/student/:id", studentV1Controllers.UpdateStudent)
	adminRoutes.DELETE("/student/:id", studentV1Controllers.DeleteStudent)
	teacherRoutes.GET("/students", studentV1Controllers.GetStudents)
	teacherRoutes.GET("/student/:id", studentV1Controllers.GetStudent)

	// teachers
	adminRoutes.POST("/teacher", teacherV1Controller.AddTeacher)
	adminRoutes.PUT("/teacher/:empId", teacherV1Controller.UpdateTeacher)
	adminRoutes.GET("/teachers", teacherV1Controller.GetAllTeachers)
	adminRoutes.GET("/teacher/:empId", teacherV1Controller.GetTeacher)
	adminRoutes.DELETE("/teacher/:empId", teacherV1Controller.DeleteTeacher)

	// departments
	adminRoutes.POST("/department", departmentV1Controller.AddDepartment)
	adminRoutes.PUT("/department/:deptId", departmentV1Controller.UpdateDepartment)
	adminRoutes.DELETE("/department/:deptId", departmentV1Controller.DeleteDepartment)
	teacherRoutes.GET("/departments", departmentV1Controller.GetDepartments)
	teacherRoutes.GET("/department/:deptId", departmentV1Controller.GetDepartment)

	// attendance
	teacherRoutes.POST("/attendance", attendanceV1Controller.AddAttendance)

	// assignments
	teacherRoutes.POST("/assignment", assignmentV1Controller.AddAssignment)
	teacherRoutes.DELETE("/assignment/:id", assignmentV1Controller.DeleteAssignment)
	teacherRoutes.PUT("/assignment/:id", assignmentV1Controller.UpdateAssignment)

	// auth
	router.POST("/admin/auth", authV1Controllers.AuthAdmin)
	router.POST("/teacher/auth", authV1Controllers.AuthTeacher)

	// privacy and security
	teacherRoutes.PUT("/change-password-admin", adminV1Controller.ChangePassword)
}
