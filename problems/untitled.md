# Go Problems

## Go基础知识

### make和new的区别**，**值类型和引用类型（slice切片、map,管道channel），值传递&引用传递

```text
type Student1 struct{
    Age int32
    Name string
}
type Student2 map[string]int

func main(){
    var s1 Student1
    var s2 Student2
}
```

**零值不同**

* 指针类型的变量，零值都是nil。
* 值类型的变量，零值是其所在类型的零值。

**变量申明后是否需要初始化才能使用**

* 指针类型的变量，需要初始化才能使用。\(slice是一个特例，slice的零值是nil，但是可以直接append\)
* 值类型的变量，不用初始化，可以直接使用

```text
fmt.Println(s1) //{0 }
fmt.Println(s2) //map[] 
fmt.Println(s1 == nil) //panic,提示cannot convert nil to type Student1
fmt.Println(s2 == nil) //true
```

**初始化方法不同**

* 值类型的变量，其实不需要初始化就可以使用。如果有良好的代码习惯，使用前进行初始化也是非常提倡的。
  * 基本类型的初始化非常简单:
    * var i int; i = 1;
    * var b bool; b = true
    * var s string; s = ""
  * 符合类型struct的初始化有两种:
    * s1 = Student1{}
    * s1 = new\(Student1\)

```text
s1 := Student1{} //{0 }
s1 := new(Student1) //&{0 }
```

* 引用类型的变量，初始化方式也不一样：

  ```text
  //map可以用{}，make，new三种方式初始化
  s2 := Student2{} //map[]
  s2 := make(Student2) //map[]
  s2 := new(Student2) //&map[]

  //slice可以用{},make,new,但是make的时候需要带len参数
  type S3 []string
  s3 := S3{} //[]
  s3 := new(S3) //&[]
  s3 := make(S3, 10) //[         ]

  //channel只能用make或者new
  type Student4 chan string
  s4 := new(S4) //0xc000096000
  s4 := make(S4) //0xc000082008
  s4 := S4{} //编译器报错：invalid type for composite literal: Student4
  ```

  **make和new的区别**

  从上面初始化时用make和new的结果就可以看出来：

  * make返回的是对象。
    * 对值类型对象的更改，不会影响原始对象的值
    * 对引用类型对象的更改，会影响原始对象的值
  * new返回的是对象的指针，对指针所在对象的更改，会影响指针指向的原始对象的值。

  ```text
  type Student struct{
    8
    9     Age int
   10 }
   11
   12 func main(){
   13
   14     s := Student{30}
   15     modify1(s)
   16     fmt.Println(s) //{30}
   17
   18     modify2(&s)
   19     fmt.Println(s) //{32}
   20 }
   21
   22 func modify1(s Student){
   23     s.Age = 31
   24 }
   25
   26 func modify2(s *Student){
   27     s.Age = 32
   28 }
  ```

### defer执行顺序

**defer**的**执行顺序**为：后**defer**的先**执行**。 **defer**的**执行顺序**在return之后，但是在返回值返回给调用方之前，所以使用**defer**可以达到修改返回值的目的。

```text
package main

import (
    "fmt"
)

func main() {
    ret := test()
    fmt.Println("test return:", ret)
}

func test() ( int) {
    var i int

    defer func() {
        i++        //defer里面对i增1
        fmt.Println("test defer, i = ", i)
    }()

    return i
}
```

执行结果为：

```text
test defer, i =  1
test return: 0
```

test函数的返回值为0，defer里面的i++操作好像对返回值并没有什么影响。  
这是否表示“return i”执行结束以后才执行defer呢？  
非也！再看下面的例子：

```text
package main

import (
    "fmt"
)

func main() {
    ret := test()
    fmt.Println("test return:", ret)
}

//返回值改为命名返回值
func test() (i int) {
    //var i int

    defer func() {
        i++
        fmt.Println("test defer, i = ", i)
    }()

    return i
}
```

执行结果为：

```text
test defer, i =  1
test return: 1
```

