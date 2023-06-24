package controllers

import (
	"fmt"
	"html/template"
	"mtservice/utils"
	"mtslashhelper/models"
	"mtslashhelper/sub"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type SubController struct {
	beego.Controller
}

func (c *SubController) Get() {
	mtstempid := c.Ctx.GetCookie("mtstempid")

	uid := utils.TryCookie(mtstempid)
	if uid == 0 {
		// 根据cookie没能获得用户说明要不cookie过期了，要不cookie是伪造的
		// 跳回登录页面
		c.Redirect("/login", 302)
	} else {
		// 根据email获得订阅的文章列表
		subList := models.GetSubListByUid(uid)
		c.Data["subList"] = subList
	}
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "manageSubs.html"
}

func (c *SubController) Post() {
	var result string
	mtstempid := c.Ctx.GetCookie("mtstempid")

	uid := utils.TryCookie(mtstempid)
	if uid == 0 {
		// 根据cookie没能获得用户说明要不cookie过期了，要不cookie是伪造的
		// 跳回登录页面
		result = "登陆状态已过期，请求失败，请重新登录"
	} else {
		// 根据ajax的请求，删除或者添加订阅
		op := c.GetString("op")
		if op == "del" {
			tidstr := c.GetStrings("tid")
			tid := make([]int, len(tidstr))
			for i := range tidstr {
				tid[i], _ = strconv.Atoi(tidstr[i])
			}
			err := models.DelSub(uid, tid)
			if err != nil {
				result = "取消订阅失败"
			} else {
				result = "取消订阅成功"
			}

		} else if op == "add" {
			// 参数：链接
			url := c.GetString("url")
			fmt.Println("要新增订阅的链接为", url)
			// 根据链接获得
			err := sub.AddSub(uid, url)
			if err != nil {
				// 链接错误或者后端获取网页失败了
				result = "新增订阅失败" + err.Error()
			} else {
				result = "新增订阅成功"
			}

		}
	}

	c.Ctx.WriteString(result)

}
