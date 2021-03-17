# 内存一致性模型-RMO



目前为止我们已经接受了两种乱序类型，分别是store-store和store-load。还剩下最后2种操作组合，分别是load-load和load-store。我们对内存模型已经有了一定的认知，所以按照常理来说，理应有一种内存模型允许4种操作都可以乱序。这就是Relaxed Memory Order，简称RMO。例如aarch64就是典型的RMO模型。

> 文章测试代码已经开源，可以点击[这里](https://link.zhihu.com/?target=https%3A//github.com/smcdef/memory-reordering)查看。

### 不想提及的乱序原因

总是有些朋友希望探究真理或者硬件的细节，希望知道乱序导致的原因。但是这些真的是硬件细节的东西，有些硬件细节我们是查不到的。包括芯片的数据手册，因为这些硬件相关，和软件无关。而手册更多的是阐述软件应该注意的事项，应该注意什么。而不是解释硬件如何实现的。所以有些东西或许不是很值得深究。深究下去或许反而会影响自己的理解。我们作为软件从业者，更应该的是从软件的角度思考，内存模型的提出正是这一目的。我们需要知道的是某一种内存模型允许的乱序类型，而不是关注为什么引起乱序（或者从硬件的角度思考为什么会这样）。导致乱序的原因很多，并不是单一的硬件细节。我从aarch64的手册找到如下可能导致乱序的原因。但是这些原因也没用提及硬件设计如何导致的。所以算是了解就好。

* Multiple issue of instructions
* Out-of-order execution
* Speculation
* Speculative loads
* Load and store optimizations
* External memory systems
* Cache coherent multi-core processing
* Optimizing compilers

可以看出硬件做出了足够多的优化，这一切综合的结果。而我们之前提及store buffer也是硬件细节实现之一，为什么我会提到store buffer了？我们不是说不关注硬件细节吗？确实不应该关注硬件细节，但是知道store buffer的存在很重要。后续会有文章提到store buffer是如何影响驱动代码编写的。不仅仅是乱序这点，当然这是后话。

### 旧事重提

还记得PSO文章提及的例子吗？当我们以RMO的模型考虑问题是，事情就不再是那么简单了。当我们在write\_data\_cpu0\(\)加入smp\_wmb\(\)后还不足以解决问题。我还需要在读取测加入smp\_rmb\(\)屏障。因此完整修改如下。

```text
static int flag, data;


static void write_data_cpu0(int value)
{
        data = value;
        smp_wmb();
        flag = 1;
}

static void read_data_cpu1(void)
{
        while (flag == 0);
        smp_rmb();
        return data;
}
```

在之前的文章中，我们已经见识了smp\_mb\(\)和smp\_wmb\(\)的作用，这里引入一个新的读屏障smp\_rmb\(\)。smp\_rmb\(\)保证屏障后的读操作不会乱序到屏障前的读操作，只针对读操作，不影响写操作。

```text
store--------------------------+
load----------+                |
              |                |   ^
              v                |   |
------------smp_rmb()----------|---|------
              ^                |   |
              |                v   |
load----------+                    |
store------------------------------+
```

### 看得见的乱序

我们以aarch64测试乱序。我们可以使用常见的手机或者aarch64服务器测试。我们使用3个CPU测试，CPU0负责更新计数。

```text
static atomic_t count = ATOMIC_INIT(0);


static void ordering_thread_fn_cpu0(void)
{
        atomic_inc(&count);
}
```

CPU1负责顺序写入a和b，a和b会被顺序赋值为count。假设不溢出的情况下，**理论上a一定大于等于b**。

```text
static unsigned int a, b;


static void ordering_thread_fn_cpu1(void)
{
        int temp = atomic_read(&count);

        a = temp;
        /* Prevent compiler reordering. */
        barrier();
        b = temp;
}
```

现在引出我们观察者CPU2，CPU2负责顺序读取b和a。理论上来说，**d一定小于等于c**。当d大于c时，就说明我们观察到了乱序。

```text
static void ordering_thread_fn_cpu2(void)
{
        unsigned int c, d;

        d = b;
        /* Prevent compiler reordering. */
        barrier();
        c = a;

        if ((int)(d - c) > 0)
                pr_info("reorders detected, a = %d, b = %d\n", c, d);
}
```

我们该如何理解这段这个结果呢？我们知道RMO模型允许store-store乱序。因此，CPU2看到CPU1的执行顺序可能是先写b再写a。这是可能造成结果的原因之一。除此以外，RMO模型同样允许load-load乱序，因此CPU2读取b的c的值可能也是反过来的。这是造成结果的第二个可能原因。

### 如何保证顺序

以上的例子中，如果我们不希望看到d比c大的情况出现（如果不介意出现的话，就无所谓了），我们该如何做呢？为了防止CPU1出现store-store乱序，我们可以插入smp\_wmb\(\)。同样，为了防止CPU2出现load-load乱序，我们还应该插入smp\_rmb\(\)。

```text
static void ordering_thread_fn_cpu1(void)
{
        int temp = atomic_read(&count);


        a = temp;
        smp_wmb();
        b = temp;
}


static void ordering_thread_fn_cpu2(void)
{
        unsigned int c, d;


        d = b;
        smp_rmb();
        c = a;


        if ((int)(d - c) > 0)
                pr_info("reorders detected, a = %d, b = %d\n", c, d);
}
```

经过如此修改后，就不会出现d比c大的情况了。

### 总结

1. 如上面所述，我们知道硬件做了足够的优化可能导致乱序。在单核系统下，这种乱序的影响对程序员是透明的。例如上述的ordering\_thread\_fn\_cpu2\(\)。CPU2来说，a和b先load哪个变量，其实是无所谓的。因为CPU并不知道a和b之间的这种特殊依赖性。CPU认为是a和b是没有任何关系的。
2. 如果访问的数据不存在竞争，根本不用考虑内存乱序。这是大多数局变量的场景。
3. 如果访问数据存在竞争，但是可以保证该数据竞争只会出现在单核上（例如竞争访问的线程都是绑定一个CPU上）。同样不需要考虑CPU乱序。这就是硬件最基本的保证，对单核透明。
4. 写任何一个通用的程序我们都应该假设CPU类型是RMO模型。因为你不知道什么时候这段代码就可能在RMO模型上运行。
5. 内存屏障的使用应该成对，例如smp\_wmb必须配对smp\_rmb或者smp\_mb，单独使用smp\_wmb是达不到顺序效果的。同样smp\_rmb必须配对使用smp\_wmb或者smp\_mb。

