package cmdutil

import (
	"strconv"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type InboxContent struct {
	Rows    []table.Row
	Content api.EmailContent
}

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func LoadEmail() tea.Msg {
	params, err := GenerateRandomEmail()
	if err != nil {
		return errMsg{err}
	}

	return *params
}

func LoadEmailsList(hash string) (*InboxContent, error) {
	emails, err := api.GetEmails(hash)
	if err != nil {
		if err == api.ErrEmptyInbox {
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
