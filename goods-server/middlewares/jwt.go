/**
 * @Author: Robby
 * @File name: jwt.go
 * @Create date: 2021-06-14
 * @Function:  业务上认为 Token类型的错误，定义为http协议的错误
 **/

package middlewares

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"goods-server/controllers/response"
	redisconnect "goods-server/dao/redis"
	"goods-server/model"
	"goods-server/settings"
	"goods-server/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)


const SSOAuthorize = "http://localhost:8888/api/v1/authorize"
const SSOToken = "http://localhost:10541/api/v1/token"
const ClientId = "goods_id"
const SecretId = "goods_secret"


// JWTAuth token校验，如果客户端携带了token，先查看Redis中是否存在当前token，如果存在并且没有超时，那么直接验证通过。如果Redis中不存在，那么需要发送给SSO服务器，让SSO服务器验证，然后获取userid，最后存储token和userid到Redis中
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 尝试获取token
		accessToken := c.Request.Header.Get("Authorization")
		// 如果token为空，有两种情况
		if accessToken == "" {

			// 1：先获取请求的查询，判断是否有code字段，如果有code字段，说明是SSO执行回调的请求
			rawQuery := c.Request.URL.RawQuery
			values, err := url.ParseQuery(rawQuery)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
				c.Abort()
				return
			}

			if _, ok := values["code"]; !ok {
				// code不存在，说明是用户请求，那么让用户重定向到SSO，完成用户登录
				ssoUrlForCode := fmt.Sprintf("%s?client_id=%s&response_type=code&scope=all&state=xyz&redirect_uri=http://localhost:%d%s",SSOAuthorize, ClientId, 8888, c.Request.URL.Path)
				fmt.Println("ssoUrl = " + ssoUrlForCode)
				c.Redirect(http.StatusFound, ssoUrlForCode)
				c.Abort()
				return
			}
			// 获取SSO的code授权码
			code := values["code"][0]
			// 基于code授权码，请求SSO获取token。下面是构建HTTP客户端请求
			client := http.Client{Timeout: 3 * time.Second}
			payload := strings.NewReader(fmt.Sprintf("grant_type=authorization_code&code=%s&redirect_uri=http://localhost:%d%s", code, 8888, c.Request.URL.Path) )
			req, err := http.NewRequest("POST", SSOToken, payload)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
				c.Abort()
				return
			}
			// 基于clientid和secretid生成base auth
			basicAuth := utils.GetBase64(ClientId, SecretId)
			req.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			// 请求sso，获取token
			resp, err := client.Do(req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
				c.Abort()
				return
			}
			content, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
				c.Abort()
				return
			}

			ssoToken := &model.SSOToken{}
			err = json.Unmarshal([]byte(content), ssoToken)
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error" + err.Error()})
				c.Abort()
				return
			}

			if ssoToken.ExpiresIn == 0 {
				response.ResponseErrorWithMsg(c, response.CodeBadRequest, "code is invalid")
				c.Abort()
				return
			}

			// 写入access_token和refresh_token到redis中, 如果写入错误，不处理
			_ = redisconnect.CreateAccessRefreshToken(ssoToken.AccessToken, ssoToken.RefreshToken, ssoToken.ExpiresIn)

			// 这里直接返回，加上一个redirect字段，浏览器获取到这redirect字段，说明需要重新发起请求给对应的url
			redirectUrl := fmt.Sprintf("http://localhost:%d%s", settings.Conf.Port, c.Request.URL.Path)
			// 在这里可以给ssotoken的返回添加用户信息和uid信息
			ssoToken.RedirectUrl = redirectUrl
			// 获取用户的uid
			accessToken := ssoToken.AccessToken
			// base64解码
			base64Code := strings.Split(accessToken, ".")[1]
			claimUpload, err := base64.RawStdEncoding.DecodeString(base64Code)
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error" + err.Error()})
				c.Abort()
				return
			}
			ssoJwtClaim := &model.SSOJWTClaim{}
			if err = json.Unmarshal(claimUpload, ssoJwtClaim); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error" + err.Error()})
				c.Abort()
				return
			}

			uid := ssoJwtClaim.Subject
			ssoToken.Uid = uid

			response.ResponseSuccess(c, ssoToken)
			c.Abort()
			return
		}

		// 获取accessToken
		accessTokens := strings.Split(accessToken," ")
		if len(accessTokens) > 1 {
			accessToken = accessTokens[1]
		}

		// 验证accessToken
		isAccessTokenExist := redisconnect.GetAccessToken(accessToken)

		// 如果存在access_token，表示token验证通过, 获取Uid写入到context中
		if isAccessTokenExist {
			// base64解码
			base64Code := strings.Split(accessToken, ".")[1]
			claimUpload, err := base64.RawStdEncoding.DecodeString(base64Code)
			if err != nil{
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error" + err.Error()})
				c.Abort()
				return
			}
			ssoJwtClaim := &model.SSOJWTClaim{}
			if err = json.Unmarshal(claimUpload, ssoJwtClaim); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error" + err.Error()})
				c.Abort()
				return
			}
			c.Set("userId", ssoJwtClaim.Subject)
			c.Next()

		}else { // 如果不存在access token，说明access token过期，那么获取refresh token
			refreshToken := c.Request.Header.Get("Refresh-Token")
			isRefreshTokenExist := redisconnect.GetRefreshToken(refreshToken)
			// 如果isRefreshTokenExist存在，说明refreshtoken有效，那么可以直接基于refresh token更新access token
			if isRefreshTokenExist {
				// 构建HTTP请求，更新access token
				client := http.Client{Timeout: 30000 * time.Second}
				payload := strings.NewReader(fmt.Sprintf("grant_type=refresh_token&refresh_token=%s", refreshToken))
				req, err := http.NewRequest("POST", SSOToken, payload)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
					c.Abort()
					return
				}
				// 基于clientid和secretid生成base auth
				basicAuth := utils.GetBase64(ClientId, SecretId)
				req.Header.Set("Authorization", fmt.Sprintf("Basic %s", basicAuth))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				// 请求sso，获取token
				resp, err := client.Do(req)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
					c.Abort()
					return
				}
				content, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error: " + err.Error()})
					c.Abort()
					return
				}
				ssoToken := &model.SSOToken{}
				err = json.Unmarshal([]byte(content), ssoToken)
				if err != nil{
					c.JSON(http.StatusInternalServerError, gin.H{"msg": "server error" + err.Error()})
					c.Abort()
					return
				}

				if ssoToken.ExpiresIn == 0 {
					response.ResponseErrorWithMsg(c, response.CodeBadRequest, "code is invalid")
					c.Abort()
					return
				}

				_ = redisconnect.CreateAccessRefreshToken(ssoToken.AccessToken, ssoToken.RefreshToken, ssoToken.ExpiresIn)

				response.Response302(c, ssoToken)
				c.Abort()
				return
			}else { // 如果refresh_token也过期了，那么让浏览器清除access_token和refresh_token，重新登录即可
				response.Response303(c)
				c.Abort()
				return
			}
		}
	}
}