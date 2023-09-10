package main

import (
	"github.com/TMkoch/gofarm/constant"
	"github.com/TMkoch/gofarm/router"
	"github.com/TMkoch/gofarm/server"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	//load static
	constant.LoadStatic(app)

	//load template folder
	app.Renderer = constant.LoadTemplate()

	//router
	router.LoadAllRouters(app)

	//server
	server.SetServer(app)
}
