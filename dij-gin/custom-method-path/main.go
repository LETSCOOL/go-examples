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
}

// SayHi a http request with "get" method.
// Url should like this in local: http://localhost:8000/hello.
func (s *TWebServer) SayHi(ctx struct {
	WebContext `http:"hello,method=get"`
}) {
	ctx.IndentedJSON(http.StatusOK, "hello")
}

func main() {
	if err := LaunchGin(&TWebServer{}); err != nil {
		log.Fatalln(err)
	}
}
