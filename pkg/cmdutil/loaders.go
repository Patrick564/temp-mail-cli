package cmdutil

import tea "github.com/charmbracelet/bubbletea"

type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

func LoadEmail() tea.Msg {
	params, err := GenerateRandomEmail()
	if err != nil {
		return errMsg{err}
	}

	return *params
}
