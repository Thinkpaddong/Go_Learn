/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-25 20:13:18
 * @LastEditTime: 2021-09-27 23:43:14
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson2/http服务器.go
 */

package main

import (
	"net/http"
)

/**
 * @description: 用于获取请求和响应的
 * @param {http.ResponseWriter}
 * @param {*http.Request} req   接口
 * @return {*}
 */
func HandConn(w http.ResponseWriter, req *http.Request) {
	fmt.Pintln("r.Method:", r.Method)
	fmt.Pintln("r.URL:", r.URL)
	fmt.Pintln("r.Header:", r.Header)
	fmt.Pintln("r.Body:", r.Body)
	w.Write([]byte("hello.go")) //给客户端回复数据
}
func main() {
	http.HandleFunc("/", HandConn)

	http.ListenAndServe(":8000", nil)

}
