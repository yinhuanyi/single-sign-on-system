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
