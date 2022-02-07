/**
 * @Author: Robby
 * @File name: oanth2.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package oauth2

import (
	"github.com/go-oauth2/oauth2/v4/errors"
	"go.uber.org/zap"
	"log"
	"net/http"
	"sso/session"
	"sso/settings"

	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

// Manager 提供认证管理
var Manager *manage.Manager
// Srv 提供认证服务
var Srv *server.Server

func Init(cfg *settings.Oauth2Config) (err error) {

	Manager = manage.NewDefaultManager()

	// 指定token存储的位置为Redis的1号库
	Manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   1,
	}))

	// 配置access_token和refresh_token的过期时间
	Manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)
	// 生成token，这里传递的kid是添加到jwt的头部的，如果是""，那么不会添加，ipfs字符串是加密的key
	Manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("ipfs"), jwt.SigningMethodHS256))
	// 创建一个map，存储客户端的信息
	clientStore := store.NewClientStore()

	// 将客户端信息写入到本地内存中
	for _, v := range cfg.Client {
		if err = clientStore.Set(v.ClientId, &models.Client{
			ID:     v.ClientId,
			Secret: v.ClientSecret,
			Domain: v.ClientDomain,
		}); err != nil {
			log.Printf("客户端注册SSO失败：%s\n", err.Error())
			return
		}
	}
	// 添加客户端信息到manager中
	Manager.MapClientStorage(clientStore)
	// 创建 Authorization server
	Srv = server.NewDefaultServer(Manager)
	//// 基于用户名和密码，验证用户
	//Srv.SetPasswordAuthorizationHandler(passwordAuthorizationHandler)
	// 基于用户的请求，验证用户，这里如果从session获取到useid为空，那么会重定向跳转到登录接口
	Srv.SetUserAuthorizationHandler(userAuthorizeHandler)
	// 基于用户请求，获取scope
	//Srv.SetAuthorizeScopeHandler(authorizeScopeHandler)
	Srv.SetInternalErrorHandler(internalErrorHandler)
	Srv.SetResponseErrorHandler(responseErrorHandler)

	return
}

func userAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userId string, err error) {

	if userId, err = session.Get(r, "LoggedInUserID"); err != nil {
		return
	}

	if userId == "" {

		if r.Form == nil {
			if err = r.ParseForm(); err != nil {
				return
			}
		}

		if err = session.Set(w, r, "RequestForm", r.Form.Encode()); err != nil {
			return
		}


		http.Redirect(w, r, "/api/v1/login", http.StatusFound)


	}

	// 如果代码userId有值，那么这里return后，就会发起 http://localhost:10001/api/v1/goods/get?code=NWE4MDG5N2QTNDRMOC0ZMDGWLWI2NMYTZTNINDAYOWNHZDKX&state=xyz 请求到 client端
	return
}

func internalErrorHandler(err error) (re *errors.Response) {
	zap.L().Error("Oauth2.0 Internal Error", zap.Error(err))

	return
}

func responseErrorHandler(re *errors.Response) {
	zap.L().Error("Oauth2.0 Response Error", zap.Error(re.Error))

	return
}
