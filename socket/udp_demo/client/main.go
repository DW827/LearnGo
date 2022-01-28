package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接服务端
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 30000,
	})
	if err != nil {
		fmt.Printf("dial faild,err: %v\n", err)
		return
	}
	defer socket.Close() // 关闭连接

	inputReader := bufio.NewReader(os.Stdin)    // 从终端读取数据
	sendDate, _ := inputReader.ReadString('\n') // 每次读一行
	_, err = socket.Write([]byte(sendDate))     // 发送数据
	if err != nil {
		fmt.Printf("send to server faild, err:%v\n", err)
		return
	}

	buf := make([]byte, 1024)
	n, remoteAddr, err := socket.ReadFromUDP(buf) // 接收数据
	if err != nil {
		fmt.Printf("recvStr from udp faild, err:%v\n", err)
		return
	}
	fmt.Printf("recvStr:%v from addr:%v count:%v\n", string(buf[:n]), remoteAddr, n)
}
