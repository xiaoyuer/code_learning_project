# MicroService

{% embed url="https://www.cnblogs.com/skabyy/tag/" %}



## [Redis集群](https://www.cnblogs.com/skabyy/p/10013322.html)

这几天工作需要研究了一下Redis集群，将其原理的核心内容记录下来以便以后查阅。

### 集群原理 <a id="&#x96C6;&#x7FA4;&#x539F;&#x7406;"></a>

一个系统建立集群主要需要解决两个问题：数据同步问题和集群容错问题。

#### Naive方案 <a id="naive&#x65B9;&#x6848;"></a>

一个简单粗暴的方案是部署多台一模一样的Redis服务，再用负载均衡来分摊压力以及监控服务状态。这种方案的优势在于容错简单，只要有一台存活，整个集群就仍然可用。但是它的问题在于保证这些Redis服务的数据一致时，会导致大量数据同步操作，反而影响性能和稳定性。

#### Redis集群方案 <a id="redis&#x96C6;&#x7FA4;&#x65B9;&#x6848;"></a>

Redis集群方案基于分而治之的思想。Redis中数据都是以Key-Value形式存储的，而不同Key的数据之间是相互独立的。因此可以将Key按照某种规则划分成多个分区，将不同分区的数据存放在不同的节点上。这个方案类似数据结构中哈希表的结构。在Redis集群的实现中，使用哈希算法（公式是`CRC16(Key) mod 16383`）将Key映射到0~16383范围的整数。这样每个整数对应存储了若干个Key-Value数据，这样一个整数对应的抽象存储称为一个槽（slot）。每个Redis Cluster的节点——准确讲是master节点——负责一定范围的槽，所有节点组成的集群覆盖了0~16383整个范围的槽。

据说任何计算机问题都可以通过增加一个中间层来解决。槽的概念也是这么一层。它介于数据和节点之间，简化了扩容和收缩操作的难度。数据和槽的映射关系由固定算法完成，不需要维护，节点只需维护自身和槽的映射关系。

