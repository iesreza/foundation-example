package sample

import (
	"github.com/iesreza/foundation/lib/router"
)

func helloWorld() interface{} {
	return struct {
		Success bool
		Message string
	}{
		true, "Hello World",
	}
}

func echo(req router.Request) interface{} {
	req.Req().ParseForm()
	return req.Req().Form
}
