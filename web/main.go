package main

import (
	"xormdemo191226/bootstrap"
	"xormdemo191226/web/middleware/identity"
	"xormdemo191226/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Movie database", "小杜同学")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
