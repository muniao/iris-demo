package controller

import (
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"org.mm/iris-demo/models"
	"org.mm/iris-demo/services"
	"org.mm/iris-demo/datasource"
)

type IndexController struct {
	Ctx     iris.Context
	Service services.MovieService
}

func (c *IndexController) Get() mvc.Result {

	datalist := c.Service.GetAll()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"info":  data,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title":    "球星库",
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.Instance().ClearCache(&models.Movie{})
	if err != nil {
		log.Fatal(err)
	}
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}