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

type BookController struct {
	beego.Controller
}

func (this *BookController) BookLst() {
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

	var books []models.Book //查询的结果是集合的话，这里需要加上[]
	qs := o.QueryTable("book")

	_, err = qs.Offset(pageIndex * pageSize).Limit(pageSize).All(&books)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"data": books, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		logs.Info(err)
		this.Data["json"] = map[string]interface{}{"data": []byte{}, "msg": err.Error(), "code": 1} // 设置返回值
		this.ServeJSON()
	}

}

func (this *BookController) BookAdd() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到book对象中
	var book models.Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(&book)
	if err == nil {
		var data = map[string]interface{}{"id": id}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": book, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *BookController) BookDel() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到book对象中
	var book models.Book
	err := json.Unmarshal(data, &book)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	n, err := o.Delete(&book)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": book, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *BookController) BookMod() {
	data := this.Ctx.Input.RequestBody
	//json数据封装到book对象中
	var book models.Book
	err := json.Unmarshal(data, &book)
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

	n, err := o.QueryTable("book").Filter("id", book.ID).Update(mapResult)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": book, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func init() {
	beego.Router("/book/lst/:pi/:ps", &BookController{}, "get:BookLst")
	beego.Router("/book/add", &BookController{}, "post:BookAdd")
	beego.Router("/book/del", &BookController{}, "delete:BookDel")
	beego.Router("/book/mod", &BookController{}, "put:BookMod")
}
