package menu

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

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
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter", " ":
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
