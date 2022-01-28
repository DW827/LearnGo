## 一.字符串
+ #### 例1：维吉尼亚加密算法
+ ##### 知识点：字符类型(byte=uint8；rune=int32)，make([]type,len,cap)，slice切片，字符编码。。。
+ code
```go
import (
	"fmt"
	"strings"
)

func main() {
	message := "CSOITEUIWUIZNSROCNKFD" //明文
	keyword := "GOLANG"                //密钥
	var start int32 = 5
	//加密
	message = VirginiaEncode(message, keyword, start)
	fmt.Println(message)
	//解密
	message = VirginiaDecode(message, keyword, start)
	fmt.Println(message)
}

//keyvalue计算密钥中每个字母对应的数值
func keyvalue(keyword string, start int32) []int32 {
	//转换为大写字母
	keyword = strings.ToUpper(keyword)
	//创建长度为密钥长度的数组切片
	keynum := make([]int32, len(keyword))

	for i := 0; i < len(keyword); i++ {
		keynum[i] = int32(keyword[i]) - 65 + start
	}
	return keynum //返回密钥对应转换值
}

//VirginiaEncode维吉尼亚加密算法-加密
func VirginiaEncode(msg, keyword string, start int32) (message string) {
	//加密前密文转换为大写字母
	msg = strings.ToUpper(msg)
	//调用keyvalue获取密钥对应转换值
	keynum := keyvalue(msg, start)

	for i, c := range msg {
		if c >= 'A' && c <= 'Z' {
			c += keynum[i%len(keynum)]
			if c > 'Z' {
				c -= 26
			}
		}
		message += string(c)
	}
	return //返回加密后的密文message
}

//解密函数
func VirginiaDecode(msg, keyword string, start int32) (message string) {
	msg = strings.ToUpper(msg)
	keynum := keyvalue(keyword, start)

	for i, c := range msg {
		if c >= 'A' && c <= 'Z' {
			c -= keynum[i%len(keynum)]
			if c < 'A' {
				c += 26
			}
		}
		message += string(c)
	}
	return
}
```
+ #### 例2：统计文本单词频率
+ ##### 知识点：map[key type]value type，分隔字符串，去除字符串的标点符号，统计字符个数。。。
+ code
```go
import (
	"fmt"
	"strings"
)

func main() {
	text := "As far as eye could reach he saw nothing" +
		" but the stems of the great plants about him receding " +
		" in the violet shade, and far overhead the multiple " +
		" transparency of huge leaves filtering the sunshine " +
		" to the solemn splendour of twilight in which he walked."

	//frequency := make(map[string]int)
	var frequency map[string]int
	frequency = make(map[string]int)
	frequency = count(text)
	for k, v := range frequency {
		fmt.Printf("%s: %d\n", k, v)
	}
}

//count统计文本每个单词出现频率并返回一个频率map
func count(text string) (frequency map[string]int) {
	//化为小写字母
	text = strings.ToLower(text)
	//以空白为界划分字符，生成字符的切片
	text2 := strings.Fields(text)
	//去除标点符号
	text3 := make([]string, 0, len(text2))
	for _, c := range text2 {
		c = strings.Trim(c, ",")
		c = strings.Trim(c, ".")
		c = strings.Trim(c, ";")
		text3 = append(text3, c)
	}
	//创建一个map用来计数
	frequency = make(map[string]int)
	for _, t := range text3 {
		frequency[t]++
	}
	return
}

```
# 二.struct
+ #### 例1：坐标转换
+ ##### 知识点：定义结构体，类型绑定方法，定义构造函数。。。
+ code
```go
import "fmt"

//经纬度坐标
type location struct {
	lat, long float64
}

//度分秒坐标
type coordinate struct {
	d, m, s float64
	h       rune
}

//decimal converts a d/m/s coordinate to decimal degrees.
//定义一个方法decimal与结构体coordinate绑定
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

//newlocation from latitude, longtitude d/m/s coordinates.
//定义一个构造函数用来构造location类型的变量
func newlocation(lat, long coordinate) location {
	//struct符合字面值初始化location
	return location{lat.decimal(), long.decimal()}
}

func main() {
	//Bradbury Landing: 4°35′22.2″ S，137°26′30.12″ E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	//将度分秒表示的坐标转换成经纬度小数表示的坐标
	fmt.Println(lat.decimal(), long.decimal())

	//或构造变量curiosity再打印
	curiosity := newlocation(lat, long)
	fmt.Println(curiosity)
}
```
+ #### 例2：行星温度与坐标
+ ##### 知识点：结构体组合，方法转发，结构体嵌入。。。
+ code
```go
import "fmt"

/*普通定义
type report struct {
	sol         int         //日期
	temperature temperature //温度
	location    location    //坐标
}
*/

//优化定义，通过结构体嵌入
type report struct {
	sol         int
	temperature //report可以只有类型名，没有字段名，会自动分配相同名作为字段
	location    //依旧可以访问该字段以及字段的方法
}

//用已有类型自定义新类型，新类型再组合成新类型
type location struct {
	lat, long float64
}
type temperature struct {
	high, low celsius
}
type celsius float64

//绑定方法,方法转发
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	//复合字面值初始化
	bradbury := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -78.0}
	//实现结构体复用
	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury}
	//打印当天温度
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.temperature.high)

	//调用方法打印当天温度
	fmt.Println(report.average())
}
result:
{sol:15 temperature:{high:-1 low:-78} location:{lat:-4.5895 long:137.4417}}
a balmy -1° C
-39.5
```
* #### 例3：激光发射
* ##### 知识点：接口定义，接口与结构体嵌入，方法组合使用
* code
```go
import (
	"fmt"
	"strings"
)

//接口类型，只要某个类型的方法签名同接口一样就满足接口
type talker interface {
	talk() string //方法签名
}

//定义martian和laser两个结构体，均满足或者说实现了上述接口方法
type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew", int(l))
}

//定义函数，满足接口类型都可以传入
func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//定义结构体，嵌入laser类型，也实现laser实现的接口和方法
type starship struct {
	laser
}

func main() {
	//直接调用
	shout(martian{})
	shout(laser(2))
	//或定义变量，类型满足接口或者说实现了接口
	s := starship{laser(3)}
	shout(s)
} 
```
# 三.错误
* #### 例1.九宫格数独
* ##### 知识点：error包, error返回和处理，错误类型断言，接口。。。
* code
```go
import (
	"errors"
	"fmt"
	"os"
	"strings"
)

//九宫格常量
const rows, columns = 9, 9

//Grid is a Sudoku grid
type Grid [rows][columns]int8

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

//SudukuError...
type SudokuError []error

//Error returns one or more errors
//SudokuError类型实现了返回string的Error方法(签名相同)，则满足了内置的error接口
func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

//set...
//SudokuError类型满足error接口，因此errs可以作为error接口类型返回
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	//没错，set成功，返回nil
	g[row][column] = digit
	return nil
}

//判断坐标是否越界
func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

//判断填入数值是否合法
func validDigit(digit int8) bool {
	if digit < 1 || digit > 9 {
		return false
	}
	return true
}

func main() {
	var g Grid

	//err是error接口类型
	err := g.Set(10, 0, 10)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit: //或
			fmt.Println("Les errors de parametres hors limites.")
		default:
			//会调用SudokuError的Error()方法
			fmt.Println(err)
		}
		//os.Exit(1)
	}
	//改进，用类型断言
	if err != nil {
		//err.(errortype)返回对应错误类型，err是否是对应错误类型
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("-%v\n", e)
			}
		}
		os.Exit(1)
	}
}
```
# 四.并发与死锁
* #### 例1.睡觉
* ##### 知识点：goroutine的使用，通道channel，select case。。。
* code
```go
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//创建一个channel通道c
	c := make(chan int)
	//创建5个goroutine,并发执行
	for i := 0; i < 5; i++ {
		go sleepygopher(i, c)
	}

	//创建通道timeout
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		//等0~4秒后从通道接收值并执行
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, "has finished sleeping")
		//等2秒后执行
		case <-timeout:
			fmt.Println("my patience ran out")
			return	//一旦执行，main返回，关闭所有goroutine
		}
	}
}

func sleepygopher(id int, c chan int) {
	//随机等待0~4秒时间
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("... ", id, " snore ...")
	c <- id
}
```
* #### 例2.流水线
* ##### 知识点：通道的关闭，通道的读取和写入。。。
* code
```go
import (
	"fmt"
	"strings"
)

//上游
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downstream <- v
	}
	//发送完最后一个，关闭通往中间的通道
	close(downstream)
}

//中间
func filterGopher(upstream, downstream chan string) {
	/*for {
		item, ok := <-upstream
		if !ok {
			//ok为false，item接收到nil，说明upstream已经关闭
			close(downstream)
			return
		}
		//非次品发送到下游
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}*/
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}

//下游
func printGopher(upstream chan string) {
	//用range从通道读取值，直到通道关闭为止
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}
```
* #### 例3.驾驶
* ##### 知识点：坐标表示，更新坐标，通道，构造函数，包含select的for循环。。。
* code
```go
import (
	"image"
	"log"
	"time"
)

/*
func worker() {
	pos := image.Point{X: 10, Y: 10} //类型
	direction := image.Point{X: 1, Y: 0}
	//等待1秒生成一个通道给next
	next := time.After(time.Second)
	for {
		select {
		case <-next:
			//Add方法更新坐标，返回point{pos.X + direction.X, pos.Y + direction.Y}
			pos = pos.Add(direction)
			fmt.Println("current position is ", pos)
			next = time.After(time.Second)
		}
	}
}*/

type command int

//定义常量right=0，left=1
const (
	right = command(0)
	left  = command(1)
)

//声明驾驶员类型，发出行驶命令
type RoverDriver struct {
	commandc chan command //有通道类型字段
}

//RoverDriver的构造函数
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive()
	return r
}

func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0}
	direction := image.Point{X: 1, Y: 0}

	//等250ms生成一个通道给nextMove
	updateInterval := 250 * time.Millisecond
	nextMove := time.After(updateInterval)

	for {
		select {
		case c := <-r.commandc:
			switch c {
			case right:
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				direction = image.Point{
					X: direction.X,
					Y: -direction.Y,
				}
			}
		//每等到250ms就执行
		case <-nextMove:
			pos = pos.Add(direction)
			log.Printf("move to %v", pos)
			nextMove = time.After(updateInterval)
		}
	}
}

//定义命令
func (r *RoverDriver) Left() {
	r.commandc <- left
}
func (r *RoverDriver) Right() {
	r.commandc <- right
}
func main() {
	//调用构造函数
	r := NewRoverDriver()
	//每隔3秒执行命令
	time.Sleep(3 * time.Second)
	r.Left()
	time.Sleep(3 * time.Second)
	r.Right()
	time.Sleep(3 * time.Second)
}
```