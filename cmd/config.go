package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// define data source name for mysql
type ServerConfig struct {
	User           string `json:user`
	Password       string `json:password`
	Protocol       string `json:protocol`
	Dialect        string `json:dialect`
	SQLIp          string `json:sqlip`
	SQLPort        string `json:sqlport`
	SQLDescription string `json:sqldescription`
	DBname         string `json:dbname`
	Attributes     string `json:attributes`
	ServerHTTPPort string `json:serverhttpport`
	ServerPortDesc string `json:serverportdesc`
}

func (sc ServerConfig) DialectInfo() string {
	return sc.Dialect
}

func (sc ServerConfig) ConnectionInfo() string {
	return fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s", sc.User, sc.Password, sc.Protocol, sc.SQLIp, sc.SQLPort, sc.DBname, sc.Attributes)
}

func GetConfig() ServerConfig {
	f, err := os.Open(".config")
	if err != nil {
		panic(err)
	}
	var sc ServerConfig
	dec := json.NewDecoder(f)
	err = dec.Decode(&sc)
	if err != nil {
		panic(err)
	}
	fmt.Println("loaded .config")
	return sc
}
