package note

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	statusMessageStyle = lipgloss.NewStyle().
		Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
		Render
)

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
	s.WriteString("\n" + m.help.ShortHelpView(m.keys.ShortHelp()))

	return s.String()
}
