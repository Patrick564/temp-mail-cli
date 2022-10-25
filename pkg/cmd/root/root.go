package root

import (
	"github.com/Patrick564/temp-mail-cli/pkg/cmd/inbox"
	"github.com/Patrick564/temp-mail-cli/pkg/cmd/menu"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	menuStyle = lipgloss.NewStyle().
			Margin(2, 3).
			Width(45).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground()
	inboxStyle = lipgloss.NewStyle().
			Margin(2, 3).
			MaxHeight(40).
			MaxWidth(96).
			Align(lipgloss.Center, lipgloss.Center)
	selectedStyle = lipgloss.NewStyle().
			MarginTop(3).
			MarginLeft(20).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#7D56F4"))
	noSelectedStyle = lipgloss.NewStyle().
			MarginTop(3).
			MarginLeft(20).
			BorderStyle(lipgloss.HiddenBorder())
	helpStyle = lipgloss.NewStyle().
			MarginTop(5).
			MarginLeft(20).
			Foreground(lipgloss.Color("241"))
)

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
	var sm string

	menuRender := menuStyle.Render(m.menu.View())
	inboxRender := inboxStyle.Render(m.inbox.View())

	if m.state == menuView {
		sm += lipgloss.JoinHorizontal(
			lipgloss.Center,
			selectedStyle.Render(menuRender),
			noSelectedStyle.Render(inboxRender),
		)
	} else {
		sm += lipgloss.JoinHorizontal(
			lipgloss.Center,
			noSelectedStyle.Render(menuRender),
			selectedStyle.Render(inboxRender),
		)
	}
	sm += helpStyle.Render("\ntab: focus next • q: exit\n")

	return sm
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
