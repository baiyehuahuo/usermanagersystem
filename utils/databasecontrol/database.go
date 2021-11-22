package databasecontrol

import (
	"fmt"
	"usermanagersystem/utils/configread"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() error {
	var err error
	config := configread.Config.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserAccount, config.Password, config.Host, config.Port, config.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func GetDB() *gorm.DB {
	return db
}
