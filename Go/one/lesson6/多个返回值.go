/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 22:21:49
 * @LastEditTime: 2021-10-17 22:25:37
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/多个返回值.go
 */
package main

import "fmt"

func myfun01() (int, int, int) {
	return 11, 22, 33
}

func myfun03() (a int, b int, c int) {

	return 111, 222, 333
}

func myfun02() (a, b, c int) {
	a, b, c = 11, 22, 33
	return
}
func main() {
	a, b, c := myfun01()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	a, b, c = myfun02()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
	a, b, c = myfun03()
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}
