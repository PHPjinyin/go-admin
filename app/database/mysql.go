package database

import (
	"fmt"
	"go-admin/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	sync2 "sync"
)

var One sync2.Once
var Db *gorm.DB

func GetMysqlDb(Conf config.Mysql) {
	One.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8", Conf.User, Conf.Password, Conf.Host, Conf.Port, Conf.DbName)
		var err error
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		var u []users
		Db.Raw("SELECT id FROM account WHERE age > ?", 18).Scan(&u)
		fmt.Println(u)
	})
}

type users struct {
	id int
}
