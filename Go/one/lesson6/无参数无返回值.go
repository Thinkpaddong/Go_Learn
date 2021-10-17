/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 21:50:55
 * @LastEditTime: 2021-10-17 21:50:57
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/无参数无返回值.go
 */
package main //必须

import "fmt"

func main() {
	//无参无返回值函数的调用： 函数名()
	MyFunc()
}

//无参无返回值函数的定义
func MyFunc() {
	a := 666
	fmt.Println("a = ", a)
}
