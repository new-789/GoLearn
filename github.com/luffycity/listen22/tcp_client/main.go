package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"strings"
)

func getInput() (data string, err error) {
	inputReader:= bufio.NewReader(os.Stdin)
	for {
		// 此处需要注意使用　ReadString　方法去掉空格时 \n 必须使用单引号
		data, err = inputReader.ReadString('\n')
		// 去掉内容结尾的空格
		data = strings.TrimSpace(data)
		if data == "Q" || data == "q" {
			return
		}
		// 去掉输入内容结尾的 \n 
		return
	}
	return
}

func send(conn net.Conn, strByte []byte) (err error) {
	// 发送数据给服务端
	_, err = conn.Write(strByte)
	return
}

func main() {
	// 建立连接
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil{
		fmt.Println("connect failed, err:", err)
		return
	}
	for {
		// 获取用户输入的数据
		data, err := getInput()
		if err != nil{
			fmt.Println("read from console failed, err:", err)
			break
		}
		if data == "Q" || data == "q" {
			break
		}
		// 发送数据
		er := send(conn, []byte(data))
		if err != nil{
			fmt.Println("send data failed,err:", er)
			break
		}
	}
}