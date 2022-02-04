/**
 * @Author: Robby
 * @File name: user.go
 * @Create date: 2021-11-03
 * @Function:
 **/

package user

import (
	"database/sql"
	mysqlconnect "sso/dao/mysql"
	"sso/model"
	"sso/utils"
)

func GetUserIdByNamePwd(user *model.User) (userId int64, err error) {

	username := user.Username
	currentPassword := user.Password

	sqlStr := "select user_id, username, password from user where username = ?"
	err = mysqlconnect.Db.Get(user, sqlStr, username)
	if err == sql.ErrNoRows {
		return 0, mysqlconnect.ErrorUserNotExist
	}
	if err != nil {
		return
	}

	password := utils.EncryptPassword(currentPassword)
	if password != user.Password {
		return 0, mysqlconnect.ErrorInvalidPassword
	}

	return user.UserId, nil

}
