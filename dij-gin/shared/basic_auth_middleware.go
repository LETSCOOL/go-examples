// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package shared

import (
	. "github.com/letscool/dij-gin"
	"net/http"
	"strconv"
)

const RefKeyForBasicAuthAccountCenter = "mdl.basic_auth_account"
const BasicAuthUserKey = "BasicAuthUserKey"

// BasicAuthAccountCenter
//
//	     ac := &AccountCenter{} // the struct must implement BasicAuthAccountCenter interface
//		 config := NewWebConfig().SetDependentRef(RefKeyForBasicAuthAccountCenter, ac)
type BasicAuthAccountCenter interface {
	GetRealm() string
	SearchCredential(credential string) (account any, found bool)
}

type BasicAuthMiddleware struct {
	WebMiddleware

	account BasicAuthAccountCenter `di:"mdl.basic_auth_account"`

	realm string
}

func (b *BasicAuthMiddleware) DidDependencyInitialization() {
	realm := b.account.GetRealm()
	b.realm = "Basic realm=" + strconv.Quote(realm)
}

func (b *BasicAuthMiddleware) Authorize(ctx struct {
	WebContext `http:"basic_auth,method=handle"`
}) {
	authorization := ctx.GetRequestHeader("Authorization")
	user, found := b.account.SearchCredential(authorization)
	if !found {
		ctx.Header("WWW-Authenticate", b.realm)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.Set(BasicAuthUserKey, user)
}
