package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
//默认请求为get
func (c *MainController) Get() {
	c.Data["Website"] = "www.baidu.com"
	c.Data["Email"] = "1454050489@qq.com"
	c.TplName = "index.tpl"
}
