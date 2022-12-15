package shared

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type Account struct {
	User  string `json:"user"`
	Email string `json:"email"`
	pass  string
	realm string
}

// FakeAccountDb should implement BasicAuthAccountCenter interface
type FakeAccountDb struct {
	accounts []Account
	creds    map[string]*Account
}

func (a *FakeAccountDb) InitFakeDb() []Account {
	a.accounts = []Account{
		{"john", "john@fake.com", "abc", ""},
		{"wayne", "wayne@fake.com", "abc", ""},
	}

	a.creds = map[string]*Account{}
	for i := range a.accounts {
		account := &a.accounts[i]
		base := account.User + ":" + account.pass
		cred := "Basic " + base64.StdEncoding.EncodeToString([]byte(base))
		a.creds[cred] = account
	}
	return a.accounts
}

// GetRealm implements one function of BasicAuthAccountCenter interface.
func (a *FakeAccountDb) GetRealm() string {
	return "Authorization Required"
}

// SearchCredential implements one function of BasicAuthAccountCenter interface.
func (a *FakeAccountDb) SearchCredential(credential string) (account any, found bool) {
	for key, value := range a.creds {
		if subtle.ConstantTimeCompare([]byte(key), []byte(credential)) == 1 {
			return value, true
		}
	}
	return nil, false
}

// BearerSecret implements one function of BearerValidator interface.
func (a *FakeAccountDb) BearerSecret() string {
	return "my_secret"
}

// ValidateClaims implements one function of BearerValidator interface.
func (a *FakeAccountDb) ValidateClaims(claims any) (userData any, err error) {
	registeredClams := claims.(jwt.MapClaims)
	for _, account := range a.accounts {
		if account.User == registeredClams["jti"] {
			return &account, nil
		}
	}
	return nil, errors.New("invalid claims")
}
