/**
 * @Author: Robby
 * @File name: settings.go
 * @Create date: 2021-11-02
 * @Function: SSO 配置信息
 **/

package settings

import (
	"log"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

type SSOConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Version string `mapstructure:"version"`
	Port    int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineId  int64 `mapstructure:"machine_id"`

	*LogConfig     `mapstructure:"log"`
	*MysqlConfig   `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
	*JWTConfig	   `mapstructure:"jwt"`

}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host         string `mapstructure:"host"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	DbName       string `mapstructure:"dbname"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"max_open_connection"`
	MaxIdleConns int    `mapstructure:"max_idle_connection"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Db       int    `mapstructure:"dbname"`
	PoolSize int    `mapstructure:"port"`
}

type JWTConfig struct {
	Key string `mapstructure:"key"`
	ExpireTime int64 `mapstructure:"expire_time"`
	Issuer string `mapstructure:"issuer"`
}

var Conf = new(SSOConfig)

func Init(filePath string) (err error) {

	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("config file parse error: %s\n", err.Error())
		return
	}

	if err = viper.Unmarshal(Conf); err != nil {
		log.Printf("config file parse to struct error：%s\n", err.Error())
		return
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("config file modified")
		if err = viper.Unmarshal(Conf); err != nil {
			log.Fatalf("config file reparse error：%s\n", err.Error())
			return
		}
	})

	return

}
