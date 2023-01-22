package request

type PostColaParameter struct {
	ColaId int    `form:"cola_id" json:"cola_id" binding:"required,min=1,max=2147483647"`
	Date   int64  `form:"result_date" json:"result_date" binding:"required,min=0,max=9223372036854775807"`
	Note   string `form:"note" json:"note" binding:""`
}
