package main

// https://github.com/astaxie/beego/

import (
	"fmt"
	_ "libsys/models"

	_ "libsys/controllers"

	"github.com/astaxie/beego"
)

func main() {
	fmt.Println("hello world!")
	//beego.AutoRender = false
	beego.BConfig.CopyRequestBody = true
	beego.Run()
}
