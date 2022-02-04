/**
 * @Author: Robby
 * @File name: user.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package model

import "sso/settings"

type User struct {
	UserId   int64  `json:"user_id"  db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UserLoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ClientScope struct {
	Client settings.ClientConfig
	Scope  []settings.ScopeConfig
	Error  string
}

/*
   data := map[string]interface{}{
       "expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
       "user_id": token.GetUserID(),
       "client_id": token.GetClientID(),
       "scope": token.GetScope(),
       "domain": cli.GetDomain(),
   }
*/

type ClientInfo struct {
	Expire   int64  `json:"expire" binding:"required"`
	UserId   string `json:"user_id" binding:"required"`
	ClientId string `json:"client_id" binding:"required"`
	Scope    string
}
