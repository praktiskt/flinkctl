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
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/spf13/cobra"
)

type Jobs struct {
	Jobs []struct {
		Name  string `json:"name" header:"name"`
		State string `json:"state" header:"state"`
		ID    string `json:"jid" header:"id"`
	} `json:"jobs" header:"jobs"`
}

func GetJobs() (Jobs, error) {
	url := tools.GetCurrentCluster() + jobPath + "/overview"
	re, err := http.Get(url)
	if err != nil {
		return Jobs{}, err
	}
	defer re.Body.Close()

	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		return Jobs{}, err
	}

	j := Jobs{}
	json.Unmarshal(body, &j)
	return j, nil

}

var (
	flagJobStatus []string
)

// jobsCmd represents the jobs command
var jobsCmd = &cobra.Command{
	Use:   "jobs",
	Short: "A brief description of your command",
	RunE: func(cmd *cobra.Command, args []string) error {
		tools.CheckCurrentClusterExists()
		re, err := GetJobs()
		if err != nil {
			return err
		}
		tools.TablePrint(re)
		return nil
	},
}

func init() {
	getCmd.AddCommand(jobsCmd)

	jobsCmd.Flags().StringSliceVarP(&flagJobStatus, "status", "s", []string{"RUNNING", "CANCELLED"}, "a comma separated list of statuses of jobs")
}
