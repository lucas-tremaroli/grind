package cmd

import (
	"fmt"
	"os"

	"github.com/lucas-tremaroli/grind/cmd/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grind",
	Short: "grind is a simple CLI tool",
	Long:  `grind is a simple CLI tool to manage tasks and notes`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{
		HiddenDefaultCmd: true,
	},
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(task.TaskCmd)
}
