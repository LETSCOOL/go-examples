// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	. "github.com/letscool/dij-gin"
	"log"
)

type TWebServer struct {
	WebServer
}

// GetResp a http request with "get" method.
// Url should like this in local: http://localhost:8000/resp?select=1 .
// Use *curl -v* command to see response code.
func (s *TWebServer) GetResp(ctx struct {
	WebContext
	Select int `http:"select"`
}) (result struct {
	Ok200       *string // the range of last three characters is between 2xx and 5xx, so the response code = 200
	Ok          *string `http:"201"` // force response code to 201
	Redirect302 *string // redirect data should be string type, because it is redirect location.
	Error       error   // default response code for error is 400
}) {
	switch ctx.Select {
	case 1:
		data := "ok"
		result.Ok200 = &data
	case 2:
		data := "ok"
		result.Ok = &data
	case 3:
		url := "https://github.com/letscool"
		result.Redirect302 = &url
	default:
		result.Error = errors.New("an error")
	}
	return
}

func main() {
	if err := LaunchGin(&TWebServer{}); err != nil {
		log.Fatalln(err)
	}
}
