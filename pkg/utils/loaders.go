package utils

import (
	"github.com/Patrick564/temp-mail-cli/pkg/user"
	tea "github.com/charmbracelet/bubbletea"
)

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func InitNewUser() tea.Msg {
	user, err := user.New()
	if err != nil {
		return errMsg{err}
	}

	return *user
}
