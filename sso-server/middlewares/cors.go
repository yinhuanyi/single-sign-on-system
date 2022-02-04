/**
 * @Author: Robby
 * @File name: cors.go
 * @Create date: 2021-06-14
 * @Function:
 **/

package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.Header("Access-Control-Allow-Origin", "*")
		// , localhost:10001
		c.Header("Access-Control-Allow-Origin", "http://localhost:8081")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token, Set-Cookie")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Set-Cookie")
		c.Header("Access-Control-Allow-Credentials", "true")

		method := c.Request.Method
		if method == "OPTIONS" {
			zap.L().Info("OPTIONS Request")
			// 设置响应码, 204不需要设置响应体
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}