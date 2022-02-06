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
	KeyAccessTokenString  = "sso:access_token"  // string类型，存储access_token
	KeyRefreshTokenString = "sso:refresh_token" // string类型，存储access_token
)

// 获取带前缀的key: ipfsmain:sso:refresh_token:MZI4MWEYMZITNGRLZI01ZTM5LTKZMTQTYTJJYJE2MGYZYJHL
func getRedisKey(key string) (newkey string) {
	return Prefix + key
}
