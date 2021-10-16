/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-16 22:03:01
 * @LastEditTime: 2021-10-16 22:03:02
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/bool类型.go
 */
package main //必须有一个main包

import "fmt"

func main() {
	//1、声明变量, 没有初始化，零值（初始值）为false
	var a bool
	fmt.Println("a0 = ", a)

	a = true
	fmt.Println("a = ", a)

	//2、自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c = ", c)
}
