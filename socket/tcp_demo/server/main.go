package main

// tcp server demo

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 新建读取信息对象
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取客户端发来的数据
		if err != nil {
			fmt.Println("read from client faild, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 给客户端发送数据
	}
}

func main() {
	// 1. 开启服务监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("listen failed,err:%v\n", err)
		return
	}
	defer listener.Close()

	for {
		// 2. 不断尝试接收客户端连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept faild, err:%v\n", err)
			continue
		}
		// 3. 启动一个单独的goroutine处理连接
		go process(conn)
	}
}
