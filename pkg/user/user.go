package user

import (
	"fmt"
	"time"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/Patrick564/temp-mail-cli/pkg/random"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/glamour"
)

var (
	emtpyContent string = `
+++++++++++++++++++++++++++++++++++++++++++++++++++++++

                No mail open yet...

+++++++++++++++++++++++++++++++++++++++++++++++++++++++
`
	baseContent string = `
## Email: %s | Time: %s
---
## Subject: %s
---
%s
`
)

type UserModel struct {
	Email        string
	Hash         string
	Inbox        api.Emails
	InboxTable   []table.Row
	ActiveMail   api.EmailContent
	RenderedMail string
}

func (u *UserModel) RefreshInbox(hash string) error {
	emails, err := api.GetEmails(hash)
	if err != nil {
		if err == api.ErrEmptyEmails {
			return nil
		}
		return err
	}

	var rows []table.Row

	for _, e := range emails {
		rows = append(rows, []string{
			e.MailFrom,
			e.MailSubject,
			" ▶ ",
		})
	}

	u.Inbox = emails
	u.InboxTable = rows

	return nil
}

func (u *UserModel) setActiveMail(idx int) {
	u.ActiveMail = u.Inbox[idx]
}

func (u *UserModel) RenderActiveMail(idx int, t *glamour.TermRenderer) error {
	u.setActiveMail(idx)

	time := time.Unix(int64(u.ActiveMail.MailTimestamp), 0)

	var rawContent = fmt.Sprintf(
		baseContent,
		u.ActiveMail.MailFrom,
		time.Format("02-01-2006 05:04:03"),
		u.ActiveMail.MailSubject,
		u.ActiveMail.MailText,
	)
	content, err := t.Render(rawContent)
	if err != nil {
		return err
	}

	u.RenderedMail = content

	return nil
}

func New() (*UserModel, error) {
	email, hash, err := random.RandomUserEmail()
	if err != nil {
		return nil, err
	}
	return &UserModel{Email: email, Hash: hash, RenderedMail: emtpyContent}, nil
}
