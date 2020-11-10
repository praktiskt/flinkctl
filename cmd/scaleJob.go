/*
Copyright © 2020 Magnus Furugård <magnus.furugard@gmail.com>

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

	"github.com/magnusfurugard/flinkctl/tools"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
)

var newParallelism int

// scaleJobCmd represents the scaleJob command
var scaleJobCmd = &cobra.Command{
	Use:    "job <job id> --parallelism <N>",
	Short:  "rescale a currently running job",
	Long:   `AS OF 2019-04 THIS ACTION IS DISABLED IN FLINK`,
	Args:   cobra.ExactValidArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) { InitCluster() },
	RunE: func(cmd *cobra.Command, args []string) error {
		jid := args[0]
		if len(jid) != 32 {
			return fmt.Errorf("`%v` is not a valid job id", jid)
		}

		if newParallelism < 0 {
			return fmt.Errorf("parallelism must be 0 or greater")
		}

		scaleURL := fmt.Sprintf("%v/%v/rescaling?parallelism=%v", cl.Jobs.URL.String(), jid, newParallelism)
		resp, body, _ := tools.ApplyHeadersToRequest(gorequest.New().Patch(scaleURL)).End()
		fmt.Println(body)
		if resp.StatusCode == 200 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}

		return nil
	},
}

func init() {
	scaleCmd.AddCommand(scaleJobCmd)
	scaleJobCmd.Flags().IntVar(&newParallelism, "parallelism", -1, "the new parallelism for your job")
}
