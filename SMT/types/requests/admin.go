package requestTypes

type AdminAuth struct {
	AdminId  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminChangePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
