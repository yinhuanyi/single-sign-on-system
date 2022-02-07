/**
 * @Author: Robby
 * @File name: keys.go
 * @Create date: 2021-05-27
 * @Function:
 **/

package redis

// redis key 使用命名空间的方式进行区分，使用:作为名称空间的分隔符
const (
	Prefix                = "ipfsmain:"         // 前缀
	KeyLoginFailedUsernameString  = "login_failed:username:"  // string类型，存储用户登录失败的次数
	KeyLoginBlockUsernameString   = "login_block:username:"   // string类型，存储用户登录锁定的username
)

// 获取带前缀的key: ipfsmain:login_failed:username:Robby
func getRedisKey(key string) (newkey string) {
	return Prefix + key
}
