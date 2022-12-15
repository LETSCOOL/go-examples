// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package shared

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/letscool/dij-gin"
	"log"
	"net/http"
	"reflect"
	"strings"
)

const RefKeyForBearerValidator = "mdl.bearer_validator"
const BearerUserKey = "BearerUserKey"

// BearerValidator
//
//	     av := &AccountValidator{} // the struct must implement BearerValidator interface
//		 config := NewWebConfig().SetDependentRef(RefKeyForBearerValidator, av)
type BearerValidator interface {
	BearerSecret() string
	ValidateClaims(claims any) (userData any, err error)
}

type BearerMiddleware struct {
	WebMiddleware

	validator BearerValidator `di:"mdl.bearer_validator"`
}

func (b *BearerMiddleware) ParseBearerHeader(ctx struct {
	WebContext `http:"bearer,method=handle"`
}) {
	authorization := ctx.GetRequestHeader("Authorization")
	if !strings.HasPrefix(authorization, "Bearer ") {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token, err := jwt.Parse(authorization[7:], func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(b.validator.BearerSecret()), nil
	})
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok {
		log.Printf("Not a MapClaims, %v\n%v\n", reflect.TypeOf(token.Claims), token.Claims)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	} else {
		if user, err := b.validator.ValidateClaims(claims); err != nil {
			log.Println(err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			ctx.Set(BearerUserKey, user)
		}
	}
}
