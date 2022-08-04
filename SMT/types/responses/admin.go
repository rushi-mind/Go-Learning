package responseTypes

type SuccessResponseData struct {
	Base
	Data any `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

const (
	INVALID_JWT_TOKEN       = "invalid JWT token"
	AUTHENTICATION_FAILED   = "Authenticaion failed"
	FAILED_TO_HASH          = "failed to hash password"
	SUCCESS_LOGIN           = "Logged in successfully"
	LOGIN_FAILED            = "Login failed"
	PASSWORD_HASHING_FAILED = "Failed to hash this password"
	INVALID_PASSWORD        = "Password entered is Invalid"
	INVALID_ADMINID         = "Admin ID is invalid"
	PASSWORD_CHANGED        = "Password updated successfully"
	DEPARTMENT_ADDED        = "Department added successfully"
	DEPARTMENT_UPDATED      = "Department updated successfully"
	DEPARTMENT_NOT_FOUND    = "Department not found"
	DEPARTMENT_DELETED      = "Department deleted successfully"
	STUDENT_CREATED         = "Student created successfully"
	STUDENT_CREATE_ERROR    = "Failed to create a student"
	INVALID_STUDENT_ID      = "Invalid student id"
	STUDENT_UPDATED         = "Student updated successfully"
	STUDENT_DELETED         = "Student deleted successfully"
	STUDENTS_FETCHED        = "Students fetched successfully"
	STUDENT_FETCHED         = "Student fetched successfully"
	INVALID_DEPARTMENT_ID   = "Department id is invalid"
	TEACHERS_FETCHED        = "Teachers fetched successfully"
	TEACHER_FETCHED         = "Teacher fetched successfully"
	INVALID_TEACHER_ID      = "Invalid teacher id"
	TEACHER_UPDATE          = "Teacher updated successfully"
	TEACHER_DELETED         = "Teacher deleted successfully"
	TEACHER_CREATE_ERROR    = "Failed to create new teacher"
	TEACHER_CREATED         = "Teacher created successfully"
	INVALID_INPUT_JSON      = "Invalid input JSON"
)
