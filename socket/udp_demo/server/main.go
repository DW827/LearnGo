package main

// udp server demo

import (
	"fmt"
	"net"
)

func main() {
	// 1.开启服务监听端口
	listener, err := net.ListenUDP("udp", &net.UDPAddr{
		IP: net.IPv4(127,0,0,1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("listen failed, err:%v\n", err)
		return
	}
	defer listener.Close()

	// 2.不断尝试收发数据
	for {
		var buf[1024]byte
		n, addr, err := listener.ReadFromUDP(buf[:])	// 接收数据
		if err != nil {
			fmt.Printf("read from udp faild, err:%v\n", err)
			continue
		}
		fmt.Println("接收到的数据：", string(buf[:n]))

		_, err = listener.WriteToUDP(buf[:n], addr)		// 发送数据
		if err != nil {
			fmt.Printf("write to %v faild, err:%v\n", addr, err)
			continue
		}
	}
}