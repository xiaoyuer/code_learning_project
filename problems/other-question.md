# Other Question

## 项目架构介绍



## Mysql

[https://zronghui.github.io/%E8%AF%BB%E4%B9%A6%E7%AC%94%E8%AE%B0/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4-MYSQL45%E8%AE%B2-%E7%AC%94%E8%AE%B0.html](https://zronghui.github.io/%E8%AF%BB%E4%B9%A6%E7%AC%94%E8%AE%B0/%E6%9E%81%E5%AE%A2%E6%97%B6%E9%97%B4-MYSQL45%E8%AE%B2-%E7%AC%94%E8%AE%B0.html)

### 索引怎么优化

合索引的列是出现在where子句中的列，或者连接子句中指定的列； 2）基数较小的类，索引效果较差，没有必要在此列建立索引； 3）使用短索引，如果对长字符串列进行索引，应该指定一个前缀长度，这样能够节省大量索引空间； 4）不要过度索引。索引需要额外的磁盘空间，并降低写操作的性能。在修改表内容的时候，索引会进行更新甚至重构，索引列越多，这个时间就会越长。所以只保持需要的索引有利于查询即可。

前导模糊查询不能命中索引： EXPLAIN SELECT \* FROM user WHERE name LIKE '%s%';

1. InnoDB Buffer size 足够的情况下，即所有数据与索引能完成全加载进内存，查询不会有大问题。如不能会引发内存-硬盘数据交换，看硬盘IO就知道交换频率了。 一般IO在50%以下，性能也不会有大问题。如果 IO到达 100%，整个MySQL不响应的很可能会发生。当然如有其它数据在另一个硬盘，旧连接是不影响的--前提是没有使用swap。当然，带索引的同一语句，也有可能出现非常小量小量的一次慢--原因不明。2. 单表数据过大，维护确实是非常痛苦，alter就别想了；基本上不可能热备，甚至冷备也不可能-- Innodbbackex除外。 3.  大量的sum与count操作 ，即使where条件有索引，对IO也会比小表大1倍。原因不太清楚，曾经有位有装13经验的人跟我说是索引引起IO--- 我想不明白，索引都load到内存了，如何会有IO？一个行数13亿的表，不一定比1亿的表大。还要看表结构。 总体来讲，我认为影响 应该是表+索引 是否能完全装裁到内存。 而当13E行的表分成1024个小表后，这时sum落到是1/1024表中，我理解是单次的IO时间小，并发qps不变的情况下，总IO降了1倍--这个是实际生产环境数据。4.  高IO会引发雪崩，搞不好整个MySQL挂掉。单硬盘安全峰会IO最好在30%以下。

### 一亿条数据查询优化

以主键水平分割表！

如果一年前的只是备份待查,分离出来另存. 如果一年前的会用到,但用得少,用分区. 如果一年前的仍然要频繁使用,用分区,但要加一个磁盘.

### 横向纵向分表

首先存储引擎的使用不同，冷数据使用MyIsam 可以有更好的查询数据。活跃数据，可以使用Innodb,可以有更好的更新速度。 其次，对冷数据进行更多的从库配置，因为更多的操作时查询，这样来加快查询速度。对热数据，可以相对有更多的主库的横向分表处理。 其实，对于一些特殊的活跃数据，也可以考虑使用memcache ,redis 之类的缓存，等累计到一定量再去更新数据库。或者mongodb 一类的nosql数据库，这里只是举例，就先不说这个。 横向分表 字面意思，就可以看出来，是把大的表结构，横向切割为同样结构的不同表，如，用户信息表，user\_1,user\_2等。表结构是完全一样，但是，根据某些特定的规则来划分的表，如根据用户ID来取模划分。 分表理由：根据数据量的规模来划分，保证单表的容量不会太大，从而来保证单表的查询等处理能力。

### 存储引擎对比

### 学号姓名课程成绩计算总分后倒序排序，找出各科及格的学号姓名

### 事务隔离级别

1. 读未提交 （脏读）
2. 读提交  （不可重复读）
3. 可重复读
4. 串行

MySQL原理（事务隔离级别、mvcc机制原理、B+树（每个节点一般存一页4K数据，4K是因为一次加载数据入内存，局部性原理，性能高，从而决定了一张表最大只能存多少数据），推荐看极客时间的《 MySQL45讲》）、MySQL主从同步原理及详细过程、MySQL两阶段提交，写入数据详细过程，事务详细过程、binlog、redolog、undolog分别作用；

### limit 优化

\#\#where

**第一次优化**

根据数据库这种查找的特性，就有了一种想当然的方法，利用自增索引（假设为id）：

```text
select * from table_name where (id >= 10000) limit 10
```

由于普通搜索是全表搜索，适当的添加 WHERE 条件就能把搜索从全表搜索转化为范围搜索，大大缩小搜索的范围，从而提高搜索效率。

这个优化思路就是告诉数据库：「你别数了，我告诉你，第10001条数据是这样的，你直接去拿吧。」

但是！！！你可能已经注意到了，这个查询太简单了，没有任何的附加查询条件，如果我需要一些额外的查询条件，比如我只要某个用户的数据 ，这种方法就行不通了。

可以见到这种思路是有局限性的，首先必须要有自增索引列，而且数据在逻辑上必须是连续的，其次，你还必须知道特征值。

如此苛刻的要求，在实际应用中是不可能满足的。

**第二次优化**

说起数据库查询优化，第一时间想到的就是索引，所以便有了第二次优化：先查找出需要数据的索引列（假设为 id），再通过索引列查找出需要的数据。

```text
Select * From table_name Where id in (Select id From table_name where ( user = xxx )) limit 10000, 10;

select * from table_name where( user = xxx ) limit 10000,10
```

相比较结果是（500w条数据）：第一条花费平均耗时约为第二条的 1/3 左右。

同样是较大的 offset，第一条的查询更为复杂，为什么性能反而得到了提升？

这涉及到 mysql 主索引的数据结构 b+Tree ，这里不展开，基本原理就是：

* 子查询只用到了索引列，没有取实际的数据，所以不涉及到磁盘IO，所以即使是比较大的 offset 查询速度也不会太差。
* 利用子查询的方式，把原来的基于 user 的搜索转化为基于主键（id）的搜索，主查询因为已经获得了准确的索引值，所以查询过程也相对较快。

**第三次优化**

在数据量大的时候 in 操作的效率就不怎么样了，我们需要把 in 操作替换掉，使用 join 就是一个不错的选择。

```text
select * from table_name inner join ( select id from table_name where (user = xxx) limit 10000,10) b using (id)
```

至此 limit 在查询上的优化就告一段落了。如果还有更好的优化方式，欢迎留言告知  
  


### mongodb groupby 

### MVCC 隔离机制的原理

### mysql 主从同步

### 并发写入

将评论直接推送到 rabbitMq之类的消息中间件，然后再开多进程进行消费插入mysql\(开几个根据服务器性能\)，这样负载就是可控的了

监控主从延迟的方法有多种：

1. Slave 使用本机当前时间，跟 Master 上 binlog 的时间戳比较
2. `pt-heartbeat`、`mt-heartbeat`

**本质**：同一条 SQL，`Master` 上`执行结束`的时间 vs. `Slave` 上`执行结束`的时间。

原因

* Master 上，`大事务`，耗时长：优化业务，拆分为小事务
  * Master 上，SQL 执行速度慢：优化索引，提升索引区分度（事务内部有查询操作）
  * Master 上，批量 DML 操作：建议延迟至业务低峰期操作
* Master 上，`多线程写入频繁`， Slave 单线程速度跟不上：提升 Slave 硬件性能、借助中间件，改善主从复制的单线程模式

### mysql锁  **InnoDB加锁方法：**

* 意向锁是 InnoDB 自动加的， 不需用户干预。
* 对于 UPDATE、 DELETE 和 INSERT 语句， InnoDB 会自动给涉及数据集加排他锁（X\)；
* 对于普通 SELECT 语句，InnoDB 不会加任何锁； 事务可以通过以下语句显式给记录集加共享锁或排他锁：
  * 共享锁（S）：SELECT \* FROM table\_name WHERE ... LOCK IN SHARE MODE。 其他 session 仍然可以查询记录，并也可以对该记录加 share mode 的共享锁。但是如果当前事务需要对该记录进行更新操作，则很有可能造成死锁。
  * 排他锁（X\)：SELECT \* FROM table\_name WHERE ... FOR UPDATE。其他 session 可以查询该记录，但是不能对该记录加共享锁或排他锁，而是等待获得锁
