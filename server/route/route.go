package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	v1controller "github.com/tuki9ko/cola_aketa/api/v1/controller"
	"github.com/tuki9ko/cola_aketa/middleware"
)

var (
	cc v1controller.ColaController
	lc v1controller.LoginController
	sc v1controller.SignupController
	uc v1controller.UserController
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("cola_session", store))

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/login", lc.GetLogin)
			v1.POST("/login", lc.PostLogin)
			v1.POST("/logout", lc.GetLogout)
			v1.GET("/signup", sc.GetSignup)
			v1.POST("/signup", sc.PostSignup)

			root := v1.Group("/")
			root.Use(middleware.HasValidLoginSession())
			{
				root.GET("/cola", cc.GetCola)
				root.POST("/cola", cc.PostCola)
				root.GET("/colas", cc.GetColas)
				root.GET("/user", uc.GetUser)
				root.PUT("/user", uc.PutUser)
			}
		}
	}
	return r
}
