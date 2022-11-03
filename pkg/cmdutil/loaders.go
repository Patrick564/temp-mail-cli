package cmdutil

import (
	"strconv"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type InboxContent struct {
	Rows    []table.Row
	Content api.Emails
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

// loadEmailUser
func LoadEmail() tea.Msg {
	params, err := GenerateUserEmail()
	if err != nil {
		return errMsg{err}
	}

	return *params
}

// loadEmails
func LoadEmailsList(hash string) (*InboxContent, error) {
	emails, err := api.GetEmails(hash)
	if err != nil {
		if err == api.ErrEmptyEmails {
			return &InboxContent{}, nil
		}
		return nil, err
	}

	var rows []table.Row

	for idx, e := range emails {
		rows = append(rows, []string{
			strconv.Itoa(idx + 1),
			e.MailFrom,
			e.MailSubject,
			" ▶ ",
		})
	}

	return &InboxContent{rows, emails}, nil
}
