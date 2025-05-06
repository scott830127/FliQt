package di

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

func NewMainDB(cfg *config.Config) (*gorm.DB, func(), error) {
	db, cleanup, err := gormx.NewDB(&gormx.DBConfig{
		Enabled: cfg.MySQL.Enabled,
		Type:    cfg.MySQL.Type,
		Master: gormx.Connection{
			Host:        cfg.MySQL.Host,
			Port:        cfg.MySQL.Port,
			User:        cfg.MySQL.User,
			Password:    cfg.MySQL.Password,
			DBName:      cfg.MySQL.DBName,
			MaxIdleConn: cfg.MySQL.MaxIdleConn,
			MaxOpenConn: cfg.MySQL.MaxOpenConn,
		},
		Debug:         cfg.Gorm.Debug,
		AutoMigration: true,
		GormConfig:    cfg.Gorm.ToGormConfig(),
	}, entity.AutoMigrate())
	if err != nil {
		return nil, nil, err
	}

	// 執行 seed
	if seedErr := entity.SeedInitialData(db); seedErr != nil {
		return nil, nil, seedErr
	}

	return db, cleanup, nil
}
