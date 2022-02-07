/**
 * @Author: Robby
 * @File name: main.go
 * @Create date: 2021-11-02
 * @Function:
 **/

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	mysqlconnect "sso/dao/mysql"
	redisconnect "sso/dao/redis"
	"sso/logger"
	"sso/oauth2"
	"sso/route"
	"sso/session"
	"sso/settings"
	"sso/utils"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("参数个数为：%d, 需要带上配置文件，例如运行 ./sso config/config.yaml\n", len(os.Args)-1)
	}
	log.Println("命令行参数解析成功")

	if err := settings.Init(os.Args[1]); err != nil {
		log.Fatalf("配置文件解析失败：%s\n", err.Error())
	}
	log.Println("配置文件读取成功")

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		log.Fatalf("日志配置失败：%s\n", err.Error())
	}
	log.Println("日志配置成功")

	defer zap.L().Sync()

	if err := mysqlconnect.Init(settings.Conf.MysqlConfig); err != nil {
		log.Fatalf("MySQL连接失败：%s\n", err.Error())
	}
	log.Println("MySQL连接成功")
	defer mysqlconnect.Close()

	if err := redisconnect.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("Redis init faied：%v", err)
	}
	defer redisconnect.Close()

	if err := session.Init(settings.Conf.SessionConfig); err != nil {
		log.Fatalf("初始化session失败：%s\n", err.Error())
	}
	log.Println("初始化session成功")

	// 将应用站点信息注册到Oauth2.0 Server中
	if err := oauth2.Init(settings.Conf.Oauth2Config); err != nil {
		log.Fatalf("初始化Oauth2.0失败：%s\n", err.Error())
	}
	log.Println("初始化Oauth2.0成功")

	r := route.Init(settings.Conf.Mode)

	if err := utils.InitTrans("zh"); err != nil {
		fmt.Printf("翻译器获取失败：%v", err)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("SSO服务启动失败: %s\n", zap.Error(err))
		} else {
			zap.L().Info("SSO服务启动成功：", zap.Int("Port", settings.Conf.Port))
		}
	}()

	fmt.Printf("Server start at localhost:%d\n", settings.Conf.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("SSO服务停止失败: ", zap.Error(err))
	}

	zap.L().Info("SSO 服务停止")
}
