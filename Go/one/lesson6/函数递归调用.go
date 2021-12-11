/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 22:33:13
 * @LastEditTime: 2021-10-17 22:33:14
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/函数递归调用.go
 */
package main //必须

import "fmt"

func test(a int) {
	if a == 1 { //函数终止调用的条件，非常重要
		fmt.Println("a = ", a)
		return //终止函数调用
	}

	//函数调用自身
	test(a - 1)

	fmt.Println("a = ", a)
}

func main() {
	test(3)
	fmt.Println("main")
}
