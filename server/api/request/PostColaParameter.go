package request

type PostColaParameter struct {
	ColaId int    `form:"colaId" json:"colaId" binding:"required,min=1,max=2147483647"`
	Date   int64  `form:"resultDate" json:"resultDate" binding:"required,min=0,max=9223372036854775807"`
	Note   string `form:"note" json:"note" binding:""`
}
