package controller

import (
	"github.com/gin-gonic/gin"
)

type ColaController struct{}

func (cc ColaController) GetCola(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cola results!",
	})
}
