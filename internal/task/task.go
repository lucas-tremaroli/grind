package task

import "github.com/google/uuid"

type Task struct {
	id          string `json:"id"`
	status      status `json:"status"`
	title       string `json:"title"`
	description string `json:"description"`
}

func NewTask(status status, title, description string) Task {
	return Task{
		id:          uuid.New().String(),
		status:      status,
		title:       title,
		description: description,
	}
}

func NewTaskWithID(id string, status status, title, description string) Task {
	return Task{
		id:          id,
		status:      status,
		title:       title,
		description: description,
	}
}

func (t *Task) Next() {
	if t.status == done {
		t.status = todo
	} else {
		t.status++
	}
}

func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

func (t Task) ID() string {
	return t.id
}

func (t Task) Status() status {
	return t.status
}

type status int

func (s status) getNext() status {
	if s == done {
		return todo
	}
	return s + 1
}

func (s status) getPrev() status {
	if s == todo {
		return done
	}
	return s - 1
}

const (
	todo status = iota
	inProgress
	done
)
