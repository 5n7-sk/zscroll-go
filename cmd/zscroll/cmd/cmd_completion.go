package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var shell string

func runCompletion(cmd *cobra.Command, args []string) error {
	switch shell {
	case "bash":
		return rootCmd.GenBashCompletion(os.Stdout)
	case "fish":
		return rootCmd.GenFishCompletion(os.Stdout, true)
	case "powershell":
		return rootCmd.GenPowerShellCompletion(os.Stdout)
	case "zsh":
		return rootCmd.GenZshCompletion(os.Stdout)
	}
	return fmt.Errorf("invalid shell: %s", shell)
}

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Output shell completion (bash/fish/powershell/zsh)",
	Long:  "Output shell completion (bash/fish/powershell/zsh).",
	RunE:  runCompletion,
}

func init() { //nolint:gochecknoinits
	completionCmd.Flags().StringVarP(&shell, "shell", "s", "bash", "shell type (bash/fish/powershell/zsh)")

	rootCmd.AddCommand(completionCmd)
}
