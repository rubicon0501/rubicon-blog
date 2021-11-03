package global

import (
	"rubicon-blog/pkg/config"
	"time"
)

var (
	ServerSetting   *config.ServerSettingS
	AppSetting      *config.AppSettingS
	DatabaseSetting *config.DatabaseSettingS
)

func LoadConfig() error {
	config, err := config.NewConfig()
	if err != nil {
		return err
	}
	err = config.ReadSection("Server", &ServerSetting)
	if err != nil {
		return err
	}
	err = config.ReadSection("App", &AppSetting)
	if err != nil {
		return err
	}
	err = config.ReadSection("Database", &DatabaseSetting)
	if err != nil {
		return err
	}

	ServerSetting.ReadTimeout *= time.Second
	ServerSetting.WriteTimeout *= time.Second
	return nil
}
