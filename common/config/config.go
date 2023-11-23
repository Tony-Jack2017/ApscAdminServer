package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf Config

type Config struct {
	System        SystemConfig   `json:"system"`
	Database      DatabaseConfig `json:"database"`
	ObjectStorage ObjectStorage  `json:"object_storage"`
}

type SystemConfig struct {
	Name     string `json:"name" yaml:"name"`
	Host     string `json:"host" yaml:"host"`
	HttpPort int    `json:"http_port" yaml:"http_port" mapstructure:"http_port"`
	Mode     string `json:"mode" yaml:"mode"`
	Secret   string `json:"secret" yaml:"secret"`
	AuthCode string `json:"auth_code" yaml:"auth_code" mapstructure:"auth_code"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver" yaml:"driver"`
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	Dbname   string `json:"dbname" yaml:"dbname"`
	Username string `json:"username" json:"username"`
	Password string `json:"password" yaml:"password"`
}

type ObjectStorage struct {
	EndPoint        string `json:"end_point" yaml:"end_point" mapstructure:"end_point"`
	Port            int    `json:"port" yaml:"port"`
	AccessKeyID     string `json:"access_key_id" yaml:"access_key_id" mapstructure:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key" yaml:"secret_access_key" mapstructure:"secret_access_key"`
	UseSSL          bool   `json:"use_ssl" yaml:"use_ssl" mapstructure:"use_ssl"`
}

func ReadConfig() {
	viper.SetConfigName("config.dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Couldn't found the config file")
		} else {
			panic(err)
		}
	}

	if errUnmarshal := viper.Unmarshal(&Conf); errUnmarshal != nil {
		fmt.Printf("Unable to decode into 'Config' struct, %s\n", errUnmarshal)
		panic("program is end >>>>")
	}
}
