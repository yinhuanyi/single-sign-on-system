/**
 * @Author: Robby
 * @File name: user.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package model

import "github.com/dgrijalva/jwt-go"

type CustomClaims struct {
	Id       int64
	NickName string
	Role     string

	jwt.StandardClaims
}

type SSOJWTClaim struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

// User user_user 表用户信息
type User struct {
	Id       int    `json:"id" db:"id"`
	UserId   int    `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Title    string `json:"title" db:"title"`
	Avatar   string `json:"avatar" db:"avatar"`
	RoleId   string `json:"role_id" db:"role_id"`
}

// Role user_role表角色信息
type Role struct {
	Id int `json:"id" db:"id"`
	Title string `json:"title" db:"title"`
	Describe string `json:"describe,omitempty" db:"describe"`
}
