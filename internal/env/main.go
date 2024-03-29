package internal

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.New()

	viper.SetConfigName("pluto")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	var err error = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	return viper.GetString(key)
}

func GetMySQlConfig() string {
	var config string

	config += GetConfig("mysql.username")
	config += ":" + GetConfig("mysql.password")
	config += "@tcp(" + GetConfig("mysql.host") + ":" + GetConfig("mysql.port") + ")"
	config += "/" + GetConfig("mysql.database")
	config += "?multiStatements=true"

	return config
}

func GetPostgreConfig() string {
	var config string

	config += GetConfig("postgre.username")
	config += ":" + GetConfig("postgre.password")
	config += "@(" + GetConfig("postgre.host") + ":" + GetConfig("postgre.port") + ")"
	config += "/" + GetConfig("postgre.database")

	return config
}

func GetSource() string {
	return GetConfig("source")
}
