package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type DownloadController struct {
	beego.Controller
}

func (c *DownloadController) Get() {
	filename := c.GetString("filename")
	c.Ctx.Output.Download("txt/"+filename+".txt", `"`+filename+`".txt`)
}
