package service

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type LoginService struct{}

func (s LoginService) Login(c *gin.Context, userId string) {
	session := sessions.Default(c)
	session.Set("userId", userId)
	session.Save()
}

func (s LoginService) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}
