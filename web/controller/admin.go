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

func (c *AdminController) Get() mvc.Result {
	datalist := c.Service.GetAll()
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":    "小杜的管理后台",
			"DataList": datalist,
		},
	}
}
