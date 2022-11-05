package user

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/glamour"
	"github.com/google/uuid"
)

var baseContent string = `
## Email: %s | Time: %s
---
## Subject: %s
---
%s
`

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

func randomID() string {
	rawId := uuid.New()

	return strings.Split(rawId.String(), "-")[0]
}

func randomDomain() (string, error) {
	dl, err := api.GetDomains()
	if err != nil {
		return "", err
	}
	return dl[rand.Intn(len(dl))], nil
}

func New() (*UserModel, error) {
	id := randomID()
	domain, err := randomDomain()
	if err != nil {
		return nil, err
	}

	email := fmt.Sprintf("%s%s", id, domain)
	hasher := md5.Sum([]byte(email))
	emailHash := hex.EncodeToString(hasher[:])

	return &UserModel{Email: email, Hash: emailHash}, nil
}
