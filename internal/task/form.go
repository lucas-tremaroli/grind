package task

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Form struct {
	help        help.Model
	title       textinput.Model
	description textarea.Model
	col         column
	index       int
	board       *Board
}

func NewForm(title, description string, board *Board) *Form {
	form := Form{
		help:        help.New(),
		title:       textinput.New(),
		description: textarea.New(),
		board:       board,
	}

	if title == "" {
		title = "task name"
	}

	form.title.Placeholder = title
	form.description.Placeholder = description
	form.title.SetValue(title)
	form.description.SetValue(description)
	form.title.Focus()
	return &form
}

func (f Form) CreateTask() Task {
	return NewTask(f.col.status, f.title.Value(), f.description.Value())
}

func (f Form) Init() tea.Cmd {
	return nil
}

func (f Form) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case column:
		f.col = msg
		f.col.list.Index()
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, formKeys.Quit):
			return f, tea.Quit
		case key.Matches(msg, formKeys.Back):
			if f.board != nil {
				return f.board.Update(nil)
			}
			return f, nil
		case key.Matches(msg, formKeys.Save):
			if f.board != nil {
				return f.board.Update(f)
			}
			return f, nil
		case key.Matches(msg, formKeys.Help):
			if f.title.Focused() {
				f.title.Blur()
				f.description.Focus()
				return f, textarea.Blink
			} else {
				f.description.Blur()
				f.title.Focus()
				return f, textinput.Blink
			}
		}
	}
	if f.title.Focused() {
		f.title, cmd = f.title.Update(msg)
		return f, cmd
	}
	f.description, cmd = f.description.Update(msg)
	return f, cmd
}

func (f Form) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		"Create a new task",
		f.title.View(),
		f.description.View(),
		f.help.View(formKeys))
}
