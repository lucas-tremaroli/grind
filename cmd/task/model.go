package task

import tea "github.com/charmbracelet/bubbletea"

type Model struct {
	tasks  []string
	cursor int
	quitting bool
}

func NewModel() Model {
	return Model{
		tasks: []string{},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.tasks)-1 {
				m.cursor++
			}
		case "enter":
			// Handle task selection/editing here
		}
	}
	return m, nil
}