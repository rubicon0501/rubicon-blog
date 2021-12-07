package logger

import (
	"errors"
	"log"
	"rubicon-blog/global"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	LoggerSetting *LogSettingS
	logger        *Logger
)

type LogSettingS struct {
	FileName  string
	MaxSizeMB int
	MaxAgeDay int
}

func SetupLogger() error {
	if LoggerSetting == nil {
		return errors.New("Invalid LoggerSetting pointer")
	}
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	logger = NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   LoggerSetting.MaxSizeMB,
		MaxAge:    LoggerSetting.MaxAgeDay,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	if logger != nil {
		return errors.New("Invalid logger pointer")
	}
	return nil
}
