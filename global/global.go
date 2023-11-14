package global

import (
	"fmt"
	"github.com/device-server/internal/repository/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	_ "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Configs struct {
	DatabaseCfg DatabaseConfig
}
type DatabaseConfig struct {
	Host         string
	Port         int
	UserName     string
	Password     string
	DatabaseName string
}
type RedisConfig struct {
	Host     string
	Port     int
	UserName string
	Password string
	Db       int
	PoolSize int
}
type ServerConfig struct {
	Port int
}

var (
	Cfg *Configs
	Db  *gorm.DB
	Rdb *redis.Client
)

func init() {
	Cfg = &Configs{}
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(Cfg)
	if err != nil {
		panic(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC&interpolateParams=true",
		Cfg.DatabaseCfg.UserName,
		Cfg.DatabaseCfg.Password,
		Cfg.DatabaseCfg.Host,
		Cfg.DatabaseCfg.Port,
		Cfg.DatabaseCfg.DatabaseName,
	)
	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         uint(256),
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{SkipDefaultTransaction: true})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&entity.Account{}, &entity.User{})
}
