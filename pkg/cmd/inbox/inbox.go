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

type InboxModel struct {
	Table table.Model
}

func (m InboxModel) Init() tea.Cmd { return nil }

func (m InboxModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m InboxModel) View() string {
	return m.Table.View() + "\n"
}
