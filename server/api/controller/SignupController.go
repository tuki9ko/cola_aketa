package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/service"
)

type SignupController struct{}

var ss service.SignupService

func (sc SignupController) GetSignup(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET signup",
	})
}

func (sc SignupController) PostSignup(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role_id := 2
	biography := c.PostForm("biography")
	user_icon := c.PostForm("user_icon") // アップロード機構をどうにかする

	ss.Signup(c, name, email, password, role_id, biography, user_icon)

	c.JSON(200, gin.H{
		"message": "signup ${name}",
	})
}
