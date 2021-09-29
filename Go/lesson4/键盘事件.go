/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 17:01:46
 * @LastEditTime: 2021-09-29 17:22:14
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/键盘事件.go
 */
package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade上的控件
	builder := gtk.NewBuilder()
	builder.AddFromFile("./UI.glade")
	//获取glade上的控件
	win := gtk.WindowFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})
	//“configure_event" 窗口大小改变时触发
	win.Connect("configure_event", func() {
		var w, h int
		win.GetSize(&w, &h)
		fmt.Printf("w= %v, h = %v\n", w, h)
	})

	win.ShowAll()

	gtk.Main()
}
