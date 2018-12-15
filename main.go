package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main(){
	fmt.Println("程序启动")
	defer fmt.Println("程序停止")

	app:= iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	user := app.Party("/users",myAuthMiddlewareHandler)

	user.Get("/{id:int}/profile",userProfileHandler)

	user.Get("/inbox/{id:int}",userMessageHandler)

	app.Run(iris.Addr(":8080"))
}

func myAuthMiddlewareHandler(ctx iris.Context){
	ctx.WriteString("header\n")
	ctx.Next()
	ctx.WriteString("\nfoot")
}
func userProfileHandler(ctx iris.Context) {//
	id:=ctx.Params().Get("id")
	fmt.Println("AAA - " + id)
	ctx.WriteString(id)
}
func userMessageHandler(ctx iris.Context){
	id:=ctx.Params().Get("id")
	fmt.Println("BBB - " + id)
	ctx.WriteString(id)
}