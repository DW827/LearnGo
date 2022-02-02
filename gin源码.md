### [gin](https://www.liwenzhou.com/posts/Go/read_gin_sourcecode/#autoid-0-1-1)
1. gin框架路由使用前缀树，路由注册的过程是构造前缀树的过程，路由匹配的过程就是查找前缀树的过程。
2. gin框架的中间件函数和处理函数是以切片形式的调用链条存在的，我们可以顺序调用也可以借助c.Next()方法实现嵌套调用。
3. 借助c.Set()和c.Get()方法我们能够在不同的中间件函数中传递数据。