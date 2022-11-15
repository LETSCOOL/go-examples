package main

import (
	. "github.com/letscool/dij-gin"
	"log"
	"net/http"
	"reflect"
)

type TWebServer struct {
	WebServer
}

// GetHello a http request with "get" method.
// Url should like this in local: http://localhost:8000/hello
func (s *TWebServer) GetHello(ctx WebContext) {
	ctx.IndentedJSON(http.StatusOK, "/hello")
}

func main() {
	wsTyp := reflect.TypeOf(TWebServer{})
	//dij.EnableLog()
	if err := LaunchGin(wsTyp); err != nil {
		log.Fatalln(err)
	}
}
