package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/api/request"
	"github.com/tuki9ko/cola_aketa/api/service"
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

	var param request.PutUserParameter

	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(), // TODO: エラーメッセージを直す
		})
	}

	us.UpdateUser(userId, param.Name, param.Email, param.Password, 2, param.Biography, param.UserIcon)

	c.JSON(200, gin.H{
		"message": "update success.",
	})
}

func (uc UserController) GetColas(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	colas, err := cs.GetColas(userId)

	msg := "success"

	if err != nil {
		msg = err.Error()
	}

	c.JSON(200, gin.H{
		"message": msg,
		"colas":   colas,
	})
}
