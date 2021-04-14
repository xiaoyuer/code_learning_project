# 进程和线程

1、程序

程序是一个按指定格式存储了一系列指令的编码序列。

打个比方的话，程序就好像一张菜谱，它原原本本精确记录了某道菜的整个制作流程。

2、操作系统

操作系统也是程序的一种。

它的作用是管理硬件，忽略厂商实现差异，给程序员提供一个统一的访问界面；管理其他程序，在用户需要时加载、运行它们。

操作系统可以从不同侧面划分成很多种类型。比如从预期响应时间可以分为实时系统和非实时系统；从是否允许多个程序同时被执行分可以分为单进程系统、多进程系统，多进程系统还可以分为非抢夺式多任务和抢夺式多任务等；从内核设计思路可以分微内核系统、宏内核系统……

总之，这里面的道道很多，不然CS专业的操作系统原理这门课也不会是那么厚的一本书、并且每年还要挂那么多人。

3、进程

我们把一个被载入内存、正在执行的程序叫做一个进程。

注意进程的关键点是“正在被执行”。比如你可以同时打开50个QQ，这50个QQ是同一个程序（QQ.exe），但它在内存的50份拷贝是50个不同的进程。

还是用菜谱来打比方的话，“西红柿炒鸡蛋”这张菜谱是“程序”；你照着菜谱做这道菜是一个进程。这道菜每天中午神州大地都有成千上万厨师在做，他们做这道菜的每一次实践也都是一个不同的进程。

4、线程

“传统”观念下，一个程序只有一个执行点，就好像一张菜谱是一个厨师炒出来的一样。

但事实上，和一张菜谱可以让多个厨师分头执行它的不同部分一样，一个程序也完全可以包含两个以上的执行点，从而利用多CPU/核心以及不同硬件同时做几件事。

比如说，完全可以在等待网络报文的同时把已有数据先排个序、建个索引什么的，不至于网络包没过来整个程序都没法动，把其他硬件晾一边。

一个程序允许多个执行点（执行现场）就叫多线程。

线程可以由操作系统直接调度，也可以由用户自己写一段代码，自己管理多个代码段的执行切换动作。后者就是所谓的“用户态多线程”——有人混淆了“有特殊硬件时某种OS的线程的具体实现”和“线程概念”本身，这是一知半解的典型表现。想想nobody权限下用户态线程库如何实现，或可帮他把搅来搅去乱成一团的糊涂认识分离开。

换句话说，线程既可以由操作系统实现，也可以自己写程序实现。前者的好处是可以利用CPU里的多个核心（或多颗CPU）；后者虽然无法利用多核心/多CPU，但可以通过灵活的执行现场切换，更方便某些有特定需要的程序的设计。

线程和进程的区别在于，进程拥有自己的资源，而线程直接使用分配给进程的资源，它自己不能占有资源。
