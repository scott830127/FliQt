package gormx

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Connection struct {
	Host        string
	Port        int
	User        string
	Password    string
	DBName      string
	MaxIdleConn int
	MaxOpenConn int
}

func (c Connection) DSN(driver string, withDB bool) string {
	if driver == "mysql" {
		if withDB {
			return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port, c.DBName)
		}
		return fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.Password, c.Host, c.Port)
	}
	if driver == "postgres" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", c.Host, c.Port, c.User, c.DBName, c.Password)
	}
	return ""
}

type DBConfig struct {
	Enabled       bool
	Type          string // mysql or postgres
	Master        Connection
	Debug         bool
	AutoMigration bool
	GormConfig    gorm.Config
}

func NewDB(cfg *DBConfig, autoMigrateModels []any) (db *gorm.DB, clean func(), err error) {
	clean = func() {}
	if !cfg.Enabled {
		return
	}
	if cfg.Type == "" {
		err = errors.New("db type must be specified")
		return
	}
	dial, err := newDial(cfg.Type, cfg.Master, true)
	if err != nil {
		return
	}
	db, err = gorm.Open(dial, &cfg.GormConfig)
	if err != nil {
		return
	}
	if cfg.AutoMigration {
		err = db.AutoMigrate(autoMigrateModels...)
		if err != nil {
			return
		}
	}
	return
}

func newDial(dbType string, conn Connection, withDB bool) (gorm.Dialector, error) {
	switch dbType {
	case "mysql":
		return mysql.Open(conn.DSN("mysql", withDB)), nil
	case "postgres":
		return postgres.Open(conn.DSN("postgres", withDB)), nil
	default:
		return nil, fmt.Errorf("unsupported db type: %s", dbType)
	}
}
