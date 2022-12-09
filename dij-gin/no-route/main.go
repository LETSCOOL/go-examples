// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	. "github.com/letscool/dij-gin"
	"log"
)

type TWebServer struct {
	WebServer
}

// NoRoute is entry for page not found, is only supported in root path currently.
// Any query will show the log.
func (s *TWebServer) NoRoute(ctx WebContext) {
	log.Printf("no route: %s\n", ctx.Request.RequestURI)
}

func main() {
	if err := LaunchGin(&TWebServer{}); err != nil {
		log.Fatalln(err)
	}
}
