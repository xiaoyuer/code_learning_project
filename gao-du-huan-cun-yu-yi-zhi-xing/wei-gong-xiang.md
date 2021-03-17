# 伪共享

我们知道kernel地址空间是所有进程共享的，所以kernel空间的全局变量，任何进程都可以访问。假设有2个全局变量global\_A和global\_B\(类型是long\)，它们在内存上紧挨在一起，假设cache line size是64Bytes，并且global\_A是cache line size对齐。所以global\_A和global\_B如果同时load到Cache中，一定是落在同一行cache line。就像下面这样。![](https://pic3.zhimg.com/80/v2-59f86dc04a371b4ed975e88605e2555a_1440w.png)

现在我们知道多核Cache一致性由MESI协议保证。有了这些基础之后，我们现在来思考一个问题，如果我们的系统有2个CPU，每个CPU上运行完全不相干的两个进程task\_A和task\_B。task\_A只会修改global\_A变量，task\_B只会修改global\_B变量。会有什么问题吗？

### 我们遇到什么问题

最初全局变量global\_A和global\_B都不在cache中缓存，如下图示意。task\_A绑定CPU0运行，task\_B绑定CPU1运行。task\_A和task\_B按照下面的次序分别修改或读取全局变量global\_A和global\_B。![](https://pic1.zhimg.com/80/v2-70d3eab08c3d7b69a9c801256d370ba0_1440w.jpg)

a\) CPU0读取global\_A，global\_A的数据被缓存到CPU0的私有L1 Cache。由于Cache控制器是以cache line为单位从内存读取数据，所以顺便就会把global\_B变量也缓存到Cache。并将cache line置为Exclusive状态。![](https://pic2.zhimg.com/80/v2-4a18fa78dc5aaf06d30a04416efe6441_1440w.jpg)

b\) CPU1读取global\_B变量，由于global\_B被CPU0私有Cache缓存，所以CPU0的L1 Cache负责返回global\_B数据到CPU1的L1 Cache。同样global\_A也被缓存。此时CPU0和CPU1的cache line状态变成Shared状态。![](https://pic4.zhimg.com/80/v2-d405954b8647d44998db2479522e17e7_1440w.jpg)

c\) CPU0现在需要修改global\_A变量。CPU0发现cache line状态是Shared，所以需要发送invalid消息给CPU1。CPU1将global\_A对应的cache line无效。然后CPU0的cache line状态变成Modified并且修改global\_A。![](https://pic2.zhimg.com/80/v2-4a18fa78dc5aaf06d30a04416efe6441_1440w.jpg)

d\) CPU1现在需要修改global\_B变量。此时global\_B变量并没有缓存在CPU1私有Cache。所以CPU1会发消息给CPU0，CPU0将global\_B数据返回给CPU1。并且会invalid CPU0的cache line。然后global\_B对应的CPU1 cache line变成Modified状态，此时CPU1就可以修改global\_B了。![](https://pic2.zhimg.com/80/v2-60440e07d78b6307a2c81ac4288c17a9_1440w.jpg)

如果CPU0和CPU1就这样持续交替的分别修改全局变量global\_A和global\_B，就会重复c\)和d\)。意识到问题所在了吗？这就是典型的cache颠簸问题。我们仔细想想，global\_A和global\_B其实并没有任何的关系，却由于落在同一行cache line的原因导致cache颠簸。我们称这种现象为伪共享\(false sharing\)。global\_A和global\_B之间就是伪共享关系，实际并没有共享。我们如何解决伪共享问题呢？

### 如何解决伪共享

既然global\_A和global\_B由于在同一行cache line导致了伪共享问题，那么解决方案也很显然易见，我们可以让global\_A和global\_B不落在一个cache line，这样就可以解决问题。不落在同一行cache line的方法很简单，使global\_A和global\_B的内存地址都按照cache line size对齐，相当于以空间换时间。浪费一部分内存，换来了性能的提升。当我们把global\_A和global\_B都cache line size对齐后，我们再思考上面的问题。此时CPU0和CPU1分别修改global\_A和global\_B互不影响。global\_A和global\_B对应的cache line状态可以一直维持Modified状态。这样MESI协议就不会在两个CPU间不停的发送消息。降低了带宽压力。

### 实际应用

在Linux kernel中存在\_\_cacheline\_aligned\_in\_smp宏定义用于解决false sharing问题。

```text
#ifdef CONFIG_SMP
#define __cacheline_aligned_in_smp __cacheline_aligned
#else
#define __cacheline_aligned_in_smp
#endif
```

我们可以看到在UP\(单核\)系统上，宏定义为空。在MP\(多核\)系统下，该宏是L1 cach line size。针对静态定义的全局变量，如果在多核之间竞争比较严重，为了避免影响其他全局变量，可以采用上面的宏使变量cache line对齐，避免false sharing问题。

