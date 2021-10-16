/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-16 22:20:55
 * @LastEditTime: 2021-10-16 22:21:07
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/if_else.go
 */
package main //必须有一个main包

import "fmt"

func main() {
	s := "屌丝"

	//if和{就是条件，条件通常都是关系运算符
	if s == "王思聪" { //左括号和if在同一行
		fmt.Println("左手一个妹子，右手一个大妈")
	}

	//if支持1个初始化语句, 初始化语句和判断条件以分号分隔
	if a := 10; a == 10 { //条件为真，指向{}语句
		fmt.Println("a == 10")
	}
}

/*
package main //必须有一个main包

import "fmt"

func main() {
	//1
	a := 10
	if a == 10 {
		fmt.Println("a == 10")
	} else { //else后面没有条件
		fmt.Println("a != 10")
	}

	//2
	if a := 10; a == 10 {
		fmt.Println("a == 10")
	} else { //else后面没有条件
		fmt.Println("a != 10")
	}

	//3
	a = 8
	if a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}

	//4
	if a := 8; a == 10 {
		fmt.Println("a == 10")
	} else if a > 10 {
		fmt.Println("a > 10")
	} else if a < 10 {
		fmt.Println("a < 10")
	} else {
		fmt.Println("这是不可能的")
	}

}
*/
