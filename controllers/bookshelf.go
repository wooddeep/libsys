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

type BookShelfController struct {
	beego.Controller
}

func (this *BookShelfController) BookShelfLst() {
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

	var shelfs []models.BookShelf //查询的结果是集合的话，这里需要加上[]
	qs := o.QueryTable("book_shelf")

	_, err = qs.Offset(pageIndex * pageSize).Limit(pageSize).All(&shelfs)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"data": shelfs, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		logs.Info(err)
		//this.Ctx.WriteString(tools.Response(3, err.Error(), nil))
		this.Data["json"] = map[string]interface{}{"data": []byte{}, "msg": err.Error(), "code": 1} // 设置返回值
		this.ServeJSON()
	}

}

func (this *BookShelfController) BookShelfAdd() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	var shelf models.BookShelf
	err := json.Unmarshal(data, &shelf)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(&shelf)
	if err == nil {
		var data = map[string]interface{}{"id": id}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": shelf, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *BookShelfController) BookShelfDel() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到user对象中
	var shelf models.BookShelf
	err := json.Unmarshal(data, &shelf)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	n, err := o.Delete(&shelf)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": shelf, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *BookShelfController) BookShelfMod() {
	data := this.Ctx.Input.RequestBody
	var shelf models.BookShelf
	err := json.Unmarshal(data, &shelf)
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

	n, err := o.QueryTable("book_shelf").Filter("id", shelf.ID).Update(mapResult)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": shelf, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func init() {
	beego.Router("/bookshelf/lst/:pi/:ps", &BookShelfController{}, "get:BookShelfLst")
	beego.Router("/bookshelf/add", &BookShelfController{}, "post:BookShelfAdd")
	beego.Router("/bookshelf/del", &BookShelfController{}, "delete:BookShelfDel")
	beego.Router("/bookshelf/mod", &BookShelfController{}, "put:BookShelfMod")
}
