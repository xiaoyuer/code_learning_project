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

1. defer执行顺序
2. foreach的拷贝问题
3. go执行的随机性和闭包
4. go的继承与组合
5. select随机性
6. defer中插入函数的执行顺序
7. make默认值和append
8. map线程安全
9. chan缓存池
10. golang的方法集
11. interface内部结构
12. type类型断言
13. 函数返回值命名
14. defer和函数返回值
15. new和make的问题
16. append切片加上...
17. 结构体比较
18. interface内部解构
19. 函数返回值类型
20. iota和const系列变量复制
21. 变量简短模式
22. 常量在预处理阶段直接展开
23. goto不能跳转到其他函数或者内层代码
24. Type Alias和Type definition
25. Type Alias ，引用方法的区别
26. Type Alias ，结构体内部字段的区别
27. 变量作用域
28. 闭包延迟求值
29. 闭包引用相同变量
30. panic仅有最后一个可以被recover捕获
31. 计算结构体大小
32. 字符串转成byte数组，会发生内存拷贝吗？
33. 拷贝大切片一定比小切片代价大吗
34. 能说说unitptr和unsafe.Pointer的区别吗？
35. reflect（反射包）如何获取字段tag？为什么json包不能导出私有变量的tag？
36. 怎么避免内存逃逸？
37. 反转含有中文，数字，英文字母的字符串
38. 知道golang的内存逃逸吗？什么情况下会发生内存逃逸？
39. 悬挂指针的问题

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





