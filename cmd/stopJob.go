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

	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

// stopJobCmd represents the stopJob command
var stopJobCmd = &cobra.Command{
	Use:    "job <job id>",
	Short:  "Stop a currently running job",
	Args:   cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		jid := args[0]
		if len(jid) != 32 {
			return fmt.Errorf("`%v` is not a valid job id", jid)
		}

		//TODO: Currently uses yarn-cancel as opposed to just /stop (which doesn't seem to work)
		stopURL := fmt.Sprintf("%v/%v/yarn-cancel", cl.Jobs.URL.String(), jid)
		resp, body, _ := gorequest.
			New().
			Get(stopURL).
			End()

		if resp.StatusCode == 202 {
			fmt.Printf("Successfully cancelled job %v\n", jid)
		} else {
			fmt.Println("Failed to cancelled job: " + body)
		}

		return nil
	},
}

func init() {
	stopCmd.AddCommand(stopJobCmd)
}
