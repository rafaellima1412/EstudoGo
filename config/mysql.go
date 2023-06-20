package config

import (
	"os"

	"example.com/estudoGo/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMysql() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	_, err := os.Stat("mysql")

	if os.IsNotExist(err) {
		logger.Info("database not found, creating....")
	}

	dsn := "root:root@tcp(127.0.0.1:3306)/mysqldb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Err("Mysql initialization error")
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Err("Mysql automigration error")
		return nil, err
	}

	return db, nil
}
