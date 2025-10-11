package task

import (
	"fmt"

	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Long:  `Add a new task to your list`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Added task: %s\n", args[0])
	},
}
