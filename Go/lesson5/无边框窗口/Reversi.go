/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-29 20:45:00
 * @LastEditTime: 2021-09-29 23:43:57
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson5/无边框窗口/Reversi.go
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

//控制结构体
type ChessWidge struct {
	window *gtk.Window
}

//控制属性结构体
type ChessInfo struct {
	w, h int
	x, y int
}

//黑白棋结构体
type Chessboard struct {
	ChessWidge
	ChessInfo
}

/**
 * @description: 方法 创建控件，设置控件属性
 * @param {*}
 * @return {*}
 */
func (obj *Chessboard) CreateWindow() {
	//加载glade上的控件
	builder := gtk.NewBuilder()
	builder.AddFromFile("ui.glade")

	//窗口相关
	obj.window = gtk.WindowFromObject(builder.GetObject("window1")) //获取控件
	obj.window.SetAppPaintable(true)                                //可以绘图
	obj.window.SetPosition(gtk.WIN_POS_CENTER)                      //居中显示
	obj.w, obj.h = 800, 480                                         //窗口的宽度和高度
	obj.window.SetSizeRequest(800, 480)                             //设置窗口的宽高
	obj.window.SetDecorated(false)                                  //去边框

	//设置事件，让窗口可以捕获鼠标点击和移动
	obj.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
}

/**
 * @description: 鼠标点击事件函数
 * @param {*}
 * @return {*}
 */
func MousePressEvent(ctx *glib.CallbackContext) {
	// 获取鼠标按下属性结构体变量，系统内部的变量，不是用户传参变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))
	// 获取用户传递参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if !ok {
		fmt.Println("MousePressEvent Chessboard err")
		return
	}

	// 保存点击的x,y坐标
	obj.x, obj.y = int(event.X), int(event.Y)
	fmt.Println("x= ", obj.x, ", y =", obj.y)

}

/**
 * @description: 鼠标移动事件函数
 * @param {*}
 * @return {*}
 */
func MouseMoveEvent(ctx *glib.CallbackContext) {

	// 获取用户传递参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if !ok {
		fmt.Println("MouseMoveEvent Chessboard err")
		return
	}

	// 获取鼠标按下属性结构体变量，系统内部的变量，不是用户传参变量
	arg := ctx.Args(0)
	event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

	x, y := int(event.XRoot)-obj.x, int(event.YRoot)-obj.y
	obj.window.Move(x, y) //移动窗口

}

/**
 * @description: 方法 事件、信号处理
 * @param {*}
 * @return {*}
 */
func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件 button-press-event
	obj.window.Connect("button-press-event", MousePressEvent, obj)

	//鼠标移动事件
}

func main() {
	//初始化
	gtk.Init(&os.Args)

	//创建结构体变量
	var obj Chessboard

	//创建控件，设置控件属性
	obj.CreateWindow()
	//显示控件
	obj.window.Show()

	gtk.Main()
}
