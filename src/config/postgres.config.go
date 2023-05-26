package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SSL      string `mapstructure:"ssl"`
}

type PostgresConfig struct {
	Database Postgres `mapstructure:"database"`
}

func LoadDatabaseConfig() (config *Postgres, err error) {
	viper.AddConfigPath("./config")
	viper.SetConfigName("postgres")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while reading the config")
	}

	conf := PostgresConfig{}

	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "error occurs while unmarshal the config")
	}

	return &conf.Database, nil
}
