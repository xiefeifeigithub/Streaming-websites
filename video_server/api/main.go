package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}


func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}

/*
listen->RegisterHandlers->handlers
1. go语言对于每一个handler都用一个go routine来处理（默认）

API请求过程：handler->validation{1.request, 2.user}->business logic->response.
1. data model
2. error handling
注意：对于request的处理采用这种分层架构对于编写test case是很容易的，
而且更能照顾到它的可扩展性，对工程上的效率也是非常高的。

流程
main->middleware(校验、鉴权、流控等httpmiddleware)->defs(message, err)->handlers->dbops->response
 */