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

	openapi *libs.SwaggerController `di:""` // Bind OpenApi controller in root.
}

// GetResp a http request with "get" method.
// Url should like this in local: http://localhost:8000/resp?select=1 .
// Use *curl -v* command to see response code.
func (s *TWebServer) GetResp(ctx struct {
	WebContext
	Select int `http:"select"`
}) (result struct {
	Ok200 *string // the range of last three characters is between 2xx and 5xx, so the response code = 200
	Ok    *string `http:"201"` // force response code to 201
	Error error   // default response code for error is 400
}) {
	switch ctx.Select {
	case 1:
		data := "ok"
		result.Ok200 = &data
	case 2:
		data := "ok"
		result.Ok = &data
	default:
		result.Error = errors.New("an error")
	}
	return
}

// The OpenAPI page will be enabled in location: http://localhost:8000/doc.
func main() {
	config := NewWebConfig().
		SetOpenApi(func(o *OpenApiConfig) {
			o.SetEnabled(true).UseHttpOnly().SetDocPath("doc")
		})
	if err := LaunchGin(&TWebServer{}, config); err != nil {
		log.Fatalln(err)
	}
}
