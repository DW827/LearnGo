package main

// 客户端

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://127.0.0.1:9090/post"
	// 表单数据
	//contentType := "application/x-www-form-urlencoded"
	//data := "name=小王子&age=18"
	// json
	contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}

// GET请求示例
/* func main() {
	// 取得回应
	resp, err := http.Get("https://www.bilibili.com/")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	// 读取信息
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	// 处理信息
	fmt.Println(string(body))
	
	defer resp.Body.Close()
} */

// 执行之后就能在终端打印https://www.bilibili.com网站首页的内容了，
// 我们的浏览器其实就是一个发送和接收HTTP协议数据的客户端，
// 我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，
// 然后浏览器会按照HTML、CSS等规则将网页渲染展示出来。