package task

import (
	"fmt"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Viewer struct {
	help  help.Model
	task  Task
	board *Board
}

func NewViewer(task Task, board *Board) *Viewer {
	return &Viewer{
		help:  help.New(),
		task:  task,
		board: board,
	}
}

func (v Viewer) Init() tea.Cmd {
	return nil
}

func (v Viewer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Back):
			if v.board != nil {
				return v.board.Update(nil)
			}
			return v, nil
		case key.Matches(msg, keys.Quit):
			return v, tea.Quit
		}
	}
	return v, nil
}

func (v Viewer) View() string {
	statusText := ""
	switch v.task.Status() {
	case todo:
		statusText = "To Do"
	case inProgress:
		statusText = "In Progress"
	case done:
		statusText = "Done"
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)

	statusStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		MarginBottom(1)

	descriptionStyle := lipgloss.NewStyle().
		MarginBottom(2).
		Width(80)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render(v.task.Title()),
		statusStyle.Render(fmt.Sprintf("Status: %s", statusText)),
		descriptionStyle.Render(v.task.Description()),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		content,
		"Press ESC to go back, q/Ctrl+C to quit",
	)
}
