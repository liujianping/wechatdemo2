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