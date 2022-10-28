package root

import (
	"github.com/Patrick564/temp-mail-cli/pkg/cmd/inbox"
	"github.com/Patrick564/temp-mail-cli/pkg/cmd/menu"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var helpStyle = lipgloss.NewStyle().
	MarginTop(5).
	MarginLeft(20).
	Foreground(lipgloss.Color("241"))

const (
	menuView sessionState = iota
	inboxView
)

type sessionState uint

type model struct {
	menu  tea.Model
	inbox tea.Model
	state sessionState
}

func (m model) Init() tea.Cmd {
	return tea.Batch(m.menu.Init(), m.inbox.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == menuView {
				m.state = inboxView
			} else {
				m.state = menuView
			}
		}

		switch m.state {
		case menuView:
			m.menu, cmd = m.menu.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.inbox, cmd = m.inbox.Update(msg)
			cmds = append(cmds, cmd)
		}
	default:
		switch m.state {
		case menuView:
			m.menu, cmd = m.menu.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.inbox, cmd = m.inbox.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var s string

	selected := lipgloss.Color("#7D56F4")
	noSelected := lipgloss.Color("240")

	view := func(state bool) string {
		if state {
			return lipgloss.JoinHorizontal(
				lipgloss.Center,
				menu.BaseStyle.BorderForeground(selected).Render(m.menu.View()),
				inbox.BaseStyle.BorderForeground(noSelected).Render(m.inbox.View()),
			)
		}

		return lipgloss.JoinHorizontal(
			lipgloss.Center,
			menu.BaseStyle.BorderForeground(noSelected).Render(m.menu.View()),
			inbox.BaseStyle.BorderForeground(selected).Render(m.inbox.View()),
		)
	}

	s += view(m.state == menuView) + helpStyle.Render("\ntab: focus next • q: exit\n")

	return s
}

func InitialModel() tea.Model {
	m := model{state: menuView}

	m.menu = menu.New()
	m.inbox = inbox.New()

	return m
}
