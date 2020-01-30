package sample

import (
	"github.com/iesreza/foundation/lib/router"
	"github.com/iesreza/foundation/system"
)

func (component component) Routers() {


	//matches api/*
	system.Router.Group("api",nil, func(handle *router.Route) {

		//matches GET api/hello
		handle.Match("hello","GET", func(req router.Request) {
			req.WriteObject( helloWorld() )
		})

		//matches POST api/echo
		handle.Match("echo","POST", func(req router.Request) {
			req.WriteObject( echo(req) )
		})


	})

	//matches static files inside /components/sample/assets
	system.Router.Static(component.Assets,component.Assets,nil)

	//return error if nothing matched
	system.Router.Fallback = func(req router.Request) {
		req.WriteObject(`{"status":false,"error":"Invalid API"}`)
	}
}
