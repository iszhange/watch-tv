package utils

import (
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	mariadb *gorm.DB
	once    sync.Once
)

func mariaDBConnect() {
	once.Do(func() {
		cfg := Config.Databases.MariaDB
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.Charset)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err.Error())
		}

		// 启用连接池
		sqlDB, err := db.DB()
		if err != nil {
			panic(err.Error())
		}
		sqlDB.SetMaxIdleConns(10)
		sqlDB.SetMaxOpenConns(100)

		mariadb = db
	})
}

func GetMariaDBInstance() *gorm.DB {
	return mariadb
}
