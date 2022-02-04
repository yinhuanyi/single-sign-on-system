/**
 * @Author: Robby
 * @File name: route.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package route

import (
	"github.com/gin-gonic/gin"
	"sso/controllers"
	"sso/middlewares"
)

func Init(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	r.Use(middlewares.RequestLogger(), middlewares.GinRecovery(true),)
	//r.Use(middlewares.RequestLogger(), middlewares.GinRecovery(true), middlewares.Cors())

	v1 := r.Group("/api/v1")
	v1.GET("/authorize", controllers.AuthorizeHandler)     // 获取授权码
	v1.GET("/reauthorize", controllers.ReAuthorizeHandler) // 获取授权码
	v1.Any("/login", controllers.LoginHandler)             // 用户登录同时处理GET和POST请求
	v1.POST("/token", controllers.TokenHandler)            // 获取token，刷新token，refresh也会被刷新
	v1.GET("/verify", controllers.VerifyHandler)           // 验证token
	v1.GET("/logout", controllers.LogoutHandler)           // 用户登出

	return r
}
