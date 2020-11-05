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

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

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
	Use:     "job <submitted filename> [flags]",
	Short:   "submit an uploaded file to the jobmanager (starts the job)",
	Example: "flinkctl submit job 6f52c1fc-e116-4497-9037-68927ae0db6f_DummyApp-1.0-SNAPSHOT.jar --parallelism 2",
	PreRun:  func(cmd *cobra.Command, args []string) { InitCluster() },
	Args:    cobra.ExactValidArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		appb, _ := json.Marshal(ApplicationBody{
			AllowNonRestoredState: StringOrNil(allowNonRestoredState),
			Parallelism:           StringOrNil(parallelism),
			ProgramArgs:           StringOrNil(programArgs),
			SavepointPath:         StringOrNil(savepointPath),
			EntryClass:            StringOrNil(entryClass),
		})
		fileToSubmitURL := fmt.Sprintf("%v/%v/run", cl.Jars.URL.String(), args[0])
		resp, _, _ := tools.ApplyHeadersToRequest(
			gorequest.
				New().
				Post(fileToSubmitURL).
				Type("json").
				Send(string(appb))).
			End()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("failed to start job: %v", string(body))
		}

		re := RunResponse{}
		json.Unmarshal(body, &re)
		Print(re)

		return nil
	},
}

func init() {
	submitCmd.AddCommand(submitJobCmd)
	submitJobCmd.PersistentFlags().StringVar(&allowNonRestoredState, "allowNonRestoredState", "", "Allow non restored state")
	submitJobCmd.PersistentFlags().StringVar(&parallelism, "parallelism", "", "set parallelism for the submitted job")
	submitJobCmd.PersistentFlags().StringVar(&programArgs, "programArgs", "", `a string of program arguments, e.g. "-A=B -C=D"`)
	submitJobCmd.PersistentFlags().StringVar(&savepointPath, "savepointPath", "", "if specified, a save point path")
	submitJobCmd.PersistentFlags().StringVar(&entryClass, "entryClass", "", "the entry class of the submitted jar")
}

func StringOrNil(a string) *string {
	if len(a) == 0 {
		return nil
	}
	return &a
}
