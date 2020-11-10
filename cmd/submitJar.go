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
	"os"
	"strings"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type SubmitResponse struct {
	Filename string `json:"filename" header:"filename"`
	Status   string `json:"status" header:"status"`
}

var jarFilePath string

// submitJarCmd represents the submitJar command
var submitJarCmd = &cobra.Command{
	Use:    "jar -f <filename>",
	Short:  "submit a jar file to your Flink cluster",
	Args:   cobra.ExactArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat(jarFilePath); os.IsNotExist(err) || jarFilePath == "" {
			return fmt.Errorf("selected fat jar does not exist")
		}

		u := cl.Jars.UploadURL.String()
		s := SubmitResponse{}
		resp, body, _ := tools.ApplyHeadersToRequest(
			gorequest.
				New().
				Post(u).
				Type("multipart").
				SendFile(jarFilePath)).
			EndStruct(&s)

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to submit jarfile: %v", string(body))
		}
		s.Filename = ExtractSubmittedFilename(s.Filename)
		Print(s)
		return nil
	},
}

func init() {
	submitCmd.AddCommand(submitJarCmd)
	submitJarCmd.Flags().StringVarP(&jarFilePath, "filename", "f", "", "the jar file to submit to the cluster")
}

func ExtractSubmittedFilename(filename string) string {
	fn := strings.Split(filename, "/")
	return fn[len(fn)-1]
}
