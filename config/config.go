package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Env   string
	Gin   *GinConfig
	Gpt3  *Gpt3Config
	Line  *LineConfig
	Redis *RedisConfig
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
	DevS   string
	DevT   string
}

type RedisConfig struct {
	Host           string
	Port           int
	Database       int
	Auth           string
	Max_idle       int
	Max_active     int
	Idle_timeout   int
	Notify_active  int
	Polling_active int
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
	// configPath += "config"
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
			DevS:   viper.GetString("line.devs"),
			DevT:   viper.GetString("line.devt"),
		},
		Redis: &RedisConfig{
			Host:           viper.GetString("redis.host"),
			Port:           viper.GetInt("redis.port"),
			Database:       viper.GetInt("redis.database"),
			Auth:           viper.GetString("redis.auth"),
			Max_idle:       viper.GetInt("redis.max_idle"),
			Max_active:     viper.GetInt("redis.max_active"),
			Idle_timeout:   viper.GetInt("redis.idle_timeout"),
			Notify_active:  viper.GetInt("redis.notify_active"),
			Polling_active: viper.GetInt("redis.polling_active"),
		},
	}
}
