/**
 * @Author：Robby
 * @Date：2021/11/21 6:06 下午
 * @Function：
 **/

package service

import (
	"go.uber.org/zap"
	"strconv"
	"strings"
	mysqlUser "user-server/dao/mysql/user"
	"user-server/model"
)

// GetUser 获取用户信息
func GetUser(user *model.User) (userInfo *model.UserInfo, err error) {

	// 获取用户信息
	userInstance, err := mysqlUser.GetUser(user)

	// 获取用户角色
	roleIds := strings.Split(userInstance.RoleId, ",")
	//fmt.Println(roleIds)
	var roles []*model.Role
	for _, roleIdStr := range roleIds {
		roleId, err := strconv.Atoi(roleIdStr)
		if err != nil {
			zap.L().Error("roleId assert error")
			continue
		}
		role := &model.Role{Id: roleId}
		role , err = mysqlUser.GetRole(role)
		if err != nil {
			zap.L().Error("Get Role based on RoleId Error", zap.Int("roleId", roleId))
			continue
		}
		roles = append(roles, role)
	}

	// 获取用户权限（基于roleId）
	var menus []string
	var points []string
	menusMap := map[string]struct{}{}
	pointsMap := map[string]struct{}{}

	// 基于角色列表获取权限列表
	for _, roleId := range roleIds {

		// 获取菜单权限
		permissionMenus, err := mysqlUser.GetPermissionMenus(roleId)
		if err != nil {
			zap.L().Error("Get Permisson Menus based on RoleId Error", zap.String("roleId", roleId))

		}

		// 菜单权限列表去重
		for _, permissionMenu := range permissionMenus {
			if _, ok := menusMap[permissionMenu]; !ok {
				menusMap[permissionMenu] = struct{}{}
			}
		}

		// 获取操作权限
		permissionPoints, err := mysqlUser.GetPermissionPoint(roleId)
		if err != nil {
			zap.L().Error("Get Permisson Menus based on RoleId Error", zap.String("roleId", roleId))
			continue
		}

		for _, permissionPoint := range permissionPoints {
			if _, ok := pointsMap[permissionPoint]; !ok {
				pointsMap[permissionPoint] = struct{}{}
			}
		}

	}

	// 获取menus
	menus = make([]string, 0, len(menusMap))
	for menu := range menusMap {
		menus = append(menus, menu)
	}

	// 获取points
	points = make([]string, 0, len(pointsMap))
	for point := range pointsMap {
		points = append(points, point)
	}

	// 创建userInfo
	userInfo = &model.UserInfo{
		Id: userInstance.Id,
		UserId: userInstance.UserId,
		Username: userInstance.Username,
		Title: userInstance.Title,
		Avatar: userInstance.Avatar,
		Role: roles,
		Permission: &model.Permission{
			Menus:  menus,
			Points: points,
		},
	}

	return

}