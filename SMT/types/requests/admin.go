package requestTypes

type AdminAuth struct {
	AdminId  string `json:"admin_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
