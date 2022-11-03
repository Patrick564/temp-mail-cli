package ui

import (
	"fmt"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/Patrick564/temp-mail-cli/pkg/cmdutil"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.Color("#265C7E")).
			Foreground(lipgloss.Color("#45b245")).
			MarginTop(3).
			MarginLeft(6)
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#265C7E")).
			MarginTop(1).
			MarginLeft(5).
			Padding(1, 2).
			MaxHeight(40).
			MaxWidth(120).
			Align(lipgloss.Center, lipgloss.Center)
	headerStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Padding(0, 1).
			Bold(true)
	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("#265C7E")).
			Bold(true)
	cellStyle = lipgloss.NewStyle().
			Padding(0, 1)
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(2).
			MarginLeft(5).
			MarginBottom(2).
			PaddingLeft(1)
)

type model struct {
	User         cmdutil.UserEmail
	Inbox        api.Emails
	SelectedMail api.EmailContent
	Rows         []table.Row
	Table        table.Model
}

func (m model) Init() tea.Cmd {
	// return cmdutil.LoadEmail
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case cmdutil.UserEmail:
		m.User = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "enter":
			m.Inbox = api.Emails{{MailID: "a", MailFrom: "a", MailText: "a", MailSubject: "a", MailTimestamp: 12.3123}}
			if len(m.Inbox) == 0 {
				break
			}

			m.SelectedMail = m.Inbox[m.Table.Cursor()]
		case "r":
			list, err := cmdutil.LoadEmailsList(m.User.Hash)
			if err != nil {
				return m, tea.Quit
			}

			m.Inbox = list.Content
			m.Table.SetRows(list.Rows)
		case "q":
			return m, tea.Quit
		}
	}

	m.Table, cmd = m.Table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	s := titleStyle.Render(fmt.Sprintf("Email: %s", m.User.Email))
	s += baseStyle.Render(m.Table.View()) + fmt.Sprintf("From :%s", m.SelectedMail.MailFrom)
	s += helpStyle.Render("enter: Open mail • n: New temp. email • r: Refresh inbox • q: quit\n")

	return s
}

func New() tea.Model {
	columns := []table.Column{
		{Title: "#", Width: 3},
		{Title: "Sender", Width: 35},
		{Title: "Subject", Width: 30},
		{Title: "Open", Width: 8},
	}
	rows := []table.Row{
		{"0", "-", "-", "-"},
	}

	newTable := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	newTable.SetStyles(table.Styles{
		Header:   headerStyle,
		Selected: selectedStyle,
		Cell:     cellStyle,
	})

	return model{Table: newTable}
}
