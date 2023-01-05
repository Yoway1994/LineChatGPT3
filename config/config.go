package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env  string
	Gin  *GinConfig
	Gpt3 *Gpt3Config
	Line *LineConfig
}

type GinConfig struct {
	Host string
	Port string
	Mode string
}

type Gpt3Config struct {
	ApiKey string
}

type LineConfig struct {
	Secret string
	Token  string
}

func NewConfig() *Config {
	configPath := "./"
	runPath, _ := os.Getwd()
	matchPathStatus := false
	pathArr := strings.Split(runPath, "/")
	for i := len(pathArr) - 1; i > 0; i-- {
		configPath += "../"
		if pathArr[i] == "cmd" || pathArr[i] == "test" || pathArr[i] == "migration" {
			matchPathStatus = true
			break
		}
	}
	if !matchPathStatus {
		configPath = "./"
	}
	configPath += "config"
	//
	viper.SetConfigName("config")
	viper.AddConfigPath(configPath)
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		Env: viper.GetString("env"),
		Gin: &GinConfig{
			Host: viper.GetString("server.host"),
			Port: viper.GetString("server.port"),
			Mode: viper.GetString("server.mode"),
		},
		Gpt3: &Gpt3Config{
			ApiKey: viper.GetString("gpt3.apikey"),
		},
		Line: &LineConfig{
			Secret: viper.GetString("line.secret"),
			Token:  viper.GetString("line.token"),
		},
	}
}
