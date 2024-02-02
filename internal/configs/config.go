package configs

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"

	"REST/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug"`
	Listen  `yaml:"listen"`
	MongoDB `yaml:"mongoDB"`
}

type Listen struct {
	Type   string `yaml:"type"`
	BindIp string `yaml:"bind_ip"`
	Port   string `yaml:"port"`
}

type MongoDB struct {
	Host       string `json:"host"`
	Port       string `json:"port"`
	Database   string `json:"database"`
	AuthDB     string `json:"auth_db"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Collection string `json:"collection"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config { // singleton
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")

		instance = &Config{} // зачем здесь опять присвоивать
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatalf("error:  %v", err)
		}

	})
	return instance

}
