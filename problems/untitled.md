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

### foreach的拷贝问题

for range是值拷贝出来的副本

可以用指针数组，value值用"\_"舍弃了元素的复制，用下标去访问\(效率更高）

```text
for i, _ := range t {
    t[i].Num += 100
}
```

### go执行的随机性和闭包



1. go的继承与组合
2. select随机性
3. defer中插入函数的执行顺序
4. make默认值和append
5. map线程安全
6. chan缓存池
7. golang的方法集
8. interface内部结构
9. type类型断言
10. 函数返回值命名
11. defer和函数返回值
12. new和make的问题
13. append切片加上...
14. 结构体比较
15. interface内部解构
16. 函数返回值类型
17. iota和const系列变量复制
18. 变量简短模式
19. 常量在预处理阶段直接展开
20. goto不能跳转到其他函数或者内层代码
21. Type Alias和Type definition
22. Type Alias ，引用方法的区别
23. Type Alias ，结构体内部字段的区别
24. 变量作用域
25. 闭包延迟求值
26. 闭包引用相同变量
27. panic仅有最后一个可以被recover捕获
28. 计算结构体大小
29. 字符串转成byte数组，会发生内存拷贝吗？
30. 拷贝大切片一定比小切片代价大吗
31. 能说说unitptr和unsafe.Pointer的区别吗？
32. reflect（反射包）如何获取字段tag？为什么json包不能导出私有变量的tag？
33. 怎么避免内存逃逸？
34. 反转含有中文，数字，英文字母的字符串
35. 知道golang的内存逃逸吗？什么情况下会发生内存逃逸？
36. 悬挂指针的问题

## Go advance

### 垃圾回收机制

1. 写屏障
2. 读屏障
3. 并发垃圾回收
4. go heap对象分布
5. 混合读写之后为什么能减少一次栈重扫

### channel的本质

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

闭包问题

反射问题

几种算法

单例模式

类型元数据

OS和网络 容器 k8s 

Kubernetes in Action 再看写的很烂的 cloneset advance stateful set源码（openkruise项目里）

Kavya Joshi这位印度姐姐的4部讲座真是厉害，直接讲清了Go里面的4大难点： Mutext机制 [https://www.bilibili.com/video/BV1kz411e7dL/](https://www.bilibili.com/video/BV1kz411e7dL/) Go Race检测机制 [https://www.bilibili.com/video/BV1Ta4y1a7Wd/](https://www.bilibili.com/video/BV1Ta4y1a7Wd/) Channel原理 [https://www.bilibili.com/video/BV1eT4y177HN](https://www.bilibili.com/video/BV1eT4y177HN) goroutine 调度机制GMP模型 [https://www.bilibili.com/video/BV1vT4y177tA](https://www.bilibili.com/video/BV1vT4y177tA)





