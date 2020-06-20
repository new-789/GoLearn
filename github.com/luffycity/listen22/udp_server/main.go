package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("wait of clident connect......")
	/* 
	**net.LIstenUDP**　方法说明: 创建UDP协议的监听方法； 
		该方法接收两个参数：
			参数一："udp" 表示监听的协议类型，此处为　udp类型
			参数二：地址类型的参数，参数为:　**``&net.UDPAddr``**　该参数用来设置监听的 IP 地址和端口号,
				   设置的 IP 地址和端口号必须为　key value 的格式，格式如下:  
				   IP格式
				   ```
				   IP: net.IPv4(x,x,x,x) // x 表示 ip 地址的数据  
				   ```  　
				   端口号格式:  
				   ```
					Port:8000  
				   ```	
		监听整体格式如下：
		```
		socker, err := net.ListenUDP("udp", &net.UDPAddr{
			IP:   net.IPv4(0,0,0,0),  // 0,0,0,0 表示监听本机所有 IP
			POrt: 8000,
		})  
		```
		返回值说明：  
			该方法返回一个监听对象，即语法中第一个内容　socker，和一个错误信息;  
	*/
	socket, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 6000,
	})
	if err != nil {
		fmt.Println("listen Failed, err:", err)
	}

	defer socket.Close()  // 关闭连接

	for {
		var data = make([]byte, 4096)
		/* 
		**ReadFromUDP** 方法说明:  用来读取客户段的数据;  
		该方法接收一个　[]byte 切片类型的数据作为参数，切定义的 []byte 类型的切片应指定长度，表示一次性读取内容的大小;  
		读取　UDP 包中的数据,返回三个内容，  
			第一个为读取到的数据内容总长度，  
			第二个为客户端的地址，  
			第三个则为错误信息  
		语法如下:
		```go
		count, arrd, err := socker.ReadFromUDP(make([]byte, 4096))
		```
		*/
		count, addr, er := socket.ReadFromUDP(data[:])
		if er != nil {
			fmt.Println("readUdp failed,err:", er)
			continue
		}
		fmt.Printf("from client add:%v, message:%v, count:%v\n", addr, string(data[:count]), count)

		/* 
		**WriteToUDP** 方法说明：该方法用来返回数据给　ＵDP　连接的客户端;  
			该方法接受两个参数，
				参数一为需要发送的内容(注发送的内容必须为　byte 类型的数据)
				参数二为客户端的地址,表示需要返回的内容发送到那个客户端
		*/
		_, e := socket.WriteToUDP([]byte("hello client"), addr)
		if e != nil {
			fmt.Println("send data failed, err:", e)
			continue
		}
	}
}