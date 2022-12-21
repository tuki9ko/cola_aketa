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

func (uc UserController) PutUser(c *gin.Context) {
	contextUserId, _ := c.Get("userId")
	userId := contextUserId.(int)

	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	role_id := 2
	biography := c.PostForm("biography")
	user_icon := c.PostForm("user_icon")

	us.UpdateUser(userId, name, email, password, role_id, biography, user_icon)

	c.JSON(200, gin.H{
		"message": "update success.",
	})
}
