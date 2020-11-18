package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	this.Ctx.WriteString("hello world")
}

func (this *UserController) UserGet() {
	this.Ctx.WriteString("shit!!")
}

func init() {
	beego.Router("/user", &UserController{})

	beego.Router("/all/:key", &UserController{}, "get:UserGet")
}
