package main

import (
	// "github.com/astaxie/beego"

	_ "mtservice/routers"

	"github.com/beego/beego/v2/server/web/context"

	beego "github.com/beego/beego/v2/server/web"
)

// 过滤器
var FilterUser = func(ctx *context.Context) {
	// 没有登陆的时候将所有路由(除了注册)重定向到登录页面
	// 检测uid是否合法
	mtstempid := ctx.GetCookie("mtstempid")

	if ctx.Request.RequestURI != "/signup" && ctx.Request.RequestURI != "/sendVeriCode" && (mtstempid == "" && ctx.Request.RequestURI != "/login") {
		ctx.Redirect(302, "/login")
	}

}

func main() {
	// beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	beego.Run()
	// 在beego.Run()上方添加如下三行配置
	beego.BConfig.WebConfig.EnableXSRF = true                                   // 开启xsrf防护
	beego.BConfig.WebConfig.XSRFKey = "61oEzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o" // beego默认的xsrfkey是：“beegoxsrf”，这里将key改为基于当前的key进行加密
	beego.BConfig.WebConfig.XSRFExpire = 3600                                   //过期时间，默认1小时，单位秒
	// utils.SetVeriCode("704081912@qq.com", "123456")
	// vc := utils.GetVariCode("704081912@qq.com")
	// fmt.Println(vc)
	// flag := utils.AlreadyVari("704081912@qq.com")
	// fmt.Println(flag)
	// runtime.GOMAXPROCS(1)
	// utils.NewFile(1, "a")
	// utils.NewFile(2, "b")
	// utils.NewFile(3, "c")

	// wg := sync.WaitGroup{}
	// for i := 1; i < 10; i++ {
	// 	// wg.Add(1)
	// 	go func(x int) {
	// 		// defer wg.Done()
	// 		fmt.Println(x)
	// 	}(i)
	// }
	// // wg.Wait()
	// time.Sleep(time.Duration(3) * time.Second)
	// go func() {
	// }()

}
