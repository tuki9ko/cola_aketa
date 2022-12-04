package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/tuki9ko/cola_aketa/session"
)

var LoginInfo session.SessionInfo

func HasValidLoginSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		LoginInfo.UserId = session.Get("userId")

		if LoginInfo.UserId == nil {
			c.JSON(200, gin.H{
				"message": "Your session expired. Please login again.",
			})
			c.Abort()
		} else {
			c.Set("userId", LoginInfo.UserId)
			c.Next()
		}
	}
}
