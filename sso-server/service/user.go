/**
 * @Author: Robby
 * @File name: user.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package service

import (
	"sso/dao/mysql/user"
	"sso/model"
	"strconv"
)

func GetUserIdByNamePwd(param *model.UserLoginParam) (userId string, err error) {

	userInstance := model.User{
		Username: param.Username,
		Password: param.Password,
	}

	userIdInt, err := user.GetUserIdByNamePwd(&userInstance)
	if err != nil {
		return
	}

	userId = strconv.Itoa(int(userIdInt))
	return
}
