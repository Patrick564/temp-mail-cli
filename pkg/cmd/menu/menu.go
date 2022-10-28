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

// TODO: Better name for struct to contain EmailValues
type model struct {
	Content  cmdutil.EmailValues
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case cmdutil.EmailValues:
		m.Content = msg
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

func (m model) View() string {
	s := fmt.Sprintf("Email: %s\n\n", m.Content.Email)

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
