// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	. "github.com/letscool/dij-gin"
	"log"
	"net/http"
	"reflect"
)

type TWebServer struct {
	WebServer
}

// GetHello a http request with "get" method.
// Url should like this in local: http://localhost:8000/hello?name=wayne&age=123
func (s *TWebServer) GetHello(ctx struct {
	WebContext
	name string
	age  int
}) {
	//fmt.Printf("%s", ctx.Query("name"))
	ctx.IndentedJSON(http.StatusOK, fmt.Sprintf("/hello %s, %d years old", ctx.name, ctx.age))
}

func main() {
	wsTyp := reflect.TypeOf(TWebServer{})
	//dij.EnableLog()
	if err := LaunchGin(wsTyp); err != nil {
		log.Fatalln(err)
	}
}
