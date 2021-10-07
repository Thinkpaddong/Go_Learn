/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 20:10:14
 * @LastEditTime: 2021-09-29 20:31:59
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/无边框窗口移动.go
 */
package main

import (
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

	win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	x, y := 0, 0
	//button-press-event 鼠标按下时触发
	win.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取鼠标按下属性结构体变量，系统内部变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		x, y = int(event.X), int(event.Y)

	})
	//“motion-notify-event" 按住鼠标移动时触发
	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		win.Move(int(event.XRoot)-x, int(event.YRoot)-y)
	})

	win.ShowAll()

	gtk.Main()
}
