# 编译乱序

经过前面三篇文章的介绍。我们应该对CPU乱序有了一定的认知。但是CPU乱序只是乱序的一部分，另一部分就是编译器乱序。我们之前提到的smp\_wmb/smp\_rmb/smp\_mb都是CPU内存屏障指令的封装，目的是防止CPU乱序。而编译器乱序是编译器优化代码的结果。我们同样需要编译器屏障阻止编译器优化导致的乱序。

编译器（compiler）的工作之一是优化我们的代码以提高性能。这包括在不改变程序行为的情况下重新排列指令。因为compiler不知道什么样的代码需要线程安全（thread-safe），所以compiler假设我们的代码都是单线程执行（single-threaded），并且进行指令重排优化并保证是单线程安全的。因此，当你不需要compiler重新排序指令的时候，你需要显式告编译器，我不需要重排。否则，它可不会听你的。本篇文章中，我们一起探究compiler关于指令重排的优化规则。

> 注：测试使用aarch64-linux-gnu-gcc编译器，版本：7.3.0

### 编译器指令重排

compiler的主要工作就是将对人们可读的源码转化成机器语言，机器语言就是对CPU可读的代码。因此，compiler可以在背后做些不为人知的事情。我们考虑下面的C语言代码：

```text
int a, b;

void foo(void)
{
        a = b + 1;
        b = 0;
}
```

使用aarch64-linux-gnu-gcc在不优化代码的情况下编译上述代码，使用objdump工具查看foo\(\)反汇编结果：

```text
foo:
        adrp    x0, b
        add     x0, x0, :lo12:b
        ldr     w0, [x0]          # w0 = b
        add     w1, w0, 1         # w1 = w0 + 1 = b + 1
        adrp    x0, a
        add     x0, x0, :lo12:a
        str     w1, [x0]          # a = b + 1
        adrp    x0, b
        add     x0, x0, :lo12:b
        str     wzr, [x0]         # b = 0
        nop
        ret
```

我们应该知道Linux默认编译优化选项是-O2（默认情况下O0编译kernel也编译不过）。因此我们采用-O2优化选项编译上述代码，并反汇编得到如下汇编结果：

```text
foo:
        adrp    x1, b
        adrp    x2, a
        ldr     w0, [x1, #:lo12:b]    # w0 = b
        str     wzr, [x1, #:lo12:b]   # b = 0
        add     w0, w0, 1             # w0 = w0 + 1 = b + 1
        str     w0, [x2, #:lo12:a]    # a = w0 = b + 1
        ret
```

比较优化和不优化的结果，我们可以发现。在不优化的情况下，a 和 b 的写入内存顺序符合代码顺序（program order）。但是-O2优化后，a 和 b 的写入顺序和program order是相反的。-O2优化后的代码转换成C语言可以看作如下形式：

```text
int a, b;

void foo(void)
{
        register int reg = b;

        b = 0;
        a = reg + 1;
}
```

这就是compiler reordering（编译器重排）。为什么可以这么做呢？对于单线程来说，a 和 b 的写入顺序，compiler认为没有任何问题。并且最终的结果也是正确的（a == 1 && b == 0），因此编译器才会这么做。

这种compiler reordering在大部分情况下是没有问题的。但是在某些情况下可能会引入问题。例如我们使用一个全局变量`flag`标记共享数据`data`是否就绪。由于compiler reordering，可能会引入问题。考虑下面的代码（无锁编程）：

```text
int flag, data;

void write_data(int value)
{
        data = value;
        flag = 1;
}

void read_data(void)
{
        int res;

        while (flag == 0);
        res = data;
        flag = 0;
        return res;
}
```

先来简单的介绍代码功能。我们拥有2个线程，一个用来更新数据，也就是更新data的值。使用flag标志data数据已经准备就绪，其他线程可以读取。另一个线程一直调用read\_data\(\)，等待flag被置位，然后返回读取的数据data。

如果compiler产生的汇编代码是flag比data先写入内存。那么，即使是单核系统上，我们也会有问题。在flag置1之后，data写45之前，系统发生抢占。另一个进程发现flag已经置1，认为data的数据已经准别就绪。但是实际上读取data的值并不是45（可能是上次的历史数据或者非法数据）。为什么compiler还会这么操作呢？因为，compiler是不知道data和flag之间有严格的依赖关系。这种逻辑关系是我们人为强加的。那么我们如何避免这种优化呢？

### 显式编译器屏障

