package main

import (
	"github.com/astaxie/beego"
	"io"
	"net/http"
	"net/rpc"
	"net"
)

type Panda int
type  S string

func (this *Panda)GetInfo(argType int, replyType *int ) error  {
	beego.Info(argType)
	*replyType = 1 + argType
	return nil
}

func (this *S) GetStr(s string , reply *string) error {
	beego.Info(s)
	*reply = s + "string Test"
	return nil
}

func main() {
	http.HandleFunc("/panda", pandatext)

	pd := new(Panda)
	ss := new(S)

	rpc.Register(pd)
	rpc.Register(ss)
	rpc.HandleHTTP()

	ln , err := net.Listen("tcp", "127.0.0.1:10087")
	if err != nil {
		beego.Info(err)
		beego.Info("网络连接失败")
	}
	beego.Info("正在监听10086")

	http.Serve(ln,nil)
}

func pandatext(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"panda")
}