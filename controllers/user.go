package controllers

import (
	//"encoding/json"

	"fmt"
	"strconv"

	"libsys/models"
	"libsys/tools"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

// beego orm 的使用：https://www.cnblogs.com/yangmingxianshen/p/10122427.html
// beego orm 高级使用：https://blog.csdn.net/yang731227/article/details/82503059

// 查询所有：_, qs := o.QueryTable("user").All(&users)
// 查询单个：_, qs := o.QueryTable("Student").Filter("Id", 1).One(&student)
// 插入对象：id, err := o.Insert(user)
// 获取url中的参数 var pageIndex = this.Ctx.Input.Param(":pi")
func (this *UserController) UserList() {
	pageIndex, err := strconv.Atoi(this.Ctx.Input.Param(":pi"))
	if err == nil {
		fmt.Printf("%T, %v", pageIndex, pageIndex)
	} else {
		this.Ctx.WriteString(tools.Response(1, "input pi error!", []byte{'n'}))
		return
	}

	pageSize, err := strconv.Atoi(this.Ctx.Input.Param(":ps"))
	if err == nil {
		fmt.Printf("%T, %v", pageIndex, pageIndex)
	} else {
		this.Ctx.WriteString(tools.Response(1, "input ps error!", []byte{'n'}))
		return
	}

	o := orm.NewOrm()

	var users []models.User //查询的结果是集合的话，这里需要加上[]
	qs := o.QueryTable("user")

	_, err = qs.Offset(pageIndex * pageSize).Limit(pageSize).All(&users)
	if err == nil {
		// for _, user := range users {
		// 	fmt.Println(user)
		// }
		this.Ctx.WriteString(tools.Response(1, "success", users))
	} else {
		logs.Info(err)
		this.Ctx.WriteString(tools.Response(1, err.Error(), nil))
	}

}

func (this *UserController) UserGet() {
	this.Ctx.WriteString("shit!!")
}

func init() {
	beego.Router("/user/list/:pi/:ps", &UserController{}, "get:UserList")

	beego.Router("/all/:key", &UserController{}, "get:UserGet")
}
