package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type AdminConf struct {
	Port string `yaml:"server_port" json:"Port,8080"`
	Db   DB     `yaml:"DB" json:"DB" `
}

type DB struct {
	MysqlConf Mysql `yaml:"MYSQL" json:"MYSQL"`
}

type Mysql struct {
	Host     string `yaml:"HOST" json:"HOST"`
	Port     string `yaml:"PORT" json:"PORT"`
	User     string `yaml:"USER" default:"root"`
	DbName   string `yaml:"DBNAME"`
	Password string `yaml:"PASSWORD" json:"PASSWORD" default:""`
}

func (c *AdminConf) SetUp(yamlFile string) error {
	file := getPath() + string(os.PathSeparator) + yamlFile
	fileDetail, err := os.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(fileDetail, c)
	if err != nil {
		panic(err.Error())
	}
	return err

}

func getPath() string {
	path, _ := os.Getwd()
	return path
}
