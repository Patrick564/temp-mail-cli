package random

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"

	"github.com/Patrick564/temp-mail-cli/api"
	"github.com/google/uuid"
)

func randomID() string {
	rawId := uuid.New()

	return strings.Split(rawId.String(), "-")[0]
}

func randomDomain() (string, error) {
	d, err := api.GetDomains()
	if err != nil {
		return "", err
	}
	return d[rand.Intn(len(d))], nil
}

func RandomUserEmail() (string, string, error) {
	id := randomID()
	domain, err := randomDomain()
	if err != nil {
		return "", "", err
	}

	email := fmt.Sprintf("%s%s", id, domain)
	hasher := md5.Sum([]byte(email))
	hash := hex.EncodeToString(hasher[:])

	return email, hash, nil
}
