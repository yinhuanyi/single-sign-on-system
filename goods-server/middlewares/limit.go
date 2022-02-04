/**
 * @Author: Robby
 * @File name: limit.go
 * @Create date: 2021-05-30
 * @Function:
 **/

package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit..., 取到是0，那么就被限制了
		if bucket.TakeAvailable(1) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{ "msg": "请求过于频繁，请稍后重试" })
			c.Abort()
			return
		}
		c.Next()
	}
}
