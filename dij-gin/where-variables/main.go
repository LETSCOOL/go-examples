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

// PutUserById a http request with "get" method.
// Curl this url should like this in local:
//
//	curl -d "age=34&name=wayne" -X PUT http://localhost:8000/user/2345/profile
//
// The result should be:
//
//	"update user(#2345)'s name wayne and age 34"
func (u *TUserController) PutUserById(ctx struct {
	WebContext `http:":id/profile"`
	Id         int    `http:"id,in=path"`
	Name       string `http:"name,in=body"`
	Age        int    `http:"age,in=body"`
}) {
	ctx.IndentedJSON(http.StatusOK, fmt.Sprintf("update user(#%d)'s name %s and age %d", ctx.Id, ctx.Name, ctx.Age))
}

func main() {
	if err := LaunchGin(&TWebServer{}); err != nil {
		log.Fatalln(err)
	}
}
