package note

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	help         help.Model
	filename     textinput.Model
	content      textarea.Model
	focusedField int
	quitting     bool
	err          error
	keys         keyMap
	service      NoteService
	config       Config
}

// NewNoteEditor creates a new note editor model
func NewNoteEditor() model {
	config := DefaultConfig()

	filename := textinput.New()
	filename.Placeholder = "note-name"
	filename.Focus()
	filename.Width = config.FilenameWidth
	filename.CharLimit = config.FilenameCharLimit

	content := textarea.New()
	content.Placeholder = "Write your markdown note here..."
	content.SetWidth(config.ContentWidth)
	content.SetHeight(config.ContentHeight)

	return model{
		help:         help.New(),
		filename:     filename,
		content:      content,
		focusedField: FieldFilename,
		keys:         NewKeyMap(),
		service:      NewFileNoteService(),
		config:       config,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			m.quitting = true
			return m, tea.Quit

		case key.Matches(msg, m.keys.Save):
			if m.filename.Value() != "" && m.content.Value() != "" {
				err := m.service.SaveNote(m.filename.Value(), m.content.Value())
				if err != nil {
					m.err = err
				} else {
					return m, tea.Quit
				}
			}
			return m, nil

		case key.Matches(msg, m.keys.Tab):
			return m.switchFocus()
		}
	}

	if m.focusedField == FieldFilename {
		m.filename, cmd = m.filename.Update(msg)
	} else {
		m.content, cmd = m.content.Update(msg)
	}

	return m, cmd
}

// switchFocus switches focus between filename and content fields
func (m model) switchFocus() (tea.Model, tea.Cmd) {
	if m.focusedField == FieldFilename {
		m.filename.Blur()
		m.content.Focus()
		m.focusedField = FieldContent
		return m, textarea.Blink
	} else if m.focusedField == FieldContent {
		m.content.Blur()
		m.filename.Focus()
		m.focusedField = FieldFilename
		return m, textinput.Blink
	}
	return m, nil
}
