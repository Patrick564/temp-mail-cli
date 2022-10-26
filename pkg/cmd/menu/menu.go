package menu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BaseStyle = lipgloss.NewStyle().
	MarginTop(3).
	MarginLeft(20).
	Width(45).
	Align(lipgloss.Center, lipgloss.Center).
	BorderStyle(lipgloss.RoundedBorder())

type MenuModel struct {
	Header   string
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (m MenuModel) Init() tea.Cmd { return nil }

func (m MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.Cursor > 0 {
				m.Cursor--
			}
		case tea.KeyDown:
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case tea.KeyEnter:
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m MenuModel) View() string {
	s := fmt.Sprintf("Email: %s\n\n", m.Header)

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = "➤"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}
