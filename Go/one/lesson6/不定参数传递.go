/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 22:06:14
 * @LastEditTime: 2021-10-17 22:13:47
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/不定参数传递.go
 */
package main

import (
	"fmt"
)

func MyFunc(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}

func MyFunc02(tmp ...int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}

}
func main() {
	//var list []int = []int{1, 2, 3, 4}
	test(1, 2, 3, 4)
}

func test(args ...int) {
	//只想把后2个参数传递给另外一个函数使用
	MyFunc(args...)

	MyFunc02(args[:2]...) //args[0]~args[2]（不包括数字args[2]）， 传递过去
	MyFunc02(args[2:]...) //从args[2]开始(包括本身)，把后面所有元素传递过去
}
