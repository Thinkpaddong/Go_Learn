/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 00:15:10
 * @LastEditTime: 2021-09-29 00:28:46
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/定时器.go
 */
package main

import (
	"os"
	"strconv"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("./UI.glade")

	win := gtk.WindowFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})

	buttonStart := gtk.ButtonFromObject(builder.GetObject("buttonStart"))
	buttonEnd := gtk.ButtonFromObject(builder.GetObject("buttonEnd"))
	//获取标签
	label := gtk.LabelFromObject((builder.GetObject("label")))
	label.SetText("0")       //设置内容
	label.ModifyFontSize(30) //设置字体大小

	buttonEnd.SetSensitive(false) //按钮变灰
	num := 0
	id := 0
	buttonStart.Clicked(func() {
		id = glib.TimeoutAdd(500, func() bool {
			num++
			label.SetText(strconv.Itoa(num)) //设置内容
			return true
		})

		buttonEnd.SetSensitive(true)
		buttonStart.SetSensitive(false)
	})

	buttonEnd.Clicked(func() {
		glib.TimeoutRemove(id)

		buttonEnd.SetSensitive(false)
		buttonStart.SetSensitive(true)
	})

	win.ShowAll()

	gtk.Main()
}
