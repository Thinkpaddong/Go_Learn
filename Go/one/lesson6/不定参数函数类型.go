/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-17 21:54:52
 * @LastEditTime: 2021-10-17 22:06:31
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/不定参数函数类型.go
 */
package main

import (
	"fmt"
)

//...int类型这样的类型， ...type不定参数类型
//注意：不定参数，一定（只能）放在形参中的最后一个参数
func MyFunc02(test ...int) { //这个名字随意取得  得到的是一个切片类型
	fmt.Println("len(test) = ", len(test))
	for i := 0; i < len(test); i++ {
		fmt.Printf("test[%d]=%d\n", i, test[i])
	}

	fmt.Println("===================================")
	for i, data := range test {
		fmt.Printf("test[%d]=%d\n", i, data)
	}

}
func main() {
	MyFunc02(66, 1111, 222)
}

//固定参数一定要传参，不定参数根据需求传递
func MyFunc03(a int, args ...int) {
}

//注意：不定参数，一定（只能）放在形参中的最后一个参数
//func MyFunc04(args ...int, a int) {
//}

func main() {
	MyFunc03(111, 1, 2, 3)
}
