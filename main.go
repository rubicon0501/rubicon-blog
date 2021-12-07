package main

import (
	"log"
	"net/http"
	"rubicon-blog/global"
	"rubicon-blog/internal/model"
	"rubicon-blog/internal/routers"
	"rubicon-blog/pkg/config"
	"rubicon-blog/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func init() {
	err := LoadConfig()
	if err != nil {
		log.Fatalf("global.LoadConfig err: %v", err)
	}
	log.Printf("Server:%+v", global.ServerSetting)
	log.Printf("App:%+v", global.AppSetting)
	log.Printf("DB:%+v", model.DatabaseSetting)
}

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
	s.ListenAndServe()
}

func LoadConfig() error {
	config, err := config.NewConfig()
	if err != nil {
		return err
	}
	err = config.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	err = config.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = config.ReadSection("Database", &model.DatabaseSetting)
	if err != nil {
		return err
	}
	err = model.SetupDBEngine()
	if err != nil {
		log.Fatalf("model.SetupDBEngine err: %v", err)
		return err
	}

	err = config.ReadSection("Logger", &logger.LoggerSetting)
	if err != nil {
		return err
	}
	err = logger.SetupLogger()
	if err != nil {
		log.Fatalf("logger.SetupLogger err: %v", err)
		return err
	}

	return nil
}
