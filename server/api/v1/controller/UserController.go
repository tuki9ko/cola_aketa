package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/v1/service"
)

type UserController struct{}

var us service.UserService

func (uc UserController) GetUser(c *gin.Context) {
	contextUserId, _ := c.Get("userId")
	userId := contextUserId.(int)

	user, err := us.GetUserById(userId)

	msg := "success"

	if err != nil {
		msg = err.Error()
	}

	if user != nil {
		user.HashedPassword = "xxxxxxxxxxxxxxxx"
	}

	c.JSON(200, gin.H{
		"message": msg,
		"user":    user,
	})

}
