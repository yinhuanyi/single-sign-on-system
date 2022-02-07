/**
 * @Author: Robby
 * @File name: sso.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package controllers

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/url"
	"sso/dao/redis"
	"sso/model"
	"sso/oauth2"
	"sso/service"
	"time"

	"sso/session"
	"sso/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

/*
认证的调用顺序：
AuthorizeHandler.....................
LoginHandler GET.....................
LoginHandler POST.....................
ReAuthorizeHandler.....................
*/

// AuthorizeHandler Get接口
func AuthorizeHandler(c *gin.Context) {

	if err := session.Delete(c.Writer, c.Request, "RequestForm"); err != nil {
		zap.L().Error("[AuthorizeHandler]：session.Delete", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	if err := oauth2.Srv.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
		zap.L().Error("[AuthorizeHandler]：Srv.HandleAuthorizeRequest", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

}

// ReAuthorizeHandler Get接口，第二次调用，数据从session中取出来
func ReAuthorizeHandler(c *gin.Context) {
	var err error
	var requestFormString string
	var requestForm url.Values

	if requestFormString, err = session.Get(c.Request, "RequestForm"); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：session.Get", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	if requestForm, err = url.ParseQuery(requestFormString); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：url.ParseQuery", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	// 给请求的form赋值
	c.Request.Form = requestForm

	if err = session.Delete(c.Writer, c.Request, "RequestForm"); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：session.Delete", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}


	if err = oauth2.Srv.HandleAuthorizeRequest(c.Writer, c.Request); err != nil {
		zap.L().Error("[ReAuthorizeHandler]：oauth2.Srv.HandleAuthorizeRequest", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}
}

// 获取requestForm的数据, 为LoginHandler服务
func getRequestForm(c *gin.Context) (data *model.ClientScope, err error) {

	requestForm, err := session.Get(c.Request, "RequestForm")
	if err != nil {
		zap.L().Error("[LoginHandler]：session.Get", zap.Error(err))
		return nil, errors.New(CodeBadRequest.ToString())
	}

	if requestForm == "" {
		zap.L().Info("[LoginHandler]：requestForm == '' ")
		return nil, errors.New(CodeBadRequest.ToString())
	}

	decodeForm, err := url.ParseQuery(requestForm)
	if err != nil {
		return nil, errors.New(CodeServerInternalError.ToString())
	}

	// Get client_id and scope_name from user
	clientID := decodeForm.Get("client_id")
	scope := decodeForm.Get("scope")
	clientObj := utils.GetClientObj(clientID)
	scopeObj := utils.GetClientScope(clientID, scope)
	if scopeObj == nil {
		zap.L().Error("[LoginHandler]：bad scope")
		return nil, errors.New(CodeInvalidParam.ToString())
	}

	return &model.ClientScope{
		Client: clientObj,
		Scope:  scopeObj,
	}, nil

}

// LoginHandler 登录
func LoginHandler(c *gin.Context) {

	switch c.Request.Method {

	case http.MethodGet:
		// 让vue来显示登录页面
		ResponseToLogin(c, "")


	case http.MethodPost:


		// 获取用户提交的用户名和密码
		params := new(model.UserLoginParam)
		if err := c.ShouldBindJSON(params); err != nil {
			zap.L().Error("登录参数校验错误", zap.Error(err))
			errs, ok := err.(validator.ValidationErrors)
			if !ok {
				ResponseError(c, CodeInvalidParam)
				return
			}

			ResponseErrorWithMsg(c, CodeInvalidParam, utils.RemoveTopStruct(errs.Translate(utils.Trans)))
			return
		}

		username := params.Username
		// 判断用户否被锁定
		if redis.IsIpBlock(username) {
			zap.L().Info("username is blocked", zap.String("username", username))
			ResponseErrorWithMsg(c, CodeUsernameIsBlocked, errors.New(fmt.Sprintf("Blocked Username: %s", username)))
			return
		}

		// 基于用户名密码获取userid
		userID, err := service.GetUserIdByNamePwd(params)

		if err != nil || userID == "" {
			zap.L().Error("[LoginHandler]：service.GetUserIdByNamePwd", zap.Error(err))

			// 登录失败，判断当前username是否需要block，如果block了，那么失败次数就不需要再加1了，如果没有block，让失败次数加1
			if isBlocked := redis.BlockFailedLoginUsername(username); !isBlocked {
				redis.IncreaseFailedLoginUsername(username)
			}

			ResponseErrorWithMsg(c, CodeInvalidPassword, err.Error())
			return
		}

		// 将userid写入到session中
		if err = session.Set(c.Writer, c.Request, "LoggedInUserID", userID); err != nil {
			zap.L().Error("[LoginHandler]：session.Set", zap.Error(err))
			ResponseError(c, CodeServerInternalError)
			return
		}

		c.Redirect(http.StatusFound, "/api/v1/reauthorize")
		return

	}
}

// LogoutHandler 登出
func LogoutHandler(c *gin.Context) {

	var redirectUri string

	if redirectUri = c.Query("redirect_uri"); redirectUri == "" {
		zap.L().Error("[LogoutHandler]：c.Query", zap.Error(errors.New("No RedirectUri")))
		ResponseError(c, CodeInvalidParam)
		return
	}

	if err := session.Delete(c.Writer, c.Request, "LoggedInUserID"); err != nil {
		zap.L().Error("[LogoutHandler]：session.Delete", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	c.Redirect(http.StatusFound, redirectUri)

}

// TokenHandler 获取token或刷新token
func TokenHandler(c *gin.Context) {

	if err := oauth2.Srv.HandleTokenRequest(c.Writer, c.Request); err != nil {
		zap.L().Error("[TokenHandler]：oauth2.Srv.HandleTokenRequest", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

}

// VerifyHandler 验证token
func VerifyHandler(c *gin.Context) {

	token, err := oauth2.Srv.ValidationBearerToken(c.Request)
	if err != nil {
		zap.L().Error("[VerifyHandler]：oauth2.Srv.ValidationBearerToken", zap.Error(err))
		ResponseError(c, CodeInvalidToken)
		return
	}

	clientInfo, err := oauth2.Manager.GetClient(context.Background(), token.GetClientID())
	if err != nil {
		zap.L().Error("[VerifyHandler]：oauth2.Manager.GetClient", zap.Error(err))
		ResponseError(c, CodeServerInternalError)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"user_id":    token.GetUserID(),
		"client_id":  token.GetClientID(),
		"scope":      token.GetScope(),
		"domain":     clientInfo.GetDomain(),
	}

	ResponseSuccess(c, data)

}
