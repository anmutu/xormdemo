package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"xormdemo191226/service"
)

type AdminController struct {
	Ctx     iris.Context
	Service service.MovieService
}

func (C *AdminController) Get() mvc.Result {
	//todo
	return nil
}