* **隐式锁定：**

InnoDB在事务执行过程中，使用两阶段锁协议：

随时都可以执行锁定，InnoDB会根据隔离级别在需要的时候自动加锁；

锁只有在执行commit或者rollback的时候才会释放，并且所有的锁都是在**同一时刻**被释放。



```text
select * from table_name limit 10000,10
```

这句 SQL 的执行逻辑是  
 1.从数据表中读取第N条数据添加到数据集中  
 2.重复第一步直到 N = 10000 + 10  
 3.根据 offset 抛弃前面 10000 条数  
 4.返回剩余的 10 条数据  
  


## **GMP调度**

\*\*\*\*[**https://blog.csdn.net/weixin\_44879611/article/details/105374175**](https://blog.csdn.net/weixin_44879611/article/details/105374175)\*\*\*\*

### **调度算法**

当一个Goroutine创建被创建时，Goroutine对象被压入Processor的本地队列或者Go运行时 全局Goroutine队列。Processor唤醒一个Machine，如果Machine的waiting队列没有等待被 唤醒的Machine，则创建一个（只要不超过Machine的最大值，10000），Processor获取到Machine后，与此Machine绑定，并执行此Goroutine。Machine执行过程中，随时会发生上下文切换。当发生上下文切换时，需要对执行现场进行保护，以便下次被调度执行时进行现场恢复。Go调度器中Machine的栈保存在Goroutine对象上，只需要将Machine所需要的寄存器\(堆栈指针、程序计数器等\)保存到Goroutine对象上即可。如果此时Goroutine任务还没有执行完，Machine可以将Goroutine重新压入Processor的队列，等待下一次被调度执行。 如果执行过程遇到阻塞并阻塞超时（调度器检测这种超时），Machine会与Processor分离，并等待阻塞结束。此时Processor可以继续唤醒Machine执行其它的Goroutine，当阻塞结束时，Machine会尝试”偷取”一个Processor，如果失败，这个Goroutine会被加入到全局队列中，然后Machine将自己转入Waiting队列，等待被再次唤醒。

在各个Processor运行完本地队列的任务时，会从全局队列中获取任务，调度器也会定期检查全局队列，否则在并发较高的情况下，全局队列的Goroutine会因为得不到调度而”饿死”。如果全局队列也为空的时候，会去分担其它Processor的任务，一次分一半任务，比如，ProcessorA任务完成了，ProcessorB还有10个任务待运行，Processor在获取任务的时候，会一次性拿走5个。（是不是觉得Processor相互之间很友爱啊 \_）。

### goroutine切换条件 <a id="goroutine&#x5207;&#x6362;&#x6761;&#x4EF6;"></a>

在正常情况下，scheduler（调度器）会按照上面的流程进行调度，当一个G（goroutine）的时间片结束后将P（Processor）分配给下一个G，但是线程会发生阻塞等情况，看一下goroutine对线程阻塞等的处理。  


## Redis

### Redis底层数据结构

1. 字符串
2. List 列表
3. Hash 字典
4. Set集合
5. Sorted Set 有序集合

### 跳表原理

### redis集群：主从，哨兵模式

### **什么是乐观锁和悲观锁**

1）乐观锁：就像它的名字一样，对于并发间操作产生的线程安全问题持乐观状态，乐观锁认为竞争不总是会发生，因此它不需要持有锁，将**比较-替换**这两个动作作为一个原子操作尝试去修改内存中的变量，如果失败则表示发生冲突，那么就应该有相应的重试逻辑。

2）悲观锁：还是像它的名字一样，对于并发间操作产生的线程安全问题持悲观状态，悲观锁认为竞争总是会发生，因此每次对某资源进行操作时，都会持有一个独占的锁，就像synchronized，不管三七二十一，直接上了锁就操作资源了。

### aof、RDB详细过程及线上怎么配置（一般AOF+RDB，此时原理是什么也要知道）

### 主从详细过程

### 分布式方案（TWproxy\codis）等原理（推荐看 黄建宏写的redis书）

### lru算法实现

### redis缓存雪崩问题

锁加数据库同步

本地锁，分布式锁，数据库锁，大规模线程阻塞，服务器线程被耗尽，服务不可用

### redis和mysql数据同步



## HashTable

### Hash碰撞

哈希（Hash）算法，即散列函数。它是一种单向密码体制，即它是一个从明文到密文的不可逆的映射。哈希函数可以将任意长度的输入经过变化以后得到固定长度的输出。哈希函数的这种单向特征和输出数据长度固定的特征使得它可以生成消息或者数据。

### Hash算法用途

1.数据校验

2.唯一标识

3.哈希表

4.负载均衡

{% embed url="https://5.分布式存储" %}

### 避免Hash碰撞策略

#### 1.开放地址法\(再散列法\)

开放地执法有一个公式:Hi=\(H\(key\)+di\) MOD m i=1,2,…,k\(k&lt;=m-1\) 其中，m为哈希表的表长。di 是产生冲突的时候的增量序列。如果di值可能为1,2,3,…m-1，称线性探测再散列。如果di取1，则每次冲突之后，向后移动1个位置.如果di取值可能为1,-1,2,-2,4,-4,9,-9,16,-16,…kk,-kk\(k&lt;=m/2\)，称二次探测再散列。如果di取值可能为伪随机数列。称伪随机探测再散列。

#### 2.再哈希法Rehash

当发生冲突时，使用第二个、第三个、哈希函数计算地址，直到无冲突时。缺点：计算时间增加。比如上面第一次按照姓首字母进行哈希，如果产生冲突可以按照姓字母首字母第二位进行哈希，再冲突，第三位，直到不冲突为止.这种方法不易产生聚集，但增加了计算时间。

#### 3.链地址法（拉链法）

将所有关键字为同义词的记录存储在同一线性链表中.基本思想:将所有哈希地址为i的元素构成一个称为同义词链的单链表，并将单链表的头指针存在哈希表的第i个单元中，因而查找、插入和删除主要在同义词链中进行。链地址法适用于经常进行插入和删除的情况。对比JDK 1.7 hashMap的存储结构是不是很好理解。至于1.8之后链表长度大于6rehash 为树形结构不在此处讨论。

#### 拉链法的优缺点

#### 优点：

①拉链法处理冲突简单，且无堆积现象，即非同义词决不会发生冲突，因此平均查找长度较短；②由于拉链法中各链表上的结点空间是动态申请的，故它更适合于造表前无法确定表长的情况；③开放定址法为减少冲突，要求装填因子α较小，故当结点规模较大时会浪费很多空间。而拉链法中可取α≥1，且结点较大时，拉链法中增加的指针域可忽略不计，因此节省空间；④在用拉链法构造的散列表中，删除结点的操作易于实现。只要简单地删去链表上相应的结点即可。而对开放地址法构造的散列表，删除结点不能简单地将被删结 点的空间置为空，否则将截断在它之后填人散列表的同义词结点的查找路径。这是因为各种开放地址法中，空地址单元\(即开放地址\)都是查找失败的条件。因此在 用开放地址法处理冲突的散列表上执行删除操作，只能在被删结点上做删除标记，而不能真正删除结点。

#### 缺点：

指针需要额外的空间，故当结点规模较小时，开放定址法较为节省空间，而若将节省的指针空间用来扩大散列表的规模，可使装填因子变小，这又减少了开放定址法中的冲突，从而提高平均查找速度。

#### 4.建立一个公共溢出区

假设哈希函数的值域为\[0,m-1\],则设向量HashTable\[0..m-1\]为基本表，另外设立存储空间向量OverTable\[0..v\]用以存储发生冲突的记录。

## Http Tcp

### 网页从输入链接到返回页面经历了什么？

简单说说，浏览器根据请求的url交给dns域名解析，查找真正的ip地址，向服务器发起请求；服务器交给后台处理后，返回数据，浏览器会接收到文件数据，比如，html,js，css，图像等；然后浏览器会对加载到的资源进行语法解析，建立相应的内部数据结构；载入解析到得资源文件，渲染页面，完成显示页面效果。

### 三次握手四次挥手

> 第一次挥手

主动关闭的一方，发送一个FIN\(上述讲过---当FIN=1，表示此报文段是一个释放连接的请求报文\)，传送数据，用来告诉对方（被动关闭方），说不会再给你发送数据了。---主动关闭的一方可以接受数据。

> 第二次挥手

被动关闭方 收到 FIN 包，发送 ACK 给对方，确认序号。

> 第三次挥手

被动关闭方 发送一个 FIN，关闭方，说我不会再给你发数据了。（你不给我发送数据，我也不给你发送数据了）

> 第四次挥手

主动关闭一方收到 FIN ，发送要给 ACK ，用来确认序号

###  TCP：滑动窗口

### 快重传

### 慢启动（二进制退避算法）

### 分包、拆包问题

### MAC头、IP头、tcp头

### http的keepalive--http1.1默认开启了keepalive

### https握手过程

### http2原理

### DNS原理

### CDN原理 

### http状态码含义，出现4XX，5XX如何定位问题）

### http头包含哪些东西

### http的chunk模式是啥

当客户端向服务器请求一个静态页面或者一张图片时，服务器可以很清楚的知道内容大小，然后通过Content-Length消息首部字段告诉客户端需要接收多少数据。但是如果是动态页面等时，服务器是不可能预先知道内容大小，这时就可以使用Transfer-Encoding：chunk模式来传输数据了。即如果要一边产生数据，一边发给客户端，服务器就需要使用"Transfer-Encoding: chunked"这样的方式来代替Content-Length。

在进行chunked编码传输时，在回复消息的头部有Transfer-Encoding: chunked

编码使用若干个chunk组成，由一个标明长度为0的chunk结束。每个chunk有两部分组成，第一部分是该chunk的长度，第二部分就是指定长度的内容，每个部分用CRLF隔开。在最后一个长度为0的chunk中的内容是称为footer的内容，是一些没有写的头部内容。



## IO多路复用

### poll / epoll

I/O多路复用这个概念被提出来以后， select是第一个实现 \(1983 左右在BSD里面实现的\)。select 被实现以后，很快就暴露出了很多问题。

* select 会修改传入的参数数组，这个对于一个需要调用很多次的函数，是非常不友好的。 
*  select 如果任何一个sock\(I/O stream\)出现了数据，select 仅仅会返回，但是并不会告诉你是那个sock上有数据，于是你只能自己一个一个的找，10几个sock可能还好，要是几万的sock每次都找一遍，这个无谓的开销就颇有海天盛筵的豪气了。 
* select 只能监视1024个链接， 这个跟草榴没啥关系哦，linux 定义在头文件中的，参见_FD\_SETSIZE。_
* select 不是线程安全的，如果你把一个sock加入到select, 然后突然另外一个线程发现，尼玛，这个sock不用，要收回。对不起，这个select 不支持的，如果你丧心病狂的竟然关掉这个sock, select的标准行为是。。呃。。不可预测的， 这个可是写在文档中的哦.

 “If a file descriptor being monitored by select\(\) is closed in another thread, the result is unspecified”  
 霸不霸气于是14年以后\(1997年）一帮人又实现了poll, poll 修复了select的很多问题，比如

* poll 去掉了1024个链接的限制，于是要多少链接呢， 主人你开心就好。 
*  poll 从设计上来说，不再修改传入数组，不过这个要看你的平台了，所以行走江湖，还是小心为妙。

**其实拖14年那么久也不是效率问题， 而是那个时代的硬件实在太弱，一台服务器处理1千多个链接简直就是神一样的存在了，select很长段时间已经满足需求。**

但是poll仍然不是线程安全的， 这就意味着，不管服务器有多强悍，你也只能在一个线程里面处理一组I/O流。你当然可以那多进程来配合了，不过然后你就有了多进程的各种问题。

于是5年以后, 在2002, 大神 Davide Libenzi 实现了epoll.epoll 可以说是I/O 多路复用最新的一个实现，epoll 修复了poll 和select绝大部分问题, 比如：

* epoll 现在是线程安全的。 
* epoll 现在不仅告诉你sock组里面数据，还会告诉你具体哪个sock有数据，你不用自己去找了。

### 线程与进程区别

### Hash表实现原理与冲突解决

## Nginx 

nginx高性能原因（IO多路复用技术 epoll原理与使用到的数据结构），senfile快的原因、文件事件、时间事件概念、nginx进程如何管理（master/worker分别干嘛的，惊群问题如何处理，进程管理的信号、数据同步的信号量两者区别）、平滑重启原理（go如何做）、平滑升级、平滑重启如何做、nginx11阶段是、keepalive及对应配置含义（应用很广泛，可能会问）、负载均衡有哪些方式及对应算法、限流算法、openresty原理及应用、nginx连接池；

## rabbitmq

### 作用：削峰，异步，解耦

消息 =&gt; 交换机 =&gt; 通过绑定的`key` =&gt; 队列 =&gt; 消费者

### 持久化

RabbitMQ 持久化包含3个部分

* exchange 持久化，在声明时指定 durable 为 true
* queue 持久化，在声明时指定 durable 为 true
* message 持久化，在投递时指定 delivery\_mode=2（1是非持久化）

queue 的持久化能保证本身的元数据不会因异常而丢失，但是不能保证内部的 message 不会丢失。要确保 message 不丢失，还需要将 message 也持久化

如果 exchange 和 queue 都是持久化的，那么它们之间的 binding 也是持久化的。

如果 exchange 和 queue 两者之间有一个持久化，一个非持久化，就不允许建立绑定。

> 注意：一旦确定了 exchange 和 queue 的 durable，就不能修改了。如果非要修改，唯一的办法就是删除原来的 exchange 或 queue 后，重现创建

### mq队列如何保证都push了

①：可以选择使用rabbitmq提供是事物功能，就是生产者在发送数据之前开启事物，然后发送消息，如果消息没有成功被rabbitmq接收到，那么生产者会受到异常报错，这时就可以回滚事物，然后尝试重新发送；如果收到了消息，那么就可以提交事物。

```text
  channel.txSelect();//开启事物
  try{
      //发送消息
  }catch(Exection e){
      channel.txRollback()；//回滚事物
      //重新提交
  }
复制代码
```

**缺点：** rabbitmq事物已开启，就会变为同步阻塞操作，生产者会阻塞等待是否发送成功，太耗性能会造成吞吐量的下降。

②：可以开启confirm模式。在生产者哪里设置开启了confirm模式之后，每次写的消息都会分配一个唯一的id，然后如何写入了rabbitmq之中，rabbitmq会给你回传一个ack消息，告诉你这个消息发送OK了；

### rabbitmq过程

{% embed url="https://1.声明MessageQueue" %}

a\)消费者是无法订阅或者获取不存在的MessageQueue中信息。

b\)消息被Exchange接受以后，如果没有匹配的Queue，则会被丢弃。  
a\) Exclusive：排他队列，如果一个队列被声明为排他队列，该队列仅对首次声明它的连接可见，并在连接断开时自动删除。这里需要注意三点：其一，排他队列是基于连接可见的，同一连接的不同信道是可以同时访问同一个连接创建的排他队列的。其二，“首次”，如果一个连接已经声明了一个排他队列，其他连接是不允许建立同名的排他队列的，这个与普通队列不同。其三，即使该队列是持久化的，一旦连接关闭或者客户端退出，该排他队列都会被自动删除的。这种队列适用于只限于一个客户端发送读取消息的应用场景。

