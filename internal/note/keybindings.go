package note

import (
	"github.com/charmbracelet/bubbles/key"
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

// NewKeyMap returns the default key bindings
func NewKeyMap() keyMap {
	return keyMap{
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
}
