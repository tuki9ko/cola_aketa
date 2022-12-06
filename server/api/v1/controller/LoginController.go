package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/v1/service"
)

type LoginController struct{}

var ls service.LoginService

func (lc LoginController) GetLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET login",
	})
}

func (lc LoginController) PostLogin(c *gin.Context) {
	userId := c.PostForm("userId")

	ls.Login(c, userId)

	c.JSON(200, gin.H{
		"message": "login successful",
	})
}

func (lc LoginController) GetLogout(c *gin.Context) {
	ls.Logout(c)

	c.JSON(200, gin.H{
		"message": "logout successful",
	})
}
