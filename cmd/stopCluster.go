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

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

var iKnowWhatImDoing bool

// stopClusterCmd represents the stopCluster command
var stopClusterCmd = &cobra.Command{
	Use:    "cluster",
	Short:  "shut down the cluster",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if !iKnowWhatImDoing {
			return fmt.Errorf("you don't know what you're doing")
		}
		resp, _, _ := tools.ApplyHeadersToRequest(gorequest.
			New().
			Delete(cl.ClusterURL.String())).
			End()
		//TODO: Error management
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		msg := string(body)
		Print(msg)
		return nil
	},
}

func init() {
	stopCmd.AddCommand(stopClusterCmd)
	stopClusterCmd.Flags().BoolVar(&iKnowWhatImDoing, "i-know-what-im-doing", false, "you need to pass this flag for the call to work :)")
}
