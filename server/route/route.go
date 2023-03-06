package route

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	controller "github.com/tuki9ko/cola_aketa/api/controller"
	"github.com/tuki9ko/cola_aketa/middleware"
)

var (
	cc controller.ColaController
	lc controller.LoginController
	sc controller.SignupController
	uc controller.UserController
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("cola_session", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:13000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge: 300,
	}))

	api := r.Group("/api")
	{
		api.GET("/login", lc.GetLogin)
		api.POST("/login", lc.PostLogin)
		api.POST("/logout", lc.GetLogout)
		api.GET("/signup", sc.GetSignup)
		api.POST("/signup", sc.PostSignup)

		root := api.Group("/")
		root.Use(middleware.HasValidLoginSession())
		{
			root.GET("/cola", cc.GetCola)
			root.POST("/cola", cc.PostCola)
			root.GET("/colas", cc.GetColas)
			root.GET("/user", uc.GetUser)
			root.PUT("/user", uc.PutUser)
			root.GET("/user/:id/cola", uc.GetColas)
		}
	}
	return r
}
