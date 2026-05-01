package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf // go-zero 的配置类 可以指定: Name/ Host/ Port/ Mode -> 都在configs/config.yaml配置
	Database      DatabaseConfig
	Redis         RedisConfig
	JWT           JWTConfig
	AppLog        LogConfig
	Upload        UploadConfig
}

type DatabaseConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DBName       string
	Charset      string
	MaxIdleConns int
	MaxOpenConns int
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type JWTConfig struct {
	Secret      string
	ExpireHours int
}

type LogConfig struct {
	Level string
}

type UploadConfig struct {
	Path    string
	MaxSize int64
	BaseURL string
}
