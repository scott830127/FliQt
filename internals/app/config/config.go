package config

import (
	"github.com/BurntSushi/toml"
	"gorm.io/gorm"
	"os"
)

var C = new(Config)

type Config struct {
	Server ServerConfig
	MySQL  MySQLConfig
	Redis  RedisConfig
	Gorm   GormConfig
}

type ServerConfig struct {
	Port string `toml:"port"`
}

type MySQLConfig struct {
	Enabled     bool   `toml:"enabled"`
	Type        string `toml:"type"`
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	User        string `toml:"user"`
	Password    string `toml:"password"`
	DBName      string `toml:"dbname"`
	MaxIdleConn int    `toml:"max_idle_conn"`
	MaxOpenConn int    `toml:"max_open_conn"`
}

type GormConfig struct {
	Debug bool `toml:"debug"`
}

type RedisConfig struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

func (g GormConfig) ToGormConfig() gorm.Config {
	return gorm.Config{}
}

func Load(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	if _, err := toml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
