/**
 * @Author：Robby
 * @Date：2022/2/4 23:40
 * @Function：
 **/

package redis

import (
	"strconv"
	"time"
)

// IsIpBlock 判断username是否被锁定
func IsIpBlock(username string) bool {
	if err := Client.Get(getRedisKey(KeyLoginBlockUsernameString+username)).Err(); err != nil {
		return false
	}
	return true
}

// IncreaseFailedLoginUsername 记录用户登录失败的次数
func IncreaseFailedLoginUsername(username string)  {
	// 判断登录失败的IP是否存在
	if err := Client.Get(getRedisKey(KeyLoginFailedUsernameString+username)).Err(); err != nil {
		Client.Set(getRedisKey(KeyLoginFailedUsernameString+username), 1, 60 * time.Second)
	}else {
		Client.Incr(getRedisKey(KeyLoginFailedUsernameString+username))
	}
}

// BlockFailedLoginUsername 判断登用户录失败的次数，如果超过了5次，直接block掉
func BlockFailedLoginUsername(username string) bool {
	// 查询登录失败IP的次数
	if numStr , err := Client.Get(getRedisKey(KeyLoginFailedUsernameString+username)).Result(); err == nil {

		// 如果可以获取到值，判断num是否大于5
		num, err := strconv.Atoi(numStr)
		if err != nil {
			num = 0
		}
		// 如果失败次数大于3，锁定IP 5分钟
		if num > 3 {
			Client.Set(getRedisKey(KeyLoginBlockUsernameString+username), 1, 5 * time.Minute)
			return true
		}
	}
	return false
}