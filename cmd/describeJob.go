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

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/spf13/cobra"
)

var describeJobCmd = &cobra.Command{
	Use:    "job <jid>",
	Short:  "Describe a job in your cluster.",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("bad args, got: %v", args)
		}
		for _, jid := range args {
			if len(jid) != 32 {
				return fmt.Errorf("%v is not a valid jid", jid)
			}
			re, err := cl.DescribeJob(jid)
			if err != nil {
				return err
			}
			tools.JSONPrint(re)
		}
		return nil
	},
}

func init() {
	describeCmd.AddCommand(describeJobCmd)
}
