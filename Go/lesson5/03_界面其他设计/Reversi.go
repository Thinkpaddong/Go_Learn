/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-30 11:25:55
 * @LastEditTime: 2021-09-30 14:06:19
 * @Description:
 * @FilePath: /lesson5/03_界面其他设计/Reversi.go
 */

package main

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

//控制结构体
type ChessWidge struct {
	window      *gtk.Window //窗口
	buttonMin   *gtk.Button //最小化按钮
	buttonClose *gtk.Button //关闭按钮
	labelBlack  *gtk.Label  //记录黑棋个数
	labelWhite  *gtk.Label  //记录白棋个数
	labelTime   *gtk.Label  //记录倒计时
	imageBlack  *gtk.Image  //提示该黑子落子
	imageWhite  *gtk.Image  //提示该白子落子
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
 * @description: 给按钮设置图标
 * @param {*}
 * @return {*}
 */
func ButtonSetImageFromFile(button *gtk.Button, filename string) {
	//获取按钮的大小
	w, h := 0, 0
	button.GetSizeRequest(&w, &h)

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)

	//创建image
	image := gtk.NewImageFromPixbuf(pixbuf)

	//释放批pixbuf
	pixbuf.Unref()

	//给按钮设置图片
	button.SetImage(image)

	//去掉按钮的焦距
	button.SetCanFocus(false)
}

/**
 * @description:
 * @param {*}
 * @return {*}
 */
func ImageSetPicFromFile(image *gtk.Image, filename string) {
	//获取image的大小
	w, h := 0, 0

	//创建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale(filename, w-10, h-10, false)

	//给image设置图片
	image.SetFromPixbuf(pixbuf)

	pixbuf.Unref()
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

	//获取按钮控件
	obj.buttonMin = gtk.ButtonFromObject(builder.GetObject("buttonMin"))
	obj.buttonClose = gtk.ButtonFromObject(builder.GetObject("buttonClose"))

	//给按钮设置图片
	ButtonSetImageFromFile(obj.buttonMin, "../image/min.png")
	ButtonSetImageFromFile(obj.buttonClose, "../image/close.png")

	//标签相关
	obj.labelBlack = gtk.LabelFromObject(builder.GetObject("labelBlack"))
	obj.labelWhite = gtk.LabelFromObject(builder.GetObject("labelWhite"))
	obj.labelTime = gtk.LabelFromObject(builder.GetObject("labelTime"))

	//设置字体
	obj.labelBlack.ModifyFontSize(50)
	obj.labelWhite.ModifyFontSize(50)
	obj.labelTime.ModifyFontSize(30)

	//设置内容
	obj.labelBlack.SetText("2")
	obj.labelWhite.SetText("2")
	obj.labelTime.SetText("20")

	//改变字体颜色
	obj.labelBlack.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelWhite.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("white"))
	obj.labelTime.ModifyBG(gtk.STATE_NORMAL, gdk.NewColor("white"))

	//image相关
	obj.imageBlack = gtk.ImageFromObject(builder.GetObject("imageBlack"))
	obj.imageWhite = gtk.ImageFromObject(builder.GetObject("imageWhite"))

	//设置图片
	ImageSetPicFromFile(obj.imageBlack, "../image/black.png")
	ImageSetPicFromFile(obj.imageWhite, "../image/white.png")
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

func PainEvent(ctx *glib.CallbackContext) {

	// 获取用户传递参数
	data := ctx.Data()
	obj, ok := data.(*Chessboard) //类型断言
	if !ok {
		fmt.Println("PainEvent Chessboard err")
		return
	}

	// 获取画家，设置绘图区域
	painter := obj.window.GetWindow().GetDrawable()
	gc := gdk.NewGC(painter)

	//新建pixbuf
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("../image/bg.jpg", obj.w, obj.h, false)

	//画图
	painter.DrawPixbuf(gc, pixbuf, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	//释放资源
	pixbuf.Unref()
}

/**
 * @description: 方法 事件、信号处理
 * @param {*}
 * @return {*}
 */
func (obj *Chessboard) HandleSignal() {
	//鼠标点击事件 button-press-event
	obj.window.Connect("button-press-event", MousePressEvent, obj)

	//鼠标移动事件 montion-notify-event
	obj.buttonClose.Clicked("motion-notify-event", MouseMoveEvent, obj)

	// 按钮的信号处理
	obj.buttonClose.Clicked(func() {
		gtk.MainQuit() //关闭窗口
	})

	obj.buttonMin.Clicked(func() {
		obj.window.Iconify() //最小化窗口
	})

	//绘图相关
	//大小改变事件 configure_event
	obj.window.Connect("configure_event", func() {
		obj.window.QueueDraw()
	})

	//绘图事件 expose-event
	obj.window.Connect("expose-event", PainEvent, obj)

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
