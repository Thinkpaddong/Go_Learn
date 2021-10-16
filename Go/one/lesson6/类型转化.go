/*
 * @Author: Thinkpaddong
 * @Date: 2021-10-16 22:10:17
 * @LastEditTime: 2021-10-16 22:19:19
 * @Description:
 * @FilePath: /Test-for-github/Go/one/lesson6/类型转化.go
 */
package main //必须有一个main包

import "fmt"

func main() {

	//这种不能转换的类型，叫不兼容类型
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)

	//bool类型不能转换为int
	//fmt.Printf("flag = %d\n", int(flag))

	//0就是假，非0就是真
	//整型也不能转换为bool
	//flag = bool(1)

	var ch byte
	ch = 'a' //字符类型本质上就是整型
	var t int
	t = int(ch) //类型转换，把ch的值取出来后，转成int再给t赋值
	fmt.Println("t = ", t)

}

/*
使用type其实也是不同的类型
s, ok := x.(T) 不会出panic
*/
/*switch i := x.(type) {
case nil:
	printString("x is nil")                // type of i is type of x (interface{})
case int:
	printInt(i)                            // type of i is int
case float64:
	printFloat64(i)                        // type of i is float64
case func(int) float64:
	printFunction(i)                       // type of i is func(int) float64
case bool, string:
	printString("type is bool or string")  // type of i is type of x (interface{})
default:
	printString("don't know the type")     // type of i is type of x (interface{})
}
*/
