package controller

import (
	"github.com/gin-gonic/gin"
)

func GetCola(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "cola results!",
	})
}
