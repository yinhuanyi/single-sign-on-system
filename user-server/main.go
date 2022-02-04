/**
 * @Author：Robby
 * @Date：2022/1/26 11:15
 * @Function：
 **/

package main

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	mysqlconnect "user-server/dao/mysql"
	"user-server/logger"
	"user-server/pkg/snowflake"
	"user-server/route"
	"user-server/settings"
	"user-server/utils"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("args num：%d, need config file，for example ./goods config/config.yaml\n", len(os.Args)-1)
	}
	log.Println("cmd args success")

	if err := settings.Init(os.Args[1]); err != nil {
		log.Fatalf("config file parse error：%s\n", err.Error())
	}
	log.Println("config file parse success")

	if err := logger.Init(settings.Conf.LogConfig, settings.Conf.Mode); err != nil {
		log.Fatalf("logger init failed：%s\n", err.Error())
	}
	log.Println("logger init success")
	defer zap.L().Sync()

	if err := mysqlconnect.Init(settings.Conf.MysqlConfig); err != nil {
		log.Fatalf("MySQL init failed：%s\n", err.Error())
	}
	log.Println("MySQL init success")
	defer mysqlconnect.Close()

	if err := utils.InitTrans("zh"); err != nil {
		fmt.Printf("translate init error：%v", err)

	}

	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineId); err != nil {
		fmt.Printf("初始化雪花算法失败：%v", err)
	}

	r := route.Init(settings.Conf.Mode)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", settings.Conf.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("Goods server start failed: %s\n", zap.Error(err))
		} else {
			zap.L().Info("Goods server start success\n", zap.Int("Port", settings.Conf.Port))
			fmt.Printf("Goods server start at localhost:%d\n", settings.Conf.Port)
		}
	}()

	fmt.Printf("Server start at localhost:%d\n", settings.Conf.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("Goods server stop error: ", zap.Error(err))
	}

	zap.L().Info("Goods server stop success")
}