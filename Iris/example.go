package main

import "github.com/kataras/iris/v12"

// func main() {
// 	app := iris.Default()
// 	app.Use(myMiddleware)

// 	app.Handle("GET", "/ping", func(ctx iris.Context) {
// 		ctx.JSON(iris.Map{"msg": "pong"})
// 	})

// 	app.Run(iris.Addr(":9001"))
// }

// func myMiddleware(ctx iris.Context) {
// 	ctx.Application().Logger().Infof("Runs before %s", ctx.Path())
// 	ctx.Next()
// }

// func main() {
// 	app := iris.New()

// 	app.RegisterView(iris.HTML("./views", ".html"))
// 	app.Get("/", func(ctx iris.Context) {
// 		ctx.ViewData("msg", "Hello Iris")
// 		ctx.View("hello.html")
// 	})

// 	app.Get("/user/{id:uint64}", func(ctx iris.Context) {
// 		userId, _ := ctx.Params().GetUint64("id")
// 		ctx.Writef("User ID: %d", userId)
// 	})

// 	app.Run(iris.Addr(":9001"))

// }

func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello World!</h1>")
	})

	app.Run(iris.Addr(":9001"), iris.WithTunneling)
}
