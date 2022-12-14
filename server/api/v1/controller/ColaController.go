package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/v1/service"
)

type ColaController struct{}

var cs service.ColaService

func (cc ColaController) GetCola(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cola results!",
	})
}

func (cc ColaController) PostCola(c *gin.Context) {
	contextUserId, _ := c.Get("userId")
	userId := contextUserId.(int)

	cola_id, _ := strconv.Atoi(c.PostForm("cola_id"))
	result_date, _ := strconv.ParseInt(c.PostForm("result_date"), 10, 64)
	note := c.PostForm("note")

	cs.AddResult(userId, cola_id, result_date, note)

	c.JSON(200, gin.H{
		"message": "cola post done!",
	})
}
