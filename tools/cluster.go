package tools

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

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
