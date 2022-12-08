// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"crypto/subtle"
	"encoding/base64"
	. "github.com/letscool/dij-gin"
	"github.com/letscool/dij-gin/libs"
	"log"
)

type TWebServer struct {
	WebServer

	_ *TUserController `di:""`
}

type TUserController struct {
	WebController `http:"user"`

	_ *libs.BasicAuthMiddleware `di:""`
}

// GetMe a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/me
func (u *TUserController) GetMe(ctx struct {
	WebContext `http:",middleware=basic_auth"`
}) (result struct {
	Account *Account `http:"200,json"`
}) {
	result.Account = ctx.MustGet(libs.BasicAuthUserKey).(*Account)
	return
}

func main() {
	ac := &FakeAccountDb{}
	ac.initFakeDb()
	config := NewWebConfig().
		SetDependentRef(libs.RefKeyForBasicAuthAccountCenter, ac)
	if err := LaunchGin(&TWebServer{}, config); err != nil {
		log.Fatalln(err)
	}
}

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

func (a *FakeAccountDb) initFakeDb() {
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
