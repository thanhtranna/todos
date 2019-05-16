package storage

import (
	"errors"
	"fmt"

	"todo-lists/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
)

func NewDb(c *config.Config) (*gorm.DB, error) {
	if c.DB.Use == "mysql" {
		return newMySQL(c)
	}
	return nil, errors.New("Not supported db")
}

func newMySQL(c *config.Config) (*gorm.DB, error) {
	mySQLInfo := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		c.DB.Mysql.UserName,
		c.DB.Mysql.Password,
		c.DB.Mysql.Database)
	db, err := gorm.Open("mysql", mySQLInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}