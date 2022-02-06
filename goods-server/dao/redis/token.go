/**
 * @Author：Robby
 * @Date：2022/2/4 23:40
 * @Function：
 **/

package redis

import (
	"go.uber.org/zap"
	"time"
)

// CreateAccessToken 写入access token到Redis中
func CreateAccessToken(accessToken string, expiration time.Duration) (err error) {
	// 写入access token到Redis中
	boolCmd, err := Client.SetNX(getRedisKey(KeyAccessTokenString+accessToken), 1, expiration).Result()
	if err != nil || boolCmd == false { // 如果设置失败
		zap.L().Error("Client.SetNX CreateAccessToken failed", zap.Error(err))
		return
	}
	return

}

func GetAccessToken(accessToken string) bool {
	if err := Client.Get(getRedisKey(KeyAccessTokenString+accessToken)).Err(); err != nil {
		return false
	}
	return true
}

// CreateRefreshToken 写入refresh token到Redis中
func CreateRefreshToken(refreshToken string) (err error) {
	// 写入access token到Redis中
	boolCmd, err := Client.SetNX(getRedisKey(KeyRefreshTokenString+refreshToken), 1, time.Hour * 24 * 3).Result()
	if err != nil || boolCmd == false { // 如果设置失败
		zap.L().Error("Client.SetNX CreateRefreshToken failed", zap.Error(err))
		return
	}
	return
}

// GetRefreshToken 判断accessToken是否存在
func GetRefreshToken(refreshToken string) bool {
	if err := Client.Get(getRedisKey(KeyRefreshTokenString+refreshToken)).Err(); err != nil {
		return false
	}
	return true
}

// CreateAccessRefreshToken 同时创建access_token和refresh_token
func CreateAccessRefreshToken(accessToken, refreshToken string, expiration time.Duration) (err error) {
	pipeline := Client.Pipeline()
	pipeline.SetNX(getRedisKey(KeyAccessTokenString+accessToken), 1, time.Second * expiration)
	pipeline.SetNX(getRedisKey(KeyRefreshTokenString+refreshToken), 1, time.Hour * 3 * 24)
	_, err = pipeline.Exec()
	if err != nil {
		zap.L().Error("pipeline.Exec CreateAccessRefreshToken failed", zap.Error(err))
		return
	}
	zap.L().Info("write access token and refresh token in redis", zap.String("access_token", accessToken), zap.String("refresh_token", refreshToken))
	return
}