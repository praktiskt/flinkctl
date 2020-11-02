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
	"github.com/spf13/cobra"
)

var (
	flagJobStatus  []string
	showJobDetails bool
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:    "jobs",
	Short:  "A brief description of your command",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if showJobDetails {
			re, err := cl.GetJobsOverview()
			if err != nil {
				return err
			}
			Print(re.Jobs)
		} else {
			re, err := cl.GetJobs()
			if err != nil {
				return err
			}
			Print(re.Jobs)
		}
		return nil
	},
}

func init() {
	getCmd.AddCommand(jobsCmd)

	//TODO: Implement filtering on these flags, add all available statuses to default
	jobsCmd.Flags().StringSliceVarP(&flagJobStatus, "status", "s", []string{"RUNNING", "CANCELLED"}, "a comma separated list of statuses of jobs")
	jobsCmd.Flags().BoolVarP(&showJobDetails, "details", "d", false, "show details of all running jobs")
}
