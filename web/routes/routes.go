package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"xormdemo191226/bootstrap"
	"xormdemo191226/service"
	"xormdemo191226/web/controller"
	"xormdemo191226/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	//superstarService := services.NewSuperstarService()
	movieService := service.NewMovieService()

	//index := mvc.New(b.Party("/"))
	//index.Register(movieService)
	//index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(movieService)
	admin.Handle(new(controller.AdminController))

}
