package main

import (
	"io"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Println("dial failed, err:", err)
		return
	}

	defer conn.Close()
	// 构造 HTTP 协议的数据包,GET 表示请求你类型，http/1.1 表示http协议的版本号
	msg := "GET/HTTP/1.1\r\n"
	// 请求的主机
	msg += "Host:www.baidu.com\r\n"
	// Connection:close　表示短连接，即请求成功则关闭连接
	msg += "Connection:close\r\n"
	// 表示请求的数据包结束
	msg += "\r\n\r\n"

	/* io.WriteString 用来发送数据，接收两个参数，
			参数一为建立连接返回的连接对象 conn ，
			参数二为需要发送的数据(字符串类型),
		当然此处也可以使用　conn 对象直接发送，不过需要对发送的数据转为　byte 类型做相应处理
	*/
	_, e := io.WriteString(conn, msg)
	if err != nil {
		fmt.Println("write string failed:", e)
		return
	}

	buf := make([]byte, 4096)
	// 循环读取返回的内容
	for {
		// 接收百度服务器返回的数据
		count, er := conn.Read(buf[:])
		if er != nil || count == 0{
			break
		}
		fmt.Println(string(buf[:count]))
	}
}