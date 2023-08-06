package src

import (
	"gopkg.in/yaml.v3"
	"os"
)

func SetUp(c interface{}, yamlFile string) {
	file := getPath() + string(os.PathSeparator) + yamlFile
	fileDetail, err := os.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(fileDetail, &c)
	if err != nil {
		panic(err.Error())
	}
}

func getPath() string {
	path, _ := os.Getwd()
	return path
}
