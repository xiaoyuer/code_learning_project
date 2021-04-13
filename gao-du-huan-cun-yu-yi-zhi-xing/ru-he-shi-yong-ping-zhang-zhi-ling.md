# 如何使用屏障指令

曾经我也想过一个问题，在现实中真的会遇到由于屏障缺失导致的BUG吗？而我们又该如何使用屏障呢？怎么判断是否应该使用屏障呢？当需要的时候，又该如何选择哪种屏障类型呢？这不，巧了，消失的半年中，我还真的fix了一个屏障缺失的BUG。正好借助这个例子实战一下吧。

#### \[w,r\]mb和smp\_\[w,r\]mb的选择

Linux内核中很容易简单两种类型的屏障指令的封装。smp\_\[w,r\]mb\(\)和\[w,r\]mb\(\)。代码中我们经常看到都是"smp\_"开头的屏障。当我们写代码时如何在这两者之间选择呢？我简单说下我的选择标准（可以保持怀疑态度）。如果内存操作的顺序的观察者全是CPU，那么请使用smp\_\[w,r\]mb\(\)。如果内存操作顺序涉及CPU和硬件设备（例如网络设备），请使用\[w,r\]mb\(\)。另外，\[w,r\]mb\(\)比smp\_\[w,r\]mb\(\)的要求更严格。因此性能方面也是\[w,r\]mb\(\)比smp\_\[w,r\]mb\(\)差。如果我们不写驱动的话，其实很少和设备打交道。因此一般都是选择smp\_\[w,r\]mb\(\)。

我们看一个wmb\(\)和rmb\(\)的使用例子。我们需要到设备驱动中寻找，就顺便选一个我也不清楚的网卡驱动吧（drivers/net/8139too.c）。

```text
static netdev_tx_t rtl8139_start_xmit (struct sk_buff *skb,
				       struct net_device *dev)
{
	/*
	 * Writing to TxStatus triggers a DMA transfer of the data
	 * copied to tp->tx_buf[entry] above. Use a memory barrier
	 * to make sure that the device sees the updated data.
	 */
	wmb();
	RTL_W32_F (TxStatus0 + (entry * sizeof (u32)),
		   tp->tx_flag | max(len, (unsigned int)ETH_ZLEN));
}
```

从这里的注释我们其实可以看出RTL\_W32\_F\(\)操作应该是发起一次设备DMA操作（通过写硬件寄存器实现）。wmb\(\)的作用是保证写入DMA发起操作命令之前写入内存的数据都已经写入内存，保证DMA操作时可以看见最新的数据。简单思考就是保证寄存器操作必须是最后发生的。我们继续找个rmb\(\)的例子（drivers/net/bnx2.c）。

```text
static int bnx2_rx_int(struct bnx2 *bp, struct bnx2_napi *bnapi, int budget)
{
	sw_cons = rxr->rx_cons;
	sw_prod = rxr->rx_prod;
	/* Memory barrier necessary as speculative reads of the rx
	 * buffer can be ahead of the index in the status block
	 */
	rmb();
	while (sw_cons != hw_cons) {
}
```

这里的rmb\(\)就是为了防止while循环里面的load操作在rmb\(\)之前发生。也就说这里是有顺序要求的。设备方面的例子就不多说了，毕竟不是我们关注的重点。

#### 编译器屏障barrier\(\)的选择

编译器屏障barrier\(\)不涉及任何硬件指令。是最弱的一种屏障，只对编译器有效。我们什么情况下应该选择barrier\(\)呢？

还记得之前的文章提到如何判断是否需要使用硬件屏障指令的原则吗？你写的这段代码（包含内存操作）在真正意义上存在并行执行，也就是可能存在不止一个CPU观察者。如果你能保证不可能同时存在多个CPU观察者（例如执行这段代码的进程都是绑定到一个CPU），实际上你就不需要硬件指令级别屏障。此时你需要考虑的问题只有：是否需要使用编译屏障。

```text
void queued_spin_lock_slowpath(struct qspinlock *lock, u32 val)
{
	node = this_cpu_ptr(&qnodes[0].mcs);
	idx = node->count++;

	/*
	 * Ensure that we increment the head node->count before initialising
	 * the actual node. If the compiler is kind enough to reorder these
	 * stores, then an IRQ could overwrite our assignments.
	 */
	barrier();

	node->locked = 0;
	node->next = NULL;
}
```

这里以qspinlock的代码为例说明。这里的node是一个per cpu变量。按照per cpu变量的使用规则，同一时间只会存在一个CPU写这块内存。因此这里不存在多个CPU观察者。因此不需要CPU指令级别的屏障。而这里同样有顺序的要求，因此编译器屏障barrier\(\)足以。

#### smp\_\[w,r\]mb\(\)的选择

现在回归到我们的重点关注对象了。当同时存在多个CPU观察者，并且我们又有内存操作顺序的要求时，我们就应该考虑smp\_\[w,r\]mb\(\)的使用。在之前的文章中我们提到下面代码示例用来说明TSO乱序。

```text
	initial: X = Y = 0

        CPU 1				CPU 2
	===============================	===============================
	X = 1;				Y = 1;
	smp_mb();			smp_mb();
	LOAD Y				LOAD X
```

由于smp\_mb\(\)的加持，这里CPU1和CPU2至少有一个CPU可以看到其中一个变量的值是1。不可能发生CPU1看到Y的值是0，同时CPU2看到X的值是0的情况。这里抽象的是一个模型，一个非常常见的模型。例如下面的代码你肯定经常遇见。

我们经常看到有个进程尝试等待一个变量event\_indicated为true，然后退出循环。否则继续schedule\(\)。

```text
	for (;;) {
		set_current_state(TASK_UNINTERRUPTIBLE);
		if (event_indicated)
			break;
		schedule();
	}
```

