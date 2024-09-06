package shared_configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	Database struct {
		Driver string `yaml:"driver"`
		Mysql  struct {
			Host     string `yaml:"host"`   //, envconfig:"DB_HOST"`
			Port     int    `yaml:"port"`   //, envconfig:"DB_PORT"`
			DBName   string `yaml:"dbname"` //, envconfig:"DB_NAME"`
			Username string `yaml:"dbuser"` //, envconfig:"DB_USER"`
			Password string `yaml:"dbpass"` //, envconfig:"DB_PASS"`
		} `yaml:"mysql"`
	}

	Jwt struct {
		Secret string `yaml: "secret"`
	} `yaml: "jwt"`
}

var cfg *Config

func GetConfig() (*Config, error) {

	if cfg != nil {
		return cfg, nil
	}

	cfg = new(Config)
	conf, err := os.Open("./configs/config.yaml")
	if err != nil {
		return nil, err
	}
	defer conf.Close()

	decoder := yaml.NewDecoder(conf)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
