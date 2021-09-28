/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 00:03:37
 * @LastEditTime: 2021-09-29 00:14:22
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson4/对话框.go
 */
package main

import (
	"fmt"
	"os"

	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("./UI.glade") //glade文件

	win := gtk.WindowFromObject(builder.GetObject("window1"))

	win.Connect("destroy", func() {
		gtk.MainQuit() //关闭程序
	})
	//获取按钮
	b1 := gtk.ButtonFromObject(builder.GetObject("buttton1"))
	b2 := gtk.ButtonFromObject(builder.GetObject("buttton2"))

	b1.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_QUESTION,
			gtk.BUTTONS_YES_NO,
			"这是消息对话框")
		dialog.SetTitle("问题对话框")

		ret := dialog.Run()
		if ret == gtk.RESPONSE_YES {
			fmt.Print(("yes"))
		} else if ret == gtk.RESPONSE_NO {
			fmt.Println("no")
		} else {
			fmt.Println("close")
		}
		dialog.Destroy()
	})

	b2.Clicked(func() {
		dialog := gtk.NewMessageDialog(
			win,
			gtk.DIALOG_MODAL,
			gtk.MESSAGE_INFO,
			gtk.BUTTONS_OK,
			"这是消息对话框")
		dialog.SetTitle("消息对话框")

		dialog.Run()
		dialog.Destroy()
	})

	win.Show()

	gtk.Main()
}
