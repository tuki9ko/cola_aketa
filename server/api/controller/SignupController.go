package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/request"
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
	var param request.PostSignupParameter

	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(), // TODO: エラーメッセージを直す
		})
		return
	}

	ss.Signup(c, param.Name, param.Email, param.Password, 2, param.Biography, param.UserIcon)

	c.JSON(200, gin.H{
		"message": "signup ${name}",
	})
}
