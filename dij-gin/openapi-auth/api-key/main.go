// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	. "github.com/letscool/dij-gin"
	"github.com/letscool/dij-gin/libs"
	"log"
)

type TWebServer struct {
	WebServer
	_ *libs.SwaggerController `di:""` // Bind OpenApi controller in root.
	_ *TUserController        `di:""`
}

type TUserController struct {
	WebController `http:"user"`
}

// GetMe a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/me.
// And login with username "john" and password "abc".
func (u *TUserController) GetMe(ctx struct {
	WebContext `http:"" security:"ApiId,ApiKey" description:"Set ApiId=abc and ApiKey=EFG"`
	apiId      string `http:"api_id"`
	apiKey     string `http:"api_key"`
}) (result struct {
	Message *string `http:"200"`
	Error   error   `http:"401"`
}) {
	if ctx.apiId == "abc" && ctx.apiKey == "EFG" {
		msg := "You got permission to access this api"
		result.Message = &msg
	} else {
		result.Error = errors.New("unauthorized api key")
	}
	return
}

func main() {
	// launch a web server
	config := NewWebConfig().
		SetOpenApi(func(o *OpenApiConfig) {
			o.Enable().UseHttpOnly().SetDocPath("doc").
				AppendApiKeyAuth("ApiId", "query", "api_id").
				AppendApiKeyAuth("ApiKey", "query", "api_key")
		})
	if err := LaunchGin(&TWebServer{}, config); err != nil {
		log.Fatalln(err)
	}
}
