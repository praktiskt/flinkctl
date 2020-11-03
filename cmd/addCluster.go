/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// addClusterCmd represents the addCluster command
var addClusterCmd = &cobra.Command{
	Use:    "add-cluster <url:port>",
	Short:  "A brief description of your command",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("add-cluster requires exactly 1 positional argument, not %v", len(args))
		}

		u, err := url.Parse(args[0])
		if err != nil {
			return err
		}

		currentConfig := tools.GetConfig()
		newConfig := tools.ClusterConfig{URL: u.String()}
		currentConfig.Clusters = append(currentConfig.Clusters, newConfig)
		currentConfig.CurrentCluster = tools.GetCurrentCluster()

		out, err := yaml.Marshal(currentConfig)
		if err != nil {
			return err
		}
		ioutil.WriteFile(viper.ConfigFileUsed(), out, os.ModePerm)

		return nil
	},
}

func init() {
	configCmd.AddCommand(addClusterCmd)
}