这次test函数的返回值变成了1，defer里面的“i++"修改了返回值。所以defer的执行时机应该是return之后，且返回值返回给调用方之前。

1. 多个defer的执行顺序为“后进先出”；
2. 所有函数在执行RET返回指令之前，都会先检查是否存在defer语句，若存在则先逆序调用defer语句进行收尾工作再退出返回；
3. 匿名返回值是在return执行时被声明，有名返回值则是在函数声明的同时被声明，因此在defer语句中只能访问有名返回值，而不能直接访问匿名返回值；
4. return其实应该包含前后两个步骤：第一步是给返回值赋值（若为有名返回值则直接赋值，若为匿名返回值则先声明再赋值）；第二步是调用RET返回指令并传入返回值，而RET则会检查defer是否存在，若存在就先逆序插播defer语句，最后RET携带返回值退出函数；

‍‍_**因此，**_‍‍defer、return、返回值三者的执行顺序应该是：return最先给返回值赋值；接着defer开始执行一些收尾工作；最后RET指令携带返回值退出函数。

a\(\)int 函数的返回值没有被提前声明，其值来自于其他变量的赋值，而defer中修改的也是其他变量（其实该defer根本无法直接访问到返回值），因此函数退出时返回值并没有被修改。

b\(\)\(i int\) 函数的返回值被提前声明，这使得defer可以访问该返回值，因此在return赋值返回值 i 之后，defer调用返回值 i 并进行了修改，最后致使return调用RET退出函数后的返回值才会是defer修改过的值。

```text
func main() {
	c:=c()
	fmt.Println("c return:", *c, c) // 打印结果为 c return: 2 0xc082008340
}

func c() *int {
	var i int
	defer func() {
		i++
		fmt.Println("c defer2:", i, &i) // 打印结果为 c defer2: 2 0xc082008340
	}()
	defer func() {
		i++
		fmt.Println("c defer1:", i, &i) // 打印结果为 c defer1: 1 0xc082008340
	}()
	return &i
}
```

```text
func main() {
	defer P(time.Now())
	time.Sleep(5e9)
	fmt.Println("1", time.Now())
}
func P(t time.Time) {
	fmt.Println("2", t)
	fmt.Println("3", time.Now())
}

// 输出结果：
// 1 2017-08-01 14:59:47.547597041 +0800 CST
// 2 2017-08-01 14:59:42.545136374 +0800 CST
// 3 2017-08-01 14:59:47.548833586 +0800 CST
```

defer的作用域

1. defer只对当前协程有效（main可以看作是主协程）；

2. 当panic发生时依然会执行当前（主）协程中已声明的defer，但如果所有defer都未调用recover\(\)进行异常恢复，则会在执行完所有defer后引发整个进程崩溃；

3. 主动调用os.Exit\(int\)退出进程时，已声明的defer将不再被执行。

### Rocover

在我们写程序时候，想让程序错误继续运行，一般我们会容错error。但是对于数组越界，我们还想让go函数跑的话，不影响主题函数

```text
package _func

import (
	"fmt"
	"time"
)

func Client()  {
	for {
		go myPainc()
		go normal()
		time.Sleep(time.Second)
	}
}

func myPainc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
			return
		}
	}()

	arr := make([]string, 0)
	arr = append(arr, "1")
	fmt.Println(arr[8]) // 看，这是重大错误
}

func normal()  {
	fmt.Println("正常运算")
}

## 输出
正常运算
```

### foreach的拷贝问题

for range是值拷贝出来的副本

