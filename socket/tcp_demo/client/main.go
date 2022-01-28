package main

// tcp client demo

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 1. 与服务的建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Printf("dial faild, err:%v\n", err)
		return
	}
	defer conn.Close() // 关闭连接
	// 2. 进行数据与发送和接收
	inputReader := bufio.NewReader(os.Stdin)	// 从终端读取数据
	for {
		input, _ := inputReader.ReadString('\n') // 读取用户输入，读一行
		inputInfo := strings.Trim(input, "\r\n")	// 按空格划分
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}

		_, err = conn.Write([]byte(inputInfo)) // 给服务端发送数据
		if err != nil {
			return
		}
		
		buf := [512]byte{}		// 从服务端接收回复的消息
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println("收到服务端回复：",string(buf[:n]))
	}
}