package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"tinnydouban/global"
	"tinnydouban/internal/model"
	"tinnydouban/internal/routers"
	"tinnydouban/pkg/logger"
	"tinnydouban/pkg/setting"
	"tinnydouban/pkg/validator"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

/**
//init 方法常用于应用程序内的一些初始化操作，它在 main 方法之前自动执行，
//它的执行顺序是：全局变量初始化 =》init 方法 =》main 方法，
//但并不是建议滥用，因为如果 init 过多，你可能会迷失在各个库的 init 方法中，会非常麻烦。
*/
func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupValidator()
	if err != nil {
		log.Fatalf("init.setupValidator err: %v", err)
	}

}

func setupSetting() error {
	setting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}

	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}

func setupValidator() error {
	global.Validator = validator.NewCustomValidator()
	global.Validator.Engine()
	binding.Validator = global.Validator

	return nil
}

// @title 博客系统
// @version 1.0
// @description Go 语言编程之旅：一起用 Go 做项目
// @termsOfService https://github.com/go-programming-tour-book
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
	global.Logger.Infof("%s: go-programming-tour-book/%s", "eddycjy", "blog-service")
}