另一段负责唤醒的代码通常长下面这样。置位event\_indicated，然后唤醒进程。

```text
	event_indicated = 1;
	wake_up_process(event_daemon);
```

我们想达到的目的其实很简单：当event\_indicated为1时，进程一定是被唤醒的状态。我们把set\_current\_state\(\)以及wake\_up\_process\(\)部分细节展开如下：

```text
CPU 1 (Sleeper)			       CPU 2 (Waker)
===============================	       ===============================
current->state = TASK_UNINTERRUPTIBLE; event_indicate = 1
LOAD event_indicated		       if ((LOAD task->state) & TASK_NORMAL)
					    task->state = TASK_RUNNING
```

set\_current\_state的目的是设置当前进程状态为TASK\_UNINTERRUPTIBLE。然后判断event\_indicated是都满足条件。而唤醒的逻辑是先无条件设置event\_indicated。然后尝试唤醒进程，但是唤醒前会先判断被唤醒的进程的状态state是否为TASK\_UNINTERRUPTIBLE状态。如果不是TASK\_UNINTERRUPTIBLE，说明进程已经处于唤醒状态，无需再次唤醒。所以这里涉及2个数据，分别是current-&gt;state和event\_indicated。为了保证sleep进程被唤醒并退出while循环，我们必须保证要么CPU1看到CPU2设置的event\_indicate（然后退出循环），要么CPU2看到进程的state状态是TASK\_UNINTERRUPTIBLE（然后执行唤醒操作），要么CPU1即看到event\_indicate等于1并且CPU2看到进程的state状态是TASK\_UNINTERRUPTIBLE（唤醒已经running的进程没有问题）。但是我们不期望看到的结果是：CPU1看到event\_indicated的值是0，CPU看到进程的state状态不是TASK\_UNINTERRUPTIBLE。这种情况下，sleep进程就错过了被唤醒。所以这里的模型抽象出来就是上面的XY模型。

为了防止这种情况出现，我们需要插入2条smp\_mb\(\)。

```text
#define set_current_state(state_value)					\
	smp_store_mb(current->state, (state_value))

static int try_to_wake_up(struct task_struct *p, unsigned int state, int wake_flags)
{
	/*
	 * If we are going to wake up a thread waiting for CONDITION we
	 * need to ensure that CONDITION=1 done by the caller can not be
	 * reordered with p->state check below. This pairs with smp_store_mb()
	 * in set_current_state() that the waiting thread does.
	 */
	raw_spin_lock_irqsave(&p->pi_lock, flags);
	smp_mb__after_spinlock();
	if (!(p->state & state))
		goto unlock;
}
```

set\_current\_state\(\)的实现是smp\_store\_mb\(\)，相当于写完state后插入smp\_mb\(\)。wake\_up\_process\(\)最终会调用try\_to\_wake\_up\(\)。我们在读取进程state前也插入了smp\_mb\(\)（smp\_mb\_\_after\_spinlock）。

#### 实战分享

我们fix一个io\_uring的BUG。patch可以看考[Fix missing smp\_mb\(\) in io\_cancel\_async\_work\(\)](https://link.zhihu.com/?target=https%3A//patchwork.kernel.org/project/linux-block/patch/20201007031635.65295-3-songmuchun%40bytedance.com/)。

```text
diff --git a/fs/io_uring.c b/fs/io_uring.c
index 2f46def7f5832..5d9583e3d0d25 100644
--- a/fs/io_uring.c
+++ b/fs/io_uring.c
@@ -2252,6 +2252,12 @@  static void io_sq_wq_submit_work(struct work_struct *work)
 
 		if (!ret) {
 			req->work_task = current;
+
+			/*
+			 * Pairs with the smp_store_mb() (B) in
+			 * io_cancel_async_work().
+			 */
+			smp_mb(); /* A */
 			if (req->flags & REQ_F_CANCEL) {
 				ret = -ECANCELED;
 				goto end_req;
@@ -3730,7 +3736,15 @@  static void io_cancel_async_work(struct io_ring_ctx *ctx,
 
 		req = list_first_entry(&ctx->task_list, struct io_kiocb, task_list);
 		list_del_init(&req->task_list);
-		req->flags |= REQ_F_CANCEL;
+
+		/*
+		 * The below executes an smp_mb(), which matches with the
+		 * smp_mb() (A) in io_sq_wq_submit_work() such that either
+		 * we store REQ_F_CANCEL flag to req->flags or we see the
+		 * req->work_task setted in io_sq_wq_submit_work().
+		 */
+		smp_store_mb(req->flags, req->flags | REQ_F_CANCEL); /* B */
+
 		if (req->work_task && (!files || req->files == files))
 			send_sig(SIGINT, req->work_task, 1);
 	}
```

简单说一下这里的BUG，这里也同样设计2个变量，分别是req-&gt;flags和req-&gt;work\_task。模型抽象简化如下：

```text
        CPU 1				CPU 2
	===============================	===============================
	STORE req->flags		STORE req->work_task
	smp_mb();			smp_mb();
	LOAD req->work_task		LOAD req->flags
```

不期望看到的结果：req-&gt;work\_task == NULL && \(req-&gt;flags & REQ\_F\_CANCEL\) == 0。所以我们需要2个smp\_mb\(\)保证不出现这种结果。

BTW，[这个](https://link.zhihu.com/?target=https%3A//git.kernel.org/pub/scm/linux/kernel/git/axboe/linux-block.git/commit/fs/io_uring.c%3Fh%3Dfor-5.5/io_uring%26id%3Dc0e48f9dea9129aa11bec3ed13803bcc26e96e49)BUG也是这个模型，有兴趣也可以看看。

