package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

var ErrEmptyEmails = errors.New("empty inbox")

type EmailContent struct {
	MailID        string  `json:"mail_id"`
	MailFrom      string  `json:"mail_from"`
	MailSubject   string  `json:"mail_subject"`
	MailText      string  `json:"mail_text"`
	MailTimestamp float64 `json:"mail_timestamp"`
}

type Emails []EmailContent

func GetEmails(emailHash string) (Emails, error) {
	url := fmt.Sprintf("%s%s/", os.Getenv("GET_EMAILS_URL"), emailHash)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-RapidAPI-Key", os.Getenv("RAPID_API_KEY"))
	req.Header.Add("X-RapidAPI-Host", os.Getenv("RAPID_API_HOST"))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var e Emails

	err = json.Unmarshal(body, &e)
	if err != nil {
		// This conditional is for a response struct like {"error": "no new messages"}
		if string(body[0]) == "{" {
			return nil, ErrEmptyEmails
		}
		return nil, err
	}

	return e, nil
}
