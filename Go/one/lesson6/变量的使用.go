/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-16 21:28:08
 * @LastEditTime: 2021-10-16 21:35:22
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/变量的使用.go
 */
package main

import "fmt"

func main() {
	//变量，程序运行期间，可以改变的量
	//1、声明格式   var 变量名 类型, 变量声明了，必须要使用
	//2、只是声明没有初始化的变量，默认值为0
	//3、同一个{}里，声明的变量名是唯一的
	var a int
	fmt.Println("a=", a)
	//4、可以同时声明多个变量
	//	var b, c int
	//	a = 10
	//2、变量的初始化，声明变量时，同时赋值 声明的变量在Go语言中一定要使用不然是会报错的
	var b int = 10
	b = 20
	println(b)
	//3、自动推导类型，必须初始化，通过初始化的值确定类型(常用)  这个与前面是相互矛盾的
	c := 20
	fmt.Printf("c type is %T\n", c)
}
