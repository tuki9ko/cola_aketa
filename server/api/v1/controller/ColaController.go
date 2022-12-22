package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/v1/service"
)

type ColaController struct{}

var cs service.ColaService

func (cc ColaController) GetCola(c *gin.Context) {
	contextUserId, _ := c.Get("userId")
	userId := contextUserId.(int)

	cola_id, _ := strconv.Atoi(c.Query("cola_id"))

	cola, err := cs.GetCola(userId, cola_id)

	msg := "success"

	if err != nil {
		msg = err.Error()
	}

	c.JSON(200, gin.H{
		"message": msg,
		"cola":    cola,
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

func (cc ColaController) GetColas(c *gin.Context) {
	contextUserId, _ := c.Get("userId")
	userId := contextUserId.(int)

	colas, err := cs.GetColas(userId)

	msg := "success"

	if err != nil {
		msg = err.Error()
	}

	c.JSON(200, gin.H{
		"message": msg,
		"colas":   colas,
	})

}
