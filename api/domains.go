package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Domains []string

func GetDomains() (Domains, error) {
	req, err := http.NewRequest("GET", os.Getenv("DOMAINS_LIST_URL"), nil)
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

	var d Domains

	err = json.Unmarshal(body, &d)
	if err != nil {
		return nil, err
	}

	return d, nil
}
