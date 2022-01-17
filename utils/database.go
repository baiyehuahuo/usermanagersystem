package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"usermanagersystem/consts"

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
		return ErrWrapOrWithMessage(true, err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

// BackupMySQL 备份MySQL
func BackupMySQL() {
	config := Config.MysqlConfig
	var cmd *exec.Cmd

	cmd = exec.Command("mysqldump", "--opt", "-h"+config.Host, "-P"+strconv.Itoa(config.Port),
		"-u"+config.UserAccount, "-p"+config.Password, config.DbName)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Println(err)
	}

	if err := cmd.Start(); err != nil {
		log.Println(err)
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Println(err)
	}

	err = ioutil.WriteFile(consts.MySQLBackUpPath, bytes, 0644)
	if err != nil {
		log.Println(err)
	}
}

// RestoreMySQL 恢复数据库
func RestoreMySQL() {
	path, err := filepath.Abs(consts.MySQLBackUpPath)
	if err != nil {
		panic(err)
	}
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	restores := strings.Split(string(file), ";")
	for _, restore := range restores {
		db.Exec(restore)
	}
}