![&#x69FD;&#x5230;&#x8282;&#x70B9;&#x7684;&#x6620;&#x5C04;](https://img.mukewang.com/5bf9218d0001a44105590608.png)

#### Slave <a id="slave"></a>

上面的方案只是解决了性能扩展的问题，集群的故障容错能力并没有提升。提高容错能力的方法一般为使用某种备份/冗余手段。负责一定数量的槽的节点被称为master节点。为了增加集群稳定性，每个master节点可以配置若干个备份节点——称为slave节点。Slave节点一般作为冷备份保存master节点的数据，在master节点宕机时替换master节点。在一些数据访问压力比较大的情况下，slave节点也可以提供读取数据的功能，不过slave节点的数据实时性会略差一下。而写数据的操作则只能通过master节点进行。

#### 请求重定向 <a id="&#x8BF7;&#x6C42;&#x91CD;&#x5B9A;&#x5411;"></a>

当Redis节点接收到对某个key的命令时，如果这个key对应的槽不在自己的负责范围内，则返回MOVED重定向错误，通知客户端到正确的节点去访问数据。

如果频繁出现重定向错误，势必会影响访问的性能。由于从key映射到槽的算法是固定公开的，客户端可以在内部维护槽到节点的映射关系，访问数据时可以自己通过key计算出槽，然后找到正确的节点，减少重定向错误。目前大部分开发语言的Redis客户端都会实现这个策略。这个地址[https://redis.io/clients](https://redis.io/clients)可以查看主流语言的Redis客户端。

### 节点通信 <a id="&#x8282;&#x70B9;&#x901A;&#x4FE1;"></a>

尽管不同节点存储的数据相互独立，这些节点仍然需要相互通信以同步节点状态信息。Redis集群采用P2P的Gossip协议，节点之间不断地通信交换信息，最终所有节点的状态都会达成一致。常用的Gossip消息有下面几种：

* ping消息：每个节点不断地向其他节点发起ping消息，用于检测节点是否在线和交换节点状态信息。
* pong消息：收到ping、meet消息时的响应消息。
* meet消息：新节点加入消息。
* fail消息：节点下线消息。
* forget消息：忘记节点消息，使一个节点下线。这个命令必须在60秒内在所有节点执行，否则超过60秒后该节点重新参与消息交换。实践中不建议直接使用forget命令来操作节点下线。

#### 节点下线 <a id="&#x8282;&#x70B9;&#x4E0B;&#x7EBF;"></a>

当某个节点出现问题时，需要一定的传播时间让多数master节点认为该节点确实不可用，才能标记标记该节点真正下线。Redis集群的节点下线包括两个环节：主观下线（pfail）和客观下线（fail）。

* 主观下线：当节点A在cluster-node-timeout时间内和节点B通信（ping-pong消息）一直失败，则节点A认为节点B不可用，标记为主观下线，并将状态消息传播给其他节点。
* 客观下线：当一个节点被集群内多数master节点标记为主观下线后，则触发客观下线流程，标记该节点真正下线。

#### 故障恢复 <a id="&#x6545;&#x969C;&#x6062;&#x590D;"></a>

一个持有槽的master节点客观下线后，集群会从slave节点中选出一个提升为master节点来替换它。Redis集群使用选举-投票的算法来挑选slave节点。一个slave节点必须获得包括故障的master节点在内的多数master节点的投票后才能被提升为master节点。假设集群规模为3主3从，则必须至少有2个主节点存活才能执行故障恢复。如果部署时将2个主节点部署到同一台服务器上，则该服务器不幸宕机后集群无法执行故障恢复。

默认情况下，Redis集群如果有master节点不可用，即有一些槽没有负责的节点，则整个集群不可用。也就是说当一个master节点故障，到故障恢复的这段时间，整个集群都处于不可用的状态。这对于一些业务来说是不可忍受的。可以**在配置中将cluster-require-full-coverage配置为no**，那么master节点故障时只会影响访问它负责的相关槽的数据，不影响对其他节点的访问。

### 搭建集群 <a id="&#x642D;&#x5EFA;&#x96C6;&#x7FA4;"></a>

#### 启动新节点 <a id="&#x542F;&#x52A8;&#x65B0;&#x8282;&#x70B9;"></a>

修改Redis配置文件以启动集群模式：

```text
# 开启集群模式
cluster-enabled yes
# 节点超时时间，单位毫秒
cluster-node-timeout 15000
# 集群节点信息文件
cluster-config-file "nodes-6379.conf"
```

然后启动新节点。

#### 发送meet消息将节点组成集群 <a id="&#x53D1;&#x9001;meet&#x6D88;&#x606F;&#x5C06;&#x8282;&#x70B9;&#x7EC4;&#x6210;&#x96C6;&#x7FA4;"></a>

使用客户端发起命令`cluster <ip> <port>`，节点会发送meet消息将指定IP和端口的新节点加入集群。

![&#x53D1;&#x9001;meet&#x6D88;&#x606F;](https://img.mukewang.com/5bf936c70001123509030129.png)

#### 分配槽 <a id="&#x5206;&#x914D;&#x69FD;"></a>

上一步执行完后我们得到的是一个还没有负责任何槽的“空”集群。为了使集群可用，我们需要将16384个槽都分配到master节点数。

在客户端执行`cluster add addslots {<a>...<b>}`命令，将`<a>`~`<b>`范围的槽都分配给当前客户端所连接的节点。将所有的槽都分配给master节点后，执行`cluster nodes`命令，查看各个节点负责的槽，以及节点的ID。

接下来还需要分配slave节点。使用客户端连接待分配的slave节点，执行`cluster replicate <nodeId>`命令，将该节点分配为`<nodeId>`指定的master节点的备份。

#### 使用命令直接创建集群 <a id="&#x4F7F;&#x7528;&#x547D;&#x4EE4;&#x76F4;&#x63A5;&#x521B;&#x5EFA;&#x96C6;&#x7FA4;"></a>

在Redis 5版本中`redis-cli`客户端新增了集群操作命令。

如下所示，直接使用命令创建一个3主3从的集群：

```text
redis-cli --cluster create 127.0.0.1:7000 127.0.0.1:7001 \
127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005 \
--cluster-replicas 1
```

如果你用的是旧版本的Redis，可以使用官方提供的`redis-trib.rb`脚本来创建集群:

```text
./redis-trib.rb create --replicas 1 127.0.0.1:7000 127.0.0.1:7001 \
127.0.0.1:7002 127.0.0.1:7003 127.0.0.1:7004 127.0.0.1:7005
```

### 集群伸缩 <a id="&#x96C6;&#x7FA4;&#x4F38;&#x7F29;"></a>

#### 扩容 <a id="&#x6269;&#x5BB9;"></a>

扩容操作与创建集群操作类似，不同的在于最后一步是将槽从已有的节点迁移到新节点。

1. 启动新节点：同创建集群。
2. 将新节点加入到集群：使用`redis-cli --cluster add-node`命令将新节点加入集群（内部使用meet消息实现）。
3. 迁移槽和数据：添加新节点后，需要将一些槽和数据从旧节点迁移到新节点。使用命令`redis-cli --cluster reshard`进行槽迁移操作。

#### 收缩 <a id="&#x6536;&#x7F29;"></a>

为了安全删除节点，Redis集群只能下线没有负责槽的节点。因此如果要下线有负责槽的master节点，则需要先将它负责的槽迁移到其他节点。

1. 迁移槽。使用命令`redis-cli --cluster reshard`将待删除节点的槽都迁移到其他节点。
2. 忘记节点。使用命令`redis-cli --cluster del-node`删除节点（内部使用forget消息实现）。

#### 集群配置工具 <a id="&#x96C6;&#x7FA4;&#x914D;&#x7F6E;&#x5DE5;&#x5177;"></a>

如果你的redis-cli版本低于5，那么可以使用redis-trib.rb脚本来完成上面的命令。点击[这里](https://redis.io/topics/cluster-tutorial)查看`redis-cli`和`redis-trib.rb`操作集群的命令。

### 持久化 <a id="&#x6301;&#x4E45;&#x5316;"></a>

Redis有RDB和AOF两种持久化策略。[这篇文章](https://juejin.im/post/5b70dfcf518825610f1f5c16)详细讲解了RDB和AOF持久化原理。

#### 一个RDB持久化的坑 <a id="&#x4E00;&#x4E2A;rdb&#x6301;&#x4E45;&#x5316;&#x7684;&#x5751;"></a>

RDB持久化神坑：

* 即使设置了`save ""`试图关闭RDB，然而RDB持久化仍然有可能会触发。
* 从节点全量复制（比如新增从节点时），主节点触发RDB持久化产生RDB文件。然后发送RDB文件给从节点。最后该从节点和对应的主节点都会有RDB文件。
* 执行shutdown时，如果没有开启AOF，也会触发RDB持久化。
* 不管save如何设置，只要RDB文件存在，redis启动时就会去加载该文件。

后果：

* 如果关闭了RDB持久化（以及AOF持久化），那么**当Redis重启时，则会加载上一次从节点全量复制或者执行shutdown时保存的RDB文件**。而这个RDB文件很可能是一份过时已久的数据。
* Cluster模式下，Redis重启并从RDB文件恢复数据后，如果没有读取到cluster-config-file中nodes的配置，则**标记自己为单独的master并占用从RDB中恢复的数据的Key对应的槽，导致此节点无法再加入其它集群**。



## [RabbitMQ权限控制原理](https://www.cnblogs.com/skabyy/p/10035158.html)

我们在使用MQ搭建系统的时候，经常要开放队列给外接系统访问。外接系统的稳定性是不可控的。为了防止外接系统不稳定导致误操作破坏了MQ的配置或数据，需要对MQ做比较精细的权限控制。

我的需求是这样的：

我有一个数据查询服务，并且通过MQ推送数据变动消息。对接MQ的每个系统都会有自己一个独立的队列来读取消息。所有消息通过一个扇形交换机广播到所有队列。我需要这个交换机和所有队列都由管理员统一创建好。而其他系统使用的用户，均没有创建交换机和队列的权限。数据查询服务只拥有推送消息的权限，其他对接MQ的系统只拥有从自己队列读取消息的权限。

我们使用的MQ是RabbitMQ。我在网上搜了一下，大部分讲的是用户角色配置。对于MQ的资源授权管理讲的比较少。以下内容将主要讲解**RabbitMQ权限控制的基本概念和模型**。理解了这些基本概念后，应该可以愉快地使用RabbitMQ管理界面进行授权操作。如果你们只有命令行可用，也能很轻松地找到相应的命令。

### RabbitMQ初始化 <a id="rabbitmq&#x521D;&#x59CB;&#x5316;"></a>

RabbitMQ初次启动时，初始创建这两个东西：

* 一个名称为`/`的virtual host
* guest用户，拥有`/`的全部权限，只能localhost访问

### RabbitMQ授权模型 <a id="rabbitmq&#x6388;&#x6743;&#x6A21;&#x578B;"></a>

第一级控权单位是virtual host，virtual host下面第二级的控权单位是resource（包含exchange和queue）。两个相同名称的resource如果分属不同的virtual host，则算是不同的resource。

> 什么是virtual host：
>
> RabbitMQ is multi-tenant system: connections, exchanges, queues, bindings, user permissions, policies and some other things belong to virtual hosts, logical groups of entities.
>
> 就是说RabbitMQ是多租户系统，简单理解就是把多virtual host当做多个MQ系统来用就好了……

当用户访问MQ时，首先触发**第一级控权，判断用户是否有访问该virtual host的权限**。

若可访问，则进行**第二级控权，判断用户是否具有操作（operation）所请求的资源的权限**。

RabbitMQ定义了操作（operation）有这么三种：

* configure：主要对应创建exchange和queue操作；
* write：write主要对应绑定和推送消息操作；
* read：read主要对应读取消息操作。

后面有个表格列出了具体的对应关系。

当管理员对一个用户进行授权时，要配置两个元素：

1. 允许什么操作，即configure、write、read三种operation；
2. 操作什么resource。用户是否拥有某资源的权限，通过对该资源的名称与授权时配置的正则进行匹配来判断。

![](https://img2018.cnblogs.com/blog/576869/201811/576869-20181128214223220-1400600904.png)

下面这张表详细描述了operation、resource和用户可执行的操作的关系：

![](https://img2018.cnblogs.com/blog/576869/201811/576869-20181128214207372-1820049281.png)

例子：

* 如果要给用户授权可以往exchange `foo`推消息，则我们找到basic.publish这行，格子不是空的只有write这列，所以我们需要给用户授权一个write权限，其正则可以匹配字符串`foo`（比如说`^foo$`，或者`.*`等）。
* 如果要给用户授权只能从queue `bar`读取消息，则需要给用户授权一个read权限，其正则可以匹配字符串`bar`。

### 进一步了解 <a id="&#x8FDB;&#x4E00;&#x6B65;&#x4E86;&#x89E3;"></a>

本文内容基本来自官网手册，如果需要更详细的说明——比如说topic的授权，可以直接看[手册](https://www.rabbitmq.com/access-control.html)。很多时候，当你刚接触一个新工具时，比起在互联网上瞎逛，直接阅读官方手册效率会高很多。虽然手册比较冗长，而且大部分只有纯英文，但毕竟最远的路，就是最快的捷径。





## [Istio入门实战与架构原理——使用Docker Compose搭建Service Mesh](https://www.cnblogs.com/skabyy/p/10668079.html)

本文将介绍如何使用Docker Compose搭建Istio。Istio号称支持多种平台（不仅仅Kubernetes）。然而，官网上非基于Kubernetes的教程仿佛不是亲儿子，写得非常随便，不仅缺了一些内容，而且还有坑。本文希望能补实这些内容。我认为在学习Istio的过程中，相比于Kubernetes，使用Docker Compose部署更能深刻地理解Istio各个组件的用处以及他们的交互关系。在理解了这些后，可以在其他环境，甚至直接在虚拟机上部署Istio。当然，**生产环境建议使用Kubernetes等成熟的容器框架**。

本文使用官方的[Bookinfo示例](https://istio.io/zh/docs/examples/bookinfo/)。通过搭建Istio控制平面，部署Bookinfo应用，最后配置路由规则，展示Istio基本的功能和架构原理。

本文涉及的名词、用到的端口比较多。Don't panic.

> 为了防止不提供原网址的转载，特在这里加上原文链接：  
> [https://www.cnblogs.com/skabyy/p/10668079.html](https://www.cnblogs.com/skabyy/p/10668079.html)

### 准备工作 <a id="&#x51C6;&#x5907;&#x5DE5;&#x4F5C;"></a>

1. 安装Docker和Docker Compose。
2. 安装`kubectl`（Kubernetes的客户端）。
3. 下载[Istio release 1.1.0](https://github.com/istio/istio/releases/download/1.1.0/istio-1.1.0-linux.tar.gz)并解压。注意，这里下载的是Linux的版本。即使你用的Windows或者OSX操作系统，也应该下载Linux版本的Istio，因为我们要放到Docker容器里去运行的。

### Service Mesh架构 <a id="service-mesh&#x67B6;&#x6784;"></a>

在微服务架构中，通常除了实现业务功能的微服务外，我们还会部署一系列的基础组件。这些基础组件有些会入侵微服务的代码。比如服务发现需要微服务启动时注册自己，链路跟踪需要在HTTP请求的headers中插入数据，流量控制需要一整套控制流量的逻辑等。这些入侵的代码需要在所有的微服务中保持一致。这导致了开发和管理上的一些难题。

为了解决这个问题，我们再次**应用抽象和服务化的思想，将这些需要入侵的功能抽象出来，作为一个独立的服务。这个独立的服务被称为sidecar**，这种模式叫**Sidecar模式**。对每个微服务节点，都需要额外部署一个sidecar来负责业务逻辑外的公共功能。所有的出站入站的网络流量都会先经过sidecar进行各种处理或者转发。这样微服务的开发就不需要考虑业务逻辑外的问题。另外所有的sidecar都是一样的，只需要部署的时候使用合适的编排工具即可方便地为所有节点注入sidecar。

> Sidecar不会产生额外网络成本。Sidecar会和微服务节点部署在同一台主机上并且共用相同的虚拟网卡。所以sidecar和微服务节点的通信实际上都只是通过内存拷贝实现的。

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233101603-682213801.png)

> 图片来自：[Pattern: Service Mesh](http://philcalcado.com/2017/08/03/pattern_service_mesh.html)

Sidecar只负责网络通信。还需要有个组件来统一管理所有sidecar的配置。在Service Mesh中，负责网络通信的部分叫数据平面（data plane），负责配置管理的部分叫控制平面（control plane）。数据平面和控制平面构成了Service Mesh的基本架构。

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233125673-953968693.png)

> 图片来自：[Pattern: Service Mesh](http://philcalcado.com/2017/08/03/pattern_service_mesh.html)

Istio的数据平面主要由Envoy实现，控制平面则主要由Istio的Pilot组件实现。

### 部署控制平面 <a id="&#x90E8;&#x7F72;&#x63A7;&#x5236;&#x5E73;&#x9762;"></a>

如果你使用Linux操作系统，需要先配置`DOCKER_GATEWAY`环境变量。非Linux系统不要配。

```text
$ export DOCKER_GATEWAY=172.28.0.1:
```

到`install/consul`目录下，使用`istio.yaml`文件启动控制平面：

> 根据自己的网络情况（你懂得），可以把`istio.yaml`中的镜像`gcr.io/google_containers/kube-apiserver-amd64:v1.7.3`换成`mirrorgooglecontainers/kube-apiserver-amd64:v1.7.3`。

```text
$ docker-compose -f istio.yaml up -d
```

用命令`docker-compose -f istio.yaml ps`看一下是不是所有组件正常运行。你可能（大概率）会看到pilot的状态是`Exit 255`。使用命令`docker-compose -f istio.yaml logs | grep pilot`查看日志发现，`pilot`启动时访问`istio-apiserver`失败。这是因为Docker Compose是同时启动所有容器的，在`pilot`启动时，`istio-apiserver`也是处于启动状态，所以访问`istio-apiserver`就失败了。

等`istio-apiserver`启动完成后，重新运行启动命令就能成功启动`pilot`了。你也可以写一个脚本来自动跑两次命令：

```text
docker-compose -f istio.yaml up -d
# 有些依赖别人的第一次启动会挂
sec=10  # 根据你的机器性能这个时间可以修改
echo "Wait $sec seconds..."
sleep $sec
docker-compose -f istio.yaml up -d
docker-compose -f istio.yaml ps
```

配置`kubectl`，让`kubectl`使用我们刚刚部署的`istio-apiserver`作为服务端。我们后面会使用`kubectl`来执行配置管理的操作。

```text
$ kubectl config set-context istio --cluster=istio
$ kubectl config set-cluster istio --server=http://localhost:8080
$ kubectl config use-context istio
```

部署完成后，使用地址`localhost:8500`可以访问`consul`，使用地址`localhost:9411`可以访问`zipkin`。

### 控制平面架构 <a id="&#x63A7;&#x5236;&#x5E73;&#x9762;&#x67B6;&#x6784;"></a>

在下一步之前，我们先来看一下控制平面都由哪些组件组成。下面是`istio.yaml`文件的内容：

```text
# GENERATED FILE. Use with Docker-Compose and consul
# TO UPDATE, modify files in install/consul/templates and run install/updateVersion.sh
version: '2'
services:
  etcd:
    image: quay.io/coreos/etcd:latest
    networks:
      istiomesh:
        aliases:
          - etcd
    ports:
      - "4001:4001"
      - "2380:2380"
      - "2379:2379"
    environment:
      - SERVICE_IGNORE=1
    command: ["/usr/local/bin/etcd", "-advertise-client-urls=http://0.0.0.0:2379", "-listen-client-urls=http://0.0.0.0:2379"]

  istio-apiserver:
    # 如果这个镜像下载不了的话，可以换成下面的地址：
    # image: mirrorgooglecontainers/kube-apiserver-amd64:v1.7.3
    image: gcr.io/google_containers/kube-apiserver-amd64:v1.7.3
    networks:
      istiomesh:
        ipv4_address: 172.28.0.13
        aliases:
          - apiserver
    ports:
      - "8080:8080"
    privileged: true
    environment:
      - SERVICE_IGNORE=1
    command: ["kube-apiserver", "--etcd-servers", "http://etcd:2379", "--service-cluster-ip-range", "10.99.0.0/16", "--insecure-port", "8080", "-v", "2", "--insecure-bind-address", "0.0.0.0"]

  consul:
    image: consul:1.3.0
    networks:
      istiomesh:
        aliases:
          - consul
    ports:
      - "8500:8500"
      - "${DOCKER_GATEWAY}53:8600/udp"
      - "8400:8400"
      - "8502:8502"
    environment:
      - SERVICE_IGNORE=1
      - DNS_RESOLVES=consul
      - DNS_PORT=8600
      - CONSUL_DATA_DIR=/consul/data
      - CONSUL_CONFIG_DIR=/consul/config
    entrypoint:
      - "docker-entrypoint.sh"
    command: ["agent", "-bootstrap", "-server", "-ui",
              "-grpc-port", "8502"
              ]
    volumes:
      - ./consul_config:/consul/config

  registrator:
    image: gliderlabs/registrator:latest
    networks:
      istiomesh:
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock
    command: ["-internal", "-retry-attempts=-1", "consul://consul:8500"]

  pilot:
    image: docker.io/istio/pilot:1.1.0
    networks:
      istiomesh:
        aliases:
          - istio-pilot
    expose:
      - "15003"
      - "15005"
      - "15007"
    ports:
      - "8081:15007"
    command: ["discovery",
              "--httpAddr", ":15007",
              "--registries", "Consul",
              "--consulserverURL", "http://consul:8500",
              "--kubeconfig", "/etc/istio/config/kubeconfig",
              "--secureGrpcAddr", "",
              ]
    volumes:
      - ./kubeconfig:/etc/istio/config/kubeconfig

  zipkin:
    image: docker.io/openzipkin/zipkin:2.7
    networks:
      istiomesh:
        aliases:
          - zipkin
    ports:
      - "9411:9411"

networks:
  istiomesh:
    ipam:
      driver: default
      config:
        - subnet: 172.28.0.0/16
          gateway: 172.28.0.1
```

控制平面部署了这几个组件（使用`istio.yaml`里写的名称）：

* `etcd`：分布式key-value存储。Istio的配置信息存在这里。
* `istio-apiserver`：实际上是一个`kube-apiserver`，提供了Kubernetes格式数据的读写接口。
* `consul`：服务发现。
* `registrator`：监听Docker服务进程，自动将容器注册到`consul`。
* `pilot`：从`consul`和`istio-apiserver`收集主机信息与配置数据，并下发到所有的sidecar。
* `zipkin`：链路跟踪组件。与其他组件的关系相对独立。

这些组件间的关系如下图：

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233155774-572225662.png)

控制平面主要实现了以下两个功能：

* **借用Kubernetes API管理配置数据**。`etcd`和`kube-apiserver`的组合可以看作是一个对象存储系统，它提供了读写接口和变更事件，并且可以直接使用`kubectl`作为客户端方便地进行操作。Istio直接使用这个组合作为控制平面的持久化层，节省了重复开发的麻烦，另外也兼容了Kubernetes容器框架。
* **使用Pilot-discovery将主机信息与配置数据同步到Envoy**。`pilot`容器中实际执行的是`pilot-discovery`（发现服务）。它从`consul`收集各个主机的域名和IP的对应关系，从`istio-apiserver`获取流量控制配置，然后按照Envoy的xDS API规范生成Envoy配置，下发到所有sidecar。

### 部署微服务和sidecar <a id="&#x90E8;&#x7F72;&#x5FAE;&#x670D;&#x52A1;&#x548C;sidecar"></a>

接下来我们开始部署微服务。这里我们使用Istio提供的例子，一个[Bookinfo应用](https://istio.io/zh/docs/examples/bookinfo/)。

Bookinfo 应用分为四个单独的微服务：

* `productpage`：`productpage`微服务会调用`details`和`reviews`两个微服务，用来生成页面。
* `details`：这个微服务包含了书籍的信息。
* `reviews`：这个微服务包含了书籍相关的评论。它还会调用`ratings`微服务。
* `ratings`：`ratings`微服务中包含了由书籍评价组成的评级信息。

`reviews`微服务有3个版本：

* v1版本不会调用`ratings`服务。
* v2版本会调用`ratings`服务，并使用1到5个黑色星形图标来显示评分信息。
* v3版本会调用`ratings`服务，并使用1到5个红色星形图标来显示评分信息。

Bookinfo应用的架构如下图所示：

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233209220-344051003.png)

> 图片来自：[Bookinfo应用](https://istio.io/zh/docs/examples/bookinfo/)

首先，我们切换到这个示例的目录`samples/bookinfo/platform/consul`下。

使用`bookinfo.yaml`文件启动所有微服务：

```text
$ docker-compose -f bookinfo.yaml up -d
```

这里只启动了微服务，还需使用`bookinfo.sidecar.yaml`文件启动所有sidecar：

```text
$ docker-compose -f bookinfo.sidecars.yaml up -d
```

部署完毕。但是当我们访问时……

Bookinfo暴露到外面的端口是9081，使用地址`localhost:9081/productpage`访问`productpage`页面。

Emmm……出错了：

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233225026-1283798248.png)

本来应该显示`reviews`的部分报错了，而`details`还是正常的。经过一番排查，我们发现，在所有微服务的容器上，不管你访问的是`productpage`、`details`、`reviews`还是`ratings`，网络请求都会跑到`details`。

> 你的情况不一定是`details`，也有可能所有流量都跑到另外的某个服务。这是随机的。

```text
# 到reviews查reviews，返回404
$ docker exec -it consul_ratings-v1_1 curl reviews.service.consul:9080/reviews/0
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.0//EN">
<HTML>
  <HEAD><TITLE>Not Found</TITLE></HEAD>
  <BODY>
    <H1>Not Found</H1>
    `/reviews/0' not found.
    <HR>
    <ADDRESS>
     WEBrick/1.3.1 (Ruby/2.3.8/2018-10-18) at
     reviews.service.consul:9080
    </ADDRESS>
  </BODY>
</HTML>

# 到reviews查details，反倒能查出数据，诡异的路由……
$ docker exec -it consul_ratings-v1_1 curl reviews.service.consul:9080/details/0
{"id":0,"author":"William Shakespeare","year":1595,"type":"paperback","pages":200,"publisher":"PublisherA","language":"English","ISBN-10":"1234567890","ISBN-13":"123-1234567890"}
```

不用怀疑部署的时候哪里操作失误了，就是官方的部署文件有坑……

要解决这个问题，我们来看看sidecar的原理。

### Istio Sidecar模式的原理 <a id="istio-sidecar&#x6A21;&#x5F0F;&#x7684;&#x539F;&#x7406;"></a>

首先看看两个部署用的yaml文件都做了什么。由于每个微服务的部署都大同小异，这里只贴出`productpage`相关的内容。

**`bookinfo.yaml`：**

```text
version: '2'
services:
  ……
  productpage-v1:
    image: istio/examples-bookinfo-productpage-v1:1.10.1
    networks:
      istiomesh:
        ipv4_address: 172.28.0.14
    dns:
      - 172.28.0.1
      - 8.8.8.8
    dns_search:
        - service.consul
    environment:
      - SERVICE_NAME=productpage
      - SERVICE_TAGS=version|v1
      - SERVICE_PROTOCOL=http
    ports:
      - "9081:9080"
    expose:
      - "9080"
  ……
```

* `dns_search: - search.consul`。Docker Compose部署的这套样例对短服务主机名的解析可能会有问题，所以这里需要加个后缀。
* `environment`环境变量的几个设置。`registrator`会以这些环境变量为配置将服务注册到`consul`。`SERVICE_NAME`是注册的服务名，`SERVICE_TAGS`是注册服务的`ServiceTags`，而`SERVICE_PROTOCOL=http`则会将`protocol: http`加入到`ServiceMeta`。

**`bookinfo.sidecar.yaml`：**

```text
version: '2'
services:
  ……
  productpage-v1-init:
    image: docker.io/istio/proxy_init:0.7.1
    cap_add:
      - NET_ADMIN
    network_mode: "container:consul_productpage-v1_1"
    command:
      - -p
      - "15001"
      - -u
      - "1337"
  productpage-v1-sidecar:
    image: docker.io/istio/proxy_debug:1.1.0
    network_mode: "container:consul_productpage-v1_1"
    entrypoint:
      - su
      - istio-proxy
      - -c
      - "/usr/local/bin/pilot-agent proxy --serviceregistry Consul --serviceCluster productpage-v1 --zipkinAddress zipkin:9411 --configPath /var/lib/istio >/tmp/envoy.log"
  ……
```

* sidecar由两部分组成，一个是负责初始化的`proxy_init`，这个容器执行完就退出了；另一个是实际的sidecar程序`proxy_debug`。
* 注意这两个容器的`network_mode`，值为`container:consul_productpage-v1_1`。这是Docker的容器网络模式，意思是这两个容器和`productpage-v1`共用同一个虚拟网卡，即它们在相同网络栈上。

#### `proxy_init` <a id="proxy_init"></a>

sidecar的网络代理一般是将一个端口转发到另一个端口。所以微服务使用的端口就必须和对外暴露的端口不一样，这样一来sidecar就不够透明。

为了使sidecar变得透明，以Istio使用`proxy_init`设置了iptables的转发规则（`proxy_init`、`proxy_debug`和`productpage-v1`在相同的网络栈上，所以这个配置对这三个容器都生效）。添加的规则为：

1. 回环网络的流量不处理。
2. 用户ID为1337的流量不处理。1337是Envoy进程的用户ID，这条规则是为了防止流量转发死循环。
3. 所有出站入站的流量除了规则1和规则2外，都转发到15001端口——这是Envoy监听的端口。

比如`productpage`服务使用的9080端口，当其他服务通过9080端口访问`productpage`是，请求会先被iptables转发到15001端口，Envoy再根据路由规则转发到9080端口。这样访问9080的流量实际上都在15001绕了一圈，但是对外部来说，这个过程是透明的。

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233255202-1270446561.png)

#### `proxy_debug` <a id="proxy_debug"></a>

`proxy_debug`有两个进程：`pilot-agent`和`envoy`。`proxy_debug`启动时，会先启动`pilot-agent`。`pilot-agent`做的事很简单，它生成了`envoy`的初始配置文件`/var/lib/istio/envoy-rev0.json`，然后启动`envoy`。后面的事就都交给`envoy`了。

使用下面命令导出初始配置文件：

```text
$ docker exec -it consul_productpage-v1-sidecar_1 cat /var/lib/istio/envoy-rev0.json > envoy-rev0.json
```

使用你心爱的编辑器打开初始配置文件，可以看到有这么一段：

```text
……
        "name": "xds-grpc",
        "type": "STRICT_DNS",
        "connect_timeout": "10s",
        "lb_policy": "ROUND_ROBIN",
        
        "hosts": [
          {
            "socket_address": {"address": "istio-pilot", "port_value": 15010}
          }
        ],
……
```

这一段的意思是`envoy`会连接到`pilot`（控制平面的组件，忘记了请往上翻翻）的15010端口。这俩将按照xDS的API规范，使用GRPC协议实时同步配置数据。

> xDS是Envoy约定的一系列发现服务（Discovery Service）的统称。如CDS（Cluster Discovery Service），EDS（Endpoint Discovery Service），RDS（Route Discovery Service）等。Envoy动态配置需要从实现了xDS规范的接口（比如这里的`pilot-discovery`）获取配置数据。

总结一下，Envoy配置初始化流程为：

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233309448-1026141117.png)

> 图片来自：[Istio流量管理实现机制深度解析](http://www.servicemesher.com/blog/istio-traffic-management-impl-intro/)

那么说`envoy`实际使用的路由配置并不在初始配置文件中，而是`pilot`生成并推送过来的。如何查看`envoy`的当前配置呢？还好`envoy`暴露了一个管理端口15000：

```text
$ docker exec -it consul_productpage-v1-sidecar_1 curl localhost:15000/help
admin commands are:
  /: Admin home page
  /certs: print certs on machine
  /clusters: upstream cluster status
  /config_dump: dump current Envoy configs (experimental)
  /contention: dump current Envoy mutex contention stats (if enabled)
  /cpuprofiler: enable/disable the CPU profiler
  /healthcheck/fail: cause the server to fail health checks
  /healthcheck/ok: cause the server to pass health checks
  /help: print out list of admin commands
  /hot_restart_version: print the hot restart compatibility version
  /listeners: print listener addresses
  /logging: query/change logging levels
  /memory: print current allocation/heap usage
  /quitquitquit: exit the server
  /reset_counters: reset all counters to zero
  /runtime: print runtime values
  /runtime_modify: modify runtime values
  /server_info: print server version/status information
  /stats: print server stats
  /stats/prometheus: print server stats in prometheus format
```

我们可以通过`/config_dump`接口导出`envoy`的当前配置：

```text
$ docker exec -it consul_productpage-v1-sidecar_1 curl localhost:15000/config_dump > envoy.json
```

打开这个配置，看到这么一段：

```text
……
     "listener": {
      "name": "0.0.0.0_9080",
      "address": {
       "socket_address": {
        "address": "0.0.0.0",
        "port_value": 9080
       }
      },
      "filter_chains": [
       {
        "filters": [
         {
          "name": "envoy.tcp_proxy",
          "typed_config": {
           "@type": "type.googleapis.com/envoy.config.filter.network.tcp_proxy.v2.TcpProxy",
           "stat_prefix": "outbound|9080||details.service.consul",
           "cluster": "outbound|9080||details.service.consul",
……
```

猜一下也能知道，这一段的意思是，访问目标地址9080端口的出站流量，都会被路由到`details`。太坑了！！！

### 解决问题 <a id="&#x89E3;&#x51B3;&#x95EE;&#x9898;"></a>

从上面原理分析可知，这个问题的根源应该在于`pilot`给Envoy生成的配置不正确。

查看`pilot`[源码](https://github.com/istio/istio)得知，`pilot`在生成配置时，用一个`map`保存Listener信息。这个map的key为`<ip>:<port>`。如果服务注册的时候，没有指明端口`<port>`上的协议的话，默认认为TCP协议。`pilot`会将这个Listener和路由写入到这个`map`，并拒绝其他相同地址端口再来监听。于是只有第一个注册的服务的路由会生效，所有流量都会走到那个服务。如果这个端口有指定使用HTTP协议的话，Pilot-discovery这里生成的是一个RDS的监听，这个RDS则根据域名路由到正确的地址。

**简单说就是所有微服务在注册到`consul`时应该在`ServiceMeta`中说明自己9080端口的协议是`http`**。

等等，前面的`bookinfo.yaml`配置里，有指定9080端口的协议是了呀。我们访问一下`consul`的接口看下`ServiceMeta`是写入了没有：

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233619577-1446488298.png)

果然没有……看来Registrator注册的时候出了岔子。网上搜了下，确实有Issue提到了这个问题：[gliderlabs/registrator\#633](https://github.com/gliderlabs/registrator/issues/633)。**`istio.yaml`中使用的`latest`版本的Registrator不支持写入Consul的ServiceMeta。应该改为`master`版本**。

修改一下`istio.yaml`配置。按照部署倒叙关闭sidecar、微服务，重新启动控制平面，等`registrator`启动完毕后，重新部署微服务和sidecar。

```text
# /samples/bookinfo/platform/consul
$ docker-compose -f bookinfo.sidecars.yaml down
$ docker-compose -f bookinfo.yaml down
# /install/consul
$ docker-compose -f istio.yaml up -d
# /samples/bookinfo/platform/consul
$ docker-compose -f bookinfo.yaml up -d
$ docker-compose -f bookinfo.sidecars.yaml up -d
```

再访问`consul`的接口试试，有了（没有的话可能是`registrator`没启动好导致没注册到`consul`，再新部署下微服务和sidecar）：

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233607066-757257293.png)

再访问页面，OK了。目前没有配置路由规则，`reviews`的版本是随机的。多刷新几次页面，可以看到打星在“没有星星”、“黑色星星”和“红色星星”三种效果间随机切换。

使用地址`http://localhost:9411`能访问Zipkin链路跟踪系统，查看微服务请求链路调用情况。

我们来看看正确的配置是什么内容。再取出Envoy的配置，`0.0.0.0_9080`的Listener内容变为：

```text
……
     "listener": {
      "name": "0.0.0.0_9080",
      "address": {
       "socket_address": {
        "address": "0.0.0.0",
        "port_value": 9080
       }
      },
      "filter_chains": [
       {
        "filters": [
         {
          "name": "envoy.http_connection_manager",
          "typed_config": {
           "@type": "type.googleapis.com/envoy.config.filter.network.http_connection_manager.v2.HttpConnectionManager",
           "stat_prefix": "0.0.0.0_9080",
           "rds": {
            "config_source": {
             "ads": {}
            },
            "route_config_name": "9080"
           },
……
```

9080端口的出站路由规则由一个名称为`"9080"`的`route_config`定义。找一下这个`route_config`：

```text
……
     "route_config": {
      "name": "9080",
      "virtual_hosts": [
       {
        "name": "details.service.consul:9080",
        "domains": [
         "details.service.consul",
         "details.service.consul:9080",
         "details",
         "details:9080",
         "details.service",
         "details.service:9080"
        ],
        "routes": [
         {
          "match": {
           "prefix": "/"
          },
          "route": {
           "cluster": "outbound|9080|v1|details.service.consul",
……
          },
……
         }
        ]
       },
       {
        "name": "productpage.service.consul:9080",
        "domains": [
         "productpage.service.consul",
         "productpage.service.consul:9080",
         "productpage",
         "productpage:9080",
         "productpage.service",
         "productpage.service:9080"
        ],
        "routes": [
         {
          "match": {
           "prefix": "/"
          },
          "route": {
           "cluster": "outbound|9080|v1|productpage.service.consul",
……
          },
……
         }
        ]
       },
……
```

由于内容太长，这里只贴`details`和`productpage`的关键内容。可以看到，9080端口的出站流量会根据目标地址的域名正确地转发到对应的微服务。

### Istio路由控制 <a id="istio&#x8DEF;&#x7531;&#x63A7;&#x5236;"></a>

**注意：本节工作目录为`/samples/bookinfo/platform/consul`**。

最后我们尝试一下Istio的路由控制能力。在配置路由规则之前，我们要先使用DestinationRule定义各个微服务的版本：

```text
$ kubectl apply -f destination-rule-all.yaml
```

> DestinationRule：DestinationRule定义了每个服务下按照某种策略分割的子集。在本例子中按照版本来分子集，`reviews`分为v1、v2、v3三个版本的子集，其他微服务都只有v1一个子集。

使用命令`kubectl get destinationrules -o yaml`可以查看已配置的DestinationRule。

接下来我们使用VirtualService来配置路由规则。`virtual-service-all-v1.yaml`配置会让所有微服务的流量都路由到v1版本。

```text
$ kubectl apply -f virtual-service-all-v1.yaml
```

> VirtualService：定义路由规则，按照这个规则决定每次请求服务应该将流量转发到哪个子集。

使用命令`kubectl get virtualservices -o yaml`可以查看已配置的VirtualService。

再刷新页面，现在不管刷新多少次，`reviews`都会使用v1版本，也就是页面不会显示星星。

下面我们试一下基于用户身份的路由规则。配置文件`virtual-service-reviews-test-v2.yaml`配置了`reviews`的路由，让用户`jason`的流量路由到v2版本，其他情况路由到v1版本。

```text
$ kubectl apply -f virtual-service-reviews-test-v2.yaml
```

执行命令后刷新页面，可以看到`reviews`都使用的v1版本，页面不会显示星星。点击右上角的`Sign in`按钮，以jason的身份登录（密码随便），可以看到`reviews`切换到v2版本了，页面显示了黑色星星。

查看`virtual-service-reviews-test-v2.yaml`文件内容可以看到，基于身份的路由是按照匹配HTTP的headers实现的。当HTTP的headers有`end-user: jason`的内容时路由到v2版本，否则路由到v1版本。

```text
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: reviews
spec:
  hosts:
    - reviews.service.consul
  http:
  - match:
    - headers:
        end-user:
          exact: jason
    route:
    - destination:
        host: reviews.service.consul
        subset: v2
  - route:
    - destination:
        host: reviews.service.consul
        subset: v1
```

### 几点注意事项的总结 <a id="&#x51E0;&#x70B9;&#x6CE8;&#x610F;&#x4E8B;&#x9879;&#x7684;&#x603B;&#x7ED3;"></a>

1. `istio.yaml`引用的Registrator的`latest`版本不支持consul的ServiceMeta。要改为`master`版本。
2. 第一次启动`istio.yaml`后，因为启动时`pilot`连不上`istio-apiserver`，`pilot`会失败退出。等待`istio-apiserver`启动完毕后再跑一次`istio.yaml`。
3. 配置`kubectl`的`context`，让`kubectl`使用`istio-apiserver`提供的Kubernetes API接口。
4. 使用`bookinfo.yaml`启动各个微服务后，还要运行`bookinfo.sidecar.yaml`以初始化和启动sidecar。

### 整体架构图 <a id="&#x6574;&#x4F53;&#x67B6;&#x6784;&#x56FE;"></a>

![](https://img2018.cnblogs.com/blog/576869/201904/576869-20190407233401806-538299451.png)

### 参考资料 <a id="&#x53C2;&#x8003;&#x8D44;&#x6599;"></a>

* [Istio文档](https://istio.io/zh/docs/)
* [Pattern: Service Mesh](http://philcalcado.com/2017/08/03/pattern_service_mesh.html)
* [Envoy 的架构与基本术语](https://jimmysong.io/posts/envoy-archiecture-and-terminology/)
* [Understanding How Envoy Sidecar Intercept and Route Traffic in Istio Service Mesh](https://medium.com/devopslinks/understanding-how-envoy-sidecar-intercept-and-route-traffic-in-istio-service-mesh-20fea2a78833)
* [Istio Pilot与Envoy的交互机制解读](https://blog.gmem.cc/interaction-between-istio-pilot-and-envoy)
* [Istio流量管理实现机制深度解析](http://www.servicemesher.com/blog/istio-traffic-management-impl-intro/)

