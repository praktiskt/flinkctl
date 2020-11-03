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
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

type SubmitResponse struct {
	Filename string
	Status   string
}

type ApplicationBody struct {
	AllowNonRestoredState *string `json:"allowNonRestoredState"`
	Parallelism           *string `json:"parallelism"`
	ProgramArgs           *string `json:"programArgs"`
	SavepointPath         *string `json:"savepointPath"`
	EntryClass            *string `json:"entryClass"`
}

type RunResponse struct {
	JobID string `json:"jobid" header:"jobid"`
}

var (
	allowNonRestoredState string
	parallelism           string
	programArgs           string
	savepointPath         string
	entryClass            string
)

// submitJobCmd represents the submitJob command
var submitJobCmd = &cobra.Command{
	Use:     "submit-job <path to jar> [flags]",
	Short:   "Submit a packaged Flink job to your cluster.",
	Example: "flinkctl submit job ~/path/to/flinkjob.jar",
	PreRun:  func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("must specify at least one job to submit")
		}

		u := cl.Jars.UploadURL.String()
		for _, file := range args {

			resp, _, _ := tools.ApplyHeadersToRequest(
				gorequest.
					New().
					Post(u).
					Type("multipart").
					SendFile(file)).
				End()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			if resp.StatusCode != 200 {
				return fmt.Errorf("failed to submit jarfile: %v", string(body))
			}

			fileToPass := fmt.Sprintf("%v/%v/run", cl.Jars.URL.String(), ExtractSubmittedFilename(body))
			appb, _ := json.Marshal(ApplicationBody{
				AllowNonRestoredState: StringOrNil(allowNonRestoredState),
				Parallelism:           StringOrNil(parallelism),
				ProgramArgs:           StringOrNil(programArgs),
				SavepointPath:         StringOrNil(savepointPath),
				EntryClass:            StringOrNil(entryClass),
			})

			resp, _, _ = tools.ApplyHeadersToRequest(
				gorequest.
					New().
					Post(fileToPass).
					Type("json").
					Send(string(appb))).
				End()

			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			if resp.StatusCode != 200 {
				return fmt.Errorf("failed to start job: %v", string(body))
			}

			re := RunResponse{}
			json.Unmarshal(body, &re)
			Print(re)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(submitJobCmd)

	submitJobCmd.Flags().StringVar(&allowNonRestoredState, "allowNonRestoredState", "", "Allow non restored state")
	submitJobCmd.Flags().StringVar(&parallelism, "parallelism", "", "set parallelism for the submitted job")
	submitJobCmd.Flags().StringVar(&programArgs, "programArgs", "", `a string of program arguments, e.g. "-A=B -C=D"`)
	submitJobCmd.Flags().StringVar(&savepointPath, "savepointPath", "", "if specified, a save point path")
	submitJobCmd.Flags().StringVar(&entryClass, "entryClass", "", "the entry class of the submitted jar")
}

func StringOrNil(a string) *string {
	if len(a) == 0 {
		return nil
	}
	return &a
}

func ExtractSubmittedFilename(responseBody []byte) string {
	s := SubmitResponse{}
	json.Unmarshal(responseBody, &s)
	fn := strings.Split(s.Filename, "/")
	return fn[len(fn)-1]
}
