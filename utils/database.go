package utils

import (
	"fmt"

	"github.com/pkg/errors"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDatabase() (err error) {
	config := Config.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.UserAccount, config.Password, config.Host, config.Port, config.DbName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.Wrap(err, RunFuncNameWithFail())
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
