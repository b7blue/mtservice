package controllers

import (
	"html/template"
	"mtservice/utils"

	"mtslashhelper/models"

	beego "github.com/beego/beego/v2/server/web"
)

// type user struct {
// 	Id       int    `form:"-"`
// 	Email    string `form:"email"`
// 	Password string `form:"pwd"`
// }

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	result := ""
	email, pw := c.GetString("email"), c.GetString("pw")

	// 检验密码正确性
	if uid, err := models.CheckUser(email, pw); err != nil {
		// 邮箱没注册
		result = "该邮箱没注册"
	} else {
		if uid == -1 {
			// 告知用户密码错误，重新输入密码
			result = "用户输入密码错误"

		} else {
			result = "OK"
			// 假如密码正确则成功登录，生成mtstempid，设置cookie
			mtstempid := utils.NewCookie(uid)
			if mtstempid == "" {
				result = "cookie设置出错，登录失败"
			} else {
				c.Ctx.SetCookie("mtstempid", mtstempid, 21600)
			}
			// c.SetSession("sid", int(1))
		}
	}
	c.Ctx.WriteString(result)

}
