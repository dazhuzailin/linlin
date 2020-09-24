package routers

import (
<<<<<<< HEAD
	"beegoto/controllers"
=======
	"beego1/controllers"
>>>>>>> 35b2367... 第一次提交
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
