// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	. "github.com/letscool/dij-gin"
	"github.com/letscool/dij-gin/libs"
	"github.com/letscool/go-examples/dij-gin/shared"
	"log"
	"time"
)

type TWebServer struct {
	WebServer
	_ *libs.SwaggerController `di:""` // Bind OpenApi controller in root.
	_ *TUserController        `di:""`
}

type TUserController struct {
	WebController `http:"user"`

	_ *shared.BearerMiddleware `di:""`
}

// GetMe a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/me.
// And login with username "john" and password "abc".
func (u *TUserController) GetMe(ctx struct {
	WebContext `http:",middleware=bearer" security:"bearer_1"`
}) (result struct {
	Account *shared.Account `http:"200,json"`
}) {
	result.Account = ctx.MustGet(shared.BearerUserKey).(*shared.Account)
	return
}

func main() {
	// a fake account db
	ac := &shared.FakeAccountDb{} // This object must implement shared.BearerValidator interface.
	accounts := ac.InitFakeDb()
	// generate a jwt token for test.
	// this is for test only, the production should another way to get token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Issuer:    "dij-gin-samples",
		ID:        accounts[0].User,
	})
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(ac.BearerSecret()))
	if err != nil {
		log.Panicln(err)
	} else {
		fmt.Println("*************************")
		fmt.Printf("User: %s\n", accounts[0].User)
		fmt.Printf("Bearer token: %s\n", tokenString)
		fmt.Println("*************************")
	}
	// launch a web server
	config := NewWebConfig().
		SetDependentRef(shared.RefKeyForBearerValidator, ac).
		SetOpenApi(func(o *OpenApiConfig) {
			o.Enable().UseHttpOnly().SetDocPath("doc").
				AppendBearerAuth("bearer_1")
		})
	if err := LaunchGin(&TWebServer{}, config); err != nil {
		log.Fatalln(err)
	}
}
