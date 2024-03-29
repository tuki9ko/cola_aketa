package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/request"
	"github.com/tuki9ko/cola_aketa/api/service"
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

	var param request.PostColaParameter

	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(), //TODO: エラーメッセージを直す
		})
		return
	}

	cs.AddResult(userId, param.ColaId, param.Date, param.Note)

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
