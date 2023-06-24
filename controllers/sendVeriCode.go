package controllers

import (
	"fmt"
	"math/rand"
	"mtservice/utils"
	"mtslashhelper/models"
	"regexp"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

type SendVeriCodeController struct {
	beego.Controller
}

// 首先填了邮箱，，点击发送验证码的按钮
// 验证1邮箱格式正确 2没有注册过 3该邮箱十分钟之内没有发过验证码(在redis中根据email查找)
// 发送验证码，有效时间十分钟，存入redis

func (c *SendVeriCodeController) Get() {
	c.Ctx.WriteString("hello")
}

func (c *SendVeriCodeController) Post() {
	var result string
	// 用户点击了发送验证码的按钮
	// 检查邮箱格式
	email := c.GetString("email")
	emailreg := regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)

	if emailreg.MatchString(email) {
		// 检查邮箱是否被注册过
		if models.IsNewUser(email) {
			// 检查该邮箱十分钟之内没有发过验证码(在redis中根据email查找)
			if utils.AlreadyVeri(email) {
				result = "上一个验证码还未过期,请勿重复发送"

			} else {
				// 生成验证码,存入redis,发送
				rand.Seed(time.Now().UnixNano())
				code := fmt.Sprintf("%06d", rand.Intn(1000000))

				err := utils.SetVeriCode(email, code)
				if err == nil {
					err = utils.SendEmail(email, code)
					result = "验证码已发送"
				}

				if err != nil {
					result = "验证码发送失败，请重试"
				}

			}

		} else {
			result = "该邮箱已被注册过,请换一个"
		}
	} else {
		result = "邮箱格式不正确"
	}

	c.Ctx.WriteString(result)
}
