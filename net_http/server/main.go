package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 服务端

func main() {
	http.HandleFunc("/get", postHandler)
	http.HandleFunc("/post", getHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server faild, err:%v\n", err)
		return
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm)
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll((r.Body))
	if err != nil {
		fmt.Printf("read request.Body faild, err:%v\n", err)
		return
	}
	fmt.Println(string(b))

	// 给服务端回答
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))

	defer r.Body.Close()
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))

	// 给服务端回答
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
	
	defer r.Body.Close()
}