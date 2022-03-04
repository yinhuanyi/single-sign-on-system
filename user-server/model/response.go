/**
 * @Author：Robby
 * @Date：2021/11/21 1:38 上午
 * @Function：
 **/

package model

import "time"

// SSOToken sso 返回信息
type SSOToken struct {
	AccessToken string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope string `json:"scope,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	ExpiresIn time.Duration `json:"expires_in,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
	Uid string `json:"uid,omitempty"`
}


// UserInfo 用户信息(带角色和权限)

type Permission struct {
	Menus []string `json:"menus"`
	Points []string `json:"points"`
}

type UserInfo struct {
	Id int `json:"id"`
	UserId int `json:"user_id"`
	Username string `json:"username"`
	Title string `json:"title"`
	Avatar string `json:"avatar"`
	Role []*Role `json:"role"`
	Permission *Permission `json:"permission"`
}