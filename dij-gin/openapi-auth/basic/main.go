// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	. "github.com/letscool/dij-gin"
	"github.com/letscool/dij-gin/libs"
	"github.com/letscool/go-examples/dij-gin/shared"
	"log"
)

type TWebServer struct {
	WebServer
	_ *libs.SwaggerController `di:""` // Bind OpenApi controller in root.
	_ *TUserController        `di:""`
}

type TUserController struct {
	WebController `http:"user"`

	_ *libs.BasicAuthMiddleware `di:""`
}

// GetMe a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/me.
// And login with username "john" and password "abc".
func (u *TUserController) GetMe(ctx struct {
	WebContext `http:",middleware=basic_auth" security:"basic_auth_1"`
}) (result struct {
	Account *shared.Account `http:"200,json"`
}) {
	result.Account = ctx.MustGet(libs.BasicAuthUserKey).(*shared.Account)
	return
}

func main() {
	ac := &shared.FakeAccountDb{} // This object should implement libs.AccountForBasicAuth interface.
	ac.InitFakeDb()
	config := NewWebConfig().
		SetDependentRef(libs.RefKeyForBasicAuthAccountCenter, ac).
		SetOpenApi(func(o *OpenApiConfig) {
			o.Enable().UseHttpOnly().SetDocPath("doc").
				AppendBasicAuth("basic_auth_1")
		})
	if err := LaunchGin(&TWebServer{}, config); err != nil {
		log.Fatalln(err)
	}
}