b\) Auto-delete:自动删除，如果该队列没有任何订阅的消费者的话，该队列会被自动删除。这种队列适用于临时队列。

c\) Durable:持久化，这个会在后面作为专门一个章节讨论。

d\) 其他选项，例如如果用户仅仅想查询某一个队列是否已存在，如果不存在，不想建立该队列，仍然可以调用queue.declare，只不过需要将参数passive设为true，传给queue.declare，如果该队列已存在，则会返回true；如果不存在，则会返回Error，但是不会创建新的队列。

{% embed url="https://2.生产者发送消息" %}

a\) 如果是Direct类型，则会将消息中的RoutingKey与该Exchange关联的所有Binding中的BindingKey进行比较，如果相等，则发送到该Binding对应的Queue中。

如果是 Fanout 类型，则会将消息发送给所有与该 Exchange 定义过 Binding 的所有 Queues 中去，其实是一种广播行为。

\)如果是Topic类型，则会按照正则表达式，对RoutingKey与BindingKey进行匹配，如果匹配成功，则发送到对应的Queue中。

3 . 消费者订阅消息  
在RabbitMQ中消费者有2种方式获取队列中的消息:

a\) 一种是通过basic.consume命令，订阅某一个队列中的消息,channel会自动在处理完上一条消息之后，接收下一条消息。（同一个channel消息处理是串行的）。除非关闭channel或者取消订阅，否则客户端将会一直接收队列的消息。

