package routers

import (
	"hellobeego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//router:路由
	//第一个参数为请求的路径，请求的路径必须保持唯一
    beego.Router("/login", &controllers.MainController{})
}
