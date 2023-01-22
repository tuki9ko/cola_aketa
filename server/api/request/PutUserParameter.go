package request

type PutUserParameter struct {
	Name      string `form:"name" json:"name" binding:"required,max=255"`
	Email     string `form:"email" json:"email" binding:"required,email,max=255"`
	Password  string `form:"password" json:"password" binding:"required,min=8,max=255"`
	Biography string `form:"biography" json:"biography" binding:"max=255"`
	UserIcon  string `form:"user_icon" json:"user_icon" binding:"max=255"`
}
