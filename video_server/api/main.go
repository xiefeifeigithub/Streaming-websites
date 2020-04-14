package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)
	router.POST("/user/:username", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

/*
listen->RegisterHandlers->handlers
1. go语言对于每一个handler都用一个go routine来处理（默认）

API请求过程：handler->validation{1.request, 2.user}->business logic->response.
1. data model
2. error handling
注意：对于request的处理采用这种分层架构对于编写test case是很容易的，
而且更能照顾到它的可扩展性，对工程上的效率也是非常高的。
 */

