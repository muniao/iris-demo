package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	"org.mm/iris-demo/services"
	"org.mm/iris-demo/web/controller"
	"org.mm/iris-demo/web/middleware"
)

func main() {
	movieService := services.NewMovieService()
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Use(recover.New())
	app.Use(logger.New())
	
	index := mvc.New(app.Party("/"))
	index.Router.Use(middleware.BasicAuth)
	index.Register(movieService)
	index.Handle(new(controller.IndexController))
	app.RegisterView(iris.HTML("./web/views", ".html").Reload(true))
	app.StaticWeb("/public", "./web/public")

	app.Run(iris.Addr(":8080"))
}
