/*
 * @Author: Thinkpaddong
 * @Date: 2021-09-25 14:58:29
 * @LastEditTime: 2021-09-27 23:42:41
 * @Description:
 * @FilePath: /Test-for-github/Go/lesson1/main.go
 */
package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var messaage = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}
func HandleConn(conn net.Conn) {

	cliAddr := conn.RemoteAddr().String()

	cli := Client{make(chan string), cliAddr, cliAddr}

	onlineMap[cliAddr] = cli

	go WriteMsgToClient(cli, conn) //往channel写数据进去

	//messaage<-"["+cli.Addr+"]"+cli.Name+": login"

	messaage <- MakeMsg(cli, "login")

	cli.C <- MakeMsg(cli, "I am here")

	isQuit := make(chan bool) //是否主动退出

	hasData := make(chan bool)

	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}

			msg := string(buf[:n-1]) //强制转string  nc多一个换行

			//messaage<- MakeMsg(cli,msg)
			if len(msg) == 3 && msg == "who" {
				//给同用户发送当前全部成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) >= 8 && msg[:6] == "rename" {
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				messaage <- MakeMsg(cli, msg)
			}
			hasData <- true
		}

	}() //别忘记了
	for {
		//通过select检测channel的流动
		select {
		case <-isQuit:
			delete(onlineMap, cliAddr)
			messaage <- MakeMsg(cli, "login out") //广播谁退出了

			return
		case <-hasData:

		case <-time.After(30 * time.Second):
			delete(onlineMap, cliAddr)
			messaage <- MakeMsg(cli, "time out leave out")
			return
		}
	}
}

func Manager() {

	onlineMap = make(map[string]Client)
	for {
		msg := <-messaage

		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	defer listener.Close()

	go Manager()
	//主协程 循环阻塞用户等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.accpet err= ", err)
			continue
		}

		go HandleConn(conn)

	}

}
