package main

import (
	. "github.com/letscool/dij-gin"
	"log"
	"net/http"
	"reflect"
)

type TWebServer struct {
	WebServer

	_ *TUserController `di:""`
}

// GetHello a http request with "get" method.
// Url should like this in local: http://localhost:8000/hello
func (s *TWebServer) GetHello(ctx WebContext) {
	ctx.IndentedJSON(http.StatusOK, "/hello")
}

type TUserController struct {
	WebController `http:"user"`
}

// GetMe a http request with "get" method.
// Url should like this in local: http://localhost:8000/user/me
func (u *TUserController) GetMe(ctx WebContext) {
	ctx.IndentedJSON(http.StatusOK, "/user/me")
}

func main() {
	wsTyp := reflect.TypeOf(TWebServer{})
	//dij.EnableLog()
	if err := LaunchGin(wsTyp); err != nil {
		log.Fatalln(err)
	}
}
