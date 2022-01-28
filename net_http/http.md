### Go语言内置的net/http包十分的优秀，提供了HTTP客户端和服务端的实现。
HTTP协议

+ 超文本传输协议（HTTP，HyperText Transfer Protocol)是互联网上应用最为广泛的一种网络传输协议，所有的WWW文件都必须遵守这个标准。设计HTTP最初的目的是为了提供一种发布和接收HTML页面的方法。

### HTTP客户端

基本的HTTP/HTTPS请求
Get、Head、Post和PostForm函数发出HTTP/HTTPS请求。
```go
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form",
	url.Values{"key": {"Value"}, "id": {"123"}})
```

程序在使用完response后必须关闭回复的主体。
```go
resp, err := http.Get("http://example.com/")
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

// 执行之后就能在终端打印https://www.bilibili.com网站首页的内容了，
// 我们的浏览器其实就是一个发送和接收HTTP协议数据的客户端，
// 我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，
// 然后浏览器会按照HTML、CSS等规则将网页渲染展示出来。