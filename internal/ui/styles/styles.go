package styles

import "github.com/charmbracelet/lipgloss"

const (
	HelpText = `
enter: Open mail • n: New temp. email • r: Refresh inbox • q: quit
🡱/🡳: Navigate • 🡰/🡲: Change tab
`
)

var (
	TitleStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.Color("#265C7E")).
			Foreground(lipgloss.Color("#45b245")).
			MarginTop(3).
			MarginLeft(6)

	TableBaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			MarginTop(1).
			MarginLeft(5).
			Padding(1, 2).
			MaxHeight(40).
			MaxWidth(120).
			Align(lipgloss.Center, lipgloss.Center)
	TableHeaderStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				BorderBottom(true).
				Padding(0, 1).
				Bold(true)
	TableCellSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("229")).
				Background(lipgloss.Color("#265C7E")).
				Bold(true)
	TableCellStyle = lipgloss.NewStyle().
			Padding(0, 1)

	ViewportBaseStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#383838")).
				PaddingRight(2).
				MarginTop(1)

	TableFocusedStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("62")).
				MarginTop(1).
				MarginLeft(5).
				Padding(1, 2).
				Inherit(TableBaseStyle)
	TableUnFocusedStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("#383838")).
				MarginTop(1).
				MarginLeft(5).
				Padding(1, 2).
				Inherit(TableBaseStyle)

	ViewportFocusedStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("62")).
				PaddingRight(2).
				MarginTop(1).
				Inherit(ViewportBaseStyle)
	ViewportUnFocusedStyle = lipgloss.NewStyle().
				BorderForeground(lipgloss.Color("#383838")).
				PaddingRight(2).
				MarginTop(1).
				Inherit(ViewportBaseStyle)

	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#626262")).
			MarginTop(2).
			MarginLeft(5).
			MarginBottom(2).
			PaddingLeft(1)
)
