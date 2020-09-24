package controllers

import (
<<<<<<< HEAD
	"github.com/astaxie/beego"
=======
	"beego1/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"io/ioutil"
>>>>>>> 35b2367... 第一次提交
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
<<<<<<< HEAD
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
=======
	user := c.Ctx.Input.Query("user")
	pwd := c.Ctx.Input.Query("pos")
	if user != "jy" || pwd != "123456"{
		c.Ctx.WriteString("出错l")
	}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "jy@gmail.com"
	c.TplName = "index.tpl"
}

//func (c *MainController)Post(){
//	name := c.Ctx.Request.FormValue("name")
//	age := c.Ctx.Request.FormValue("age")
//	//sex := c.Ctx.Request.FormValue("sex")
//	if name != "jy" && age != "120" {
//		c.Ctx.ResponseWriter.Write([]byte("数据校验失败"))
//	}else {
//		c.Ctx.WriteString("数据校验成功")
//	}
//}

func (c *MainController)Post ()  {
	var person models.Person
	databytes ,err := ioutil.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.WriteString("发生错误")
		return
	}
	err = json.Unmarshal(databytes,&person)
	if err != nil {
		c.Ctx.WriteString("发生错误")
		return
	}
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	fmt.Println(person.Sex)
	c.Ctx.WriteString("成功")
}
>>>>>>> 35b2367... 第一次提交
