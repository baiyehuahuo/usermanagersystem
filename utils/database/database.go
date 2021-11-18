package database

import (
	"fmt"
	"usermanagersystem/utils/configReader"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() error {
	var err error
	config := configReader.Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserAccount, config.Password, config.Host, config.Port, config.DbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
