/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-25 19:49:17
 * @LastEditTime: 2021-09-27 23:43:01
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson2/测试服务器.go
 */
package main

import (
	"fmt"
	"net/http"
)

//服务端编写的业务逻辑处理程序
func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/go", myHandler)

	//在指定的地址进行监听，开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}
