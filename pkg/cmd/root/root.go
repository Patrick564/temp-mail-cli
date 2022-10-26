package root

import (
	"github.com/Patrick564/temp-mail-cli/pkg/cmd/inbox"
	"github.com/Patrick564/temp-mail-cli/pkg/cmd/menu"
	"github.com/charmbracelet/bubbles/table"
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

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
			m.menu, _ = m.menu.Update(msg)
		default:
			m.inbox, _ = m.inbox.Update(msg)
		}
	}

	return m, nil
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
	columns := []table.Column{
		{Title: "Sender", Width: 35},
		{Title: "Subject", Width: 30},
		{Title: "Open", Width: 8},
	}
	rows := []table.Row{
		{"Cosme Fulanito | cosme@gmail.com", "Good nigth", "▶"},
		{"Atlassian Money | atlassian@contact.com", "Register at Atlassian", "▶"},
		{"Octocat Register | register@github.com", "Hello from GitHub", "->"},
		{"JetBrains Comercial | comercial@jetbrains.com", "JetBrains Fleet!", "->"},
		{"JetBrains Register | register@jetbrains.com", "Save 30% at JetBrains products", "->"},
		{"Spammy Spammer | veryspammer1998@spammy.com", "Just a spam an very large email...", "->"},
		{"Birds Food Seller | food@birds.com", "Buy some birds food!", "->"},
	}
	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(15),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{state: menuView}
	m.menu = menu.MenuModel{
		Header:   "nojog41234@abudat.com",
		Choices:  []string{"Get 10 min. more", "Refresh inbox   ", "New temp. email  "},
		Selected: make(map[int]struct{}),
	}
	m.inbox = inbox.InboxModel{
		Table: t,
	}

	return m
}
