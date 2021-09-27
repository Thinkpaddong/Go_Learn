/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-26 23:38:27
 * @LastEditTime: 2021-09-27 23:43:36
 * @Description: 网站已经做了修改了，所以现在已经不在适用下面这段代码的逻辑
 * @FilePath: /Test-for-github/Go/lesson3/段子爬虫.go
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}

	defer resp.Body.Close()

	buf := make([]byte, 4*1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}

		result += string(buf[:n])
	}
	return
}
func SpiderOneJoy(url string) (title, content string, err error) {
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("Http Get err ", err)
		return
	}

	//取关键信息
	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil {
		//fmt.Println("regeXp.MustCompile err")
		err = fmt.Errorf("%s", "regeXp.MustCompile err")
		return
	}

	//取内容
	tmpTitle := re1.FindAllStringSubmatch(result, 1)
	for _, data := range tmpTitle {
		title = data[1]
		// title=strings.Replace(title,"\r","",-1)
		// title=strings.Replace(title,"\n","",-1)
		// title=strings.Replace(title," ","",-1)
		title = strings.Replace(title, "\t", "", -1)
		break
	}

	//取关键信息
	re2 := regexp.MustCompile(`<div class="content="content-txt pt10">(?s:(.*?))</h1>`)
	if re2 == nil {
		//fmt.Println("regeXp.MustCompile err")
		err = fmt.Errorf("%s", "regeXp.MustCompile err2")
		return
	}

	//取内容
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		content = strings.Replace(content, "\r", "", -1)
		content = strings.Replace(content, "\n", "", -1)
		content = strings.Replace(content, "<br />", "", -1)
		content = strings.Replace(content, "\t", "", -1)
		break
	}
	return

}
func StoreJoyToFile(i int, fileTitle, fileContent []string) {

	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println("os.Create err= ", err)
		return
	}

	defer f.Close()

	n := len(fileTitle)
	for i := 0; i < n; i++ {
		f.WriteString(fileTitle[i] + "\n")
		f.WriteString(fileContent[i] + "\n")
		f.WriteString("\n=========================================================================\n")
	}

}
func SpiderPape(i int) {
	url := "https://www.pengfue.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页:%s\n", i, url)

	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("Http Get err ", err)
		return
	}

	//解析表达式是不是合法的，采用的是正则表达式去匹配内容
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re == nil {
		fmt.Println("regeXp.MustCompile err")
		return
	}

	joyUrls := re.FindAllStringSubmatch(result, -1)

	fileTitle := make([]string, 0)
	fileContent := make([]string, 0)

	for _, data := range joyUrls {
		//fmt.Println("url=",data[1])
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpinderOnerJoy err=", err)
			continue
		}

		//fmt.Println("title=#%v#",title)
		//fmt.Println("content=#%v#",content)
		fileTitle = append(fileTitle, title)
		fileContent = append(fileContent, content)

	}
}
func DoWork(start, end int) {
	fmt.Printf("准备爬取第%d页到%d页的网址", start, end)

	for i := start; i <= end; i++ {
		SpiderPape(i)
	}
}
func main() {
	var start, end int
	fmt.Printf("请输入起始页(>=1) :")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(>=1) :")
	fmt.Scan(&end)

	DoWork(start, end)
}
