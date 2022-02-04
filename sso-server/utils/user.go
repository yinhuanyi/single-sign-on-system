/**
 * @Author: Robby
 * @File name: user.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package utils

import (
	"crypto/md5"
	"encoding/hex"
	"sso/settings"
	"strings"
)

const secret = "Md5SortKey"

func EncryptPassword(password string) string {

	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))

}

func GetClientObj(clientId string) (client settings.ClientConfig) {

	for _, clientObj := range settings.Conf.Client {
		if clientObj.ClientId == clientId {
			client = clientObj
		}
	}
	return

}

func ScopeNameJoin(scopeList []settings.ScopeConfig) string {

	var scopeNameList []string
	for _, scopeObj := range scopeList {
		scopeNameList = append(scopeNameList, scopeObj.Name)
	}
	return strings.Join(scopeNameList, ",")

}

// GetClientScope 判断client上传的scope和在SSO注册的scope，是否存在交集
func GetClientScope(clientId string, scopeNames string) (scopeList []settings.ScopeConfig) {

	clientObj := GetClientObj(clientId)

	scopeNameList := strings.Split(scopeNames, ",")

	for _, scopeName := range scopeNameList {
		for _, clientScopeObj := range clientObj.ClientScope {
			if scopeName == clientScopeObj.Name {
				scopeList = append(scopeList, clientScopeObj)
			}
		}
	}

	return
}
