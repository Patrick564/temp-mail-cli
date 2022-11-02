package ui

import (
	"fmt"

	"github.com/Patrick564/temp-mail-cli/pkg/cmdutil"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#45b245")).
			MarginTop(3).
			MarginLeft(20)
	baseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#265C7E")).
			MarginTop(2).
			MarginLeft(20).
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
			MarginLeft(21)
)

type model struct {
	Table       table.Model
	EmailParams cmdutil.RandomEmail
	EmailList   [][]string
}

func (m model) Init() tea.Cmd {
	// return cmdutil.LoadEmail
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case cmdutil.RandomEmail:
		m.EmailParams = msg
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "r":
			rows, err := cmdutil.LoadEmailsList(m.EmailParams.Hash)
			if err != nil {
				return m, tea.Quit
			}
			m.Table.SetRows(rows)
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		case "q":
			return m, tea.Quit
		}
	}

	m.Table, cmd = m.Table.Update(msg)

	return m, cmd
}

func (m model) View() string {
	s := titleStyle.Render(fmt.Sprintf("Email: %s\n", m.EmailParams.Email))
	s += baseStyle.Render(m.Table.View())
	s += helpStyle.Render("enter: Open email • n: New email • r: Refresh inbox • q: quit\n")

	return s
}

func InitialModel() tea.Model {
	columns := []table.Column{
		{Title: "Sender", Width: 35},
		{Title: "Subject", Width: 30},
		{Title: "Open", Width: 8},
	}
	rows := []table.Row{
		{"-", "-", "-"},
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
