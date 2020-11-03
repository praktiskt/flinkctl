package config

//TODO: Refactor into more relevant files

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type FlinkctlConfig struct {
	Clusters       []ClusterConfig `yaml:"clusters"`
	CurrentCluster string          `yaml:"current-cluster"`
}

type ClusterConfig struct {
	URL     string   `yaml:"url"`
	Headers []string `yaml:"headers"`
}

func Get() *FlinkctlConfig {
	conf := &FlinkctlConfig{}
	if err := viper.Unmarshal(conf); err != nil {
		fmt.Printf("Could not parse config: %v", err)
		return conf
	}
	return conf
}

func GetCurrent() (*ClusterConfig, error) {
	globalConfig := Get()
	currentCluster := GetCurrentName()
	for _, conf := range globalConfig.Clusters {
		if conf.URL == currentCluster {
			return &conf, nil
		}
	}
	return &ClusterConfig{}, fmt.Errorf("no such cluster: %v", currentCluster)
}

func ConfigExists(url string) bool {
	fullConf := Get()
	for _, conf := range fullConf.Clusters {
		if conf.URL == url {
			return true
		}
	}
	return false
}

func GetCurrentName() string {
	return viper.GetString("current-cluster")
}

func CurrentExists() bool {
	return GetCurrentName() == ""
}

func CheckCurrentExists() {
	if !CurrentExists() {
		fmt.Println("No current cluster context set")
		os.Exit(1)
	}
}

func GetHeaders() []string {
	current, _ := GetCurrent()
	return current.Headers
}