可以用指针数组，value值用"\_"舍弃了元素的复制，用下标去访问\(效率更高）

```text
for i, _ := range t {
    t[i].Num += 100
}
```

### go执行的随机性和闭包

```text
func main() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

# 输出
B:  9
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
A:  10
B:  0
B:  1
B:  2
B:  3
B:  4
B:  5
B:  6
B:  7
B:  8
```

其中A:输出完全随机，取决于goroutine执行时i的值是多少；  
而B:一定输出为0~9，但顺序不定。

第一个go func中i是外部for的一个变量，地址不变化，但是值都在改变。

第二个go func中i是函数参数，与外部for中的i完全是两个变量。  
尾部\(i\)将发生值拷贝，go func内部指向值拷贝地址。

所以在使用goroutine在处理闭包的时候，避免发生类似第一个go func中的问题。

闭包能够访问外层代码中的变量； for循环与gotoutine同时执行； 所有的goroutine操作的变量都是直接操作外层代码的变量，而外层代码中的变量的值取决于循环执行的节点。

### go的继承与组合

这是Golang的组合模式，可以实现OOP的继承。 被组合的类型People所包含的方法虽然升级成了外部类型Teacher这个组合类型的方法（一定要是匿名字段），但它们的方法\(ShowA\(\)\)调用时接受者并没有发生变化。 此时People类型并不知道自己会被什么类型组合，当然也就无法调用方法时去使用未知的组合者Teacher类型的功能。

```text
type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA-People")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB-People")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	t.ShowA()
}

# 输出
showA-People
showB-People
```

{% embed url="https://refactoringguru.cn/design-patterns/composite/go/example" %}

### select随机性

select会随机选择一个可用通用做收发操作。 单个chan如果无缓冲时，将会阻塞。但结合 select可以在多个chan间等待执行。有三点原则： _select 中只要有一个case能return，则立刻执行。_ 当如果同一时间有多个case均能return则伪随机方式抽取任意一个执行。 如果没有一个case能return则可以执行”default”块。

```text
func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
}

# 输出
都有可能
```

Go里面提供了一个关键字select，通过select可以监听channel上的数据流动。  
select的用法与switch语言非常类似，由select开始一个新的选择块，每个选择条件由case语句来描述。  
与switch语句相比， select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作，大致的结构如下

```text
func WorkHere(quit chan interface{}) {
	for {
		select {
		case <-quit:
			fmt.Println("over")
			return
		default:
			time.Sleep(time.Millisecond * 100)
			fmt.Println("work")
		}
	}
}

func main() {
	var (
		quit chan interface{}
	)
	quit = make(chan interface{})
	go WorkHere(quit)
	time.Sleep(2 * time.Second)
	quit<- "quit"
}
```

```text
select {
    case <-chan1:
        // 如果chan1成功读到数据，则进行该case处理语句
    case chan2 <- 1:
        // 如果成功向chan2写入数据，则进行该case处理语句
    default:
        // 如果上面都没有成功，则进入default处理流程
}
```

注意事项： 1. 监听的case中，没有满足监听条件，阻塞。

1. 监听的case中，有多个满足监听条件，任选一个执行。
2. 可以使用default来处理所有case都不满足监听条件的状况。 通常不用（会产生忙轮询）
3. select 自身不带有循环机制，需借助外层 for 来循环监听
4. break 跳出 select中的一个case选项 。类似于switch中的用法。

### defer中插入函数的执行顺序

defer 压入栈的是值，如果为函数，则可以修改变量值

```text
func c() (i int) {
    defer func() { i++ }()
    return 1
}//函数返回值为 2。
```

通过`defer`修改返回值，`defer`也可以用于控制恢复`panic`断言。

```text
func main() {
    f()
    fmt.Println("Returned normally from f.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered in f", r)
        }
    }()
    fmt.Println("Calling g.")
    g(0)
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer in g", i)
    fmt.Println("Printing in g", i)
    g(i + 1)
}
```

```text
Calling g.
Printing in g 0
Printing in g 1
Printing in g 2
Printing in g 3
Panicking!
Defer in g 3
Defer in g 2
Defer in g 1
Defer in g 0
Recovered in f 4
Returned normally from f.
```

### make默认值和append

```text
func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

# 输出
[0 0 0 0 0 1 2 3]
```

### map线程安全

map：只读是安全的，不是线程安全的。在同一时间段内，让不同 goroutine 中的代码，对同一个字典进行读写操作是不安全 的。

```text
// 创建一个int到int的映射
m := make(map[int]int)
// 开启一段并发代码
go func() {
    // 不停地对map进行写入
    for {
        m[1] = 1
    }
}()
// 开启一段并发代码
go func() {
    // 不停地对map进行读取
    for {
        _ = m[1]
    }
}()
// 无限循环, 让并发程序在后台执行
for {
}
```

需要并发读写时，一般的做法是加锁，但这样性能并不高，Go语言在 1.9 版本中提供了一种效率较高的并发安全的 sync.Map，sync.Map 和 map 不同，不是以语言原生形态提供，而是在 sync 包下的特殊结构。  
  
sync.Map 有以下特性：

* 无须初始化，直接声明即可。
* sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
* 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。

```text
package main
import (
      "fmt"
      "sync"
)
func main() {
    var scene sync.Map
    // 将键值对保存到sync.Map
    scene.Store("greece", 97)
    scene.Store("london", 100)
    scene.Store("egypt", 200)
    // 从sync.Map中根据键取值
    fmt.Println(scene.Load("london"))
    // 根据键删除对应的键值对
    scene.Delete("london")
    // 遍历所有sync.Map中的键值对
    scene.Range(func(k, v interface{}) bool {
        fmt.Println("iterate:", k, v)
        return true
    })
}
```

### chan缓存池

[https://blog.csdn.net/zaimeiyeshicengjing/article/details/106124095](https://blog.csdn.net/zaimeiyeshicengjing/article/details/106124095)

### golang的方法集

T类型的值的方法集只包含值接收者声明的方法。而指向T类型的指针的方法集既包含值接收者声明的方法，也包含指针接收者声明的方法。

### interface内部结构



### type类型断言

### 函数返回值命名

可以给一个函数的返回值指定名字。如果指定了一个返回值的名字，则可以视为在该函数的第一行中定义了该名字的变量。

1. append切片加上...
2. 结构体比较

### interface内部解构

空接口是 `var i interface{}`,这个不是

```text
type People interface {
    Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
    var stu *Student
    return stu
}

func main() {
    if live() == nil {
        fmt.Println("AAAAAAA")
    } else {
        fmt.Println("BBBBBBB")
    }
}

# 输出
BBBBBBB
```

### 函数返回值类型

### iota和const系列变量复制

### 变量简短模式

### 常量在预处理阶段直接展开

### goto不能跳转到其他函数或者内层代码

### Type Alias和Type definition

### Type Alias ，引用方法的区别

### Type Alias ，结构体内部字段的区别

### 变量作用域

1. 闭包延迟求值
2. 闭包引用相同变量
3. panic仅有最后一个可以被recover捕获
4. 计算结构体大小
5. 字符串转成byte数组，会发生内存拷贝吗？
6. 拷贝大切片一定比小切片代价大吗
7. 能说说unitptr和unsafe.Pointer的区别吗？
8. reflect（反射包）如何获取字段tag？为什么json包不能导出私有变量的tag？
9. 反转含有中文，数字，英文字母的字符串

### 知道golang的内存逃逸吗？什么情况下会发生内存逃逸？怎么避免内存逃逸？

### 悬挂指针的问题

### 不定参数的使用 <a id="&#x4E0D;&#x5B9A;&#x53C2;&#x6570;&#x7684;&#x4F7F;&#x7528;"></a>

Go文档的`cmd`，`exec.Command(name string, args ...string)`

在我们想要获取`linux`服务器的输出状态，获取`cpu`使用率，内存剩余时候。就需要这个，那么对于通用接口实现，也就是说实现前面输出，后面解析，只用一个函数就能解决一系列问题。

```text
func t() {
	var (
		cmd *exec.Cmd
		output []byte
		err error
	)

	a := []string{"docker", "ps", "-a"}
	a2 := a[1:]
	// 生成Cmd
	cmd = exec.Command(a[0], a2...)

	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	// 打印子进程的输出
	fmt.Println(string(output))
}

func main() {
	t()
}
```

```text
// 执行后台任务 输出到前台
// 前台的文件是以 "," 为分隔符
func CmdHandle(c *gin.Context) {
	var (
		sliceCmd []string
		sliceCmdHead string
		sliceOther []string
		output []byte
		err error
		cmd *exec.Cmd
		q string
	)
	q = c.Query("cmd")
	sliceCmd = strings.Split(q, ",")

	sliceCmdHead = sliceCmd[0]
	sliceOther = sliceCmd[1:]

	cmd = exec.Command(sliceCmdHead, sliceOther...)

	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}

	str := string(output)

	strSlice := strings.Split(str, "\n")

	c.JSON(200, gin.H{
		"msg": "cmd_ok",
		"cmd": strSlice,
	})
}
```

```text
//标准输出
func CpuStatusHandle(ctx *gin.Context) {
	var (
		cmd           *exec.Cmd
		out           bytes.Buffer
		err           error
		process       []*Process
		line          string
		itemsNoFormat []string // 没格式化的
		items         []string
		pid           string
		cpu           float64
		cpuTotal      float64
		processTmp    Process
	)

	cmd = exec.Command("ps", "u")
	cmd.Stdout = &out
	if err = cmd.Run(); err != nil {
		fmt.Println("cmd_err------------------")
		return
	}

	process = make([]*Process, 0)
	for {
		if line, err = out.ReadString('\n'); err != nil {
			break
		}

		// 格式化命令
		itemsNoFormat = strings.Split(line, " ")
		items = make([]string, 0)
		for _, v := range itemsNoFormat {
			if v != "\t" && v != "" {
				items = append(items, v)
			}
		}

		// 提取pid 和 cpu
		pid = items[1]
		if cpu, err = strconv.ParseFloat(items[2], 64); err != nil {
			fmt.Println("系统异常  解析失败哦哦")
			//break
		}

		cpuTotal += cpu

		processTmp = Process{pid: pid}
		process = append(process, &processTmp)
	}

	ctx.JSON(200, gin.H{
		"cputotal": cpuTotal,
		"cpulist":  process,
	})
}
```

### CutIf消除if\_else（不推荐）

### FuncClosure闭包函数

**闭包就是能够读取其他函数内部变量的函数。例如在javascript中，只有函数内部的子函数才能读取局部变量，所以闭包可以理解成“定义在一个函数内部的函数“。在本质上，闭包是将函数内部和函数外部连接起来的桥梁**

```text
func ClosureBase() func(x1 int, x2 int) int {
	i := 0
	return func(x1 int, x2 int) int {
		i++
		fmt.Println("此时闭包里面的i的值:", i)
		sum := i + x1 + x2
		return sum
	}
}
```

```text
func TestClosureBase(t *testing.T) {
	fmt.Println("---------f1函数测试-----------------")
	f1 := ClosureBase()
	fmt.Println(f1(1, 1))
	fmt.Println(f1(1, 1))
	fmt.Println(f1(1, 1))

	fmt.Println("------------------f2函数测试---------------")
	f2 := ClosureBase()
	fmt.Println(f2(1, 1))
	fmt.Println(f2(1, 1))
	fmt.Println(f2(1, 1))
}
我是tmp
此时闭包里面的i的值: 1
3
此时闭包里面的i的值: 2
4
此时闭包里面的i的值: 3
5
------------------f2函数测试---------------
我是tmp
此时闭包里面的i的值: 1
3
此时闭包里面的i的值: 2
4
此时闭包里面的i的值: 3
5
```

### 闭包应用 <a id="&#x95ED;&#x5305;&#x5E94;&#x7528;"></a>

对于斐波那切数列

```text
func Fbi() func() int {
	fmt.Println("函数经过跑下面这俩了吗")
	b0 := 0
	b1 := 1
	return func() int {
		tmp := b0 + b1
		b0  = b1
		b1 = tmp
		return b1
	}
}

func TestFbi(t *testing.T) {
	fmt.Println("斐波那切数列，非递归")
	f := Fbi()
	for i := 1; i <= 4; i++ {
		fmt.Println(f())
	}
)
```

### FuncHigh高阶函数

这个函数的形参列表或返回参数列表中存在数据类型为函数类型，这就是高阶函数。

闭包就是高阶函数，他只是其中的一种，因为闭包是函数返回，还有一种就是形参是函数的。

```text
func HighFunc(val int, f func(i1 int) int) (func() int) {
	biVar := 0
	fmt.Println("此时闭包的值为", biVar)
	return func() int {
		if f(biVar) <= 3 {
			fmt.Println("可以加一")
			biVar++
		}
		return val + biVar
	}
}

func TestHighFunc(t *testing.T) {
	highFunc := HighFunc(0, func(i1 int) int {
		return i1
	})

	for i := 0; i <= 9; i++ {
		fmt.Println(highFunc())
	}
}
此时闭包的值为 0
可以加一
1
可以加一
2
可以加一
3
可以加一
4
4
4
4
4
4
4
```

过滤输出

```text
type student struct{
    name string
    grade int8
}

func filter(stu []student, f func(s student) bool) []student{
    var r []student

    for _, s := range stu {
        if f(s) == true {
            r = append(r, s)
        }
    }

    return r
}

func main() {
    s1 := student{
        "zhangsan",
        90,
    }

    s2 := student{
        "lisi",
        80,
    }

    s3 := student{
        "wanggang",
        70,
    }

    s := []student{s1, s2, s3}

    fmt.Println("all student: ", s)

    var result []student

    result = filter(s, func(s student) bool {
        if s.grade < 90 {
            return true
        }
        return false
    })

    fmt.Println("less than 90: ", result)
}
```

### Go init func Order



1. 代码执行，先执行所有的`init()`，init的执行规则是最后引用的先执行，就像栈一样
2. 最后是普通代码的执行顺序



###  Go Mod（ver 1.11）

* 如何只更新直接依赖；go get -u github.com/gin-gonic/gin
* 如何只更新间接依赖；
* 如何更新所有依赖；
* 本地依赖：replace 把引用的包放入自己的项目 放到github上面再拉下来
* **查看项目中依赖和情况：**go test ./...
* 未找到包：

  ```text
  export GO111MODULE=on
  export GOPROXY=https://goproxy.io
  ```

  **go build 提示找不到路径的报错**

* 在goland中的referrence中的Go的Go Modules（vgo）要勾选Enable等选项然后apply
* 进入main函数所在的目录，然后go build

### go.sum（module名、版本和哈希组成）

本意在于提供防篡改的保障，如果拉第三方库的时候发现其实际内容和记录的校验值不同，就让构建过程报错退出。然而它能做的也就只限于此。

`go.mod`只需要记录直接依赖的`依赖包版本`，只在`依赖包版本`不包含`go.mod`文件时候才会记录间接`依赖包版本`，而`go.sum`则是要记录构建用到的所有`依赖包版本`。

`go.sum`存在的意义在于，我们希望别人或者在别的环境中构建当前项目时所使用依赖包跟`go.sum`中记录的是完全一致的，从而达到一致构建的目的。

## Go advance

### 垃圾回收机制

1. 写屏障
2. 读屏障
3. 并发垃圾回收
4. go heap对象分布
5. 混合读写之后为什么能减少一次栈重扫

### channel的本质（如何解决并发冲突，常见并发模式）

### Work pool 并发控制，线程池

* `gopool.Submit` 在令牌不足时，会阻塞当前调用\(因此Go runtime会执行其他不阻塞的代码\)
* `gopool.Wait()` 会等到回收所有令牌之后，才返回

```text
package pool

type GoPool struct {
	MaxLimit int

	tokenChan chan struct{}
}

type GoPoolOption func(*GoPool)

func WithMaxLimit(max int) GoPoolOption {
	return func(gp *GoPool) {
		gp.MaxLimit = max
		gp.tokenChan = make(chan struct{}, gp.MaxLimit)

		for i := 0; i < gp.MaxLimit; i++ {
			gp.tokenChan <- struct{}{}
		}
	}
}

func NewGoPool(options ...GoPoolOption) *GoPool {
	p := &GoPool{}
	for _, o := range options {
		o(p)
	}

	return p
}

// Submit will wait a token, and then execute fn
func (gp *GoPool) Submit(fn func()) {
	token := <-gp.tokenChan // if there are no tokens, we'll block here

	go func() {
		fn()
		gp.tokenChan <- token
	}()
}

// Wait will wait all the tasks executed, and then return
func (gp *GoPool) Wait() {
	for i := 0; i < gp.MaxLimit; i++ {
		<-gp.tokenChan
	}

	close(gp.tokenChan)
}

func (gp *GoPool) size() int {
	return len(gp.tokenChan)
}
```

### 实现原理 <a id="&#x5B9E;&#x73B0;&#x539F;&#x7406;"></a>

池子里面的属性是令牌桶，这个桶子用有缓冲chan来实现，实现了自动阻塞

### 带有接口那味的实现 <a id="&#x5E26;&#x6709;&#x63A5;&#x53E3;&#x90A3;&#x5473;&#x7684;&#x5B9E;&#x73B0;"></a>

```text
workpool.go
package work

// 什么是池子，池子就是他一直有地址空间，你上传就Ok了
// 方法:
// 上传 + 执行任务
// 当前任务空闲
// 属性： 令牌桶
type IWorkPool interface {
	Submit(work IWork)
	GetWorkCount() int
}

type WorkPool struct {
	Tocken chan int
}

func NewWorkPool(limit int) *WorkPool {
	wp := &WorkPool{
		Tocken:make(chan int, limit),
	}
	for i := 0; i < limit; i++ {
		wp.Tocken<- i
	}
	return  wp
}

func (wp *WorkPool) Submit(work IWork) {
	<- wp.Tocken
	go func() {
		work.DoWork()
		wp.Tocken<-1
	}()
}

func (wp *WorkPool) GetWorkCount() int {
	return len(wp.Tocken)
}
```

```text
work.go
package work

import "fmt"

// 属性：
// work func()
//方法：
// DoWork()
type IWork interface {
	DoWork()
}

// 具体任务
type Work struct {
	work func()
}

func NewWork(fn func()) *Work {
	return &Work{
		work: fn,
	}
}

func (w *Work) DoWork() {
	w.work()
	fmt.Println("任务ok")
}
```

```text
func main() {
	pool := work.NewWorkPool(4)
	for i := 0; i < 100; i++ {
		newWork := work.NewWork(func() {
			time.Sleep(3 * time.Second)
		})
		pool.Submit(newWork)
	}

	//time.Sleep(time.Second * 3)
	time.Sleep(10000 * time.Second)
}
```

1. 结构体源码
2. 创建channel
3. 读取数据
4. channel close关闭channel源码分析

### mutex锁

1. 基本接口
2. 两种模式
3. 信号量概念
4. 加锁
5. 解锁

### GMP模型内存管理（goroutine）

1. 协程的基本概念
2. 阻塞式IO和非阻塞式IO，IO多路复用
3. GMP模型
4. 线程池-》GM模型-》GPM模型演变-》调度细节

### 闭包问题

### 反射问题

### 几种算法

### 单例模式

### 类型元数据

### OS和网络 容器 k8s 

Kubernetes in Action 再看写的很烂的 cloneset advance stateful set源码（openkruise项目里）

设计模式、用代码写算法、PHP数组 array\_merge与‘+’区别（现在php少，可能还会问go的），gc原理，内存池等； 应用场景设计，主要分布式复杂系统设计思路，

### 难点讲座

Kavya Joshi这位印度姐姐的4部讲座真是厉害，直接讲清了Go里面的4大难点： Mutext机制 [https://www.bilibili.com/video/BV1kz411e7dL/](https://www.bilibili.com/video/BV1kz411e7dL/) Go Race检测机制 [https://www.bilibili.com/video/BV1Ta4y1a7Wd/](https://www.bilibili.com/video/BV1Ta4y1a7Wd/) Channel原理 [https://www.bilibili.com/video/BV1eT4y177HN](https://www.bilibili.com/video/BV1eT4y177HN) goroutine 调度机制GMP模型 [https://www.bilibili.com/video/BV1vT4y177tA](https://www.bilibili.com/video/BV1vT4y177tA)