b\) 另外一种方式是通过basic.get命令主动获取队列中的消息，但是绝对不可以通过循环调用basic.get来代替basic.consume，这是因为basic.get RabbitMQ在实际执行的时候，是首先consume某一个队列，然后检索第一条消息，然后再取消订阅。如果是高吞吐率的消费者，最好还是建议使用basic.consume。

4 . 持久化： Rabbit MQ默认是不持久队列、Exchange、Binding以及队列中的消息的，这意味着一旦消息服务器重启，所有已声明的队列，Exchange，Binding以及队列中的消息都会丢失。通过设置Exchange和MessageQueue的durable属性为true，可以使得队列和Exchange持久化，但是这还不能使得队列中的消息持久化，这需要生产者在发送消息的时候，

## 面向对象

**面向对象**的**三大**特性是"封装、"多态"、"继承"，五大原则是"单一职责原则"、"开放封闭原则"、"里氏替换原则"、"依赖倒置原则"、"接口分离原则"、"迪米特原则（高内聚低耦合）"。

### 微服务

其中软件由通过明确定义的 API 进行通信的小型独立服务组成。这些服务由各个小型独立团队负责。

### 微服务的特点

1. 单一职责 每个微服务都需要满足单一职责原则，微服务本身是内聚的，因此微服务通常比较小。比如示例中每个微服务按业务逻辑划分，每个微服务仅负责自己归属于自己业务领域的功能。
2. 自治 一个微服务就是一个独立的实体，它可以独立部署、升级，服务与服务之间通过REST等形式的标准接口进行通信，并且一个微服务实例可以被替换成另一种实现，而对其它的微服务不产生影响。
3. 简化部署 在一个单块系统中，只要修改了一行代码，就需要对整个系统进行重新的构建、测试，然后将整个系统进行部署。而微服务则可以对一个微服务进行部署。
4. 可扩展 应对系统业务增长的方法通常采用横向（Scale out）或纵向（Scale up）的方向进行扩展。分布式系统中通常要采用Scale out的方式进行扩展。因为不同的功能会面对不同的负荷变化，因此采用微服务的系统相对单块系统具备更好的可扩展性。
5. 灵活组合 在微服务架构中，可以通过组合已有的微服务以达到功能重用的目的。
6. 技术异构 在一个大型系统中，不同的功能具有不同的特点，并且不同的团队可能具备不同的技术能力。因为微服务间松耦合，不同的微服务可以选择不同的技术栈进行开发。 同时，在应用新技术时，可以仅针对一个微服务进行快速改造，而不会影响系统中的其它微服务，有利于系统的演进。
7. 高可靠 微服务间独立部署，一个微服务的异常不会导致其它微服务同时异常。通过隔离、融断等技术可以避免极大的提升微服务的可靠性。
8. 基础设施自动化
9. 服务组件化
10. 容错设计
11. 演进式设计

### 微服务的缺点

1. 复杂度高 微服务间通过REST、RPC等形式交互，相对于Monolithic模式下的API形式，需要考虑被调用方故障、过载、消息丢失等各种异常情况，代码逻辑更加复杂。 对于微服务间的事务性操作，因为不同的微服务采用了不同的数据库，将无法利用数据库本身的事务机制保证一致性，需要引入二阶段提交等技术。 同时，在微服务间存在少部分共用功能但又无法提取成微服务时，各个微服务对于这部分功能通常需要重复开发，或至少要做代码复制，以避免微服务间的耦合，增加了开发成本。
2. 运维复杂 在采用微服务架构时，系统由多个独立运行的微服务构成，需要一个设计良好的监控系统对各个微服务的运行状态进行监控。运维人员需要对系统有细致的了解才对够更好的运维系统。
3. 影响性能 微服务的间通过REST、RPC等形式进行交互，通信的时延会受到较大的影响。



