package cmdutil

import (
	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func LoadEmail() tea.Msg {
	params, err := GenerateRandomEmail()
	if err != nil {
		return errMsg{err}
	}

	return *params
}

func LoadEmailsList(hash string) ([]table.Row, error) {
	var row []table.Row

	emails, err := api.GetEmails(hash)
	if err != nil {
		if err == api.ErrEmptyInbox {
			row = append(row, []string{"No new messages yet.", "", ""})
			return row, nil
		}
		return nil, err
	}

	for _, e := range emails {
		row = append(row, []string{e.MailFrom, e.MailSubject, " ▶ "})
	}

	return row, nil
}
