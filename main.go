package main

import (
	"log"
	"net/http"
	"rubicon-blog/global"
	"rubicon-blog/internal/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	err := global.LoadConfig()
	if err != nil {
		log.Fatalf("config.LoadConfig err: %v", err)
	}
	log.Printf("Server:%+v", global.ServerSetting)
	log.Printf("App:%+v", global.AppSetting)
	log.Printf("DB:%+v", global.DatabaseSetting)

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
