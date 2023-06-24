package controllers

import (
	"fmt"
	"html/template"
	"mtservice/utils"
	"mtslashhelper/models"

	beego "github.com/beego/beego/v2/server/web"
)

type SignupController struct {
	beego.Controller
}

func (c *SignupController) Get() {
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "signup.html"
}

// 首先填了邮箱，，点击发送验证码的按钮
// 验证1邮箱格式正确 2没有注册过 3该邮箱十分钟之内没有发过验证码(在redis中根据email查找)
// 发送验证码，有效时间十分钟，存入redis
// 用户填写验证码
// 根据email从redis中取验证码,3种情况:1验证码已经过期 2验证码填写错误 3验证码填写正确

func (c *SignupController) Post() {
	// 从表单中获取新用户的信息
	email, pw, vericode := c.GetString("email"), c.GetString("pw"), c.GetString("vericode")
	result := ""

	// 根据注册的用户的Email生成独一无二的uid
	// 随机生成密钥
	// rand.Seed(time.Now().UnixNano())
	// key_int := rand.Intn(100000000)
	// key := strconv.Itoa(key_int)
	// cipher := goencrypt.NewDESCipher([]byte(key), []byte(""), goencrypt.ECBMode, goencrypt.Pkcs7, goencrypt.PrintBase64)
	// uid, err := cipher.DESEncrypt([]byte(u.Email))
	// if err != nil {
	// 	fmt.Println("生成uid出错", err)
	// }
	// u.Uid = uid
	// fmt.Printf("新用户：%+v\n", u)

	// 验证验证码对不对
	trueVeriCode := utils.GetVeriCode(email)
	if trueVeriCode == "" {
		result = "验证码已过期，请重新发送"
	} else {
		fmt.Println(trueVeriCode, vericode)
		if trueVeriCode == vericode {
			// 将新用户插入user表
			if uid, err := models.NewUser(email, pw); err != nil {
				// 假如新用户创建成功了，记得把redis里的验证码删了
				utils.DelVeriCode(email)
				// 给该用户创建msg表
				if err := models.NewMsgList(uid); err != nil {
					result = "数据库错误，注册失败"
					fmt.Println(err.Error())
				} else {
					result = "OK"
				}

			} else {
				result = "数据库错误，注册失败"
				fmt.Println(err.Error())
			}
		} else {
			result = "验证码错误"
		}
	}

	c.Ctx.WriteString(result)

}
