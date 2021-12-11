/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-16 22:19:55
 * @LastEditTime: 2021-10-16 22:19:56
 * @Description:  类型别名也不是同一个类型
 * @FilePath: /Test-for-github/Go/one/lesson6/类型别名.go
 */
package main //必须有一个main包

import "fmt"

func main() {
	//给int64起一个别名叫bigint
	type bigint int64

	var a bigint // 等价于var a int64
	fmt.Printf("a type is %T\n", a)

	type (
		long int64
		char byte
	)

	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c\n", b, ch)
}
