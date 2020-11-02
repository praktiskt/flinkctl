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

type JarsResponse struct {
	Address string `json:"address"`
	Files   []struct {
		Entry []struct {
			Description string `json:"description" header:"entry-description"`
			Name        string `json:"name" header:"entry-name"`
		} `json:"entry"`
		ID       string `json:"id" header:"id"`
		Name     string `json:"name" header:"name"`
		Uploaded int64  `json:"uploaded" header:"uploaded"`
	} `json:"files"`
}

var getJarsCmd = &cobra.Command{
	Use:    "jars",
	Short:  "List all uploaded jars in your cluster",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		resp, err := http.Get(cl.Jars.URL.String())
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		s := JarsResponse{}
		err = json.Unmarshal(body, &s)
		if err != nil {
			return err
		}
		tools.TablePrint(s.Files)
		return nil
	},
}

func init() {
	getCmd.AddCommand(getJarsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getJarsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getJarsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
