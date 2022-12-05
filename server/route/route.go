package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1controller "github.com/tuki9ko/cola_aketa/api/v1/controller"
	"github.com/tuki9ko/cola_aketa/middleware"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("cola_session", store))

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/login", v1controller.GetLogin)
			v1.POST("/login", v1controller.PostLogin)
			v1.POST("/logout", v1controller.GetLogout)

			root := v1.Group("/")
			root.Use(middleware.HasValidLoginSession())
			{
				root.GET("/cola", v1controller.GetCola)
			}
		}
	}
	return r
}