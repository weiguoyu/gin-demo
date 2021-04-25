// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config

import (
	"github.com/toolkits/file"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type Config struct {
	Common *Common `yaml:"common"`
	Api    *Api    `yaml:"api"`
}

type Common struct {
	Gin struct {
		Mode string `yaml:"mode"`
	}
	Region struct {
		RegionId string `yaml:"region_id"`
	}
	Log struct {
		Filepath   string `yaml:"filepath"`
		Level      string `yaml:"level"`
		MaxSize    int    `yaml:"maxsize"` // megabytes
		MaxBackups int    `yaml:"maxbackups"`
		MaxAge     int    `yaml:"maxage"` //days
	}

	mysql struct {
		Address  string `yaml:"address"`
		Dbname   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}

type Api struct {
	Port string `yaml:"port"`
}

var (
	config *Config
	lock   = new(sync.RWMutex)
)

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent.")
	}

	lock.Lock()
	defer lock.Unlock()
	config = &Config{}

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
