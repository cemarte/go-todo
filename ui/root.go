package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	// "github.com/charmbracelet/lipgloss"
)

type TodoItem struct {
	text string
}

func (todoItem TodoItem) Text() string {
	return todoItem.text
}

func (todoItem *TodoItem) SetText(t string) {
	todoItem.text = t
}

func (todoItem TodoItem) FilterValue() string {
	return todoItem.text
}

type Model struct {
	TodoList list.Model
}

func (m Model) Init() tea.Cmd {
	return nil

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}

	}
	var cmd tea.Cmd
	m.TodoList, cmd = m.TodoList.Update(msg)
	return m, cmd

}

func (m Model) View() string {

	return m.TodoList.View()
}
