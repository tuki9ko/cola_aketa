package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/v1/service"
)

func GetLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET login",
	})
}

func PostLogin(c *gin.Context) {
	userId := c.PostForm("userId")

	service.Login(c, userId)

	c.JSON(200, gin.H{
		"message": "login successful",
	})
}

func GetLogout(c *gin.Context) {
	service.Logout(c)

	c.JSON(200, gin.H{
		"message": "logout successful",
	})
}
