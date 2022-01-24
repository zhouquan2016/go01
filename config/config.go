package config

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type datasource struct {
	Host     *string `json:"host"`
	Port     *int    `json:"port"`
	Username *string `json:"username"`
	Password *string `json:"password"`
	Database *string `json:"database"`
	Type     *string `json:"type"`
	Query    *string `json:"query"`
	//最大连接数
	MaxOpenConns int `json:"maxOpenConns"`
	//最大空闲连接数
	MaxIdleConns int `json:"maxIdleConns"`
	//最大空闲
	ConnMaxIdleTime string `json:"connMaxIdleTime"`
	//连接最大时间
	ConnMaxLifetime string `json:"connMaxLifetime"`
}

var Config = &struct {
	Datasource datasource `json:"datasource"`
	PrivateKey string     `json:"private_key"`
	PublicKey  string     `json:"public_key"`
}{}

func init() {
	configFile := "config.json"
	log.Println("config:", configFile)
	path, err := filepath.Abs(configFile)
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(Config)
	if err != nil {
		panic(err)
	}
}
