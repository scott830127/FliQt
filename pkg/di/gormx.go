package di

import "C"
import (
	"FliQt/internals/app/config"
	"FliQt/internals/app/entity"
	"FliQt/pkg/gormx"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var WireGormSet = wire.NewSet(
	NewMainDB,
)

func NewMainDB() (*gorm.DB, func(), error) {
	return gormx.NewDB(&gormx.DBConfig{
		Enabled: config.C.MySQL.Enabled,
		Type:    config.C.MySQL.Type,
		Master: gormx.Connection{
			Host:        config.C.MySQL.Host,
			Port:        config.C.MySQL.Port,
			User:        config.C.MySQL.User,
			Password:    config.C.MySQL.Password,
			DBName:      config.C.MySQL.DBName,
			MaxIdleConn: config.C.MySQL.MaxIdleConn,
			MaxOpenConn: config.C.MySQL.MaxOpenConn,
		},
		Debug:         config.C.Gorm.Debug,
		AutoMigration: true,
		GormConfig:    config.C.Gorm.ToGormConfig(),
	}, entity.AutoMigrate())
}
