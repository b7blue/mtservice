package controllers

import (
	"encoding/json"
	"html/template"
	"mtservice/utils"
	"mtslashhelper/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MsgController struct {
	beego.Controller
}

func (c *MsgController) Get() {
	mtstempid := c.Ctx.GetCookie("mtstempid")

	uid := utils.TryCookie(mtstempid)
	if uid == 0 {
		// 根据cookie没能获得用户说明要不cookie过期了，要不cookie是伪造的
		// 跳回登录页面
		c.Data["toLogin"] = true
	} else {
		c.Data["toLogin"] = false

		// 根据email获得订阅列表
		msgList := models.GetMsgList(uid)
		c.Data["msgList"] = msgList
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "msgBox.html"
}

func (c *MsgController) Post() {
	var result string
	mtstempid := c.Ctx.GetCookie("mtstempid")

	uid := utils.TryCookie(mtstempid)
	if uid == 0 {
		// 根据cookie没能获得用户说明要不cookie过期了，要不cookie是伪造的
		// 跳回登录页面
		result = "登陆状态已过期，请求失败，请重新登录"
	} else {

		// 删除消息
		idstr := c.GetString("id")
		id := make([]int, len(idstr))
		json.Unmarshal([]byte(idstr), &id)

		err := models.DelMsg(uid, id)
		if err != nil {
			result = "删除消息失败"
		} else {
			result = "删除消息成功"
		}

	}

	c.Ctx.WriteString(result)
}
