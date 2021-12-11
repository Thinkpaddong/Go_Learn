/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 22:30:00
 * @LastEditTime: 2021-10-17 22:30:01
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/普通函数的调用过程.go
 */
package main //必须

import "fmt"

func funcc(c int) {
	fmt.Println("c = ", c)
}

func funcb(b int) {

	funcc(b - 1)
	fmt.Println("b = ", b)
}

func funca(a int) {
	funcb(a - 1)
	fmt.Println("a = ", a)
}

func main() {
	funca(3) //函数调用
	fmt.Println("main")
}
