package ui

import (
	"fmt"
	"strings"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/Patrick564/temp-mail-cli/internal/ui/styles"
	"github.com/Patrick564/temp-mail-cli/pkg/user"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

const (
	tableView    sessionState = 0
	viewportView sessionState = 1
	width        int          = 72
	height       int          = 20
)

type sessionState int

type model struct {
	User     user.UserModel
	Table    table.Model
	Viewport viewport.Model
	state    sessionState
	term     *glamour.TermRenderer
}

func (m model) Init() tea.Cmd {
	// return utils.InitNewUser,
	return tea.Batch(m.Viewport.Init())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case user.UserModel:
		m.User = msg
		m.Viewport.SetContent(m.User.RenderedMail)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape:
			if m.Table.Focused() {
				m.Table.Blur()
			} else {
				m.Table.Focus()
			}
		case tea.KeyEnter:
			m.User.Inbox = api.Emails{
				{MailID: "a", MailFrom: "a", MailText: "a", MailSubject: "a", MailTimestamp: 12.3123},
				{MailID: "b", MailFrom: "b", MailText: "a", MailSubject: "a", MailTimestamp: 12.3123},
				{MailID: "c", MailFrom: "c", MailText: "a", MailSubject: "a", MailTimestamp: 12.3123},
			}
			if len(m.User.Inbox) == 0 {
				break
			}
			err := m.User.RenderActiveMail(m.Table.Cursor(), m.term)
			if err != nil {
				return m, tea.Quit
			}
			m.Viewport.SetContent(m.User.RenderedMail)
		case tea.KeyLeft:
			m.state = tableView
		case tea.KeyRight:
			m.state = viewportView
		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "r":
				err := m.User.RefreshInbox(m.User.Hash)
				if err != nil {
					return m, tea.Quit
				}
				m.Table.SetRows(m.User.InboxTable)
			case "q":
				return m, tea.Quit
			}
		}

		switch m.state {
		case tableView:
			m.Table, cmd = m.Table.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.Viewport, cmd = m.Viewport.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	b.WriteString(styles.TitleStyle.Render(fmt.Sprintf("Email: %s", m.User.Email)))
	if m.state == tableView {
		b.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.TableFocusedStyle.Render(m.Table.View()),
			styles.ViewportUnFocusedStyle.Render(m.Viewport.View()),
		))
	} else {
		b.WriteString(lipgloss.JoinHorizontal(
			lipgloss.Top,
			styles.TableUnFocusedStyle.Render(m.Table.View()),
			styles.ViewportFocusedStyle.Render(m.Viewport.View()),
		))
	}
	b.WriteString(styles.HelpStyle.Render(styles.HelpText))
	b.WriteRune('\n')

	return b.String()
}

func New() (tea.Model, error) {
	columns := []table.Column{
		{Title: "Sender", Width: 35},
		{Title: "Subject", Width: 30},
		{Title: "Open", Width: 8},
	}
	rows := []table.Row{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)
	t.SetStyles(table.Styles{
		Header:   styles.TableHeaderStyle,
		Selected: styles.TableCellSelectedStyle,
		Cell:     styles.TableCellStyle,
	})

	vp := viewport.New(width, height)

	renderer, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return nil, err
	}

	return model{
		Table:    t,
		Viewport: vp,
		state:    tableView,
		term:     renderer,
	}, nil
}
