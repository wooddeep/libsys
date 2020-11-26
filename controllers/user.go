package controllers

import (
	//"encoding/json"

	"encoding/json"
	"fmt"
	"strconv"

	"libsys/models"

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
// 删除对象：n, err := o.Delete(&user)
// 更新对象：n, err:=qs.Filter("stu_id",5).Update(orm.Params{"hobby": "学习"})

// 获取url中的参数 var pageIndex = this.Ctx.Input.Param(":pi")

func (this *UserController) UserLst() {
	pageIndex, err := strconv.Atoi(this.Ctx.Input.Param(":pi"))
	if err == nil {
		fmt.Printf("%T, %v", pageIndex, pageIndex)
	} else {
		//this.Ctx.WriteString(tools.Response(1, err.Error(), []byte{'n'}))
		this.Data["json"] = map[string]interface{}{"data": []byte{}, "msg": err.Error(), "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	pageSize, err := strconv.Atoi(this.Ctx.Input.Param(":ps"))
	if err == nil {
		fmt.Printf("%T, %v", pageIndex, pageIndex)
	} else {
		//this.Ctx.WriteString(tools.Response(2, err.Error(), []byte{'n'}))
		this.Data["json"] = map[string]interface{}{"data": []byte{}, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
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
		//this.Ctx.WriteString(tools.Response(0, "success", users))
		this.Data["json"] = map[string]interface{}{"data": users, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		logs.Info(err)
		//this.Ctx.WriteString(tools.Response(3, err.Error(), nil))
		this.Data["json"] = map[string]interface{}{"data": []byte{}, "msg": err.Error(), "code": 1} // 设置返回值
		this.ServeJSON()
	}

}

// -d {"username": "lihan", "name": "李翰"}
func (this *UserController) UserAdd() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	var user models.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err == nil {
		var data = map[string]interface{}{"id": id}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": user, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *UserController) UserDel() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	var user models.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	n, err := o.Delete(&user)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": user, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *UserController) UserMod() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	var user models.User
	err := json.Unmarshal(data, &user)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()

	jsonStr := string(data)
	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)

	n, err := o.QueryTable("user").Filter("id", user.ID).Update(mapResult)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": user, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

// select book_id, book_state, b.name from book_shelf a left join book b on a.book_id = b.id where a.user_id = 16;

func init() {
	beego.Router("/user/lst/:pi/:ps", &UserController{}, "get:UserLst")
	beego.Router("/user/add", &UserController{}, "post:UserAdd")
	beego.Router("/user/del", &UserController{}, "delete:UserDel")
	beego.Router("/user/mod", &UserController{}, "put:UserMod")
}
