package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type emails struct {
	Id        string  `json:"mail_id"`
	From      string  `json:"mail_from"`
	Subject   string  `json:"mail_subject"`
	Text      string  `json:"mail_text"`
	Timestamp float64 `json:"mail_timestamp"`
}

func GetEmails(emailHash string) ([]emails, error) {
	url := fmt.Sprintf("%s/%s/", os.Getenv("GET_EMAILS_URL"), emailHash)

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

	var e []emails

	err = json.Unmarshal(body, &e)
	if err != nil {
		return nil, err
	}

	return e, nil
}
