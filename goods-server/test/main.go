/**
 * @Author：Robby
 * @Date：2022/1/27 10:11
 * @Function：
 **/

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type SSOJWTClaim struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

func main() {

	decoded, err := base64.RawStdEncoding.DecodeString("eyJhdWQiOiJnb29kc19pZCIsImV4cCI6MTY0MzI1NzIwMCwic3ViIjoiMTAwMDEifQ")
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("%s", decoded)
	//decodestr := string(decoded)
	//fmt.Println(decodestr)
	ssoJwtClaim := &SSOJWTClaim{}
	err = json.Unmarshal(decoded, ssoJwtClaim)
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Println(ssoJwtClaim)

}
