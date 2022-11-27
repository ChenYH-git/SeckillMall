package data_source

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DataBase string `json:"database"`
	LogoMode bool   `json:"logo_mode"`
}

func LoadMysqlConf() *MysqlConf {

	mysqlConf := MysqlConf{}

	file, err := os.Open("conf/mysql_conf.json")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	byteData, err2 := ioutil.ReadAll(file)

	if err2 != nil {
		panic(err2)
	}

	err3 := json.Unmarshal(byteData, &mysqlConf)

	if err3 != nil {
		panic(err3)
	}

	return &mysqlConf

}
