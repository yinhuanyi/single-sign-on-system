/**
 * @Author: Robby
 * @File name: error_code.go
 * @Create date: 2021-05-25
 * @Function:
 **/

package mysqlconnect

import "errors"

var (
	ErrorUserExist       = errors.New("用户已经存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户密码错误")
	ErrorInvalidId       = errors.New("无效的ID")
)
