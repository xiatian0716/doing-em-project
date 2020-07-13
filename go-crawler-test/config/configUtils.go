package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

//Config   系统配置配置
type Configs struct {
	Version  string `yaml:"Version"`
	StartUrl string `yaml:"StartUrl"`

	CEConfig      CE `yaml:"ConcurrentEngine"`
	ESConfig      ES `yaml:"Staticsearch"`
	FetcherConfig FC `yaml:"Fetcher"`
}

// ConcurrentEngine 配置
type CE struct {
	WorkerCount int  `yaml:"WorkerCount"`
	ESSave      bool `yaml:"ESSave"`
}

// ElstaticSearch 配置
type ES struct {
	SetURL   string `yaml:"ESConnection"`
	Index    string `yaml:"Database"`
	Type     string `yaml:"Table"`
	SetSniff bool   `yaml:"SetSniff"`
}

// Fetcher 配置
type FC struct {
}

func ConfigsParse(FilePath string) Configs {
	var configs Configs
	//config, err := ioutil.ReadFile("./config/config.yaml")
	config, err := ioutil.ReadFile(FilePath)
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &configs)

	return configs
}
