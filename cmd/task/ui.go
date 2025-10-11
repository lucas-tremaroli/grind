package task

import "fmt"

func (m Model) View() string {
	s := "Task Management\n\n"

	if len(m.tasks) == 0 {
		s += "No tasks yet. Press 'q' to quit.\n"
	} else {
		for i, task := range m.tasks {
			cursor := " "
			if m.cursor == i {
				cursor = ">"
			}
			s += fmt.Sprintf("%s %s\n", cursor, task)
		}
	}

	s += "\nPress q to quit.\n"
	return s
}
