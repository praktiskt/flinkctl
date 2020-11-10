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
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "generate-autocomplete [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

$ source <(flinkctl generate-autocomplete bash)

# To load completions for each session, execute once:
Linux:
  $ flinkctl generate-autocomplete bash > /etc/bash_completion.d/flinkctl
MacOS:
  $ flinkctl generate-autocomplete bash > /usr/local/etc/bash_completion.d/flinkctl

Zsh:

# If shell completion is not already enabled in your environment you will need
# to enable it.  You can execute the following once:

$ echo "autoload -U compinit; compinit" >> ~/.zshrc

# To load completions for each session, execute once:
$ flinkctl generate-autocomplete zsh > "${fpath[1]}/_flinkctl"

# You will need to start a new shell for this setup to take effect.

Fish:

$ flinkctl generate-autocomplete fish | source

# To load completions for each session, execute once:
$ flinkctl generate-autocomplete fish > ~/.config/fish/completions/flinkctl.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
