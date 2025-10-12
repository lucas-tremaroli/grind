package task

import (
	"log"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucas-tremaroli/clist/internal/storage"
)

type Board struct {
	help     help.Model
	loaded   bool
	focused  status
	cols     []column
	quitting bool
	db       *storage.DB
}

func NewBoard() *Board {
	help := help.New()
	help.ShowAll = true

	db, err := storage.NewDB()
	if err != nil {
		log.Printf("Failed to initialize database: %v", err)
		return nil
	}

	board := &Board{help: help, focused: todo, db: db}
	board.initLists()
	return board
}

func (m *Board) Init() tea.Cmd {
	return nil
}

func (m *Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmd tea.Cmd
		var cmds []tea.Cmd
		m.help.Width = msg.Width - margin
		for i := 0; i < len(m.cols); i++ {
			var res tea.Model
			res, cmd = m.cols[i].Update(msg)
			m.cols[i] = res.(column)
			cmds = append(cmds, cmd)
		}
		m.loaded = true
		return m, tea.Batch(cmds...)
	case Form:
		task := msg.CreateTask()
		if _, err := m.db.CreateTask(task.Title(), task.Description(), int(task.Status())); err != nil {
			log.Printf("Failed to save task to database: %v", err)
		}
		return m, m.cols[m.focused].Set(msg.index, task)
	case moveMsg:
		if err := m.db.UpdateTask(msg.Task.ID(), msg.Task.Title(), msg.Task.Description(), int(msg.Task.Status())); err != nil {
			log.Printf("Failed to update task in database: %v", err)
		}
		return m, m.cols[m.focused.getNext()].Set(APPEND, msg.Task)
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			m.quitting = true
			if m.db != nil {
				m.db.Close()
			}
			return m, tea.Quit
		case key.Matches(msg, keys.Left):
			m.cols[m.focused].Blur()
			m.focused = m.focused.getPrev()
			m.cols[m.focused].Focus()
		case key.Matches(msg, keys.Right):
			m.cols[m.focused].Blur()
			m.focused = m.focused.getNext()
			m.cols[m.focused].Focus()
		}
	}
	res, cmd := m.cols[m.focused].UpdateWithBoard(msg, m)
	if _, ok := res.(column); ok {
		m.cols[m.focused] = res.(column)
	} else {
		return res, cmd
	}
	return m, cmd
}

func (m *Board) View() string {
	if m.quitting {
		return ""
	}
	if !m.loaded {
		return "loading..."
	}
	board := lipgloss.JoinHorizontal(
		lipgloss.Left,
		m.cols[todo].View(),
		m.cols[inProgress].View(),
		m.cols[done].View(),
	)
	return lipgloss.JoinVertical(lipgloss.Left, board, m.help.View(keys))
}

func (b *Board) initLists() {
	b.cols = []column{
		newColumn(todo),
		newColumn(inProgress),
		newColumn(done),
	}
	b.cols[todo].list.Title = "To Do"
	b.cols[inProgress].list.Title = "In Progress"
	b.cols[done].list.Title = "Done"

	b.loadTasksFromDB()
}

func (b *Board) loadTasksFromDB() {
	if b.db == nil {
		b.loadDefaultTasks()
		return
	}

	taskRecords, err := b.db.GetAllTasks()
	if err != nil {
		log.Printf("Failed to load tasks from database: %v", err)
		b.loadDefaultTasks()
		return
	}

	var todoItems, inProgressItems, doneItems []list.Item

	for _, record := range taskRecords {
		task := NewTaskWithID(record.ID, status(record.Status), record.Title, record.Description)
		switch status(record.Status) {
		case todo:
			todoItems = append(todoItems, task)
		case inProgress:
			inProgressItems = append(inProgressItems, task)
		case done:
			doneItems = append(doneItems, task)
		}
	}

	b.cols[todo].list.SetItems(todoItems)
	b.cols[inProgress].list.SetItems(inProgressItems)
	b.cols[done].list.SetItems(doneItems)
}

func (b *Board) loadDefaultTasks() {
	b.cols[todo].list.SetItems([]list.Item{
		NewTask(todo, "buy milk", "strawberry milk"),
		NewTask(todo, "eat sushi", "negitoro roll, miso soup, rice"),
		NewTask(todo, "fold laundry", "or wear wrinkly t-shirts"),
	})
	b.cols[inProgress].list.SetItems([]list.Item{
		NewTask(inProgress, "write code", "don't worry, it's Go"),
	})
	b.cols[done].list.SetItems([]list.Item{
		NewTask(done, "stay cool", "as a cucumber"),
	})
}
