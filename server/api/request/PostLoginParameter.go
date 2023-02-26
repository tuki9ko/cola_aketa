package request

type PostLoginRequestParameter struct {
	UserId   string `form:"userId" json:"userId" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}
