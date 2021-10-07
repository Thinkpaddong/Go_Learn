/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-25 20:13:18
 * @LastEditTime: 2021-09-30 00:08:22
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson2/http服务器/http服务器.go
 */

package main

import (
	"fmt"
	"net/http"
)

/**
 * @description: 用于获取请求和响应的
 * @param {http.ResponseWriter}
 * @param {*http.Request} req   接口
 * @return {*}
 */
func HandConn(w http.ResponseWriter, req *http.Request) {
	fmt.Println("r.Method:", req.Method)
	fmt.Println("r.URL:", req.URL)
	fmt.Println("r.Header:", req.Header)
	fmt.Println("r.Body:", req.Body)
	w.Write([]byte("hello.go")) //给客户端回复数据
}
func main() {
	http.HandleFunc("/", HandConn)

	http.ListenAndServe(":8000", nil)

}
