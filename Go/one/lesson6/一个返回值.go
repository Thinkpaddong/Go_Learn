/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 22:14:49
 * @LastEditTime: 2021-10-17 22:20:38
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/一个返回值.go
 */
package main

import (
	"fmt"
)

//无参有返回值：只有一个返回值
//有返回值的函数需要通过return中断函数，通过return返回
func myfun01() int {
	return 666
}

//相当于定义了一定局部变量
func myfun02() (result int) {
	result = 222 //不需要重新定义
	return
}

//给返回值起一个变量名，go推荐写法
//常用写法
func myfun03() (result int) {
	return 666
}

func main() {
	var a int
	a = myfun01()
	fmt.Println(a)

	b := myfun02()
	fmt.Println(b)

	c := myfun03()
	fmt.Println(c)

}
