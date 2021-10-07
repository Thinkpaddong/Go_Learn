/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 17:13:48
 * @LastEditTime: 2021-09-30 00:11:01
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/绘图事件/绘图事件.go
 */
package main

import (
	"os"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	//初始化
	gtk.Init(&os.Args)
	//加载glade上的控件
	builder := gtk.NewBuilder()
	builder.AddFromFile("./ui.glade")
	//获取glade上的控件
	win := gtk.WindowFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	var w, h int
	//窗口大小改变的时触发  "configure_event"
	win.Connect("configure_event", func() {
		win.GetSize(&w, &h)
		win.QueueDraw()
	})
	//窗口绘图
	win.SetAppPaintable(true)

	x := 0

	win.Connect("expose-event", func() {
		//设置画家，指定绘图区域
		painter := win.GetWindow().GetDrawable()
		gc := gdk.NewGC(painter)
		//创建图片资源
		bg, _ := gdkpixbuf.NewPixbufFromFileAtScale("../image/bk.jpg", w, h, false)
		face, _ := gdkpixbuf.NewPixbufFromFileAtScale("../image/face.jpg", 80, 80, false)
		painter.DrawPixbuf(gc, bg, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		painter.DrawPixbuf(gc, face, 0, 0, x, 150, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)
		//释放图片资源
		bg.Unref()
		face.Unref()
	})

	// button := gtk.WindowFromObject(builder.GetObject("button1"))

	// button.Clicked(func() {
	// 	x += 50
	// 	if x >= w {
	// 		x = 0
	// 	}
	// 	//刷图，整个窗口区域刷图
	// 	win.QueueDraw()
	// })

	win.ShowAll()

	gtk.Main()
}
