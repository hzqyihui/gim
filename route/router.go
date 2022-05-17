package route

import (
	"gim/controller/api"
	"gim/controller/msgBoard"
	"gim/middleware"
	"github.com/gin-gonic/gin"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	//r.Use(middleware.Cors())
	//r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api")
	{
		v1.GET("ping", api.Ping)
		// 用户注册
		v1.POST("user/register", api.UserRegister)
		// 用户登录
		v1.POST("user/login", api.UserLogin)
		// 用户注销
		v1.POST("user/logout", api.UserLogout)

		v1.POST("msgBoard/save",msgBoard.InsertMsg)
		v1.GET("msgBoard/list",msgBoard.MsgList)

	}

	// swagger文档
	// 游览器打开 http://localhost:3000/swagger/index.html
	r.StaticFile("/swagger.json", "./swagger/swagger.json")
	r.Static("/swagger", "./swagger/dist")

	return r
}
