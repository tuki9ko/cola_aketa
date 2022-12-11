package controller

import (
	"fmt"

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
	password := c.PostForm("password")

	user, err := ls.Login(c, userId, password)

	msg := ""

	if user != nil {
		msg = fmt.Sprintf("Login successful, hello %s.", user.Name)
	}

	if err != nil {
		msg = err.Error()
	}

	c.JSON(200, gin.H{
		"message": msg,
	})
}

func (lc LoginController) GetLogout(c *gin.Context) {
	ls.Logout(c)

	c.JSON(200, gin.H{
		"message": "logout successful",
	})
}
