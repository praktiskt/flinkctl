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

// submitJobCmd represents the submitJob command
var submitJobCmd = &cobra.Command{
	Use:    "job <path to jar> [flags]",
	Short:  "Submit a packaged Flink job to your cluster.",
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("must specify at least one job to submit")
		}

		u := cl.Jars.UploadURL.String()
		for _, file := range args {
			resp, _, _ := gorequest.
				New().
				Post(u).
				Type("multipart").
				SendFile(file).
				End()

			defer resp.Body.Close()

			fmt.Println("response Status:", resp.Status)
			body, _ := ioutil.ReadAll(resp.Body)
			s := SubmitResponse{}
			json.Unmarshal(body, &s)
			fmt.Println("response Body:", string(body))

			fmt.Println("-- STARTING JOB --")
			fn := strings.Split(s.Filename, "/")
			f := fn[len(fn)-1]
			fmt.Println(f)
			appb, _ := json.Marshal(ApplicationBody{
				//TODO: Pass flags to each atrribute
				AllowNonRestoredState: nil,
				Parallelism:           nil,
				ProgramArgs:           nil,
				SavepointPath:         nil,
				EntryClass:            nil,
			})
			resp, _, _ = gorequest.New().
				Post(cl.Jars.URL.String() + "/" + f + "/run"). //TODO: Pass entry class if exists: ?entry-class=dummyApp.StreamingJob"
				Type("json").
				Send(string(appb)).
				End()
			fmt.Println("response Status:", resp.Status)
			body, _ = ioutil.ReadAll(resp.Body)
			fmt.Println("response Body:", string(body))

		}

		return nil
	},
}

func init() {
	submitCmd.AddCommand(submitJobCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// submitJobCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// submitJobCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
