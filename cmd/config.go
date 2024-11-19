package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// define data source name for mysql
type SQLConfig struct {
	User           string `json:user`
	Password       string `json:password`
	Protocol       string `json:protocol`
	RemoteIPstring string `json:remoteIPstring`
	RemotePort     string `json:remotePort`
	ServerPort     string `json:serverPort`
	ServerPortDesc string `json:serverPortDesc`
	DBname         string `json:dbname`
	Attributes     string `json:attributes`
	Description    string `json:description`
}

func (sc SQLConfig) Dialect() string {
	return "mySQL"
}

func (sc SQLConfig) ConnectionInfo() string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", sc.User, sc.Password, sc.Protocol, sc.RemoteIPstring, sc.RemotePort, sc.DBname, sc.Attributes)
}

func GetConfig() SQLConfig {
	f, err := os.Open(".config")
	if err != nil {
		panic(err)
	}
	var sc SQLConfig
	dec := json.NewDecoder(f)
	err = dec.Decode(&sc)
	if err != nil {
		panic(err)
	}
	fmt.Println("loaded .config")
	return sc
}
