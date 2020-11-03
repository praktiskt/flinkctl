package config

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

func GetConfig() *FlinkctlConfig {
	conf := &FlinkctlConfig{}
	if err := viper.ReadInConfig(); err != nil {
		viper.SetDefault("current-cluster", "")
		viper.SafeWriteConfig()
	}
	if err := viper.Unmarshal(conf); err != nil {
		fmt.Printf("Could not parse config: %v", err)
		return conf
	}
	fmt.Println(conf.Clusters)
	return conf
}

func GetCurrentConfig() (*ClusterConfig, error) {
	globalConfig := GetConfig()
	currentCluster := GetCurrentCluster()
	for _, conf := range globalConfig.Clusters {
		if conf.URL == currentCluster {
			return &conf, nil
		}
	}
	return &ClusterConfig{}, fmt.Errorf("no such cluster: %v", currentCluster)
}

func GetCurrentCluster() string {
	return viper.GetString("current-cluster")
}

func CurrentClusterExists() bool {
	return GetCurrentCluster() == ""
}

func CheckCurrentClusterExists() {
	if !CurrentClusterExists() {
		fmt.Println("No current cluster context set")
		os.Exit(1)
	}
}
