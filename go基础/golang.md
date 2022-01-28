# **GoLang notes**

## package&function:包和函数

## 语法规范
##### 只说明与C语言不同的地方

* ### 1.运算
+ 注意：没有++count
>
> 1. 使用big包处理特大数（超过10e18）：[很大的数](https://www.bilibili.com/video/BV1fD4y1m7TD?p=9&spm_id_from=pageDriver)
>一旦使用了 big.Int，那么等式里其它的部分也必须使用 big.Int
>>NewInt() 函数可以把 int64 转化为 big.Int 类型
>>如何把 24 x 1018 转化为 big.Int 类型？
>>首先 new 一个 big.Int
>>再通过 SetString 函数把数值的字符串形式，和几进制传递进行即可。
>
>缺点：用起来繁琐，且速度慢



* ### 2.基本语法
>package 本文件属于哪个包，每个Go应用程序都包含一个名为 main 的包。
>import "调入其他包"
>func 声明函数
>
> + "{"不能单独在一行
> + "}"必须单独在一行
>
>每行代码不需要";"结尾
>
+ golang不允许import了某个包但是不使用该包，变量也是。但是导入的包前面加"_"可以忽略导入包
+ 可以给导入的包起别名，如：import ( aa "fmt" ),可以通过aa.Println()调用。
+ 

* ### 3.变量声明
* 
> 1. 变量var声明或者短声明:=
>短声明“：”可以在if，for，switch等语句内声明同时赋值变量
>code eg:
    for num := 5; num > 0; num-- {
		fmt.Println(num)
	}
>
> 2. 常量const 支持多重声明赋值。
>
> 3. 声明变量而不赋初值时，变量初值变量类型的零值。
>
> 4. go是静态类型语言，一旦声明变量类型就不能再改变。

* ##### 声明浮点型变量
> 1. 下面三句等价：
    days := 365.365
    var days = 365.365
    var days float64 = 365.365
>
> 2. 默认
>> + 只要数字含有小数部分，那它的类型默认就是float64
>> + 只要数字是指数形式，如24e+18，那它的类型也是默认float64
>
> 3. go语言有两种浮点型，float32和float64:
>> + float32占四字节，适合大量数据时，牺牲精度节约内存；
>> + float64占八字节，精度高，由于math包的函数操作都是float64类型，故首选float64

* ##### 整型变量
>8种，包括有符号int和无符号uint
>有8,16,32,64位

* ##### 字符串类型
+ string、[]rune()、[]byte()
string 底层是一个包含多个字节（1字节=8bit）的集合。
string 类型的值是不可改变的。
string 可以被拆分为一个包含多个字节的序列，[]byte(str)
string 可以被拆分为一个包含多个字符的序列，[]rune(str)
+ 
+ 我们通常说的字符是指 Unicode 字符。
‘G’, ‘o’, ‘菜’, ‘鸟’ 都是一个字符。
一个字符可以是只包含一个字节（像：‘G’, ‘o’），
也可以是包含多个字节（像：‘菜’, ‘鸟’）。
一个英文字符对应一个byte，一个中文字符对应三个byte。
一个rune对应一个UTF-8字符，所以一个中文字符对应一个rune。
1. 类型别名
>type用来声明新类型
>eg: type celsius float64
>> + type byte=uint8
>> + type rune=int32
>
2. 字符串与字符
>> + 字符串字面值：string literal:""。可以带转义字符\n换行
>> + 原始字符串字面值：raw string literal:''
>> + 字符默认为rune类型
>> + 字符串可以赋值给变量但是字符串不可以修改
>
>凯撒密码：字母+1
>ROT13密码：字母+13
```go
    //加密
    for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
    //解密
    for i, c := range msg {
		if c >= 'A' && c <= 'Z' {
			c -= keynum[i%len(keynum)]
			if c < 'A' {
				c += 26
			}
		}
	}
```
>字符串只能与字符或字符串连接连接
3. 整数转字符串
> + strconv.Itoa(count),可以把数值count转换为code point对应的字符串。
> + fmt.Springtf("%v, count),也可以把数值count转换为字符串。

4. 字符串转整数
```go
count, err := strconv.Atoi("10")
if err != nil {     //nil相当于null
    //error
    fmt.Println(err.Error())
}
fmt.Println(count)
```
5. go的bool类型只能是true和false
6. string本身不可变，因此要改变string的字符，要先把string转换为byte(ACKLL码)或rune(中文字符)，
code eg：
```go
func main() {
    str := "Hello world"
    s := []byte(str) //中文字符需要用[]rune(str)
    s[6] = 'G'
    s = s[:8]
    s = append(s, '!')
    str = string(s)
    fmt.Println(str)
}
result：
Hello Go！
```
>[更多](https://www.topgoer.cn/docs/golang/chapter03-8)

* ##### %格式化输出

* ### 4.循环与分支
* ### 5.函数与方法
> + [函数](https://www.topgoer.cn/docs/golang/chapter05)
> + [方法](https://www.topgoer.cn/docs/golang/chapter06)
> + go的函数和方法需要大写字母开头才能导出包被其他文件调用.
> ###### go的函数和方法都是传递参数的副本.复制品。注意1：无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
> ###### 注意2：map、slice、chan、指针、interface默认以引用的方式传递。
> + 随时定义函数，随时执行函数，func(){}()，加()表示立即执行函数
> + "_"标识符，用来忽略函数的某个返回值
```go
package main

import "fmt"

func test(fn func() int) int {
	return fn()
}

// 定义函数类型。FormatFunc是一个函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)
}

```
#### 匿名函数
#### [闭包](https://www.topgoer.cn/docs/golang/chapter05-5)
+ 当函数a()的内部函数b()被函数a()外的一个变量引用的时候，就创建了一个闭包。
+ 闭包复制的是原对象指针
### 6.数组&&切片
+ [切片slice](https://www.topgoer.cn/docs/golang/chapter03-10)
> +  切片与指针
>> +  每个 slice 内部都会被表示为一个包含 3 个元素的结构，它们分别指向：
>> + 数组的指针
>> + slice 的容量
>> + slice 的长度
>> 当 slice 被直接传递至函数或方法时，slice 的内部指针就可以对底层数据进行修改。
>  + 利用指针修改切片
>> 指向 slice 的显式指针的唯一作用就是修改 slice 本身：slice 的长度、容量以及起始偏移量。指针的指针
> + 定义切片
>> + 用:=arr[low:high]，创建的切片容量等于底层数组的长度。
>> + 用:=arr[:len:cap]，限制了容量的方法创建一个slice就会自动分配一个新的底层数组，与原数组无关。
>> + 而:=arr[:len]，不限制容量的方法创建一个slice时cap等于len，但此时引用的是原数组；此时修改slice会影响原数组，此时强制扩容也会影响原数组元素但不改变元素组容量，并重新分配新的底层数组。
>> + 用append( ,...)可以追加切片的长度，若追加时超过了切片容量，会重新分配底层数组。
> + 特别的：make可以用来预先分配内存空间。
> + 总结：
>> + slice中元素个数决定了slice长度
>> + slice的底层数组长度决定了slice的容量
>> + 常规：a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x
 + slice中cap重新分配规律：
```go
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, 1)
		if n := cap(s); n > c {
			fmt.Printf("cap:%d -> %d\n", c, n)
			c = n
		}
	}
```
>分配内存的[new & make](https://www.topgoer.cn/docs/golang/chapter03-12)
```go
	var a *int		//声明
	a = new(int)	//初始化
	*a = 10
	fmt.Println(*a) //10
```
###### 或
```go
	a := new(int)
	*a = 10
	fmt.Println(*a) //10
```
* ### 7.逗号ok模式
* ，ok或者 ok,_ 两种写法是根据函数的返回值决定的，有的第一个返回的是bool,另一个是error信息，就选用ok,err进行判断，有的第一个参数是一个值或者nil，第二个参数是true/false，就选用value,ok。
* 1.在函数返回时检测错误
```go
func SomeFunc() error {
    …
    if ok, err := pack1.Func1(param1); err != nil {
        …//函数错误
        return err
    }
    …
    return nil
}

```
* 2.检测映射中是否存在一个键值：key1在映射map1中是否有值？
```go
if value, ok = map1[key1]; ok {
        Process(value)	//有值
}else{
	fmt.Println("Where is the map1[key1])
}
```
* 3.检测一个接口类型变量var是否包含了类型T：类型断言
```go
if value, ok := var.(T); ok {
	Process(value)	//包含
}else{
	//不包含
}
```
* 4.检测一个通道ch是否关闭
```go
for {
    if v, ok := <-ch; !open {
        break // 通道是关闭的
    }
    Process(v)	//通道未关闭
}
```
* ### map
* map是一种隐式指针，作为引用传递
* map[key]value
* 应用举例一：计算器
```go
func main() {
	temperature := []float64 {
		28.0, 32.0, 31.0, 29.0, 23.0, 29.0, 28.0, 33.0,
	}
	//创建一个map用来计数
	frequency := make(map9[float64]int)
	for _, t := range temperature {
		frequency[t]++
	}
	//输出计数结果
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}
result:
+33.00 occurs 1 times
+28.00 occurs 2 times
+32.00 occurs 1 times
+31.00 occurs 1 times
+29.00 occurs 2 times
+23.00 occurs 1 times
```
* 应用举例二：数据分组
```go
func main() {
	temp := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64)
	//分组,值为切片的map
	for _, t := range temp {
		//以10为跨度进行分组
		g := math.Trunc(t/10) * 10
		//将值加入各自所属分组
		groups[g] = append(groups[g], t)
	}
	//输出分组结果
	for g, temp := range groups {
		fmt.Printf("%v: %v\n", g, temp)
	}
}
result:
-20: [-28 -29 -23 -29 -28]
30: [32]
-30: [-31 -33]
```
* 应用举例三：创建集合(元素不会重复)
```go
func main() {
	temp := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	set := make(map[float64]bool)
	for _, t := range temp {
		set[t] = true
	}
	//集合unique
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	//对切片unique进行排序
	sort.Float64s(unique)
	fmt.Println(unique)
}
```
* ### 8.指针
* &和*
> + & 获得变量的地址
> + ‘ * ’ 解引用，提供内存地址指向的值
> + & 操作符无法获得字符串/数值/布尔字面值的地址。
&42，&“hello”这些都会导致编译器报错
> + 与字符串和数值不一样，复合字面量的前面可以放置 &。
&接某个struct类型不会报错
> + 和结构体一样，可以把 & 放在数组的复合字面值前面来创建指向数组的指针。
> + 和结构体不一样的是，数组在执行索引或切片操作时会自动解引用。没有必要写 (*superpower)[0] 这种形式。
* go的指针p不支持p++操作
* 尝试解引用一个 nil 指针将导致程序崩溃。
### 9.[并发编程](https://www.topgoer.cn/docs/golang/chapter09-2)
* goroutine实现并发
> + main函数退出，其他goroutine全部结束。
> + 即使已经停止等待 goroutine，但只要 main 函数还没返回，仍在运行的 goroutine 将会继续占用内存。 
* sync实现goroutine的互斥和同步
> + sync.WaitGroup 实现同步
> + sync.Mutex 上锁
```go
var wg sync.WaitGroup

func hello(i int) {
    defer wg.Done() // goroutine结束就登记-1
    fmt.Println("Hello Goroutine!", i)
}
func main() {

    for i := 0; i < 10; i++ {
        wg.Add(1) // 启动一个goroutine就登记+1
        go hello(i)
    }
    wg.Wait() // 等待所有登记的goroutine都结束
}
```
* 无缓冲通道channel会使发送和接收的goroutine同步
### 10. [反射](https://www.topgoer.cn/docs/golang/chapter11-12)
反射可以：查看类型、字段和方法，查看和修改值，调用方法，获取字段的tag。
```go
package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name string `json:"name1" db:"name2"`
}

func main() {
    var s Student
	// v = &s
    v := reflect.ValueOf(&s)

    // 类型
    t := v.Type()

    // 获取字段
	// Elem()获取地址指向的值
    f := t.Elem().Field(0)

	fmt.Println(v)
	fmt.Println(t)
    fmt.Println(f.Tag.Get("json"))
    fmt.Println(f.Tag.Get("db"))
}
//result:
&{}
*main.Student
name1
name2
```