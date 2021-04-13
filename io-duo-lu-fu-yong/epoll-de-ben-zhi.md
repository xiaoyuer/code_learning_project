# Epoll的本质

从事服务端开发，少不了要接触网络编程。epoll作为linux下高性能网络服务器的必备技术至关重要，nginx、redis、skynet和大部分游戏服务器都使用到这一多路复用技术。

**文**/罗培羽

因为epoll的重要性，不少游戏公司（如某某九九）在招聘服务端同学时，可能会问及epoll相关的问题。比如epoll和select的区别是什么？epoll高效率的原因是什么？如果只靠背诵，显然不能算上深刻的理解。

网上虽然也有不少讲解epoll的文章，但要不是过于浅显，就是陷入源码解析，很少能有通俗易懂的。于是决定编写此文，让缺乏专业背景知识的读者也能够明白epoll的原理。文章核心思想是：

### **要让读者清晰明白EPOLL为什么性能好。**

本文会从网卡接收数据的流程讲起，串联起CPU中断、操作系统进程调度等知识；再一步步分析阻塞接收数据、select到epoll的进化过程；最后探究epoll的实现细节。目录：

> 一、从网卡接收数据说起  
> 二、如何知道接收了数据？  
> 三、进程阻塞为什么不占用cpu资源？  
> 四、内核接收网络数据全过程  
> 五、同时监视多个socket的简单方法  
> 六、epoll的设计思路  
> 七、epoll的原理和流程  
> 八、epoll的实现细节  
> 九、结论

### **一、从网卡接收数据说起**

