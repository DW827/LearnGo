### 1. channel
+ 如果说goroutine是Go程序并发的执行体，channel就是它们之间的连接。channel是可以让一个goroutine发送特定值到另一个goroutine的通信机制。
+ Go 语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。每一个通道都是一个具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

+ channel类型
channel是一种类型，一种引用类型，零值为nil。声明通道类型的格式如下：

    var 变量 chan 元素类型  
+ channel 操作有：发送，接收，关闭：close(chan)
关闭后的通道有以下特点：

    1.对一个关闭的通道再发送值就会导致panic。
    2.对一个关闭的通道进行接收会一直获取值直到通道为空。
    3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
    4.关闭一个已经关闭的通道会导致panic。

+ 无缓冲通道

无缓冲通道上的发送操作会阻塞，直到另一个goroutine在该通道上执行接收操作，这时值才能发送成功，两个goroutine将继续执行。相反，如果接收操作先执行，接收方的goroutine将阻塞，直到另一个goroutine在该通道上发送一个值。

使用无缓冲通道进行通信将导致发送和接收的goroutine同步化。因此，无缓冲通道也被称为同步通道。
+ 有缓冲的通道

只要通道的容量大于零，那么该通道就是有缓冲的通道使用make函数初始化通道的时候为其指定通道的容量


### 2. 协程goroutine

Goroutine 特点：

占用内存更小（几 kb）
调度更灵活 (runtime 调度)

### 3. 调度runtime

### 4. GMP模型，调度

在 Go 中，线程是运行 goroutine 的实体，调度器的功能是把可运行的 goroutine 分配到工作线程上。

Processor，它包含了运行 goroutine 的资源，如果线程想运行 goroutine，必须先获取 P，P 中还包含了可运行的 G 队列。

1、P 的数量：

由启动时环境变量 GOMAXPROCS 或者是由 runtime 的方法 GOMAXPROCS() 决定。这意味着在程序执行的任意时刻都只有 $GOMAXPROCS 个 goroutine 在同时运行。

2、M 的数量:

go 语言本身的限制：go 程序启动时，会设置 M 的最大数量，默认 10000. 但是内核很难支持这么多的线程数，所以这个限制可以忽略。
runtime/debug 中的 SetMaxThreads 函数，设置 M 的最大数量
一个 M 阻塞了，会创建新的 M。

CPU -> M -> P -> G本地队列 -> 执行G销毁G返回 -> M

### 5. context包的用途
Go服务器的每个请求都有自己的goroutine,而有的请求为了提高性能，会经常启动额外的goroutine处理请求,当该请求被取消或超时，该请求上的所有goroutines应该退出，防止资源泄露。
那么context来了，它对该请求上的所有goroutines进行约束，然后进行取消信号，超时等操作。而context优点就是简洁的管理goroutines的生命周期。

注意: 使用时遵循context规则

1. 不要将 Context放入结构体，Context应该作为第一个参数传入，命名为ctx。
2. 即使函数允许，也不要传入nil的 Context。如果不知道用哪种Context，可以使用context.TODO()。
3. 使用context的Value相关方法,只应该用于在程序和接口中传递和请求相关数据，不能用它来传递一些可选的参数
4. 相同的 Context 可以传递给在不同的goroutine；Context 是并发安全的。

### . [context](https://www.liwenzhou.com/posts/Go/go_context/)
Context类型，专门用来简化 对于处理单个请求的多个 goroutine 之间与请求域的数据、取消信号、截止时间等相关操作，这些操作可能涉及多个 API 调用。

对服务器传入的请求应该创建上下文，而对服务器的传出调用应该接受上下文。

它们之间的函数调用链必须传递上下文，或者可以使用WithCancel、WithDeadline、WithTimeout或WithValue创建的派生上下文。

当一个上下文被取消时，它派生的所有上下文也被取消。

### Context接口
context.Context是一个接口，该接口定义了四个需要实现的方法。具体签名如下：
```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}
```

    ctx.Deadline()  // 返回当前Context被取消的时间，也就是完成工作的截止时间（deadline）;
    ctx.Done()      // 返回一个Channel，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用Done方法会返回同一个Channel；
    ctx.Err()       // 返回当前Context结束的原因，它只会在Done返回的Channel被关闭时才会返回非空的值；
    >如果当前Context被取消就会返回Canceled错误；
    >如果当前Context超时就会返回DeadlineExceeded错误；
    ctx.Value()     // 从Context中返回键对应的值，对于同一个上下文来说，多次调用Value 并传入相同的Key会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；

```go

```