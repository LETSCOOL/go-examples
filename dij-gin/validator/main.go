// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
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

// GetUserById a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/2345/profile.
// The result will be:
//
//	{"message":"Key: 'Id' Error:Field validation for 'Id' failed on the 'lte' tag","code":"400"}
func (u *TUserController) GetUserById(ctx struct {
	WebContext `http:":id/profile"`
	Id         int `http:"id,in=path" validate:"gte=100,lte=999"`
}) {
	ctx.IndentedJSON(http.StatusOK, fmt.Sprintf("get user(#%d)'s profile", ctx.Id))
}

func main() {
	if err := LaunchGin(&TWebServer{}); err != nil {
		log.Fatalln(err)
	}
}
