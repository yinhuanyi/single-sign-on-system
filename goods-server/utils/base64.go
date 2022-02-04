/**
 * @Author：Robby
 * @Date：2022/1/26 16:29
 * @Function：
 **/

package utils

import "encoding/base64"

func GetBase64(username, password string) string {
	message := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte (message))
}