为了解决上述变量之间存在依赖关系导致compiler错误优化。compiler为我们提供了编译器屏障（compiler barriers），可用来告诉compiler不要reorder。我们继续使用上面的foo\(\)函数作为演示实验，在代码之间插入compiler barriers。

```text
#define barrier() __asm__ __volatile__("": : :"memory")

int a, b;

void foo(void)
{
        a = b + 1;
        barrier();
        b = 0;
}
```

barrier\(\)就是compiler提供的屏障，作用是告诉compiler内存中的值已经改变，之前对内存的缓存（缓存到寄存器）都需要抛弃，barrier\(\)之后的内存操作需要重新从内存load，而不能使用之前寄存器缓存的值。并且可以防止compiler优化barrier\(\)前后的内存访问顺序。barrier\(\)就像是代码中的一道不可逾越的屏障，barrier前的 load/store 操作不能跑到barrier后面；同样，barrier后面的 load/store 操作不能在barrier之前。依然使用-O2优化选项编译上述代码，反汇编得到如下结果：

```text
foo:
        adrp    x1, b
        adrp    x2, a
        ldr     w0, [x1, #:lo12:b]    # w0 = b
        add     w0, w0, 1             # w0 = w0 + 1
        str     w0, [x2, #:lo12:a]    # a = b + 1
        str     wzr, [x1, #:lo12:b]   # b = 0
        ret
```

我们可以看到插入compiler barriers之后，a 和 b 的写入顺序和program order一致。因此，当我们的代码中需要严格的内存顺序，就需要考虑compiler barriers。

### 隐式编译器屏障

除了显示的插入compiler barriers之外，还有别的方法阻止compiler reordering。例如CPU barriers 指令，同样会阻止compiler reordering。后续我们再考虑CPU barriers。

除此以外，当某个函数内部包含compiler barriers时，该函数也会充当compiler barriers的作用。即使这个函数被inline，也是这样。例如上面插入barrier\(\)的foo\(\)函数，当其他函数调用foo\(\)时，foo\(\)就相当于compiler barriers。考虑下面的代码：

```text
int a, b, c;

void fun(void)
{
        c = 2;
        barrier();
}

void foo(void)
{
        a = b + 1;
        fun();      /* fun() call act as compiler barriers */
        b = 0;
}
```

fun\(\)函数包含barrier\(\)，因此foo\(\)函数中fun\(\)调用也表现出compiler barriers的作用。同样可以保证 a 和 b 的写入顺序。如果fun\(\)函数不包含barrier\(\)，结果又会怎么样呢？实际上，大多数的函数调用都表现出compiler barriers的作用。但是，这不包含inline的函数。因此，fun\(\)如果被inline进foo\(\)，那么fun\(\)就不会具有compiler barriers的作用。如果被调用的函数是一个外部函数，其副作用会比compiler barriers还要强。因为compiler不知道函数的副作用是什么。它必须忘记它对内存所作的任何假设，即使这些假设对该函数可能是可见的。我们看一下下面的代码片段，printf\(\)一定是一个外部的函数。

```text
int a, b;

void foo(void)
{
        a = 5;
        printf("smcdef");
        b = a;
}
```

同样使用-O2优化选项编译代码，objdump反汇编得到如下结果。

```text
foo:
        stp     x29, x30, [sp, -32]!
        mov     w1, 5
        adrp    x0, .LC0
        mov     x29, sp
        str     x19, [sp, 16]
        adrp    x19, a
        add     x0, x0, :lo12:.LC0
        str     w1, [x19, #:lo12:a]    # a = 5
        bl      printf                 # printf("smcdef")
        adrp    x0, b
        ldr     w1, [x19, #:lo12:a]    # w1 = a
        ldr     x19, [sp, 16]
        str     w1, [x0, #:lo12:b]     # b = w1 = a
        ldp     x29, x30, [sp], 32
        ret
.LC0:
        .string "smcdef"
```

compiler不能假设printf\(\)不会使用或者修改 a 变量。因此在调用printf\(\)之前会将 a 写5，以保证printf\(\)可能会用到新值。在printf\(\)调用之后，重新从内存中load a 的值，然后赋值给变量 b。重新load a 的原因是compiler也不知道printf\(\)会不会修改 a 的值。

因此，我们可以看到即使存在compiler reordering，但是还是有很多限制。当我们需要考虑compiler barriers时，一定要显示的插入barrier\(\)，而不是依靠函数调用附加的隐式compiler barriers。因为，谁也无法保证调用的函数不会被compiler优化成inline方式。

