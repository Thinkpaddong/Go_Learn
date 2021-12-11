/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 22:26:35
 * @LastEditTime: 2021-10-17 22:28:39
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/有参数有返回值.go
 */
package main //必须

import "fmt"

//函数定义 可以理解返回值参数其实就是函数内部定义的一组局部变量
func MaxAndMin(a, b int) (max int, min int) {
	//var min int
	//var max int
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return //max, min //有返回值的函数，必须通过return返回  这句不一定是成立的
}

func main() {
	max, min := MaxAndMin(10, 20)
	fmt.Printf("max = %d, min = %d\n", max, min)

	//通过匿名变量丢弃某个返回值
	a, _ := MaxAndMin(10, 20)
	fmt.Printf("a = %d\n", a)
}
