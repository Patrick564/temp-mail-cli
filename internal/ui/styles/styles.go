package styles

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.Color("#265C7E")).
			Foreground(lipgloss.Color("#45b245")).
			MarginTop(3).
			MarginLeft(6)
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#265C7E")).
			MarginTop(1).
			MarginLeft(5).
			Padding(1, 2).
			MaxHeight(40).
			MaxWidth(120).
			Align(lipgloss.Center, lipgloss.Center)
	HeaderStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240")).
			BorderBottom(true).
			Padding(0, 1).
			Bold(true)
	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("229")).
			Background(lipgloss.Color("#265C7E")).
			Bold(true)
	CellStyle = lipgloss.NewStyle().
			Padding(0, 1)
	ViewportStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			PaddingRight(2).
			MarginTop(1)
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(2).
			MarginLeft(5).
			MarginBottom(2).
			PaddingLeft(1)
)
