package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/request"
	"github.com/tuki9ko/cola_aketa/api/service"
)

type LoginController struct{}

var ls service.LoginService

func (lc LoginController) GetLogin(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET login",
	})
}

func (lc LoginController) PostLogin(c *gin.Context) {
	var param request.PostLoginRequestParameter

	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect user or password",
		})
		return
	}

	user, _ := ls.Login(c, param.UserId, param.Password)

	msg := ""

	if user != nil {
		msg = fmt.Sprintf("Login successful, hello %s.", user.Name)
	} else {
		msg = "Incorrect user or password"
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
