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
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"test-servers/goods/controllers/response"
	"test-servers/goods/model"
	"test-servers/goods/settings"
	"test-servers/goods/utils"
	"time"
)

//const SSOAuthorize = "http://localhost:10541/api/v1/authorize"
const SSOAuthorize = "http://localhost:8888/api/v1/authorize"
const SSOToken = "http://localhost:10541/api/v1/token"
const ClientId = "goods_id"
const SecretId = "goods_secret"




// JWTAuth token校验，如果客户端携带了token，先查看Redis中是否存在当前token，如果存在并且没有超时，那么直接验证通过。如果Redis中不存在，那么需要发送给SSO服务器，让SSO服务器验证，然后获取userid，最后存储token和userid到Redis中
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 尝试获取token
		token := c.Request.Header.Get("Authorization")
		// 如果token为空，有两种情况
		if token == "" {

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


		// 如果获取到了token, 先从Redis中判断是否存在，如果不存在，直接对sso请求，认证token
		// 此时的从x-token中获取的是 Bearer Token, 这里是获取到Token的值
		tokens := strings.Split(token, " ")
		// 如果大于1，说明有Bearer标识, 获取真实的token
		if len(tokens) > 1 {
			token = tokens[1]
		}

		j := NewJWT()
		claims, err := j.ParseToken(token)
		// 精确判断token解析的错误类型
		if err != nil {
			if err == TokenExpired {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token已过期"})
				c.Abort()
				return
			}else if err == TokenNotValidYet {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token未激活"})
				c.Abort()
				return
			}else if err == TokenInvalid {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token非法"})
				c.Abort()
				return
			}else {
				c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token无法识别"})
				c.Abort()
				return
			}
		}

		c.Set("userId", claims.Id)
		c.Next()
	}
}

// 返回的错误码
var (
	TokenExpired     = errors.New("Token过期")
	TokenNotValidYet = errors.New("Token未激活")
	TokenMalformed   = errors.New("Token错误")
	TokenInvalid     = errors.New("Token非法")
)

type JWT struct {
	SigningKey []byte
}

// NewJWT 获取jwt的服务器端的secret
func NewJWT() *JWT {
	return &JWT{
		[]byte(settings.Conf.JWTConfig.Key),
	}
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析 token
func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}