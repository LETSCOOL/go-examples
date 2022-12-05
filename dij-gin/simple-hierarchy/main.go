// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	. "github.com/letscool/dij-gin"
	"log"
	"net/http"
)

type TWebServer struct {
	WebServer

	userCtl *TUserController `di:""` // inject dependency by class/default name
}

type TUserController struct {
	WebController `http:"user"` //group by 'user' path
}

// Get a http request with "get" method.
// Url should like this in local: http://localhost:8000/user
func (u *TUserController) Get(ctx WebContext) {
	ctx.IndentedJSON(http.StatusOK, "/user")
}

// GetMe a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/me
func (u *TUserController) GetMe(ctx WebContext) {
	ctx.IndentedJSON(http.StatusOK, "/user/me")
}

func main() {
	//dij.EnableLog()
	if err := LaunchGin(&TWebServer{}); err != nil {
		log.Fatalln(err)
	}
}
