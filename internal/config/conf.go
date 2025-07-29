package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpsServer struct {
	Add string `yaml:"address"`
}

type Config struct {
	Env          string `yaml:"env" env:"ENV" env-required:"true"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HttpsServer  `yaml:"http_server"`
}

func ConfigLoad() *Config {
	var conpath string

	conpath = os.Getenv("CONFIG_PATH")

	if conpath == "" {
		conflag := flag.String("config", "", "path to config path")
		flag.Parse()
		conpath = *conflag

		if conpath == "" {
			log.Fatal("config path is not set in even in flags")
		}

	}

	_, err := os.Stat(conpath)

	if os.IsNotExist(err) {
		log.Fatalf("config file not exist %s", conpath)
	}

	var cfg Config

	err = cleanenv.ReadConfig(conpath, &cfg)

	if err != nil {
		log.Fatalf("cannot read cofig file %s", err.Error())
	}
	return &cfg
}
