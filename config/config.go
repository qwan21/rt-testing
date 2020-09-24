package config

import "github.com/spf13/viper"

//Init is initalization of config
func Init() error {
	viper.AddConfigPath("etc/rt")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("~/.rt")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
