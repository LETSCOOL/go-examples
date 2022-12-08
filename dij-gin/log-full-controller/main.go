// Copyright 2022 Yuchi Chen. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/letscool/dij-gin"
	"github.com/letscool/dij-gin/libs"
	"log"
	"net/http"
)

type TWebServer struct {
	WebServer `http:",middleware=log"`

	_ *libs.LogMiddleware `di:""`
}

// GetHello a http request with "get" method.
// Url should like this in local: http://localhost:8000/hello
func (t *TWebServer) GetHello(ctx WebContext) {
	ctx.IndentedJSON(http.StatusOK, "/hello")
}

func main() {
	//f, _ := os.Create("gin.log") // log to file
	config := NewWebConfig().
		//SetDefaultWriter(io.MultiWriter(f)).
		SetDependentRef(libs.RefKeyForLogFormatter, (gin.LogFormatter)(func(params gin.LogFormatterParams) string {
			// your custom format
			return fmt.Sprintf("[%s-%s] \"%s %s\"\n",
				params.ClientIP,
				params.TimeStamp.Format("15:04:05.000"),
				params.Method,
				params.Path,
			)
		}))
	if err := LaunchGin(&TWebServer{}, config); err != nil {
		log.Fatalln(err)
	}
}
