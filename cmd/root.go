package cmd

import (
	"fmt"
	"os"

	"github.com/lucas-tremaroli/clist/cmd/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clist",
	Short: "clist is a simple CLI tool",
	Long:  `clist is a simple CLI tool to manage tasks and notes`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
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
	rootCmd.AddCommand(task.AddCmd)
}
