package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	Mongo MongoConfig
}

type MongoConfig struct {
	Username    string `env:"MONGO_USER"`
	Password    string `env:"MONGO_PASS"`
	ClusterName string `env:"MONGO_CLUSTER_NAME"`
}

func (cfg *Config) InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = env.Parse(cfg)
	if err != nil {
		panic(err)
	}

	return cfg
}
