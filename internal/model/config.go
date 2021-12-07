package model

import (
	"rubicon-blog/global"

	"github.com/jinzhu/gorm"
)

var (
	DBEngine        *gorm.DB
	DatabaseSetting *DatabaseSettingS
)

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func SetupDBEngine() error {

	var err error
	DBEngine, err = NewDBEngine(DatabaseSetting, global.ServerSetting.RunMode)
	return err
}
