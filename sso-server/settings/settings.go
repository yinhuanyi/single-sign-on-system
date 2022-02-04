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

	*LogConfig     `mapstructure:"log"`
	*MysqlConfig   `mapstructure:"mysql"`
	*RedisConfig   `mapstructure:"redis"`
	*SessionConfig `mapstructure:"session"`
	*Oauth2Config  `mapstructure:"oauth2"`
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

type SessionConfig struct {
	SessionId string `mapstructure:"session_id"`
	HashKey   string `mapstructure:"hash_key"`
}

type Oauth2Config struct {
	Client []ClientConfig `mapstructure:"client"`
}

type ClientConfig struct {
	ClientId     string        `mapstructure:"client_id"`
	ClientSecret string        `mapstructure:"client_secret"`
	ClientName   string        `mapstructure:"client_name"`
	ClientDomain string        `mapstructure:"client_domain"`
	ClientScope  []ScopeConfig `mapstructure:"client_scope"`
}

type ScopeConfig struct {
	Name  string `mapstructure:"name"`
	Title string `mapstructure:"title"`
}

var Conf = new(SSOConfig)

func Init(filePath string) (err error) {

	viper.SetConfigFile(filePath)

	if err = viper.ReadInConfig(); err != nil {
		log.Printf("配置文件读取失败: %s\n", err.Error())
		return
	}

	if err = viper.Unmarshal(Conf); err != nil {
		log.Printf("配置文件映射到结构体失败：%s\n", err.Error())
		return
	}

	viper.WatchConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("配置文件修改了")
		if err = viper.Unmarshal(Conf); err != nil {
			log.Fatalf("配置文件重新加载失败：%s\n", err.Error())
			return
		}
	})

	return

}
