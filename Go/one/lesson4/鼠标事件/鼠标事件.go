/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 14:42:39
 * @LastEditTime: 2021-09-29 17:00:37
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/鼠标事件.go
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
	//BUTTON_PRESS_MASK:鼠标按下，触发信号“button-press-event”
	//BUTTON_RELEASE_MASK：鼠标抬起，触发“button-release-event"
	//鼠标移动都是触发“motion-notify-event"
	win.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
	// button-press-event  鼠标按下时触发
	win.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
		//为左键，2为中间键，3为右键
		//fmt.Println("鼠标按下")
		flag := event.Button
		if flag == 1 {
			fmt.Println("1为左键")
		} else if flag == 2 {
			fmt.Println("2为中间键")
		} else if flag == 3 {
			fmt.Println("3为右键")
		}

		if event.Type == int(gdk.BUTTON_PRESS) {
			fmt.Println("单击")
		} else if event.Type == int(gdk.BUTTON2_PRESS) {
			fmt.Println("双击")
		}

		fmt.Println("相对于窗口:%v,%v\n", event.X, event.Y)
		fmt.Println("相对于屏幕:%v,%v\n", event.XRoot, event.YRoot)

	})

	win.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		fmt.Printf("===============相对于窗口:%v,%v\n", event.X, event.Y)
		fmt.Printf("===============相对于屏幕：%v,%v\n", event.XRoot, event.YRoot)
	})

	win.Show()

	gtk.Main()
}
