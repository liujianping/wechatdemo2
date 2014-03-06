# wechat demo 2

## 开发

````
package main 

//import "strings"
import "github.com/liujianping/wechat"


func DemoHandle(data []byte, back chan []byte){
	wechat.Info("recieve raw msg:\n", string(data))
	//! TODO: anything you want to do
	//! send empty string back
	back <- []byte("")
}

func main() {
	wechat.SetLogger("console", "")
	app := wechat.NewWeChatApp()
	app.SetConfig("ini", "demo.ini")	
	app.SetHandle(DemoHandle)
	app.Run()
}

````

## 编译

	go get github.com/liujianping/wechat
	go get github.com/liujianping/wechatdemo2
	cd src/github.com/liujianping/wechatdemo2
	go build

## 运行

	修改demo.ini配置之前，首先注册一个测试号：

	https://mp.weixin.qq.com/debug/cgi-bin/sandbox?t=sandbox/login

	设置好测试服务号配置后，相应设置demo.ini配置

	服务程序不能在本机运行，微信服务必须是具有域名服务的HTTP服务。
