package cmdutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/google/uuid"
)

type EmailValues struct {
	Email string
	Hash  string
}

// Generate a uuid but split and return the first part.
func randomId() string {
	rawId := uuid.New()

	return strings.Split(rawId.String(), "-")[0]
}

func randomDomain() (string, error) {
	dl, err := api.GetDomainsList()
	if err != nil {
		return "", err
	}
	return dl[rand.Intn(len(dl))], nil
}

func GenerateRandomEmail() (EmailValues, error) {
	id := randomId()
	domain, err := randomDomain()
	if err != nil {
		return EmailValues{}, err
	}

	email := fmt.Sprintf("%s%s", id, domain)
	hasher := md5.Sum([]byte(email))
	emailHash := hex.EncodeToString(hasher[:])

	return EmailValues{email, emailHash}, nil
}
