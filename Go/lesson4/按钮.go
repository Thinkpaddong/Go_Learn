/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-27 22:47:07
 * @LastEditTime: 2021-09-28 22:10:49
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/按钮.go
 */

package main

import (
	"fmt"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFormFile("test.glade")

	win := gtk.WindowFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	b1 := gtk.ButtonFromObject(builder("buttton1"))
	b2 := gtk.ButtonFromObject(builder("buttton2"))

	b1.SetLabel("*_*")
	b1.SetLabelFontSize(30)
	fmt.Println("b1 text = ", b1.GetLabel())
	b1.SetSensitive(false)

	var w, h int
	b2.GetSizeRequest(&w, &h)
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../Image/face.png", w-10, h-10, false)
	image := gtk.NewImageFromPixbuf()
	pixbuf.Unref() //释放资源

	b2.SetImage(image)
	b2.SetCanFocus(false)
	win.Show()

	gtk.Main()
}
