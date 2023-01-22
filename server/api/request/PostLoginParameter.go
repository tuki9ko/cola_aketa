package request

type PostLoginRequestParameter struct {
	UserId   string `form:"user_id" json:"user_id" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
