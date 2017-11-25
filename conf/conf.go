package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type DBConf struct {
	Url    string
	DbName string
}

func getDBConf() DBConf {
	file, _ := os.Open("conf/db.json")
	decoder := json.NewDecoder(file)
	conf := DBConf{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(conf)
	return conf
}

var DB = getDBConf()

type ServerConf struct {
	Port string
}

func getServerConf() ServerConf {
	file, _ := os.Open("conf/server.json")
	decoder := json.NewDecoder(file)
	conf := ServerConf{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(conf)
	return conf
}

var Server = getServerConf()
