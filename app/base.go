package app

import (
	"fmt"
	"go-admin/app/config"
	"go-admin/app/database"
	"go-admin/app/router"
	"log"
	"os"
)

var Conf = new(config.AdminConf)

func init() {

	err := Conf.SetUp("config" + string(os.PathSeparator) + "admin.yaml")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(Conf)
}
func Init() {
	InitLog()
	database.GetMysqlDb(Conf.Db.MysqlConf)
	router.Init(Conf)

	log.Fatal("test记录")
}

func InitLog() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.SetOutput(file)
}
