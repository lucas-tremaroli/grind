package note

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	statusMessageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
		Render
)

type keyMap struct {
	Save key.Binding
	Quit key.Binding
	Tab  key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Save, k.Tab, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Save, k.Tab, k.Quit},
	}
}

var keys = keyMap{
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c"),
		key.WithHelp("ctrl+c", "quit"),
	),
	Tab: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch field"),
	),
}

type model struct {
	help         help.Model
	filename     textinput.Model
	content      textarea.Model
	focusedField int // 0 for filename, 1 for content
	quitting     bool
	err          error
}

func NewNoteEditor() model {
	filename := textinput.New()
	filename.Placeholder = "note-name"
	filename.Focus()
	filename.Width = 50

	content := textarea.New()
	content.Placeholder = "Write your markdown note here..."
	content.SetWidth(80)
	content.SetHeight(20)

	m := model{
		help:         help.New(),
		filename:     filename,
		content:      content,
		focusedField: 0,
	}
	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			return m, tea.Quit


		case key.Matches(msg, keys.Save):
			if m.filename.Value() != "" && m.content.Value() != "" {
				err := m.saveFile()
				if err != nil {
					m.err = err
				} else {
					return m, tea.Quit
				}
			}
			return m, nil

		case key.Matches(msg, keys.Tab):
			if m.focusedField == 0 {
				m.filename.Blur()
				m.content.Focus()
				m.focusedField = 1
				return m, textarea.Blink
			} else if m.focusedField == 1 {
				m.content.Blur()
				m.filename.Focus()
				m.focusedField = 0
				return m, textinput.Blink
			}
			return m, nil
		}
	}

	if m.focusedField == 0 {
		m.filename, cmd = m.filename.Update(msg)
	} else {
		m.content, cmd = m.content.Update(msg)
	}

	return m, cmd
}

func (m model) saveFile() error {
	filename := m.filename.Value()
	if !strings.HasSuffix(filename, ".md") {
		filename += ".md"
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(cwd, filename)
	return os.WriteFile(filePath, []byte(m.content.Value()), 0644)
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var s strings.Builder
	
	// Filename input
	s.WriteString(m.filename.View())
	if m.filename.Value() != "" && !strings.HasSuffix(m.filename.Value(), ".md") {
		s.WriteString(".md")
	}
	s.WriteString("\n\n")
	
	// Content textarea  
	s.WriteString(m.content.View())
	s.WriteString("\n")
	
	// Error if any
	if m.err != nil {
		s.WriteString("Error: " + m.err.Error() + "\n")
	}

	// Help - always show short help
	s.WriteString("\n" + m.help.ShortHelpView(keys.ShortHelp()))

	return s.String()
}
