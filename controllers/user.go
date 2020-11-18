package controllers

import (
	//"encoding/json"
	"fmt"
	"strconv"

	"libsys/tools"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// beego orm 的使用：https://www.cnblogs.com/yangmingxianshen/p/10122427.html

// beego orm 高级使用：https://blog.csdn.net/yang731227/article/details/82503059

// 查询所有：_, qs := o.QueryTable("user").All(&users)
// 查询单个：_, qs := o.QueryTable("Student").Filter("Id", 1).One(&student)
// 插入对象：id, err := o.Insert(user)

func (this *UserController) UserList() {

	//var pageIndex = this.Ctx.Input.Param(":pi") // 获取url中的参数
	//var pageSize = this.Ctx.Input.Param(":ps")
	//out := make(map[string]interface{})
	if pageIndex, err := strconv.Atoi(this.Ctx.Input.Param(":pi")); err == nil {
		fmt.Printf("%T, %v", pageIndex, pageIndex)
	} else {
		//out["code"] = 1
		//out["msg"] = "input pi error!"
		//json, _ := json.Marshal(out)
		this.Ctx.WriteString(tools.Response(1, "input pi error!", []byte{'n'}))
		return
	}

	// _ = pageSize

	// fmt.Println("pi:%v", pageIndex)

	// this.Ctx.Input.Param(":username")

	// o := orm.NewOrm()

	// var users []models.User //查询的结果是集合的话，这里需要加上[]
	// qs := o.QueryTable("user")

	// _, err := qs.Offset(pageIndex * pageSize).Limit(pageSize).All(&users)

	// if err == nil {
	// 	for _, user := range users {
	// 		fmt.Println(user)
	// 	}
	// } else {
	// 	logs.Info(err)
	// }

	this.Ctx.WriteString("hello world")

}

func (this *UserController) UserGet() {
	this.Ctx.WriteString("shit!!")
}

func init() {
	beego.Router("/user/list/:pi/:ps", &UserController{}, "get:UserList")

	beego.Router("/all/:key", &UserController{}, "get:UserGet")
}
