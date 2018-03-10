package utils

import (
	"github.com/spf13/viper"
	"fmt"
	"strconv"
)

func GetConfig() *viper.Viper {
	c := viper.New()
	c.SetConfigType("yaml")
	c.SetConfigName("config")
	c.AddConfigPath(".")

	c.SetDefault("database.host", "127.0.0.1")
	c.SetDefault("database.port", 3306)
	c.SetDefault("database.name", "shadowsocks")
	c.SetDefault("database.user", "root")
	c.SetDefault("database.pass", "password")

	c.SetDefault("redis.host", "localhost")
	c.SetDefault("redis.port", 4468)
	c.SetDefault("redis.enableAuth", true)
	c.SetDefault("redis.password", "password")

	c.SetDefault("cacheTTL", 60)
	c.SetDefault("verifyKey", "Hello")

	c.ReadInConfig()
	return c
}

func GetDataUrl() string {
	config := GetConfig()
	return fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true",
		config.Get("database.user"), config.Get("database.pass"), config.Get("database.host"),
		strconv.Itoa(config.Get("database.port").(int)), config.Get("database.name"))
}