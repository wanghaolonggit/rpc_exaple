package main

import (
	"github.com/astaxie/beego"
	"net/rpc"
)

func main() {
	cli, err := rpc.DialHTTP("tcp", "127.0.0.1:10086")

	if err != nil {
		beego.Info("链接不上")
	}

	var val int

	err  = cli.Call("Panda.GetInfo", 123, &val)


	if err != nil {
		beego.Info(err)
		beego.Info("call 失败")
	}
	beego.Info("返回结果",val)
}