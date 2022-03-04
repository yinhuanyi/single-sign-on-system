/**
 * @Author：Robby
 * @Date：2022/3/1 13:41
 * @Function：
 **/

package user

import (
	"database/sql"
	mysqlconnect "user-server/dao/mysql"
	"user-server/model"
)

// GetUser 查询user_user表信息
func GetUser(user *model.User) (userRes *model.User, err error) {

	sqlStr := `select id, user_id, username, title, avatar, role_id from user_user where user_id = ?`
	err = mysqlconnect.Db.Get(user, sqlStr, user.UserId)

	if err == sql.ErrNoRows {
		return nil, mysqlconnect.ErrorUserNotExist
	}
	if err != nil {
		return nil, err
	}

	userRes = user
	return
}

// GetRole 查询 user_role表信息
func GetRole(role *model.Role) (roleRes *model.Role, err error) {
	sqlStr := `select id, title from user_role where id = ?`
	err = mysqlconnect.Db.Get(role, sqlStr, role.Id)
	if err == sql.ErrNoRows {
		return nil, mysqlconnect.ErrorUserNotExist
	}
	if err != nil {
		return nil, err
	}
	roleRes = role
	return
}

type Menus struct {
	PermissionMark []string `db:"permission_mark"`
}

// GetPermissionMenus 联查 user_permission表、user_role_permission表，获取menus权限
func GetPermissionMenus(roleId string) (menus []string, err error) {
	sqlStr := `select permission_mark from user_permission left join user_role_permission on user_permission.id=user_role_permission.permission_id where permission_id not like '%-%' and role_id=?;`
	// 虽然role_id字段在表中是int类型，这里的roleId是string类型，但是作为SQL的参数可以
	err = mysqlconnect.Db.Select(&menus, sqlStr, roleId)
	if err != nil {
		return nil, err
	}
	return
}

// GetPermissionPoint 联查user_permission表、user_role_permission表，获取points权限
func GetPermissionPoint(roleId string) (points []string, err error) {
	sqlStr := `select permission_mark from user_permission left join user_role_permission on user_permission.id=user_role_permission.permission_id where permission_id like '%-%' and role_id=?;`
	err = mysqlconnect.Db.Select(&points, sqlStr, roleId)
	if err != nil {
		return nil, err
	}
	return
}