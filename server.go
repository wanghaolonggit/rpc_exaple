package main

import (
	"github.com/astaxie/beego"
	"io"
	"net/http"
	"net/rpc"
	"net"
)

type Panda int

func (this *Panda)GetInfo(argType int, replyType *int ) error  {
	beego.Info(argType)
	*replyType = 1 + argType
	return nil
}

func main() {
	http.HandleFunc("/panda", pandatext)

	pd := new(Panda)

	rpc.Register(pd)
	rpc.HandleHTTP()

	ln , err := net.Listen("tcp", "127.0.0.1:10086")
	if err != nil {
		beego.Info("网络连接失败")
	}
	beego.Info("正在监听10086")

	http.Serve(ln,nil)
}

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"panda")
}