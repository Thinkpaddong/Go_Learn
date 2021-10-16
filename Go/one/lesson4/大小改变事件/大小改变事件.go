/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 17:13:48
 * @LastEditTime: 2021-09-29 20:49:58
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/大小改变事件.go
 */
package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
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
	//添加鼠标按下事件

	win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
	// key-press-event  鼠标按下时触发
	win.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventKey)(unsafe.Pointer(&arg))
		//为左键，2为中间键，3为右键
		//fmt.Println("鼠标按下")
		key := event.Keyval

		if key == gdk.KEY_A {
			fmt.Printf("aaaaaaaaaaaaaaaaaaa")
		}
	})

	win.ShowAll()

	gtk.Main()
}