下图是一个典型的计算机结构图，计算机由CPU、存储器（内存）、网络接口等部件组成。了解epoll本质的**第一步**，要从**硬件**的角度看计算机怎样接收网络数据。![](https://pic2.zhimg.com/80/v2-e549406135abf440331de9dd8c3925e9_1440w.jpg)计算机结构图（图片来源：linux内核完全注释之微型计算机组成结构）

下图展示了网卡接收数据的过程。在①阶段，网卡收到网线传来的数据；经过②阶段的硬件电路的传输；最终将数据写入到内存中的某个地址上（③阶段）。这个过程涉及到DMA传输、IO通路选择等硬件有关的知识，但我们只需知道：**网卡会把接收到的数据写入内存。**![](https://pic2.zhimg.com/80/v2-6827b63c9fb42823dcd1913ea5433b15_1440w.jpg)网卡接收数据的过程

通过硬件传输，网卡接收的数据存放到内存中。操作系统就可以去读取它们。

### **二、如何知道接收了数据？**

了解epoll本质的**第二步**，要从**CPU**的角度来看数据接收。要理解这个问题，要先了解一个概念——中断。

计算机执行程序时，会有优先级的需求。比如，当计算机收到断电信号时（电容可以保存少许电量，供CPU运行很短的一小段时间），它应立即去保存数据，保存数据的程序具有较高的优先级。

一般而言，由硬件产生的信号需要cpu立马做出回应（不然数据可能就丢失），所以它的优先级很高。cpu理应中断掉正在执行的程序，去做出响应；当cpu完成对硬件的响应后，再重新执行用户程序。中断的过程如下图，和函数调用差不多。只不过函数调用是事先定好位置，而中断的位置由“信号”决定。![](https://pic4.zhimg.com/80/v2-89a9490f1d5c316167ff4761184239f7_1440w.jpg)中断程序调用

以键盘为例，当用户按下键盘某个按键时，键盘会给cpu的中断引脚发出一个高电平。cpu能够捕获这个信号，然后执行键盘中断程序。下图展示了各种硬件通过中断与cpu交互。![](https://pic3.zhimg.com/80/v2-c756381c0f63f9104f9102d280759d22_1440w.jpg)cpu中断（图片来源：net.pku.edu.cn）

现在可以回答本节提出的问题了：当网卡把数据写入到内存后，**网卡向cpu发出一个中断信号，操作系统便能得知有新数据到来**，再通过网卡**中断程序**去处理数据。

### **三、进程阻塞为什么不占用cpu资源？**

了解epoll本质的**第三步**，要从**操作系统进程调度**的角度来看数据接收。阻塞是进程调度的关键一环，指的是进程在等待某事件（如接收到网络数据）发生之前的等待状态，recv、select和epoll都是阻塞方法。**了解“进程阻塞为什么不占用cpu资源？”，也就能够了解这一步**。

为简单起见，我们从普通的recv接收开始分析，先看看下面代码：

```text
//创建socket
int s = socket(AF_INET, SOCK_STREAM, 0);   
//绑定
bind(s, ...)
//监听
listen(s, ...)
//接受客户端连接
int c = accept(s, ...)
//接收客户端数据
recv(c, ...);
//将数据打印出来
printf(...)
```

这是一段最基础的网络编程代码，先新建socket对象，依次调用bind、listen、accept，最后调用recv接收数据。recv是个阻塞方法，当程序运行到recv时，它会一直等待，直到接收到数据才往下执行。

> 插入：如果您还不太熟悉网络编程，欢迎阅读我编写的《Unity3D网络游戏实战\(第2版\)》，会有详细的介绍。

那么阻塞的原理是什么？

**工作队列**

操作系统为了支持多任务，实现了进程调度的功能，会把进程分为“运行”和“等待”等几种状态。运行状态是进程获得cpu使用权，正在执行代码的状态；等待状态是阻塞状态，比如上述程序运行到recv时，程序会从运行状态变为等待状态，接收到数据后又变回运行状态。操作系统会分时执行各个运行状态的进程，由于速度很快，看上去就像是同时执行多个任务。

下图中的计算机中运行着A、B、C三个进程，其中进程A执行着上述基础网络程序，一开始，这3个进程都被操作系统的工作队列所引用，处于运行状态，会分时执行。![](https://pic3.zhimg.com/80/v2-2f3b71710f1805669a780a2d634f0626_1440w.jpg)工作队列中有A、B和C三个进程

**等待队列**

当进程A执行到创建socket的语句时，操作系统会创建一个由文件系统管理的socket对象（如下图）。这个socket对象包含了发送缓冲区、接收缓冲区、等待队列等成员。等待队列是个非常重要的结构，它指向所有需要等待该socket事件的进程。![](https://pic3.zhimg.com/80/v2-7ce207c92c9dd7085fb7b823e2aa5872_1440w.jpg)创建socket

当程序执行到recv时，操作系统会将进程A从工作队列移动到该socket的等待队列中（如下图）。由于工作队列只剩下了进程B和C，依据进程调度，cpu会轮流执行这两个进程的程序，不会执行进程A的程序。**所以进程A被阻塞，不会往下执行代码，也不会占用cpu资源**。![](https://pic1.zhimg.com/80/v2-1c7a96c8da16f123388e46f88772e6d8_1440w.jpg)socket的等待队列

ps：操作系统添加等待队列只是添加了对这个“等待中”进程的引用，以便在接收到数据时获取进程对象、将其唤醒，而非直接将进程管理纳入自己之下。上图为了方便说明，直接将进程挂到等待队列之下。

**唤醒进程**

当socket接收到数据后，操作系统将该socket等待队列上的进程重新放回到工作队列，该进程变成运行状态，继续执行代码。也由于socket的接收缓冲区已经有了数据，recv可以返回接收到的数据。

**以下内容待续**

四、内核接收网络数据全过程

五、同时监视多个socket的简单方法

六、epoll的设计思路

七、epoll的原理和流程

八、epoll的实现细节

九、结论

既然说到网络编程，笔者的**《Unity3D网络游戏实战（第2版）》**是一本专门介绍如何开发**多人网络游戏**的书籍，用实例介绍开发游戏的全过程，非常实用。书中对网络编程有详细的讲解，全书用一个大例子贯穿，真正的“实战”教程。

致谢：本文力图详细说明epoll的原理，特别感谢 [@陆俊壕](https://www.zhihu.com/people/e622f8ea68620104614bcc9a4ce3855d) [@AllenKong12](https://www.zhihu.com/people/8887d646fe997ca00f7ff99b724dd230) 雄爷、堂叔 等同事审阅了文章并给予修改意见。![](https://pic1.zhimg.com/80/v2-6ab25b1164ab427bdbe7eccaff7e9570_1440w.jpg)  
**上篇回顾**

一、从网卡接收数据说起

二、如何知道接收了数据？

三、进程阻塞为什么不占用cpu资源？

**系列文章**

[罗培羽：如果这篇文章说不清epoll的本质，那就过来掐死我吧！ （1）](https://zhuanlan.zhihu.com/p/63179839)

[罗培羽：如果这篇文章说不清epoll的本质，那就过来掐死我吧！ （2）](https://zhuanlan.zhihu.com/p/64138532)

[罗培羽：如果这篇文章说不清epoll的本质，那就过来掐死我吧！ （3）](https://zhuanlan.zhihu.com/p/64746509)

### **四、内核接收网络数据全过程**

**这一步，贯穿网卡、中断、进程调度的知识，叙述阻塞recv下，内核接收数据全过程。**

如下图所示，进程在recv阻塞期间，计算机收到了对端传送的数据（步骤①）。数据经由网卡传送到内存（步骤②），然后网卡通过中断信号通知cpu有数据到达，cpu执行中断程序（步骤③）。此处的中断程序主要有两项功能，先将网络数据写入到对应socket的接收缓冲区里面（步骤④），再唤醒进程A（步骤⑤），重新将进程A放入工作队列中。![](https://pic4.zhimg.com/80/v2-696b131cae434f2a0b5ab4d6353864af_1440w.jpg)内核接收数据全过程

唤醒进程的过程如下图所示。![](https://pic3.zhimg.com/80/v2-3e1d0a82cdc86f03343994f48d938922_1440w.jpg)唤醒进程

**以上是内核接收数据全过程**

这里留有两个思考题，大家先想一想。

其一，操作系统如何知道网络数据对应于哪个socket？

其二，如何同时监视多个socket的数据？

（——我是分割线，想好了才能往下看哦~）

公布答案的时刻到了。

第一个问题：因为一个socket对应着一个端口号，而网络数据包中包含了ip和端口的信息，内核可以通过端口号找到对应的socket。当然，为了提高处理速度，操作系统会维护端口号到socket的索引结构，以快速读取。

第二个问题是**多路复用的重中之重，**是本文后半部分的重点！

### **五、同时监视多个socket的简单方法**

服务端需要管理多个客户端连接，而recv只能监视单个socket，这种矛盾下，人们开始寻找监视多个socket的方法。epoll的要义是**高效**的监视多个socket。从历史发展角度看，必然先出现一种不太高效的方法，人们再加以改进。只有先理解了不太高效的方法，才能够理解epoll的本质。

假如能够预先传入一个socket列表，**如果列表中的socket都没有数据，挂起进程，直到有一个socket收到数据，唤醒进程**。这种方法很直接，也是select的设计思想。

为方便理解，我们先复习select的用法。在如下的代码中，先准备一个数组（下面代码中的fds），让fds存放着所有需要监视的socket。然后调用select，如果fds中的所有socket都没有数据，select会阻塞，直到有一个socket接收到数据，select返回，唤醒进程。用户可以遍历fds，通过FD\_ISSET判断具体哪个socket收到数据，然后做出处理。

```text
int s = socket(AF_INET, SOCK_STREAM, 0);  
bind(s, ...)
listen(s, ...)

int fds[] =  存放需要监听的socket

while(1){
    int n = select(..., fds, ...)
    for(int i=0; i < fds.count; i++){
        if(FD_ISSET(fds[i], ...)){
            //fds[i]的数据处理
        }
    }
}
```

**select的流程**

select的实现思路很直接。假如程序同时监视如下图的sock1、sock2和sock3三个socket，那么在调用select之后，操作系统把进程A分别加入这三个socket的等待队列中。![](https://pic4.zhimg.com/80/v2-0cccb4976f8f2c2f8107f2b3a5bc46b3_1440w.jpg)操作系统把进程A分别加入这三个socket的等待队列中

当任何一个socket收到数据后，中断程序将唤起进程。下图展示了sock2接收到了数据的处理流程。

> ps：recv和select的中断回调可以设置成不同的内容。

![](https://pic1.zhimg.com/80/v2-85dba5430f3c439e4647ea4d97ba54fc_1440w.jpg)sock2接收到了数据，中断程序唤起进程A

所谓唤起进程，就是将进程从所有的等待队列中移除，加入到工作队列里面。如下图所示。![](https://pic4.zhimg.com/80/v2-a86b203b8d955466fff34211d965d9eb_1440w.jpg)将进程A从所有等待队列中移除，再加入到工作队列里面

经由这些步骤，当进程A被唤醒后，它知道至少有一个socket接收了数据。程序只需遍历一遍socket列表，就可以得到就绪的socket。

这种简单方式**行之有效**，在几乎所有操作系统都有对应的实现。

**但是简单的方法往往有缺点，主要是：**

其一，每次调用select都需要将进程加入到所有监视socket的等待队列，每次唤醒都需要从每个队列中移除。这里涉及了两次遍历，而且每次都要将整个fds列表传递给内核，有一定的开销。正是因为遍历操作开销大，出于效率的考量，才会规定select的最大监视数量，默认只能监视1024个socket。

其二，进程被唤醒后，程序并不知道哪些socket收到数据，还需要遍历一次。

那么，有没有减少遍历的方法？有没有保存就绪socket的方法？这两个问题便是epoll技术要解决的。

> 补充说明： 本节只解释了select的一种情形。当程序调用select时，内核会先遍历一遍socket，如果有一个以上的socket接收缓冲区有数据，那么select直接返回，不会阻塞。这也是为什么select的返回值有可能大于1的原因之一。如果没有socket有数据，进程才会阻塞。

### **六、epoll的设计思路**

epoll是在select出现N多年后才被发明的，是select和poll的增强版本。epoll通过以下一些措施来改进效率。

**措施一：功能分离**

select低效的原因之一是将“维护等待队列”和“阻塞进程”两个步骤合二为一。如下图所示，每次调用select都需要这两步操作，然而大多数应用场景中，需要监视的socket相对固定，并不需要每次都修改。epoll将这两个操作分开，先用epoll\_ctl维护等待队列，再调用epoll\_wait阻塞进程。显而易见的，效率就能得到提升。![](https://pic2.zhimg.com/80/v2-5ce040484bbe61df5b484730c4cf56cd_1440w.jpg)相比select，epoll拆分了功能

为方便理解后续的内容，我们先复习下epoll的用法。如下的代码中，先用epoll\_create创建一个epoll对象epfd，再通过epoll\_ctl将需要监视的socket添加到epfd中，最后调用epoll\_wait等待数据。

```text
int s = socket(AF_INET, SOCK_STREAM, 0);   
bind(s, ...)
listen(s, ...)

int epfd = epoll_create(...);
epoll_ctl(epfd, ...); //将所有需要监听的socket添加到epfd中

while(1){
    int n = epoll_wait(...)
    for(接收到数据的socket){
        //处理
    }
}
```

功能分离，使得epoll有了优化的可能。

**措施二：就绪列表**

select低效的另一个原因在于程序不知道哪些socket收到数据，只能一个个遍历。如果内核维护一个“就绪列表”，引用收到数据的socket，就能避免遍历。如下图所示，计算机共有三个socket，收到数据的sock2和sock3被rdlist（就绪列表）所引用。当进程被唤醒后，只要获取rdlist的内容，就能够知道哪些socket收到数据。![](https://pic4.zhimg.com/80/v2-5c552b74772d8dbc7287864999e32c4f_1440w.jpg)就绪列表示意图

**以下内容待续**

七、epoll的原理和流程

八、epoll的实现细节

九、结论

除了网络编程，「同步」也是网络游戏开发的核心课题。多人游戏中，玩家在世界中的位置旋转以及各种属性都会对游戏表现产生影响，需要同步给其他玩家。然而由于网络通信存在延迟和抖动，很难做到完美的同步效果。如果没能处理好，游戏将不同步和卡顿，这是玩家所不能容忍的。笔者即将开展一场知乎live**《网络游戏同步算法》**，分析同步技术，欢迎收听。[网络游戏同步算法​www.zhihu.co](https://www.zhihu.com/lives/1104162893850898432)

**上篇回顾**

四、内核接收网络数据全过程

五、同时监视多个socket的简单方法

六、epoll的设计思路

**系列文章**

[罗培羽：如果这篇文章说不清epoll的本质，那就过来掐死我吧！ （1）](https://zhuanlan.zhihu.com/p/63179839)

[罗培羽：如果这篇文章说不清epoll的本质，那就过来掐死我吧！ （2）](https://zhuanlan.zhihu.com/p/64138532)

[罗培羽：如果这篇文章说不清epoll的本质，那就过来掐死我吧！ （3）](https://zhuanlan.zhihu.com/p/64746509)

### **七、epoll的原理和流程**

本节会以示例和图表来讲解epoll的原理和流程。

**创建epoll对象**

如下图所示，当某个进程调用epoll\_create方法时，内核会创建一个eventpoll对象（也就是程序中epfd所代表的对象）。eventpoll对象也是文件系统中的一员，和socket一样，它也会有等待队列。![](https://pic4.zhimg.com/80/v2-e3467895734a9d97f0af3c7bf875aaeb_1440w.jpg)内核创建eventpoll对象

创建一个代表该epoll的eventpoll对象是必须的，因为内核要维护“就绪列表”等数据，“就绪列表”可以作为eventpoll的成员。

**维护监视列表**

创建epoll对象后，可以用epoll\_ctl添加或删除所要监听的socket。以添加socket为例，如下图，如果通过epoll\_ctl添加sock1、sock2和sock3的监视，内核会将eventpoll添加到这三个socket的等待队列中。![](https://pic2.zhimg.com/80/v2-b49bb08a6a1b7159073b71c4d6591185_1440w.jpg)添加所要监听的socket

当socket收到数据后，中断程序会操作eventpoll对象，而不是直接操作进程。

**接收数据**

当socket收到数据后，中断程序会给eventpoll的“就绪列表”添加socket引用。如下图展示的是sock2和sock3收到数据后，中断程序让rdlist引用这两个socket。![](https://pic1.zhimg.com/80/v2-18b89b221d5db3b5456ab6a0f6dc5784_1440w.jpg)给就绪列表添加引用

eventpoll对象相当于是socket和进程之间的中介，socket的数据接收并不直接影响进程，而是通过改变eventpoll的就绪列表来改变进程状态。

当程序执行到epoll\_wait时，如果rdlist已经引用了socket，那么epoll\_wait直接返回，如果rdlist为空，阻塞进程。

**阻塞和唤醒进程**

假设计算机中正在运行进程A和进程B，在某时刻进程A运行到了epoll\_wait语句。如下图所示，内核会将进程A放入eventpoll的等待队列中，阻塞进程。![](https://pic1.zhimg.com/80/v2-90632d0dc3ded7f91379b848ab53974c_1440w.jpg)epoll\_wait阻塞进程

当socket接收到数据，中断程序一方面修改rdlist，另一方面唤醒eventpoll等待队列中的进程，进程A再次进入运行状态（如下图）。也因为rdlist的存在，进程A可以知道哪些socket发生了变化。![](https://pic4.zhimg.com/80/v2-40bd5825e27cf49b7fd9a59dfcbe4d6f_1440w.jpg)epoll唤醒进程

### **八、epoll的实现细节**

至此，相信读者对epoll的本质已经有一定的了解。但我们还留有一个问题，**eventpoll的数据结构**是什么样子？

再留两个问题，**就绪队列**应该应使用什么数据结构？eventpoll应使用什么数据结构来管理通过epoll\_ctl添加或删除的socket？

（——我是分割线，想好了才能往下看哦~）

如下图所示，eventpoll包含了lock、mtx、wq（等待队列）、rdlist等成员。rdlist和rbr是我们所关心的。![](https://pic4.zhimg.com/80/v2-e63254878f67751dcc07a25b93f974bb_1440w.jpg)epoll原理示意图，图片来源：《深入理解Nginx：模块开发与架构解析\(第二版\)》，陶辉

**就绪列表的数据结构**

就绪列表引用着就绪的socket，所以它应能够快速的插入数据。

程序可能随时调用epoll\_ctl添加监视socket，也可能随时删除。当删除时，若该socket已经存放在就绪列表中，它也应该被移除。

所以就绪列表应是一种能够快速插入和删除的数据结构。双向链表就是这样一种数据结构，epoll使用双向链表来实现就绪队列（对应上图的rdllist）。

**索引结构**

既然epoll将“维护监视队列”和“进程阻塞”分离，也意味着需要有个数据结构来保存监视的socket。至少要方便的添加和移除，还要便于搜索，以避免重复添加。红黑树是一种自平衡二叉查找树，搜索、插入和删除时间复杂度都是O\(log\(N\)\)，效率较好。epoll使用了红黑树作为索引结构（对应上图的rbr）。

> ps：因为操作系统要兼顾多种功能，以及由更多需要保存的数据，rdlist并非直接引用socket，而是通过epitem间接引用，红黑树的节点也是epitem对象。同样，文件系统也并非直接引用着socket。为方便理解，本文中省略了一些间接结构。

### **九、结论**

epoll在select和poll（poll和select基本一样，有少量改进）的基础引入了eventpoll作为中间层，使用了先进的数据结构，是一种高效的多路复用技术。

再留一点**作业**！

下表是个很常见的表，描述了select、poll和epoll的区别。读完本文，读者能否解释select和epoll的时间复杂度为什么是O\(n\)和O\(1\)？![](https://pic2.zhimg.com/80/v2-14e0536d872474b0851b62572b732e39_1440w.jpg)select、poll和epoll的区别。图片来源《Linux高性能服务器编程》

笔者的**《Unity3D网络游戏实战（第2版）》**是一本专门介绍如何开发**多人网络游戏**的书籍，用实例介绍开发游戏的全过程，手把手教你如何制作一款多人开房间的坦克对战游戏。

「同步」也是网络游戏开发的核心课题。如何恰当的使用不同的同步算法？帧同步的应用场景和优越有哪些？笔者即将开展一场知乎live**《网络游戏同步算法》**，欢迎收听。[网络游戏同步算法​www.zhihu.com](https://www.zhihu.com/lives/1104162893850898432)

## Epoll使用详解

在linux的网络编程中，很长的时间都在使用select来做事件触发。在linux新的内核中，有了一种替换它的机制，就是epoll。  
相比于select，epoll最大的好处在于它不会随着监听fd数目的增长而降低效率。因为在内核中的select实现中，它是采用轮询来处理的，轮询的fd数目越多，自然耗时越多。并且，在linux/posix\_types.h头文件有这样的声明：  
\#define \_\_FD\_SETSIZE    1024  
表示select最多同时监听1024个fd，当然，可以通过修改头文件再重编译内核来扩大这个数目，但这似乎并不治本。  
  
epoll的接口非常简单，一共就三个函数：  
1. int epoll\_create\(int size\);  
创建一个epoll的句柄，size用来告诉内核这个监听的数目一共有多大。这个参数不同于select\(\)中的第一个参数，给出最大监听的fd+1的值。需要注意的是，当创建好epoll句柄后，它就是会占用一个fd值，在linux下如果查看/proc/进程id/fd/，是能够看到这个fd的，所以在使用完epoll后，必须调用close\(\)关闭，否则可能导致fd被耗尽。  
  
  
2. int epoll\_ctl\(int epfd, int op, int fd, struct epoll\_event \*event\);  
epoll的事件注册函数，它不同与select\(\)是在监听事件时告诉内核要监听什么类型的事件，而是在这里先注册要监听的事件类型。第一个参数是epoll\_create\(\)的返回值，第二个参数表示动作，用三个宏来表示：  
EPOLL\_CTL\_ADD：注册新的fd到epfd中；  
EPOLL\_CTL\_MOD：修改已经注册的fd的监听事件；  
EPOLL\_CTL\_DEL：从epfd中删除一个fd；  
第三个参数是需要监听的fd，第四个参数是告诉内核需要监听什么事，struct epoll\_event结构如下：  
  
[![&#x590D;&#x5236;&#x4EE3;&#x7801;](https://common.cnblogs.com/images/copycode.gif)](javascript:void%280%29;)

```text
 1 typedef union epoll_data {
 2     void *ptr;
 3     int fd;
 4     __uint32_t u32;
 5     __uint64_t u64;
 6 } epoll_data_t;
 7 
 8 struct epoll_event {
 9     __uint32_t events; /* Epoll events */
10     epoll_data_t data; /* User data variable */
11 };
```

[![&#x590D;&#x5236;&#x4EE3;&#x7801;](https://common.cnblogs.com/images/copycode.gif)](javascript:void%280%29;)

  
events可以是以下几个宏的集合：  
EPOLLIN ：表示对应的文件描述符可以读（包括对端SOCKET正常关闭）；  
EPOLLOUT：表示对应的文件描述符可以写；  
EPOLLPRI：表示对应的文件描述符有紧急的数据可读（这里应该表示有带外数据到来）；  
EPOLLERR：表示对应的文件描述符发生错误；  
EPOLLHUP：表示对应的文件描述符被挂断；  
EPOLLET： 将EPOLL设为边缘触发\(Edge Triggered\)模式，这是相对于水平触发\(Level Triggered\)来说的。  
EPOLLONESHOT：只监听一次事件，当监听完这次事件之后，如果还需要继续监听这个socket的话，需要再次把这个socket加入到EPOLL队列里  
  
  
3. int epoll\_wait\(int epfd, struct epoll\_event \* events, int maxevents, int timeout\);  
等待事件的产生，类似于select\(\)调用。参数events用来从内核得到事件的集合，maxevents告之内核这个events有多大，这个 maxevents的值不能大于创建epoll\_create\(\)时的size，参数timeout是超时时间（毫秒，0会立即返回，-1将不确定，也有说法说是永久阻塞）。该函数返回需要处理的事件数目，如返回0表示已超时。  
  
  
4、关于ET、LT两种工作模式：  
可以得出这样的结论:  
ET模式仅当状态发生变化的时候才获得通知,这里所谓的状态的变化并不包括缓冲区中还有未处理的数据,也就是说,如果要采用ET模式,需要一直read/write直到出错为止,很多人反映为什么采用ET模式只接收了一部分数据就再也得不到通知了,大多因为这样;而LT模式是只要有数据没有处理就会一直通知下去的.  
  
  
那么究竟如何来使用epoll呢？其实非常简单。  
通过在包含一个头文件\#include &lt;sys/epoll.h&gt; 以及几个简单的API将可以大大的提高你的网络服务器的支持人数。  
  
首先通过create\_epoll\(int maxfds\)来创建一个epoll的句柄，其中maxfds为你epoll所支持的最大句柄数。这个函数会返回一个新的epoll句柄，之后的所有操作将通过这个句柄来进行操作。在用完之后，记得用close\(\)来关闭这个创建出来的epoll句柄。  
  
之后在你的网络主循环里面，每一帧的调用epoll\_wait\(int epfd, epoll\_event events, int max events, int timeout\)来查询所有的网络接口，看哪一个可以读，哪一个可以写了。基本的语法为：  
nfds = epoll\_wait\(kdpfd, events, maxevents, -1\);  
其中kdpfd为用epoll\_create创建之后的句柄，events是一个epoll\_event\*的指针，当epoll\_wait这个函数操作成功之后，epoll\_events里面将储存所有的读写事件。max\_events是当前需要监听的所有socket句柄数。最后一个timeout是 epoll\_wait的超时，为0的时候表示马上返回，为-1的时候表示一直等下去，直到有事件范围，为任意正整数的时候表示等这么长的时间，如果一直没有事件，则范围。一般如果网络主循环是单独的线程的话，可以用-1来等，这样可以保证一些效率，如果是和主逻辑在同一个线程的话，则可以用0来保证主循环的效率。  
  
epoll\_wait范围之后应该是一个循环，遍利所有的事件。  
  
几乎所有的epoll程序都使用下面的框架：  
  
   [![&#x590D;&#x5236;&#x4EE3;&#x7801;](https://common.cnblogs.com/images/copycode.gif)](javascript:void%280%29;)

```text
 1  for( ; ; )
 2     {
 3         nfds = epoll_wait(epfd,events,20,500);
 4         for(i=0;i<nfds;++i)
 5         {
 6             if(events[i].data.fd==listenfd) //有新的连接
 7             {
 8                 connfd = accept(listenfd,(sockaddr *)&clientaddr, &clilen); //accept这个连接
 9                 ev.data.fd=connfd;
10                 ev.events=EPOLLIN|EPOLLET;
11                 epoll_ctl(epfd,EPOLL_CTL_ADD,connfd,&ev); //将新的fd添加到epoll的监听队列中
12             }
13             else if( events[i].events&EPOLLIN ) //接收到数据，读socket
14             {
15                 n = read(sockfd, line, MAXLINE)) < 0    //读
16                 ev.data.ptr = md;     //md为自定义类型，添加数据
17                 ev.events=EPOLLOUT|EPOLLET;
18                 epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev);//修改标识符，等待下一个循环时发送数据，异步处理的精髓
19             }
20             else if(events[i].events&EPOLLOUT) //有数据待发送，写socket
21             {
22                 struct myepoll_data* md = (myepoll_data*)events[i].data.ptr;    //取数据
23                 sockfd = md->fd;
24                 send( sockfd, md->ptr, strlen((char*)md->ptr), 0 );        //发送数据
25                 ev.data.fd=sockfd;
26                 ev.events=EPOLLIN|EPOLLET;
27                 epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev); //修改标识符，等待下一个循环时接收数据
28             }
29             else
30             {
31                 //其他的处理
32             }
33         }
34     }
```

[![&#x590D;&#x5236;&#x4EE3;&#x7801;](https://common.cnblogs.com/images/copycode.gif)](javascript:void%280%29;)

下面给出一个完整的服务器端例子：[![&#x590D;&#x5236;&#x4EE3;&#x7801;](https://common.cnblogs.com/images/copycode.gif)](javascript:void%280%29;)

```text
  1 #include <iostream>
  2 #include <sys/socket.h>
  3 #include <sys/epoll.h>
  4 #include <netinet/in.h>
  5 #include <arpa/inet.h>
  6 #include <fcntl.h>
  7 #include <unistd.h>
  8 #include <stdio.h>
  9 #include <errno.h>
 10 
 11 using namespace std;
 12 
 13 #define MAXLINE 5
 14 #define OPEN_MAX 100
 15 #define LISTENQ 20
 16 #define SERV_PORT 5000
 17 #define INFTIM 1000
 18 
 19 void setnonblocking(int sock)
 20 {
 21     int opts;
 22     opts=fcntl(sock,F_GETFL);
 23     if(opts<0)
 24     {
 25         perror("fcntl(sock,GETFL)");
 26         exit(1);
 27     }
 28     opts = opts|O_NONBLOCK;
 29     if(fcntl(sock,F_SETFL,opts)<0)
 30     {
 31         perror("fcntl(sock,SETFL,opts)");
 32         exit(1);
 33     }
 34 }
 35 
 36 int main(int argc, char* argv[])
 37 {
 38     int i, maxi, listenfd, connfd, sockfd,epfd,nfds, portnumber;
 39     ssize_t n;
 40     char line[MAXLINE];
 41     socklen_t clilen;
 42 
 43 
 44     if ( 2 == argc )
 45     {
 46         if( (portnumber = atoi(argv[1])) < 0 )
 47         {
 48             fprintf(stderr,"Usage:%s portnumber/a/n",argv[0]);
 49             return 1;
 50         }
 51     }
 52     else
 53     {
 54         fprintf(stderr,"Usage:%s portnumber/a/n",argv[0]);
 55         return 1;
 56     }
 57 
 58 
 59 
 60     //声明epoll_event结构体的变量,ev用于注册事件,数组用于回传要处理的事件
 61 
 62     struct epoll_event ev,events[20];
 63     //生成用于处理accept的epoll专用的文件描述符
 64 
 65     epfd=epoll_create(256);
 66     struct sockaddr_in clientaddr;
 67     struct sockaddr_in serveraddr;
 68     listenfd = socket(AF_INET, SOCK_STREAM, 0);
 69     //把socket设置为非阻塞方式
 70 
 71     //setnonblocking(listenfd);
 72 
 73     //设置与要处理的事件相关的文件描述符
 74 
 75     ev.data.fd=listenfd;
 76     //设置要处理的事件类型
 77 
 78     ev.events=EPOLLIN|EPOLLET;
 79     //ev.events=EPOLLIN;
 80 
 81     //注册epoll事件
 82 
 83     epoll_ctl(epfd,EPOLL_CTL_ADD,listenfd,&ev);
 84     bzero(&serveraddr, sizeof(serveraddr));
 85     serveraddr.sin_family = AF_INET;
 86     char *local_addr="127.0.0.1";
 87     inet_aton(local_addr,&(serveraddr.sin_addr));//htons(portnumber);
 88 
 89     serveraddr.sin_port=htons(portnumber);
 90     bind(listenfd,(sockaddr *)&serveraddr, sizeof(serveraddr));
 91     listen(listenfd, LISTENQ);
 92     maxi = 0;
 93     for ( ; ; ) {
 94         //等待epoll事件的发生
 95 
 96         nfds=epoll_wait(epfd,events,20,500);
 97         //处理所发生的所有事件
 98 
 99         for(i=0;i<nfds;++i)
100         {
101             if(events[i].data.fd==listenfd)//如果新监测到一个SOCKET用户连接到了绑定的SOCKET端口，建立新的连接。
102 
103             {
104                 connfd = accept(listenfd,(sockaddr *)&clientaddr, &clilen);
105                 if(connfd<0){
106                     perror("connfd<0");
107                     exit(1);
108                 }
109                 //setnonblocking(connfd);
110 
111                 char *str = inet_ntoa(clientaddr.sin_addr);
112                 cout << "accapt a connection from " << str << endl;
113                 //设置用于读操作的文件描述符
114 
115                 ev.data.fd=connfd;
116                 //设置用于注测的读操作事件
117 
118                 ev.events=EPOLLIN|EPOLLET;
119                 //ev.events=EPOLLIN;
120 
121                 //注册ev
122 
123                 epoll_ctl(epfd,EPOLL_CTL_ADD,connfd,&ev);
124             }
125             else if(events[i].events&EPOLLIN)//如果是已经连接的用户，并且收到数据，那么进行读入。
126 
127             {
128                 cout << "EPOLLIN" << endl;
129                 if ( (sockfd = events[i].data.fd) < 0)
130                     continue;
131                 if ( (n = read(sockfd, line, MAXLINE)) < 0) {
132                     if (errno == ECONNRESET) {
133                         close(sockfd);
134                         events[i].data.fd = -1;
135                     } else
136                         std::cout<<"readline error"<<std::endl;
137                 } else if (n == 0) {
138                     close(sockfd);
139                     events[i].data.fd = -1;
140                 }
141                 line[n] = '/0';
142                 cout << "read " << line << endl;
143                 //设置用于写操作的文件描述符
144 
145                 ev.data.fd=sockfd;
146                 //设置用于注测的写操作事件
147 
148                 ev.events=EPOLLOUT|EPOLLET;
149                 //修改sockfd上要处理的事件为EPOLLOUT
150 
151                 //epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev);
152 
153             }
154             else if(events[i].events&EPOLLOUT) // 如果有数据发送
155 
156             {
157                 sockfd = events[i].data.fd;
158                 write(sockfd, line, n);
159                 //设置用于读操作的文件描述符
160 
161                 ev.data.fd=sockfd;
162                 //设置用于注测的读操作事件
163 
164                 ev.events=EPOLLIN|EPOLLET;
165                 //修改sockfd上要处理的事件为EPOLIN
166 
167                 epoll_ctl(epfd,EPOLL_CTL_MOD,sockfd,&ev);
168             }
169         }
170     }
171     return 0;
172 
```