### RPC 简介

1. 远程过程调用（Remote Procedure Call，缩写为 RPC）
2. 可以将一些比较通用的场景抽象成微服务，然后供其他系统远程调用
3. RPC 可以基于HTTP协议 也可以基于TCP协议，基于HTTP协议的RPC像是我们访问网页一样（GET/POST/PUT/DELETE/UPDATE），大部分的RPC都是基于TPC协议的（因为基于传输层，效率稍高一些）
4. 基于TCP 的 RPC 工作过程
   1. 客户端对请求的对象序列化
   2. 客户端连接服务端，并将序列化的对象通过socket 传输给服务端，并等待接收服务端的响应
   3. 服务端收到请求对象后将其反序列化还原客户端的对象
   4. 服务端从请求对象中获取到请求的参数，然后执行对应的方法，得到返回结果
   5. 服务端将其结果序列化并传给客户端，客户端得到响应结果对象后将其反序列化，得到响应结果
5.  Golang中的RPC 注：例子参考 [golang实现RPC的几种方式](https://studygolang.com/articles/14336)

### 微服务相关，自己封的rpc使用什么协议？http？thrift？grpc？

基于TCP 的 RPC 工作过程

1. 客户端对请求的对象序列化
2. 客户端连接服务端，并将序列化的对象通过socket 传输给服务端，并等待接收服务端的响应
3. 服务端收到请求对象后将其反序列化还原客户端的对象
4. 服务端从请求对象中获取到请求的参数，然后执行对应的方法，得到返回结果
5. 服务端将其结果序列化并传给客户端，客户端得到响应结果对象后将其反序列化，得到响应结果

### 内存分析工具 内存溢出 火焰图 

### map里可以key 类型限制，value无类型限制

golang中的map，的 key 可以是很多种类型，比如 bool, 数字，string, 指针, channel , 还有 只包含前面几个类型的 interface types, structs, arrays 

显然，slice， map 还有 function 是不可以了，因为这几个没法用 == 来判断

### GMP模型 goroutine存在哪里 

### php map不是无序的，go map为什么无序 

它生成了随机数。用于决定从哪里开始循环迭代。更具体的话就是根据随机数，选择一个桶位置作为起始点进行遍历迭代

因此每次重新 `for range map`，你见到的结果都是不一样的。那是因为它的起始位置根本就不固定！

* 从已选定的桶中开始进行遍历，寻找桶中的下一个元素进行处理
* 如果桶已经遍历完，则对溢出桶 `overflow buckets` 进行遍历处理
* [https://cloud.tencent.com/developer/article/1422355](https://cloud.tencent.com/developer/article/1422355)

![](../.gitbook/assets/image%20%2837%29.png)

1. 介绍自己做过的项目，突出难点解决方案
2. 业务配置如何热更新，不重启服务
3. kafka比mq用的多，需要看，mq的模式fanout和direct区别，业务用的哪个
4. Lru是不是自己实现的（去年字节直接让手撸Lru），Lru如果有多个实例怎么配置？（之前的业务场景只用了一个）
5. 6. 开源error包和errors.New的区别

### 

1. IO多路复用：select poll epoll
2. 网校的开源go-zero框架，我没参与。。。
3. context用法
4. 内存分配（想问没问，悄悄说了一句）
5. sync和atomic，sync.pull

## 实际问题

### 数据大了分表分库

### 设计一个类似微博、微信、12306等等系统

### 边界考虑充分

### 双活建设、宏观架构

### 遇到问题的解决思路如500、499、502、504，服务治理理念等等；

## 系统：

### 内存分段分页原理、

### 协程、线程、进程概念、进程状态（结合状态问为什么出现，如僵尸进程）、

### go的协程原理、软中断硬中断概念、上下文切换概念、cpu负载概念等等、linux命名常用的柑橘够了，如分析nginx请求日志的访问top10的IP或URL;



## Gin框架 

### Static静态文件配置



### ToLearn

spark实现实时统计指标

循环队列，622 

解决问题的方法论

反转链表

3sum

linux io模型，协程是什么，redis cmd， mysql 

实现LRU

meeting room 252

求x的n次方（二分 递归）

## 进程间通信

**第一类：传统的Unix通信机制**  
 **1. 管道/匿名管道\(pipe\)**

* 管道是半双工的，数据只能向一个方向流动；需要双方通信时，需要建立起两个管道。
* 只能用于父子进程或者兄弟进程之间\(具有亲缘关系的进程\);
* 单独构成一种独立的文件系统：管道对于管道两端的进程而言，就是一个文件，但它不是普通的文件，它不属于某种文件系统，而是自立门户，单独构成一种文件系统，并且只存在与内存中。
* 数据的读出和写入：一个进程向管道中写的内容被管道另一端的进程读出。写入的内容每次都添加在管道缓冲区的末尾，并且每次都是从缓冲区的头部读出数据。 ![](//upload-images.jianshu.io/upload_images/1281379-05378521a7b41af4.png?imageMogr2/auto-orient/strip|imageView2/2/w/228/format/webp)进程间管道通信模型

**管道的实质：**  
 管道的实质是一个内核缓冲区，进程以先进先出的方式从缓冲区存取数据，管道一端的进程顺序的将数据写入缓冲区，另一端的进程则顺序的读出数据。  
 该缓冲区可以看做是一个循环队列，读和写的位置都是自动增长的，不能随意改变，一个数据只能被读一次，读出来以后在缓冲区就不复存在了。  
 当缓冲区读空或者写满时，有一定的规则控制相应的读进程或者写进程进入等待队列，当空的缓冲区有新数据写入或者满的缓冲区有数据读出来时，就唤醒等待队列中的进程继续读写。

**管道的局限：**  
 管道的主要局限性正体现在它的特点上：

* 只支持单向数据流；
* 只能用于具有亲缘关系的进程之间；
* 没有名字；
* 管道的缓冲区是有限的（管道制存在于内存中，在管道创建时，为缓冲区分配一个页面大小）；
* 管道所传送的是无格式字节流，这就要求管道的读出方和写入方必须事先约定好数据的格式，比如多少字节算作一个消息（或命令、或记录）等等；

**2. 有名管道\(FIFO\)**  
 匿名管道，由于没有名字，只能用于亲缘关系的进程间通信。为了克服这个缺点，提出了有名管道\(FIFO\)。  
 有名管道不同于匿名管道之处在于它提供了一个路径名与之关联，**以有名管道的文件形式存在于文件系统中**，这样，**即使与有名管道的创建进程不存在亲缘关系的进程，只要可以访问该路径，就能够彼此通过有名管道相互通信**，因此，通过有名管道不相关的进程也能交换数据。值的注意的是，有名管道严格遵循**先进先出\(first in first out\)**,对匿名管道及有名管道的读总是从开始处返回数据，对它们的写则把数据添加到末尾。它们不支持诸如lseek\(\)等文件定位操作。**有名管道的名字存在于文件系统中，内容存放在内存中。**

> **匿名管道和有名管道总结：**  
>  （1）管道是特殊类型的文件，在满足先入先出的原则条件下可以进行读写，但不能进行定位读写。  
>  （2）匿名管道是单向的，只能在有亲缘关系的进程间通信；有名管道以磁盘文件的方式存在，可以实现本机任意两个进程通信。  
>  （3）**无名管道阻塞问题：**无名管道无需显示打开，创建时直接返回文件描述符，在读写时需要确定对方的存在，否则将退出。如果当前进程向无名管道的一端写数据，必须确定另一端有某一进程。如果写入无名管道的数据超过其最大值，写操作将阻塞，如果管道中没有数据，读操作将阻塞，如果管道发现另一端断开，将自动退出。  
>  （4）**有名管道阻塞问题：**有名管道在打开时需要确实对方的存在，否则将阻塞。即以读方式打开某管道，在此之前必须一个进程以写方式打开管道，否则阻塞。此外，可以以读写（O\_RDWR）模式打开有名管道，即当前进程读，当前进程写，不会阻塞。

[延伸阅读：该博客有匿名管道和有名管道的C语言实践](https://link.jianshu.com?t=http://blog.chinaunix.net/uid-26833883-id-3227144.html)

**3. 信号\(Signal\)**

* 信号是Linux系统中用于进程间互相通信或者操作的一种机制，信号可以在任何时候发给某一进程，而无需知道该进程的状态。
* 如果该进程当前并未处于执行状态，则该信号就有内核保存起来，知道该进程回复执行并传递给它为止。
* 如果一个信号被进程设置为阻塞，则该信号的传递被延迟，直到其阻塞被取消是才被传递给进程。

> **Linux系统中常用信号：**  
>  （1）**SIGHUP：**用户从终端注销，所有已启动进程都将收到该进程。系统缺省状态下对该信号的处理是终止进程。  
>  （2）**SIGINT：**程序终止信号。程序运行过程中，按`Ctrl+C`键将产生该信号。  
>  （3）**SIGQUIT：**程序退出信号。程序运行过程中，按`Ctrl+\\`键将产生该信号。  
>  （4）**SIGBUS和SIGSEGV：**进程访问非法地址。  
>  （5）**SIGFPE：**运算中出现致命错误，如除零操作、数据溢出等。  
>  （6）**SIGKILL：**用户终止进程执行信号。shell下执行`kill -9`发送该信号。  
>  （7）**SIGTERM：**结束进程信号。shell下执行`kill 进程pid`发送该信号。  
>  （8）**SIGALRM：**定时器信号。  
>  （9）**SIGCLD：**子进程退出信号。如果其父进程没有忽略该信号也没有处理该信号，则子进程退出后将形成僵尸进程。

**信号来源**  
 信号是软件层次上对中断机制的一种模拟，是一种异步通信方式，，信号可以在用户空间进程和内核之间直接交互，内核可以利用信号来通知用户空间的进程发生了哪些系统事件，信号事件主要有两个来源：

* 硬件来源：用户按键输入`Ctrl+C`退出、硬件异常如无效的存储访问等。
* 软件终止：终止进程信号、其他进程调用kill函数、软件异常产生信号。

**信号生命周期和处理流程**  
 （1）信号被某个进程产生，并设置此信号传递的对象（一般为对应进程的pid），然后传递给操作系统；  
 （2）操作系统根据接收进程的设置（是否阻塞）而选择性的发送给接收者，如果接收者阻塞该信号（且该信号是可以阻塞的），操作系统将暂时保留该信号，而不传递，直到该进程解除了对此信号的阻塞（如果对应进程已经退出，则丢弃此信号），如果对应进程没有阻塞，操作系统将传递此信号。  
 （3）目的进程接收到此信号后，将根据当前进程对此信号设置的预处理方式，暂时终止当前代码的执行，保护上下文（主要包括临时寄存器数据，当前程序位置以及当前CPU的状态）、转而执行中断服务程序，执行完成后在回复到中断的位置。当然，对于抢占式内核，在中断返回时还将引发新的调度。  
![](//upload-images.jianshu.io/upload_images/1281379-3eed8cca67aa9f55.png?imageMogr2/auto-orient/strip|imageView2/2/w/889/format/webp)信号的生命周期

**4. 消息\(Message\)队列**

* 消息队列是存放在内核中的消息链表，每个消息队列由消息队列标识符表示。
* 与管道（无名管道：只存在于内存中的文件；命名管道：存在于实际的磁盘介质或者文件系统）不同的是消息队列存放在内核中，只有在内核重启\(即，操作系统重启\)或者显示地删除一个消息队列时，该消息队列才会被真正的删除。
* 另外与管道不同的是，消息队列在某个进程往一个队列写入消息之前，并不需要另外某个进程在该队列上等待消息的到达。[延伸阅读：消息队列C语言的实践](https://link.jianshu.com?t=http://blog.csdn.net/yang_yulei/article/details/19772649)

> **消息队列特点总结：**  
>  （1）消息队列是消息的链表,具有特定的格式,存放在内存中并由消息队列标识符标识.  
>  （2）消息队列允许一个或多个进程向它写入与读取消息.  
>  （3）管道和消息队列的通信数据都是先进先出的原则。  
>  （4）消息队列可以实现消息的随机查询,消息不一定要以先进先出的次序读取,也可以按消息的类型读取.比FIFO更有优势。  
>  （5）消息队列克服了信号承载信息量少，管道只能承载无格式字 节流以及缓冲区大小受限等缺。  
>  （6）目前主要有两种类型的消息队列：POSIX消息队列以及System V消息队列，系统V消息队列目前被大量使用。系统V消息队列是随内核持续的，只有在内核重起或者人工删除时，该消息队列才会被删除。

**5. 共享内存\(share memory\)**

* 使得多个进程可以可以直接读写同一块内存空间，是最快的可用IPC形式。是针对其他通信机制运行效率较低而设计的。
* 为了在多个进程间交换信息，内核专门留出了一块内存区，可以由需要访问的进程将其映射到自己的私有地址空间。进程就可以直接读写这一块内存而不需要进行数据的拷贝，从而大大提高效率。
* 由于多个进程共享一段内存，因此需要依靠某种同步机制（如信号量）来达到进程间的同步及互斥。  [延伸阅读：Linux支持的主要三种共享内存方式：mmap\(\)系统调用、Posix共享内存，以及System V共享内存实践](https://link.jianshu.com?t=http://www.cnblogs.com/linuxbug/p/4882776.html) ![](//upload-images.jianshu.io/upload_images/1281379-adfde0d80334c1f8.png?imageMogr2/auto-orient/strip|imageView2/2/w/538/format/webp)共享内存原理图

**6. 信号量\(semaphore\)**  
 信号量是一个计数器，用于多进程对共享数据的访问，信号量的意图在于进程间同步。  
 为了获得共享资源，进程需要执行下列操作：  
 （1）**创建一个信号量**：这要求调用者指定初始值，对于二值信号量来说，它通常是1，也可是0。  
 （2）**等待一个信号量**：该操作会测试这个信号量的值，如果小于0，就阻塞。也称为P操作。  
 （3）**挂出一个信号量**：该操作将信号量的值加1，也称为V操作。

为了正确地实现信号量，信号量值的测试及减1操作应当是原子操作。为此，信号量通常是在内核中实现的。Linux环境中，有三种类型：**Posix（**[**可移植性操作系统接口**](https://link.jianshu.com?t=http://baike.baidu.com/link?url=hYEo6ngm9MlqsQHT3h28baIDxEooeSPX6wr_FdGF-F8mf7wDp2xJWIDtQWGEDxthtPNiJtlsw460g1_N0txJYa)**）有名信号量（使用Posix IPC名字标识）**、**Posix基于内存的信号量（存放在共享内存区中）**、**System V信号量（在内核中维护）**。这三种信号量都可用于进程间或线程间的同步。  
![](//upload-images.jianshu.io/upload_images/1281379-376528c40d03717e.png?imageMogr2/auto-orient/strip|imageView2/2/w/635/format/webp)两个进程使用一个二值信号量

![](//upload-images.jianshu.io/upload_images/1281379-a72c8fbe22340031.png?imageMogr2/auto-orient/strip|imageView2/2/w/613/format/webp)两个进程所以用一个Posix有名二值信号量![](//upload-images.jianshu.io/upload_images/1281379-a1b276fae9db985d.png?imageMogr2/auto-orient/strip|imageView2/2/w/284/format/webp)一个进程两个线程共享基于内存的信号量

> **信号量与普通整型变量的区别：**  
>  （1）信号量是非负整型变量，除了初始化之外，它只能通过两个标准原子操作：wait\(semap\) , signal\(semap\) ; 来进行访问；  
>  （2）操作也被成为PV原语（P来源于荷兰语proberen"测试"，V来源于荷兰语verhogen"增加"，P表示通过的意思，V表示释放的意思），而普通整型变量则可以在任何语句块中被访问；

> **信号量与互斥量之间的区别：**  
>  （1）互斥量用于线程的互斥，信号量用于线程的同步。这是互斥量和信号量的根本区别，也就是互斥和同步之间的区别。  
>  **互斥：**是指某一资源同时只允许一个访问者对其进行访问，具有唯一性和排它性。但互斥无法限制访问者对资源的访问顺序，即访问是无序的。  
>  **同步：**是指在互斥的基础上（大多数情况），通过其它机制实现访问者对资源的有序访问。  
>  在大多数情况下，同步已经实现了互斥，特别是所有写入资源的情况必定是互斥的。少数情况是指可以允许多个访问者同时访问资源  
>  （2）互斥量值只能为0/1，信号量值可以为非负整数。  
>  也就是说，一个互斥量只能用于一个资源的互斥访问，它不能实现多个资源的多线程互斥问题。信号量可以实现多个同类资源的多线程互斥和同步。当信号量为单值信号量是，也可以完成一个资源的互斥访问。  
>  （3）互斥量的加锁和解锁必须由同一线程分别对应使用，信号量可以由一个线程释放，另一个线程得到。

**7. 套接字\(socket\)**  
 套接字是一种通信机制，凭借这种机制，客户/服务器（即要进行通信的进程）系统的开发工作既可以在本地单机上进行，也可以跨网络进行。也就是说它可以让不在同一台计算机但通过网络连接计算机上的进程进行通信。![](//upload-images.jianshu.io/upload_images/1281379-2db1deb0115ec4f2.png?imageMogr2/auto-orient/strip|imageView2/2/w/319/format/webp)Socket是应用层和传输层之间的桥梁  


套接字是支持TCP/IP的网络通信的基本操作单元，可以看做是不同主机之间的进程进行双向通信的端点，简单的说就是通信的两方的一种约定，用套接字中的相关函数来完成通信过程。

**套接字特性**  
 套接字的特性由3个属性确定，它们分别是：域、端口号、协议类型。  
 **（1）套接字的域**  
 它指定套接字通信中使用的网络介质，最常见的套接字域有两种：  
 **一是AF\_INET，它指的是Internet网络。**当客户使用套接字进行跨网络的连接时，它就需要用到服务器计算机的IP地址和端口来指定一台联网机器上的某个特定服务，所以在使用socket作为通信的终点，服务器应用程序必须在开始通信之前绑定一个端口，服务器在指定的端口等待客户的连接。  
 **另一个域AF\_UNIX，表示UNIX文件系统，**它就是文件输入/输出，而它的地址就是文件名。  
 **（2）套接字的端口号**  
 每一个基于TCP/IP网络通讯的程序\(进程\)都被赋予了唯一的端口和端口号，端口是一个信息缓冲区，用于保留Socket中的输入/输出信息，端口号是一个16位无符号整数，范围是0-65535，以区别主机上的每一个程序（端口号就像房屋中的房间号），低于256的端口号保留给标准应用程序，比如pop3的端口号就是110，每一个套接字都组合进了IP地址、端口，这样形成的整体就可以区别每一个套接字。  
 **（3）套接字协议类型**  
 因特网提供三种通信机制，  
 **一是流套接字，**流套接字在域中通过TCP/IP连接实现，同时也是AF\_UNIX中常用的套接字类型。流套接字提供的是一个有序、可靠、双向字节流的连接，因此发送的数据可以确保不会丢失、重复或乱序到达，而且它还有一定的出错后重新发送的机制。  
 **二个是数据报套接字，**它不需要建立连接和维持一个连接，它们在域中通常是通过UDP/IP协议实现的。它对可以发送的数据的长度有限制，数据报作为一个单独的网络消息被传输,它可能会丢失、复制或错乱到达，UDP不是一个可靠的协议，但是它的速度比较高，因为它并一需要总是要建立和维持一个连接。  
 **三是原始套接字，**原始套接字允许对较低层次的协议直接访问，比如IP、 ICMP协议，它常用于检验新的协议实现，或者访问现有服务中配置的新设备，因为RAW SOCKET可以自如地控制Windows下的多种协议，能够对网络底层的传输机制进行控制，所以可以应用原始套接字来操纵网络层和传输层应用。比如，我们可以通过RAW SOCKET来接收发向本机的ICMP、IGMP协议包，或者接收TCP/IP栈不能够处理的IP包，也可以用来发送一些自定包头或自定协议的IP包。网络监听技术很大程度上依赖于SOCKET\_RAW。

> **原始套接字与标准套接字的区别在于：**  
>  原始套接字可以读写内核没有处理的IP数据包，而流套接字只能读取TCP协议的数据，数据报套接字只能读取UDP协议的数据。因此，如果要访问其他协议发送数据必须使用原始套接字。

**套接字通信的建立**![](//upload-images.jianshu.io/upload_images/1281379-2575b81bbab6b67b.png?imageMogr2/auto-orient/strip|imageView2/2/w/437/format/webp)Socket通信基本流程  


\*\* 服务器端\*\*  
 （1）首先服务器应用程序用系统调用socket来创建一个套接字，它是系统分配给该服务器进程的类似文件描述符的资源，它不能与其他的进程共享。  
 （2）然后，服务器进程会给套接字起个名字，我们使用系统调用bind来给套接字命名。然后服务器进程就开始等待客户连接到这个套接字。  
 （3）接下来，系统调用listen来创建一个队列并将其用于存放来自客户的进入连接。  
 （4）最后，服务器通过系统调用accept来接受客户的连接。它会创建一个与原有的命名套接不同的新套接字，这个套接字只用于与这个特定客户端进行通信，而命名套接字（即原先的套接字）则被保留下来继续处理来自其他客户的连接（建立客户端和服务端的用于通信的流，进行通信）。

**客户端**  
 （1）客户应用程序首先调用socket来创建一个未命名的套接字，然后将服务器的命名套接字作为一个地址来调用connect与服务器建立连接。  
 （2）一旦连接建立，我们就可以像使用底层的文件描述符那样用套接字来实现双向数据的通信（通过流进行数据传输）。  


## CPU密集型（CPU-bound）

CPU密集型也叫计算密集型，指的是系统的硬盘、内存性能相对CPU要好很多，此时，系统运作大部分的状况是CPU Loading 100%，CPU要读/写I/O\(硬盘/内存\)，I/O在很短的时间就可以完成，而CPU还有许多运算要处理，CPU Loading很高。

在多重程序系统中，大部份时间用来做计算、逻辑判断等CPU动作的程序称之CPU bound。例如一个计算圆周率至小数点一千位以下的程序，在执行的过程当中绝大部份时间用在三角函数和开根号的计算，便是属于CPU bound的程序。

CPU bound的程序一般而言CPU占用率相当高。这可能是因为任务本身不太需要访问I/O设备，也可能是因为程序是多线程实现因此屏蔽掉了等待I/O的时间。

#### IO密集型（I/O bound）

IO密集型指的是系统的CPU性能相对硬盘、内存要好很多，此时，系统运作，大部分的状况是CPU在等I/O \(硬盘/内存\) 的读/写操作，此时CPU Loading并不高。

I/O bound的程序一般在达到性能极限时，CPU占用率仍然较低。这可能是因为任务本身需要大量I/O操作，而pipeline做得不是很好，没有充分利用处理器能力。  
[https://jiang.ma/2019/01/28/golang-advantags-when-facing-networking-io-bound-situation.html](https://jiang.ma/2019/01/28/golang-advantags-when-facing-networking-io-bound-situation.html)

go中，io密集型的应用，比如有很多文件io，磁盘io，网络io，调大GOMAXPROCS，会不会对性能有帮助？为什么？  
 福哥答案2021-03-05：

这是面试中被问到的。实力有限，真正的答案还不知道。

答案1：  
 调节这个参数影响的是P的个数，也就影响了M（线程）干活的个数。相当于你可以有更多的执行线程。  
 先以网络io来说，网络io 在golang 里面是异步的，用epoll池做的io复用。每个网络调用其实都是异步的，发数据给到内存，调度权就可以让给其他goroutine了，所以，其实一个线程能处理过来的话，性能是不会差的，这个时候你加多P其实提升不大。只有你单线程处理不过来这些网络io的时候（每个都处理很慢），加多P才有明显提升  
 如果是磁盘io的话，这个有点特殊，磁盘io不是异步的，没有aio这种方式。所以你的磁盘io调用下去就卡住M了，这个时候等sysmon发现系统调用超时才会抢占M，这一来回就耗费时间了，所以，这种情况下你干活的M多一点确实能带来一些性能的提升，相当于并行干活的M多一些。  
 无论哪种情况，P的个数都不建议超过本机cpu的个数。因为多个cpu才是真正的并行执行，上层都是通过调度切换模拟出来的。

答案2：  
 GOMAXPROCS 用默认的，就是CPU的硬件线程数目，  
 对于大部分IO密集的应用是不合适的。  
 至少应该配置到硬件线程数目的5倍以上, 最大256。  
 GO的调度器是迟钝的，它很可能什么时都没做，直到M阻塞了想当长时间以后，才会发现有一个P/M被syscall阻塞了。然后，才会用空闲的M来强这个P。  
 补充说明：调度器迟钝不是M迟钝，M也就是操作系统线程，是非常的敏感的，只要阻塞就会被操作系统调度（除了极少数自旋的情况）。但是GO的调度器会等待一个时间间隔才会行动，这也是为了减少调度器干预的次数。也就是说，如果一个M调用了什么API导致了操作系统线程阻塞了，操作系统立刻会把这个线程M调度走，挂起等阻塞解除。这时候，Go调度器不会马上把这个M持有的P抢走。这就会导致一定的P被浪费了。  
 这就是为何，GOMAXPROCS 太小，也就是P的数量太少，会导致IO密集\(或者syscall较多\)的go程序运行缓慢的原因。  
 那么，GOMAXPROCS 很大，超过硬件线程的8倍，会不会有开销呢？  
 答案是，开销是有的，但是远小于Go运行时迟钝的调度M来抢夺P而导致CPU利用不足的开销。  
  
[https://jiang.ma/2019/01/28/golang-advantags-when-facing-networking-io-bound-situation.html](https://jiang.ma/2019/01/28/golang-advantags-when-facing-networking-io-bound-situation.html)

## gin路由

定义了两个路由 `/user/get`，`/user/delete`，则会构造出拥有三个节点的路由树，根节点是 `user`，两个子节点分别是 `get` `delete`。

上述是一种实现路由树的方式，且比较直观，容易理解。对 url 进行切分、比较，时间复杂度是 `O(2n)`。

Gin的路由实现使用了类似前缀树的数据结构，只需遍历一遍字符串即可，时间复杂度为`O(n)`。

当然，对于一次 http 请求来说，这点路由寻址优化可以忽略不计。

#### Engine <a id="item-1"></a>

`Gin` 的 `Engine` 结构体内嵌了 `RouterGroup` 结构体，定义了 `GET`，`POST` 等路由注册方法。

`Engine` 中的 `trees` 字段定义了路由逻辑。`trees` 是 `methodTrees` 类型（其实就是 `[]methodTree`），`trees` 是一个数组，不同请求方法的路由在不同的树（`methodTree`）中。

最后，`methodTree` 中的 `root` 字段（`*node`类型）是路由树的根节点。树的构造与寻址都是在 `*node`的方法中完成的。

UML 结构图  
![engine&#x7ED3;&#x6784;&#x56FE;](https://segmentfault.com/img/bVbh2X1?w=1374&h=774)

`trees` 是个数组，数组里会有不同请求方法的路由树。

![tree&#x7ED3;&#x6784;](https://segmentfault.com/img/bVbh2Yc?w=1356&h=320)

#### node <a id="item-2"></a>

node 结构体定义如下

```text
type node struct {
    path      string           // 当前节点相对路径（与祖先节点的 path 拼接可得到完整路径）
    indices   string           // 所以孩子节点的path[0]组成的字符串
    children  []*node          // 孩子节点
    handlers  HandlersChain    // 当前节点的处理函数（包括中间件）
    priority  uint32           // 当前节点及子孙节点的实际路由数量
    nType     nodeType         // 节点类型
    maxParams uint8            // 子孙节点的最大参数数量
    wildChild bool             // 孩子节点是否有通配符（wildcard）
}
```

**path 和 indices**

关于 `path` 和 `indices`，其实是使用了前缀树的逻辑。

举个栗子：  
如果我们有两个路由，分别是 `/index`，`/inter`，则根节点为 `{path: "/in", indices: "dt"...}`，两个子节点为`{path: "dex", indices: ""}，{path: "ter", indices: ""}`  














