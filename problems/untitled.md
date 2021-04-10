# Go Problems

1. make和new的区别
2. defer执行顺序
3. foreach的拷贝问题
4. go执行的随机性和闭包
5. go的继承与组合
6. select随机性
7. defer中插入函数的执行顺序
8. make默认值和append
9. map线程安全
10. chan缓存池
11. golang的方法集
12. interface内部结构
13. type类型断言
14. 函数返回值命名
15. defer和函数返回值
16. new和make的问题
17. append切片加上...
18. 结构体比较
19. interface内部解构
20. 函数返回值类型
21. iota和const系列变量复制
22. 变量简短模式
23. 常量在预处理阶段直接展开
24. goto不能跳转到其他函数或者内层代码
25. Type Alias和Type definition
26. Type Alias ，引用方法的区别
27. Type Alias ，结构体内部字段的区别
28. 变量作用域
29. 闭包延迟求值
30. 闭包引用相同变量
31. panic仅有最后一个可以被recover捕获
32. 计算结构体大小
33. 字符串转成byte数组，会发生内存拷贝吗？
34. 拷贝大切片一定比小切片代价大吗
35. 能说说unitptr和unsafe.Pointer的区别吗？
36. reflect（反射包）如何获取字段tag？为什么json包不能导出私有变量的tag？
37. 怎么避免内存逃逸？
38. 反转含有中文，数字，英文字母的字符串
39. 知道golang的内存逃逸吗？什么情况下会发生内存逃逸？
40. 悬挂指针的问题

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





