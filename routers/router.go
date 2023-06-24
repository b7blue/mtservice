package routers

import (
	"mtservice/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signup", &controllers.SignupController{})
	beego.Router("/txtdownload", &controllers.TXTController{})
	beego.Router("/gettxt", &controllers.DownloadController{})
	beego.Router("/sendVeriCode", &controllers.SendVeriCodeController{})
	beego.Router("/msgBox", &controllers.MsgController{})
	beego.Router("/manageSubs", &controllers.SubController{})
}
