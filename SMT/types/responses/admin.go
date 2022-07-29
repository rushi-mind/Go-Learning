package responseTypes

type AdminAuthSuccess struct {
	Base
	Data TokenResponse `json:"data"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type SuccessResponse struct {
	Base
}

type SuccessResponseData struct {
	Base
	Data any
}

type AuthError struct {
	Base
}

const (
	INVALID_JWT_TOKEN                        = "invalid JWT token"
	AUTHENTICATION_FAILED                    = "Authenticaion failed"
	FAILED_TO_HASH                           = "failed to hash password"
	SUCCESS_LOGIN                            = "Logged in successfully"
	ADMIN_AUTH_INVALID_JSON_INPUT            = "Please enter valid Admin ID and Password"
	LOGIN_FAILED                             = "Login failed"
	PASSWORD_HASHING_FAILED                  = "Failed to hash this password"
	INVALID_PASSWORD                         = "Password entered is Invalid"
	INVALID_ADMINID                          = "Admin ID is invalid"
	INVALID_ADMIN_CHANGE_PASSWORD_JSON_INPUT = "Please enter old password and new password correctly"
	PASSWORD_CHANGED                         = "Password updated successfully"
	INVALID_ADD_DEPARTMENT_JSON_INPUT        = "Please enter valid Deparment-Code and Name. Code should be of 6 digit only"
	DEPARTMENT_ADDED                         = "Department added successfully"
)
