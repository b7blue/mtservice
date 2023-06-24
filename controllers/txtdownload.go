package controllers

import (
	"encoding/json"
	"html/template"
	"mtservice/utils"
	conn "mtslashhelper/connector"
	mtdownload "mtslashhelper/download"

	beego "github.com/beego/beego/v2/server/web"
)

type TXTController struct {
	beego.Controller
}

func (c *TXTController) Get() {
	// c.Data["IsSubmit"] = false
	c.Data["xsrfdata"] = template.HTML(c.XSRFFormHTML())
	c.TplName = "txtdownload.html"
}

func (c *TXTController) Post() {
	href := c.GetString("href")
	var tid int
	var err error

	var result struct {
		Err      string `json:"err"`
		Filename string `json:"filename"`
	}

	// 用链接获取tid
	tid, err = conn.GetTid(href)
	if err != nil {
		result.Err = err.Error()
	} else {

		// 缓存中没有文件
		if filename := utils.GetFilebyTid(tid); filename == "" {
			// 运行咱们的mtslash_download，下载txt到服务器上
			filename, err = mtdownload.GetTXT(tid)
			if err != nil {
				result.Err = err.Error()
			} else {
				// 然后填文件名字
				result.Filename = filename
				// 在缓存中添加一条记录
				utils.NewFile(tid, filename)
			}
			// 缓存中有文件的话，就不用重新下载文件了
		} else {
			result.Filename = filename
		}
	}
	b, _ := json.Marshal(result)
	c.Ctx.WriteString(string(b))

}
