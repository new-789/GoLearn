package main

import (
	"fmt"
	"net"
)

func main() {
	// 建立连接
	socket, e := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127,0,0,1),
		Port: 6000,
	})
	if e != nil {
		fmt.Println("dial failed, err:, e", e)
		return
	}

	// 关闭连接
	defer socket.Close()

	// 发送数据给服务端
	sendData := []byte("hello server")
	_, er := socket.Write(sendData)
	if er != nil {
		fmt.Println("write data failed, err:", er)
		return
	}

	// 从服务端读取数据
	data := make([]byte, 4096)
	count, addr, err := socket.ReadFromUDP(data[:])
	if err != nil {
		fmt.Println("read data failed, err:", err)
		return
	}
	fmt.Printf("data:%v, addr:%v, count:%v\n", string(data), addr, count)
}