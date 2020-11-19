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

type PunchController struct {
	beego.Controller
}

func (this *PunchController) Lst() {
	const tabelName = "punch"
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

	var list []models.Punch //查询的结果是集合的话，这里需要加上[]
	qs := o.QueryTable(tabelName)

	_, err = qs.Offset(pageIndex * pageSize).Limit(pageSize).All(&list)
	if err == nil {
		this.Data["json"] = map[string]interface{}{"data": list, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		logs.Info(err)
		//this.Ctx.WriteString(tools.Response(3, err.Error(), nil))
		this.Data["json"] = map[string]interface{}{"data": []byte{}, "msg": err.Error(), "code": 1} // 设置返回值
		this.ServeJSON()
	}

}

func (this *PunchController) Add() {
	data := this.Ctx.Input.RequestBody
	var object models.Punch
	err := json.Unmarshal(data, &object)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	id, err := o.Insert(&object)
	if err == nil {
		var data = map[string]interface{}{"id": id}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": object, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *PunchController) Del() {
	data := this.Ctx.Input.RequestBody
	var object models.Punch
	err := json.Unmarshal(data, &object)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"data": string(data), "msg": "输入数据错误", "code": 1} // 设置返回值
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	n, err := o.Delete(&object)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": object, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func (this *PunchController) Mod() {
	const tabelName = "punch"
	data := this.Ctx.Input.RequestBody
	var object models.Punch
	err := json.Unmarshal(data, &object)
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

	n, err := o.QueryTable(tabelName).Filter("id", object.ID).Update(mapResult)
	if err == nil {
		var data = map[string]interface{}{"num": n}
		this.Data["json"] = map[string]interface{}{"data": data, "msg": "success", "code": 0} // 设置返回值
		this.ServeJSON()
	} else {
		this.Data["json"] = map[string]interface{}{"data": object, "msg": err.Error(), "code": 2} // 设置返回值
		this.ServeJSON()
	}
}

func init() {

	const rootPath = "punch"
	beego.Router(fmt.Sprintf("/%s/lst/:pi/:ps", rootPath), &PunchController{}, "get:Lst")
	beego.Router(fmt.Sprintf("/%s/add", rootPath), &PunchController{}, "post:Add")
	beego.Router(fmt.Sprintf("/%s/del", rootPath), &PunchController{}, "delete:Del")
	beego.Router(fmt.Sprintf("/%s/mod", rootPath), &PunchController{}, "put:Mod")
}
