package config

import "github.com/jinzhu/configor"

type Config struct {
	AppName string `default:"todos"`
	Port    string `default:"8282"`
	Logger  struct {
		Use         string `default:"zapLogger"`
		Environment string `default:"dev"`
		LogLevel    string `default:"info"`
		FileName    string `default:"main.log"`
	}
	DB struct {
		Use      string `default:"mysql"`
		Mysql struct {
			UserName string `default:"root"`
			Password string `default:"password"`
			Database string `default:"todos"`
		}
	}
}

func NewConfig() (*Config, error) {
	c := &Config{}
	err := configor.Load(c, "pkg/config/config.yml")
	if err != nil {
		return nil, err
	}
	return c, nil
}