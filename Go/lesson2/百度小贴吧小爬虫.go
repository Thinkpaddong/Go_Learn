/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-25 20:36:57
 * @LastEditTime: 2021-09-27 23:42:54
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson2/百度小贴吧小爬虫.go
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string) (res string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	buf := make([]byte, 1024*4)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err: ", err)
			break
		}
		res += string(buf[:n])
	}

	return
}
func DoWork(start, end int) {
	fmt.Println("正在爬取%d 到%d的页面\n", start, end)
	for i := start; i <= end; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)

		fmt.Println("url: ", url)
		res, err := HttpGet(url)
		if err != nil {
			fmt.Println("httpGet err = ", err)
			continue
		}

		fileName := strconv.Itoa((i-1)*50) + ".html"
		f, err1 := os.Create(fileName)
		if err1 != nil {
			fmt.Println("os.Create err1: ", err1)
			continue
		}

		f.WriteString(res)

		f.Close()

	}
}
func main() {
	var start, end int
	fmt.Println("请输入起始页(>=1) :")
	fmt.Scan(&start)
	fmt.Println("请输入终止页(>=1) :")
	fmt.Scan(&end)

	DoWork(start, end)

}