### barrier\(\)除了防止编译乱序，还能做什么

barriers\(\)作用除了防止compiler reordering之外，还有什么妙用吗？我们考虑下面的代码片段。

```text
int run = 1;

void foo(void)
{
        while (run)
                ;
}
```

run是个全局变量，foo\(\)在一个进程中执行，一直循环。我们期望的结果是foo\(\)一直等到其他进程修改run的值为0才推出循环。实际compiler编译的代码和我们会达到我们预期的结果吗？我们看一下汇编代码。

```text
foo:
        adrp    x0, .LANCHOR0
        ldr     w0, [x0, #:lo12:.LANCHOR0]  # w0 = run
.L2:
        cbnz    w0, .L2                     # if (w0) while (1);
        ret
run:
        .word   1
```

汇编代码可以转换成如下的C语言形式。

```text
int run = 1;

void foo(void)
{
        register int reg = run;

        if (reg)
                while (1)
                        ;
}
```

compiler首先将run加载到一个寄存器reg中，然后判断reg是否满足循环条件，如果满足就一直循环。但是循环过程中，寄存器reg的值并没有变化。因此，即使其他进程修改run的值为0，也不能使foo\(\)退出循环。很明显，这不是我们想要的结果。我们继续看一下加入barrier\(\)后的结果。

```text
#define barrier() __asm__ __volatile__("": : :"memory")

int run = 1;

void foo(void)
{
        while (run)
                barrier();
}
```

其对应的汇编结果：

```text
foo:
        adrp    x1, .LANCHOR0
        ldr     w0, [x1, #:lo12:.LANCHOR0]
        cbz     w0, .L1
        add     x1, x1, :lo12:.LANCHOR0
.L3:
        ldr     w0, [x1]    # w0 = run
        cbnz    w0, .L3     # if (w0) goto .L3
.L1:
        ret
run:
        .word   1
```

我们可以看到加入barrier\(\)后的结果真是我们想要的。每一次循环都会从内存中重新load run的值。因此，当有其他进程修改run的值为0的时候，foo\(\)可以正常退出循环。为什么加入barrier\(\)后的汇编代码就是正确的呢？因为barrier\(\)作用是告诉compiler内存中的值已经变化，后面的操作都需要重新从内存load，而不能使用寄存器缓存的值。因此，这里的run变量会从内存重新load，然后判断循环条件。这样，其他进程修改run变量，foo\(\)就可以看得见。

在Linux kernel中，提供了[cpu\_relax\(\)](https://link.zhihu.com/?target=https%3A//elixir.bootlin.com/linux/v4.15.7/source/arch/arm64/include/asm/processor.h%23L178)函数，该函数在ARM64平台定义如下：

```text
static inline void cpu_relax(void)
{
       asm volatile("yield" ::: "memory");
}
```

我们可以看出，cpu\_relax\(\)是在barrier\(\)的基础上又插入一条汇编指令yield。在linux kernel中，我们经常会看到一些类似上面举例的while循环，循环条件是个共享变量。为了避免上述所说问题，我们就会在循环中插入cpu\_relax\(\)调用。

```text
int run = 1;

void foo(void)
{
        while (run)
                cpu_relax();
}
```

当然也可以使用Linux 提供的READ\_ONCE\(\)。例如，下面的修改也同样可以达到我们预期的效果。

```text
int run = 1;

void foo(void)
{
        while (READ_ONCE(run))  /* similar to while (*(volatile int *)&run) */
                ;
}
```

当然你也可以修改run的定义为`volatile int run;`，就会得到如下代码，同样可以达到预期目的。

```text
volatile int run = 1;

void foo(void)
{
        while (run)
                ;
}
```

更多关于volatile使用建议可以参考下面的文章。[smcdef：为什么我们不应该使用volatile类型​zhuanlan.zhihu.com![&#x56FE;&#x6807;](https://pic4.zhimg.com/v2-ddb798c8abc39266198e218dd94c836f_180x120.jpg)](https://zhuanlan.zhihu.com/p/102406978)

### CPU乱序和编译器乱序的关系

smp\_wmb/smp\_rmb/smp\_mb等都是防止CPU乱序的指令封装。是不是意味着这些接口仅仅阻止CPU乱序，而允许编译器乱序呢？答案肯定是不可能的。这里有个点需要记住，**所有的CPU内存屏障封装都隐世包含了编译器屏障**。

