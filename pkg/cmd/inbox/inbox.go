package inbox

import (
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var BaseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	MarginTop(3).
	MarginLeft(20).
	MaxHeight(40).
	MaxWidth(120).
	Align(lipgloss.Center, lipgloss.Center)

type model struct {
	Table table.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.Table.SelectedRow()[1]),
			)
		}
	}
	m.Table, cmd = m.Table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return m.Table.View() + "\n"
}

func New() model {
	columns := []table.Column{
		{Title: "Sender", Width: 35},
		{Title: "Subject", Width: 30},
		{Title: "Open", Width: 8},
	}
	rows := []table.Row{{"Cosme Fulanito | cosme@gmail.com", "Good nigth", "▶"}}
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

	model := model{
		Table: t,
	}

	return model
}
