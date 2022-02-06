/**
 * @Author：Robby
 * @Date：2021/11/21 1:38 上午
 * @Function：
 **/

package model

import "time"

type SSOToken struct {
	AccessToken string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	Scope string `json:"scope,omitempty"`
	TokenType string `json:"token_type,omitempty"`
	ExpiresIn time.Duration `json:"expires_in,omitempty"`
	RedirectUrl string `json:"redirect_url,omitempty"`
	Uid string `json:"uid,omitempty"`
}
