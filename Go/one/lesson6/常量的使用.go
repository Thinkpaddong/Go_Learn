/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-16 21:44:50
 * @LastEditTime: 2021-10-16 21:47:38
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/常量的使用.go
 */
package main

import "fmt"

func main() {
	//err, 常量不允许修改
	//变量：程序运行期间，可以改变的量， 变量声明需要var
	//常量：程序运行期间，不可以改变的量，常量声明需要const
	const a int = 10
	fmt.Println("a=", a)

	const b = 11.2 //没有使用:=
	fmt.Printf("b type is %f\n", b)
	fmt.Println(b)
}
