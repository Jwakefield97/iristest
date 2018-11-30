package main

import (
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	app.StaticWeb("/static", "./views")
	app.Run(iris.Addr(":2000"))

}