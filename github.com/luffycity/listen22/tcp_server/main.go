package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建监听类型以及监听端口和　ＩＰ
	listen, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil{
		fmt.Println("listen failed,err:", err)
	}

	for {
		// 开始监听等待客户端连接
		fmt.Println("starting listen client accept.......")
		conn, err := listen.Accept()
		if err != nil{
			fmt.Println("accept failed err:", err)
			continue
		}

		// 创建　goroutine 处理连接
		go process(conn)
	}
}

func process(conn net.Conn) {
	
	// 表示处理完一个任务则关闭一个连接
	defer conn.Close()
	for {
		// 设置每次读取数据的大小
		buf := make([]byte, 512)
		// 开始读取数据
		n, err := conn.Read(buf[:]) 
		if err != nil {
			fmt.Println("read failed, err:", err)
			break
		}
		
		// 读取到的内容由于的字节类型的数据，所以需要将其转换为字符串
		str := string(buf[:n])
		fmt.Printf("recv from client, data:%v\n", str)
	}
}