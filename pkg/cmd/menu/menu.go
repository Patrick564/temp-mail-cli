package menu

import (
	"fmt"

	"github.com/Patrick564/temp-mail-cli/pkg/cmdutil"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BaseStyle = lipgloss.NewStyle().
	MarginTop(3).
	MarginLeft(20).
	Width(45).
	Align(lipgloss.Center, lipgloss.Center).
	BorderStyle(lipgloss.RoundedBorder())

type model struct {
	Email    string
	Hash     string
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmdutil.RandomEmail:
		m.Email = msg.Email
		m.Hash = msg.Hash
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
			if _, ok := m.Selected[m.Cursor]; ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := fmt.Sprintf("Email: %s\n\n", m.Email)

	for i, choice := range m.Choices {
		cursor := " "
		if m.Cursor == i {
			cursor = "➤"
		}

		s += fmt.Sprintf("%s %s\n", cursor, choice)
	}

	return s
}

func New() model {
	return model{
		Choices: []string{
			"Get 10 min. more",
			"Refresh inbox   ",
			"New temp. email  ",
		},
		Selected: make(map[int]struct{}),
	}
}
