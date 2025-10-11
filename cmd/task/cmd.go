package task

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var TaskCmd = &cobra.Command{
	Use:   "task",
	Short: "Opens the task management tool",
	Long:  `Opens the task management tool`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(NewModel())
		if _, err := p.Run(); err != nil {
			fmt.Printf("Error running program: %v", err)
			os.Exit(1)
		}
	},
}
