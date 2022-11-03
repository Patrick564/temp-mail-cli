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

type UserEmail struct {
	Email string
	Hash  string
	// Inbox api.Emails
}

// Generate a uuid but split and return the first part.
func randomID() string {
	rawId := uuid.New()

	return strings.Split(rawId.String(), "-")[0]
}

// Retrive the domain list and pick one random.
func randomDomain() (string, error) {
	dl, err := api.GetDomains()
	if err != nil {
		return "", err
	}
	return dl[rand.Intn(len(dl))], nil
}

func GenerateUserEmail() (*UserEmail, error) {
	id := randomID()
	domain, err := randomDomain()
	if err != nil {
		return nil, err
	}

	email := fmt.Sprintf("%s%s", id, domain)
	hasher := md5.Sum([]byte(email))
	emailHash := hex.EncodeToString(hasher[:])

	return &UserEmail{email, emailHash}, nil
}
