# 内存一致性模型-PSO

TSO内存一致性模型我们已经有所了解。在TSO模型中，我们说过store buffer会按照FIFO的顺序将数据写入L1 Cache。例如我们执行如下的代码：

```text
x = 10;
y = 1;
```

我们知道按照TSO模型的执行思路，首先x=10的值会写入store buffer，同样y=1的值也会写入store buffer。然后按照FIFO的次序写入L1 Cache。假设y变量已经缓存在L1 Cache。但是x变量没有缓存。理论上来说我们可以先将y写入Cache，然后将x写入Cache。因为x和y操作看起来没有任何关系，这么做看起来没什么影响。如果这样做的话，还会让store buffer腾出多余的空间以缓存后续的CPU操作。这在一定程度又可以提升性能。此时，store buffer不再以FIFO的次序写入Cache。这种硬件优化会对程序的运行又会造成什么影响呢？

### 问题引入

我们假设flag和data的初始值是0。CPU0执行`write_data_cpu0()`，更新数据`data`，以`flag`变量标识数据更新完成。CPU1执行`read_data_cpu1()`等待数据更新完成，然后读取数据`data`。我们来分析这个过程。

```text
int flag, data;


static void write_data_cpu0(int value)
{
        data = value;
        flag = 1;
}


static void read_data_cpu1(void)
{
        while (flag == 0);
        return data;
}
```

1. CPU0将data更新的值写入store buffer。
2. CPU0将flag的值写入store buffer。
3. CPU1读取flag的值，由于flag新值还在CPU0的store buffer里面，所以CPU1看到的flag值依然是0。所以CPU1继续自旋等待。
4. CPU0的store buffer将flag的值写入Cache，此时data的值依然在store buffer里面没有更新到Cache。
5. CPU1发现flag的值是1，退出循环。然后读取data的值，此时data的值是旧的数据，也就是0。因为data的新值还在CPU0的store buffer里面。
6. CPU0的store buffer将data的值写入Cache。

我们已经看到了不想要的结果。按照TSO内存模型分析这个例子的话，我们是一定能够保证CPU1读取到新值。因为TSO模型中，store buffer是按照FIFO的次序写入Cache。但是由于现在store buffer不以FIFO的次序更新Cache，所以导致CPU1读取到data的旧值0。对于CPU来说是不知道`data`和`flag`之间是有依赖关系的，CPU认为`data`和`flag`的值谁先写入Cache都无所谓。所以做出了这种优化。我们如何换个思路理解这个现象呢？

PSO模型

我们知道TSO模型可以看成store-load操作可以乱序成load-store。现在，我们提出了一种新的内存模型，部分存储定序\(Part Store Order\)，简称PSO。PSO在TSO的基础上继续允许store-store乱序。上面的例子中CPU1看到CPU0对内存的操作顺序可能是下面这样的。

```text
flag = 1;
data = value;
```

还是那句话，CPU0看自己的内存操作是顺序的，只是CPU1看CPU0内存操作可能是乱序的。现在我们可以忘记store buffer的存在，从内存模型的角度考虑这个程序。data和flag操作都是store，所以按照PSO模型来说，其他CPU是可以看到乱序操作的结果。这样就很容易理解了。尽量不去思考硬件为什么会这样，而是认为这是PSO模型理所应当存在的现象。

### 如何保证顺序一致

Linux内核中提供了`smp_wmb()`宏对不同架构的指令进行封装，`smp_wmb()`的作用是阻止它后面的写操作不会乱序到宏前面的写操作指令前执行。它就像是个屏障，写操作不容逾越。`smp_wmb()`充当的就是一个store-store屏障。但是这个屏障只针对store操作，对load操作不影响。

```text
store---------+
              |
              v
-------store-store barrier--------
              ^
              |
store---------+
```

如果我们需要上述的示例代码在PSO模型的处理器上正确运行\(按照我们期望的结果运行\)，就需要做出如下修改。

```text
static void write_data_cpu0(int value)
{
        data = value;
        smp_wmb();
        flag = 1;
}
```

`smp_wmb()`可以保证CPU1看到flag的值为1时，data的值一定是最新的值。

### 总结

当使用PSO模型时，我们知道PSO模型允许两种操作乱序。分别是store-store和store-load。我们需要根据需求选择使用`smp_mb()`或者`smp_wmb()`保证顺序一致性。

