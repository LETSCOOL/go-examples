package shared

import (
	"crypto/subtle"
	"encoding/base64"
)

type Account struct {
	User  string `json:"user"`
	Email string `json:"email"`
	pass  string
	realm string
}

// FakeAccountDb should implement libs.AccountForBasicAuth interface
type FakeAccountDb struct {
	accounts []Account
	creds    map[string]*Account
}

func (a *FakeAccountDb) InitFakeDb() {
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
}

func (a *FakeAccountDb) GetRealm() string {
	return "Authorization Required"
}

func (a *FakeAccountDb) SearchCredential(credential string) (account any, found bool) {
	for key, value := range a.creds {
		if subtle.ConstantTimeCompare([]byte(key), []byte(credential)) == 1 {
			return value, true
		}
	}
	return nil, false
}
