/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-25 17:17:05
 * @LastEditTime: 2021-09-30 00:10:13
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson2/http请求包格式/http请求包格式.go
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("listen err = ", err)
		return
	}
	defer listener.Close()

	conn, err1 := listener.Accept()
	if err != nil {
		fmt.Println("Accept err1 = ", err1)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)
	if n == 0 {
		fmt.Println("Read err2= ", err2)
		return
	}
	fmt.Printf("#%v#", string(buf[:n]))

}
