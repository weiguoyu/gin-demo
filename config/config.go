package config

import (
	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	Log *LogConfig `yaml:"log"`
}

type LogConfig struct {
	Filename   string `yaml:"filename"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"maxsize"` // megabytes
	MaxBackups int    `yaml:"maxbackups"`
	MaxAge     int    `yaml:"maxage"` //days
}

var (
	config *Config
	lock   = new(sync.RWMutex)
)

func (c *Config) LoadConfig(cfg string) error {
	configFile, err := ioutil.ReadFile(cfg)

	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(configFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}

	lock.Lock()
	defer lock.Unlock()

	return nil
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent.")
	}

	lock.Lock()
	defer lock.Unlock()
	config = &Config{
		Log: &LogConfig{},
	}

	configFile, err := ioutil.ReadFile(cfg)

	if err != nil {
		log.Fatalf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func ReadConf() *Config {
	lock.Lock()
	defer lock.Unlock()
	return config
}
