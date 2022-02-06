/**
 * @Author: Robby
 * @File name: route.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package route

import (
	"github.com/gin-gonic/gin"
	"user-server/controllers"
	"user-server/middlewares"
)

func Init(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	//r.Use(middlewares.RequestLogger(), middlewares.GinRecovery(true), middlewares.Cors())
	r.Use(middlewares.RequestLogger(), middlewares.GinRecovery(true))

	v1 := r.Group("/api/v1")

	{
		v1.GET("/sso_login", middlewares.JWTAuth(), controllers.SSOLoginHandler)
	}

	userRouter := v1.Group("/user")

	userRouter.Use(middlewares.JWTAuth())
	{
		userRouter.GET("/profile", controllers.UserProfileHandler) // 增
	}

	return r
}
