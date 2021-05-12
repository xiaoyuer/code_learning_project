# Concurrent programming

## Goroutines are lightweight threads

The **go** statement runs a func­tion in a sepa­rate thread of execu­tion.

You can start a new thread of execution, a goroutine, with the `go` statement. It runs a function in a different, newly created, goroutine. All goroutines in a single program share the same address space.

```text
go list.Sort() // Run list.Sort in parallel; don’t wait for it.
```

The following program will print “Hello from main goroutine”. It _might_ also print “Hello from another goroutine”, depending on which of the two goroutines finish first.

```text
func main() {
    go fmt.Println("Hello from another goroutine")
    fmt.Println("Hello from main goroutine")

    // At this point the program execution stops and all
    // active goroutines are killed.
}
```

The next program will, most likely, print both “Hello from main goroutine” and “Hello from another goroutine”. They may be printed in any order. Yet another possibility is that the second goroutine is extremely slow and doesn’t print its message before the program ends.

```text
func main() {
    go fmt.Println("Hello from another goroutine")
    fmt.Println("Hello from main goroutine")

    time.Sleep(time.Second) // give the other goroutine time to finish
}
```

Here is a somewhat more realistic example, where we define a function that uses concurrency to postpone an event.

```text
// Publish prints text to stdout after the given time has expired.
// It doesn’t block but returns right away.
func Publish(text string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println("BREAKING NEWS:", text)
    }() // Note the parentheses. We must call the anonymous function.
}
```

This is how you might use the `Publish` function.

```text
func main() {
    Publish("A goroutine starts a new thread.", 5*time.Second)
    fmt.Println("Let’s hope the news will published before I leave.")

    // Wait for the news to be published.
    time.Sleep(10 * time.Second)

    fmt.Println("Ten seconds later: I’m leaving now.")
}
```

The program will, most likely, print the following three lines, in the given order and with a five second break in between each line.

```text
$ go run publish1.go
Let’s hope the news will published before I leave.
BREAKING NEWS: A goroutine starts a new thread.
Ten seconds later: I’m leaving now.
```

In general it’s not possible to arrange for threads to wait for each other by sleeping. Go’s main method for synchronization is to use [channels](https://yourbasic.org/golang/channels-explained/).

### Implementation <a id="implementation"></a>

Goroutines are lightweight, costing little more than the allocation of stack space. The stacks start small and grow by allocating and freeing heap storage as required.

Internally goroutines act like coroutines that are multiplexed among multiple operating system threads. _**If one goroutine blocks an OS thread, for example waiting for input, other goroutines in this thread will migrate so that they may continue running.**_

## Channels offer synchronized communication

A channel is a mechanism for goroutines to **synchronize execution** and **communicate** by passing values.



A new channel value can be made using the built-in function `make`.

```text
// unbuffered channel of ints
ic := make(chan int)

// buffered channel with room for 10 strings
sc := make(chan string, 10)
```

**To send** a value on a channel, use `<-` as a binary operator. **To receive** a value on a channel, use it as a unary operator.

```text
ic <- 3   // Send 3 on the channel.
n := <-sc // Receive a string from the channel.
```

The `<-` operator specifies the channel direction, **send** or **receive**. If no direction is given, the channel is **bi-directional**.

```text
chan Sushi    // can be used to send and receive values of type Sushi
chan<- string // can only be used to send strings
<-chan int    // can only be used to receive ints
```

### Buffered and unbuffered channels <a id="buffered-and-unbuffered-channels"></a>

* If the capacity of a channel is zero or absent, the channel is **unbuffered** and the sender blocks until the receiver has received the value.
* If the channel **has a buffer**, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.
* Receivers always block until there is data to receive.
* Sending or receiving from a `nil` channel blocks forever.

### Closing a channel <a id="closing-a-channel"></a>

The `close` function records that no more values will be sent on a channel. Note that it is only necessary to close a channel if a receiver is looking for a close.

* After calling `close`, and after any previously sent values have been received, receive operations will return a [zero value](https://yourbasic.org/golang/default-zero-value/) without blocking.
* A multi-valued receive operation additionally returns an indication of whether the channel is closed.
* Sending to or closing a closed channel causes a run-time panic. Closing a nil channel also causes a run-time panic.

```text
ch := make(chan string)
go func() {
    ch <- "Hello!"
    close(ch)//Go program ends when the main function ends.
}()

fmt.Println(<-ch) // Print "Hello!".
fmt.Println(<-ch) // Print the zero value "" without blocking.
fmt.Println(<-ch) // Once again print "".
v, ok := <-ch     // v is "", ok is false.

// Receive values from ch until closed.
for v := range ch {
    fmt.Println(v) // Will not be executed.
}
```

### Example <a id="example"></a>

In the following example we let the `Publish` function return a channel, which is used to broadcast a message when the text has been published.

```text
// Publish prints text to stdout after the given time has expired.
// It closes the wait channel when the text has been published.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(delay)
		fmt.Println(text)
		close(ch)
	}()
	return ch
}
```

Note that we use a channel of empty structs to indicate that the channel will only be used for signalling, not for passing data. This is how you might use the function.

```text
wait := Publish("important news", 2 * time.Minute)
// Do some more work.
<-wait // Block until the text has been published.
```

通信是同步且无缓冲的。这种特性导致通道的发送/接收操作，在对方准备好之前是阻塞的。

对于同一个通道，发送操作在接收者准备好之前是阻塞的。如果通道中的数据无人接收，就无法再给通道传入其他数据。新的输入无法在通道非空的情况下传入，所以发送操作会等待channel再次变为可用状态，即通道值被接收后。 对于同一个通道，接收操作是阻塞的，直到发送者可用。如果通道中没有数据，接收者就阻塞了。

但是根据我们之前讲的，对于同一无缓冲通道，在接收者未准备好之前，发送操作是阻塞的。而此处的通道`ch1`就是缺少一个配对的接收者，因此造成了死锁。  
解决上面问题的方式有两种：第一种添加配对的接收者；第二种将默认的通道替换成缓冲通道。

```text
package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch1 <- "hello world"
	fmt.Println(<-ch1)
}
```

```text
//方法一
//在主函数中启用了一个goroutine，匿名函数用来发送数据，而在main()函数中接收通道中的数据。
package main

import "fmt"

func main() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "hello world"
	}()
	fmt.Println(<-ch1)
}
```

```text
package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch1 <- "hello world"
	fmt.Println(<-ch1)
}
```

此时的ch1通道可以称为缓冲通道，在缓冲满载\(缓冲被全部使用\)之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。定义方法如：ch:=make\(chan type, value\)。 这个value表示缓冲容量，它的大小和类型无关，所以可以给一些通道设置不同的容量，只要它们拥有相同的元素类型。 内置的cap函数可以返回缓冲区的容量，如果容量大于0，通道就是异步的了。缓冲满载\(发送\)或变空\(接收\)之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功。

这种异步channel可以减少排队阻塞，在你的请求激增的时候表现得更好，更具伸缩性。

如果发送多个值，我们如何接收。例如：

```text
package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- "hello world"
	ch1 <- "hello China"
}
```

```text
hello world
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
    D:/gotest/main.go:13 +0x97
```

出现这样的结果是因为通道实际上是类型化消息的队列，它是先进先出\(FIFO\)的结构，可以保证发送给它们的元素的顺序。所以上面代码只取出了第一次传的值，即"hello world"，而第二次传入的值没有一个配对的接收者来接收，因此就出现了deadlock。那么将代码变成这样，又会是什么结果呢？

```text
package main

import "fmt"

func main() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "hello world"
		ch1 <- "hello China"
	}()
	fmt.Println(<-ch1)
}
```

原因分析 信号量模式：协程通过在通道中放置一个值来处理结束的信号，main协程等待，直到从通道中获取到值。 上面的程序中有两个函数：main\(\)函数和一个发送操作的匿名函数。它们按独立的处理单元按顺序启动，然后开始并行运行。通常情况下，由于main\(\)函数不会等待其他非main协程的结束。但是此处的ch1相当于信号量，通过在ch1中放置一个值来处理结束的信号。main\(\)协程等待&lt;-ch，直到从中获取到值，然后程序直接退出。根本没有执行到继续往通道中传入"hello China"，也就不会出现deadlock的出现。

#### 示例三

```text
package main

import "fmt"

func main() {
	ch1 := make(chan string)
	go func() {
		ch1 <- "hello world"
		ch1 <- "hello China"
	}()
	for {
		fmt.Println(<-ch1)
	}
}
```

```text
hello world
hello China
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
    D:/gotest/main.go:12 +0x83
exit status 2
```

出现上面的结果是因为`for`循环一直在获取通道中的值，但是在读取完`hello world`和`hello China`后，通道中没有新的值传入，这样接收者就阻塞了。

## Select waits on a group of channelsSelect waits on a group of channels

The select statement waits for multiple send or receive opera­tions simul­taneously.

* The statement blocks as a whole until one of the operations becomes unblocked.
* If several cases can proceed, a single one of them will be chosen at random.

```text
// blocks until there's data available on ch1 or ch2
select {
case <-ch1:
    fmt.Println("Received from ch1")
case <-ch2:
    fmt.Println("Received from ch2")
}
```

Send and receive operations on a `nil` channel block forever. This can be used to disable a channel in a select statement:

```text
ch1 = nil // disables this channel
select {
case <-ch1:
    fmt.Println("Received from ch1") // will not happen
case <-ch2:
    fmt.Println("Received from ch2")
}
```

### Default case <a id="default-case"></a>

The `default` case is always able to proceed and runs if all other cases are blocked.

```text
// never blocks
select {
case x := <-ch:
    fmt.Println("Received", x)
default:
    fmt.Println("Nothing available")
}
```

### Examples <a id="examples"></a>

#### An infinite random binary sequence <a id="an-infinite-random-binary-sequence"></a>

As a toy example you can use the _random choice among cases that can proceed_ to generate random bits.

```text
rand := make(chan int)
for {
    select {
    case rand <- 0: // no statement
    case rand <- 1:
    }
}
```

#### A blocking operation with a timeout <a id="a-blocking-operation-with-a-timeout"></a>

The function [`time.After`](https://golang.org/pkg/time#After) is part of the standard library; it waits for a specified time to elapse and then sends the current time on the returned channel.

```text
select {
case news := <-AFP:
    fmt.Println(news)
case <-time.After(time.Minute):
    fmt.Println("Time out: No news in one minute")
}
```

#### A statement that blocks forever <a id="a-statement-that-blocks-forever"></a>

A select statement blocks until _at least one_ of it’s cases can proceed. With zero cases this will never happen.

```text
select {}
```

A typical use would be at the end of the main function in some multithreaded programs. When main returns, the program exits and it does not wait for other goroutines to complete.

## Data races explained

A data race happens when two goroutines access the same variable concur­rently, and at least one of the accesses is a write.

Data races are quite common and can be very hard to debug.

This function has a data race and it’s behavior is undefined. It may, for example, print the number 1. Try to figure out how that can happen – one possible explanation comes after the code.

```text
func race() {
    wait := make(chan struct{})
    n := 0
    go func() {
        n++ // read, increment, write
        close(wait)
    }()
    n++ // conflicting access
    <-wait
    fmt.Println(n) // Output: <unspecified>
}
```

The two goroutines, g1 and g2, participate in a race and there is no way to know in which order the operations will take place. The following is one out of many possible outcomes.

| g1 | g2 |
| :--- | :--- |
| Read the value 0 from `n`. |  |
|  | Read the value 0 from `n`. |
| Incre­ment value from 0 to 1. |  |
| Write 1 to `n`. |  |
|  | Incre­ment value from 0 to 1. |
|  | Write 1 to `n`. |
| Print `n`, which is now 1. |  |

The name ”data race” is somewhat misleading. Not only is the ordering of operations undefined – there are very few guarantees. Both compilers and hardware frequently turn code upside-down and inside-out to achieve better performance. If you look at a thread in mid-action, you might see pretty much anything.

### How to avoid data races <a id="how-to-avoid-data-races"></a>

The only way to avoid data races is to synchronize access to all mutable data that is shared between threads. There are several ways to achieve this. In Go, you would normally use a **channel** or a **lock**. \(Lower-lever mechanisms are available in the [`sync`](https://golang.org/pkg/sync/) and [`sync/atomic`](https://golang.org/pkg/sync/atomic/) packages.\)

The preferred way to handle concurrent data access in Go is to use a channel to pass the actual data from one goroutine to the next. The motto is: “Don’t communicate by sharing memory; share memory by communicating.”

```text
func sharingIsCaring() {
    ch := make(chan int)
    go func() {
        n := 0 // A local variable is only visible to one goroutine.
        n++
        ch <- n // The data leaves one goroutine...
    }()
    n := <-ch // ...and arrives safely in another.
    n++
    fmt.Println(n) // Output: 2
}
```

In this code the channel does double duty:

* it passes the data from one goroutine to another,
* and it acts as a point of synchronization.

The sending goroutine will wait for the other goroutine to receive the data and the receiving goroutine will wait for the other goroutine to send the data.

[The Go memory model](https://golang.org/ref/mem) – the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine – is quite complicated, but as long as you share all mutable data between goroutines through channels you are safe from data races.

## How to detect data races

Data races can happen easily and are hard to debug. Luckily, the Go runtime is often able to help.

Use `-race` to enable the built-in data race detector.

```text
$ go test -race [packages]
$ go run -race [packages]
```

### Example <a id="example"></a>

Here’s a program with a data race:

```text
package main
import "fmt"

func main() {
    i := 0
    go func() {
        i++ // write
    }()
    fmt.Println(i) // concurrent read
}
```

Running this program with the `-race` options tells us that there’s a race between the write at line 7 and the read at line 9:

```text
$ go run -race main.go
0
==================
WARNING: DATA RACE
Write by goroutine 6:
  main.main.func1()
      /tmp/main.go:7 +0x44

Previous read by main goroutine:
  main.main()
      /tmp/main.go:9 +0x7e

Goroutine 6 (running) created at:
  main.main()
      /tmp/main.go:8 +0x70
==================
Found 1 data race(s)
exit status 66
```

### Details <a id="details"></a>

The data race detector does not perform any static analysis. It checks the memory access in runtime and only for the code paths that are actually executed.

It runs on darwin/amd64, freebsd/amd64, linux/amd64 and windows/amd64.

The overhead varies, but typically there’s a 5-10x increase in memory usage, and 2-20x increase in execution time.

## How to debug deadlocks

A deadlock happens when a group of goroutines are waiting for each other and none of them is able to proceed.

Take a look at this simple example.

```text
func main() {
	ch := make(chan int)
	ch <- 1
	fmt.Println(<-ch)
}
```

The program will get stuck on the channel send operation waiting forever for someone to read the value. Go is able to detect situations like this at runtime. Here is the output from our program:

```text
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
	.../deadlock.go:7 +0x6c
```

### Debugging tips <a id="debugging-tips"></a>

A goroutine can get stuck

* either because it’s waiting for a **channel** or
* because it is waiting for one of the **locks** in the [sync](https://golang.org/pkg/sync/) package.

Common reasons are that

* no other goroutine has access to the channel or the lock,
* a group of goroutines are waiting for each other and none of them is able to proceed.

Currently Go only detects when the program as a whole freezes, not when a subset of goroutines get stuck.

With channels it’s often easy to figure out what caused a deadlock. Programs that make heavy use of mutexes can, on the other hand, be notoriously difficult to debug.

## Waiting for goroutines



A [`sync.WaitGroup`](https://golang.org/pkg/sync/) waits for a group of goroutines to finish.

```text
var wg sync.WaitGroup
wg.Add(2)
go func() {
    // Do work.
    wg.Done()
}()
go func() {
    // Do work.
    wg.Done()
}()
wg.Wait()
```

* First the main goroutine calls [`Add`](https://golang.org/pkg/sync/#WaitGroup.Add) to set the number of goroutines to wait for.
* Then two new goroutines run and call [`Done`](https://golang.org/pkg/sync/#WaitGroup.Done) when finished.

_At the same time_, [`Wait`](https://golang.org/pkg/sync/#WaitGroup.Wait) is used to block until these two goroutines have finished.

> **Note:** A WaitGroup must not be copied after first use.

## Broadcast a signal on a channel

All readers receive zero values on a closed channel.

In this example the `Publish` function returns a channel, which is used to broadcast a signal when a message has been published.

```text
// Print text after the given time has expired.
// When done, the wait channel is closed.
func Publish(text string, delay time.Duration) (wait <-chan struct{}) {
    ch := make(chan struct{})
    go func() {
        time.Sleep(delay)
        fmt.Println("BREAKING NEWS:", text)
        close(ch) // Broadcast to all receivers.
    }()
    return ch
}
```

Notice that we use a channel of empty structs: `struct{}`. This clearly indicates that the channel will only be used for signalling, not for passing data.

This is how you may use the function.

```text
func main() {
    wait := Publish("Channels let goroutines communicate.", 5*time.Second)
    fmt.Println("Waiting for news...")
    <-wait
    fmt.Println("Time to leave.")
}
```

```text
Waiting for news...
BREAKING NEWS: Channels let goroutines communicate.
Time to leave.
```

## How to kill a goroutine

One goroutine can't forcibly stop another.

To make a goroutine stoppable, let it listen for a stop signal on a dedicated quit channel, and check this channel at suitable points in your goroutine.

```text
quit := make(chan bool)
go func() {
    for {
        select {
        case <-quit:
            return
        default:
            // …
        }
    }
}()
// …
quit <- true
```

Here is a more complete example, where we use a single channel for both data and signalling.

```text
// Generator returns a channel that produces the numbers 1, 2, 3,…
// To stop the underlying goroutine, send a number on this channel.
func Generator() chan int {
    ch := make(chan int)
    go func() {
        n := 1
        for {
            select {
            case ch <- n:
                n++
            case <-ch:
                return
            }
        }
    }()
    return ch
}

func main() {
    number := Generator()
    fmt.Println(<-number)
    fmt.Println(<-number)
    number <- 0           // stops underlying goroutine
    fmt.Println(<-number) // error, no one is sending anymore
    // …
}
```

```text
1
2
fatal error: all goroutines are asleep - deadlock!
```

## Timer and Ticker: events in the future

Timers and Tickers let you execute code in the future, once or repeatedly.

### Timeout \(Timer\) <a id="timeout-timer"></a>

[`time.After`](https://golang.org/pkg/time/#After) waits for a specified duration and then sends the current time on the returned channel:

```text
select {
case news := <-AFP:
	fmt.Println(news)
case <-time.After(time.Hour):
	fmt.Println("No news in an hour.")
}
```

The underlying [`time.Timer`](https://golang.org/pkg/time/#Timer) will not be recovered by the garbage collector until the timer fires. If this is a concern, use [`time.NewTimer`](https://golang.org/pkg/time/#NewTimer) instead and call its [`Stop`](https://golang.org/pkg/time/#Timer.Stop) method when the timer is no longer needed:

```text
for alive := true; alive; {
	timer := time.NewTimer(time.Hour)
	select {
	case news := <-AFP:
		timer.Stop()
		fmt.Println(news)
	case <-timer.C:
		alive = false
		fmt.Println("No news in an hour. Service aborting.")
	}
}
```

### Repeat \(Ticker\) <a id="repeat-ticker"></a>

[`time.Tick`](https://golang.org/pkg/time/#Tick) returns a channel that delivers clock ticks at even intervals:

```text
go func() {
	for now := range time.Tick(time.Minute) {
		fmt.Println(now, statusUpdate())
	}
}()
```

The underlying [`time.Ticker`](https://golang.org/pkg/time/#Ticker) will not be recovered by the garbage collector. If this is a concern, use [`time.NewTicker`](https://golang.org/pkg/time/#NewTicker) instead and call its [`Stop`](https://golang.org/pkg/time/#Timer.Stop) method when the ticker is no longer needed.

### Wait, act and cancel <a id="wait-act-and-cancel"></a>

[`time.AfterFunc`](https://golang.org/pkg/time/#AfterFunc) waits for a specified duration and then calls a function in its own goroutine. It returns a [`time.Timer`](https://golang.org/pkg/time/#Timer) that can be used to cancel the call:

```text
func Foo() {
    timer = time.AfterFunc(time.Minute, func() {
        log.Println("Foo run for more than a minute.")
    })
    defer timer.Stop()

    // Do heavy work
}
```

## Mutual exclusion lock \(mutex\)

Mutexes let you synchronize data access by explicit locking, without channels.

Sometimes it’s more convenient to synchronize data access by explicit locking instead of using channels. The Go standard library offers a mutual exclusion lock, [sync.Mutex](https://golang.org/pkg/sync/#Mutex), for this purpose.

### Use with caution <a id="use-with-caution"></a>

For this type of locking to be safe, it’s crucial that all accesses to the shared data, both reads and writes, are performed only when a goroutine holds the lock. One mistake by a single goroutine is enough to introduce a data race and break the program.

Because of this you should consider designing a custom data structure with a clean API and make sure that all the synchronization is done internally.

In this example we build a safe and easy-to-use concurrent data structure, `AtomicInt`, that stores a single integer. Any number of goroutines can safely access this number through the `Add` and `Value` methods.

```text
// AtomicInt is a concurrent data structure that holds an int.
// Its zero value is 0.
type AtomicInt struct {
    mu sync.Mutex // A lock than can be held by one goroutine at a time.
    n  int
}

// Add adds n to the AtomicInt as a single atomic operation.
func (a *AtomicInt) Add(n int) {
    a.mu.Lock() // Wait for the lock to be free and then take it.
    a.n += n
    a.mu.Unlock() // Release the lock.
}

// Value returns the value of a.
func (a *AtomicInt) Value() int {
    a.mu.Lock()
    n := a.n
    a.mu.Unlock()
    return n
}

func main() {
    wait := make(chan struct{})
    var n AtomicInt
    go func() {
        n.Add(1) // one access
        close(wait)
    }()
    n.Add(1) // another concurrent access
    <-wait
    fmt.Println(n.Value()) // 2
}
```

## 3 rules for efficient parallel computation

Dividing a large compu­tation into work units for parallel pro­cessing is more of an art than a science.

Here are three rules of thumb.

* _Divide the work into units that take about 100μs to 1ms to compute._

  * If the work units are too small, the adminis­trative over­head of divi­ding the problem and sched­uling sub-problems might be too large.
  * If the units are too big, the whole computation may have to wait for a single slow work item to finish. This slowdown can happen for many reasons, such as scheduling, interrupts from other processes, and unfortunate memory layout.

  Note that the number of work units is independent of the number of CPUs.

* _Try to minimize the amount of data sharing._
  * Concurrent writes can be very costly, particularly so if goroutines execute on separate CPUs.
  * Sharing data for reading is often much less of a problem.
* _Strive for good locality when accessing data._
  * If data can be kept in cache memory, data loading and storing will be dramatically faster.
  * Once again, this is particularly important for writing.

Whatever strategies you are using, don’t forget to [benchmark](https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go) and [profile](https://blog.golang.org/profiling-go-programs) your code.

### Example <a id="example"></a>

The following example shows how to divide a costly computation and distribute it on all available CPUs. This is the code we want to optimize.

```text
type Vector []float64

// Convolve computes w = u * v, where w[k] = Σ u[i]*v[j], i + j = k.
// Precondition: len(u) > 0, len(v) > 0.
func Convolve(u, v Vector) Vector {
    n := len(u) + len(v) - 1
    w := make(Vector, n)
    for k := 0; k < n; k++ {
        w[k] = mul(u, v, k)
    }
    return w
}

// mul returns Σ u[i]*v[j], i + j = k.
func mul(u, v Vector, k int) float64 {
    var res float64
    n := min(k+1, len(u))
    j := min(k, len(v)-1)
    for i := k - j; i < n; i, j = i+1, j-1 {
        res += u[i] * v[j]
    }
    return res
}
```

The idea is simple: identify work units of suitable size and then run each work unit in a separate goroutine. Here is a parallel version of `Convolve`.

```text
func Convolve(u, v Vector) Vector {
    n := len(u) + len(v) - 1
    w := make(Vector, n)

    // Divide w into work units that take ~100μs-1ms to compute.
    size := max(1, 1000000/n)

    var wg sync.WaitGroup
    for i, j := 0, size; i < n; i, j = j, j+size {
        if j > n {
            j = n
        }
        // These goroutines share memory, but only for reading.
        wg.Add(1)
        go func(i, j int) {
            for k := i; k < j; k++ {
                w[k] = mul(u, v, k)
            }
            wg.Done()
        }(i, j)
    }
    wg.Wait()
    return w
}
```

When the work units have been defined, it’s often best to leave the scheduling to the runtime and the operating system. However, if needed, you can tell the runtime how many goroutines you want executing code simultaneously.

```text
func init() {
    numcpu := runtime.NumCPU()
    runtime.GOMAXPROCS(numcpu) // Try to use all available CPUs.
}
```

{% embed url="https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go" %}

{% embed url="https://blog.golang.org/pprof" %}

## 14.2 协程间的信道

```text
var ch1 chan string
ch1 = make(chan string)
```

### 14.2.2 通信操作符 &lt;-

`ch <- int1` 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送）

从通道流出（接收），三种方式：

`int2 = <- ch` 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设 int2 已经声明过了，如果没有的话可以写成：`int2 := <- ch`。

`<- ch` 可以单独调用获取通道的（下一个）值，当前值会被丢弃，但是可以用来验证，所以以下代码是合法的：

```text
if <- ch != 1000{
	...
}
```

同一个操作符 `<-` 既用于**发送**也用于**接收**，但Go会根据操作对象弄明白该干什么 。虽非强制要求，但为了可读性通道的命名通常以 `ch` 开头或者包含 `chan`。通道的发送和接收都是原子操作：它们总是互不干扰的完成的。下面的示例展示了通信操作符的使用。

```text
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
```

输出：

```text
Washington Tripoli London Beijing tokyo
```

如果 2 个协程需要通信，你必须给他们同一个通道作为参数才行。

尝试一下如果注释掉 `time.Sleep(1e9)` 会如何。

我们发现协程之间的同步非常重要：

* main\(\) 等待了 1 秒让两个协程完成，如果不这样，sendData\(\) 就没有机会输出。
* getData\(\) 使用了无限循环：它随着 sendData\(\) 的发送完成和 ch 变空也结束了。
* 如果我们移除一个或所有 `go` 关键字，程序无法运行，Go 运行时会抛出 panic：

```text
---- Error run E:/Go/Goboek/code examples/chapter 14/goroutine2.exe with code Crashed ---- Program exited with code -2147483645: panic: all goroutines are asleep-deadlock!
```

为什么会这样？运行时（runtime）会检查所有的协程（像本例中只有一个）是否在等待着什么东西（可从某个通道读取或者写入某个通道），这意味着程序将无法继续执行。这是死锁（deadlock）的一种形式，而运行时（runtime）可以为我们检测到这种情况。

注意：不要使用打印状态来表明通道的发送和接收顺序：由于打印状态和通道实际发生读写的时间延迟会导致和真实发生的顺序不同。

练习 14.4：解释一下为什么如果在函数 `getData()` 的一开始插入 `time.Sleep(2e9)`，不会出现错误但也没有输出呢。

### 14.2.3 通道阻塞

默认情况下，通信是同步且无缓冲的：在有接受者接收数据之前，发送不会结束。可以想象一个无缓冲的通道在没有空间来保存数据的时候：必须要一个接收者准备好接收通道的数据然后发送者可以直接把数据发送给接收者。所以通道的发送/接收操作在对方准备好之前是阻塞的：

1）对于同一个通道，发送操作（协程或者函数中的），在接收者准备好之前是阻塞的：如果ch中的数据无人接收，就无法再给通道传入其他数据：新的输入无法在通道非空的情况下传入。所以发送操作会等待 ch 再次变为可用状态：就是通道值被接收时（可以传入变量）。

2）对于同一个通道，接收操作是阻塞的（协程或函数中的），直到发送者可用：如果通道中没有数据，接收者就阻塞了。

尽管这看上去是非常严格的约束，实际在大部分情况下工作的很不错。

程序 `channel_block.go` 验证了以上理论，一个协程在无限循环中给通道发送整数数据。不过因为没有接收者，只输出了一个数字 0。

示例 14.3-[channel\_block.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/channel_block.go)

```text
package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	fmt.Println(<-ch1) // prints only 0
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}
```

输出：

```text
0
```

`pump()` 函数为通道提供数值，也被叫做生产者。

为通道解除阻塞定义了 `suck` 函数来在无限循环中读取通道，参见示例 14.4-[channel\_block2.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/channel_block2.go)：

```text
func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
```

在 `main()` 中使用协程开始它：

```text
go pump(ch1)
go suck(ch1)
time.Sleep(1e9)
```

给程序 1 秒的时间来运行：输出了上万个整数。

练习 14.1：[channel\_block3.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/channel_block3.go)：写一个通道证明它的阻塞性，开启一个协程接收通道的数据，持续 15 秒，然后给通道放入一个值。在不同的阶段打印消息并观察输出。

### 14.2.4 通过一个（或多个）通道交换数据进行协程同步。

通信是一种同步形式：通过通道，两个协程在通信（协程会和）中某刻同步交换数据。无缓冲通道成为了多个协程同步的完美工具。

甚至可以在通道两端互相阻塞对方，形成了叫做死锁的状态。Go 运行时会检查并 panic，停止程序。死锁几乎完全是由糟糕的设计导致的。

无缓冲通道会被阻塞。设计无阻塞的程序可以避免这种情况，或者使用带缓冲的通道。

练习 14.2： [blocking.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/blocking.go)

解释为什么下边这个程序会导致 panic：所有的协程都休眠了 - 死锁！

```text
package main

import (
	"fmt"
)

func f1(in chan int) {
	fmt.Println(<-in)
}

func main() {
	out := make(chan int)
	out <- 2
	go f1(out)
}
```

### 14.2.5 同步通道-使用带缓冲的通道

一个无缓冲通道只能包含 1 个元素，有时显得很局限。我们给通道提供了一个缓存，可以在扩展的 `make` 命令中设置它的容量，如下：

```text
buf := 100
ch1 := make(chan string, buf)
```

buf 是通道可以同时容纳的元素（这里是 string）个数

在缓冲满载（缓冲被全部使用）之前，给一个带缓冲的通道发送数据是不会阻塞的，而从通道读取数据也不会阻塞，直到缓冲空了。

缓冲容量和类型无关，所以可以（尽管可能导致危险）给一些通道设置不同的容量，只要他们拥有同样的元素类型。内置的 `cap` 函数可以返回缓冲区的容量。

如果容量大于 0，通道就是异步的了：缓冲满载（发送）或变空（接收）之前通信不会阻塞，元素会按照发送的顺序被接收。如果容量是0或者未设置，通信仅在收发双方准备好的情况下才可以成功。

同步：`ch :=make(chan type, value)`

* value == 0 -&gt; synchronous, unbuffered \(阻塞）
* value &gt; 0 -&gt; asynchronous, buffered（非阻塞）取决于value元素

若使用通道的缓冲，你的程序会在“请求”激增的时候表现更好：更具弹性，专业术语叫：更具有伸缩性（scalable）。在设计算法时首先考虑使用无缓冲通道，只在不确定的情况下使用缓冲。

练习 14.3：[channel\_buffer.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/channel_buffer.go)：给 [channel\_block3.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/channel_block3.go) 的通道增加缓冲并观察输出有何不同。

### 14.2.6 协程中用通道输出结果

为了知道计算何时完成，可以通过信道回报。在例子 `go sum(bigArray)` 中，要这样写：

```text
ch := make(chan int)
go sum(bigArray, ch) // bigArray puts the calculated sum on ch
// .. do something else for a while
sum := <- ch // wait for, and retrieve the sum
```

也可以使用通道来达到同步的目的，这个很有效的用法在传统计算机中称为信号量（semaphore）。或者换个方式：通过通道发送信号告知处理已经完成（在协程中）。

在其他协程运行时让 main 程序无限阻塞的通常做法是在 `main` 函数的最后放置一个 `select {}`。

也可以使用通道让 `main` 程序等待协程完成，就是所谓的信号量模式，我们会在接下来的部分讨论。



### 14.2.7 信号量模式

下边的片段阐明：协程通过在通道 `ch` 中放置一个值来处理结束的信号。`main` 协程等待 `<-ch` 直到从中获取到值。

我们期望从这个通道中获取返回的结果，像这样：

```text
func compute(ch chan int){
	ch <- someComputation() // when it completes, signal on the channel.
}

func main(){
	ch := make(chan int) 	// allocate a channel.
	go compute(ch)		// start something in a goroutines
	doSomethingElseForAWhile()
	result := <- ch
}
```

这个信号也可以是其他的，不返回结果，比如下面这个协程中的匿名函数（lambda）协程：

```text
ch := make(chan int)
go func(){
	// doSomething
	ch <- 1 // Send a signal; value does not matter
}()
doSomethingElseForAWhile()
<- ch	// Wait for goroutine to finish; discard sent value.
```

或者等待两个协程完成，每一个都会对切片s的一部分进行排序，片段如下：

```text
done := make(chan bool)
// doSort is a lambda function, so a closure which knows the channel done:
doSort := func(s []int){
	sort(s)
	done <- true
}
i := pivot(s)
go doSort(s[:i])
go doSort(s[i:])
<-done
<-done
```

下边的代码，用完整的信号量模式对长度为N的 float64 切片进行了 N 个 `doSomething()` 计算并同时完成，通道 sem 分配了相同的长度（且包含空接口类型的元素），待所有的计算都完成后，发送信号（通过放入值）。在循环中从通道 sem 不停的接收数据来等待所有的协程完成。

```text
type Empty interface {}
var empty Empty
...
data := make([]float64, N)
res := make([]float64, N)
sem := make(chan Empty, N)
...
for i, xi := range data {
	go func (i int, xi float64) {
		res[i] = doSomething(i, xi)
		sem <- empty
	} (i, xi)
}
// wait for goroutines to finish
for i := 0; i < N; i++ { <-sem }
```

注意上述代码中闭合函数的用法：`i`、`xi` 都是作为参数传入闭合函数的，这一做法使得每个协程（译者注：在其启动时）获得一份 `i` 和 `xi` 的单独拷贝，从而向闭合函数内部屏蔽了外层循环中的 `i` 和 `xi`变量；否则，for 循环的下一次迭代会更新所有协程中 `i` 和 `xi` 的值。另一方面，切片 `res` 没有传入闭合函数，因为协程不需要`res`的单独拷贝。切片 `res` 也在闭合函数中但并不是参数。

### 14.2.8 实现并行的 for 循环

在上一部分章节 [14.2.7](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md#1427-%E4%BF%A1%E5%8F%B7%E9%87%8F%E6%A8%A1%E5%BC%8F) 的代码片段中：for 循环的每一个迭代是并行完成的：

```text
for i, v := range data {
	go func (i int, v float64) {
		doSomething(i, v)
		...
	} (i, v)
}
```

在 for 循环中并行计算迭代可能带来很好的性能提升。不过所有的迭代都必须是独立完成的。有些语言比如 Fortress 或者其他并行框架以不同的结构实现了这种方式，在 Go 中用协程实现起来非常容易：

### 14.2.9 用带缓冲通道实现一个信号量

信号量是实现互斥锁（排外锁）常见的同步机制，限制对资源的访问，解决读写问题，比如没有实现信号量的 `sync` 的 Go 包，使用带缓冲的通道可以轻松实现：

* 带缓冲通道的容量和要同步的资源容量相同
* 通道的长度（当前存放的元素个数）与当前资源被使用的数量相同
* 容量减去通道的长度就是未处理的资源个数（标准信号量的整数值）

不用管通道中存放的是什么，只关注长度；因此我们创建了一个长度可变但容量为0（字节）的通道：

```text
type Empty interface {}
type semaphore chan Empty
```

将可用资源的数量N来初始化信号量 `semaphore`：`sem = make(semaphore, N)`

然后直接对信号量进行操作：

```text
// acquire n resources
func (s semaphore) P(n int) {
	e := new(Empty)
	for i := 0; i < n; i++ {
		s <- e
	}
}

// release n resources
func (s semaphore) V(n int) {
	for i:= 0; i < n; i++{
		<- s
	}
}
```

可以用来实现一个互斥的例子：

```text
/* mutexes */
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock(){
	s.V(1)
}

/* signal-wait */
func (s semaphore) Wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}
```

练习 14.5：[gosum.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/gosum.go)：用这种习惯用法写一个程序，开启一个协程来计算2个整数的和并等待计算结果并打印出来。

练习 14.6：[producer\_consumer.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/producer_consumer.go)：用这种习惯用法写一个程序，有两个协程，第一个提供数字 0，10，20，...90 并将他们放入通道，第二个协程从通道中读取并打印。`main()` 等待两个协程完成后再结束。

习惯用法：通道工厂模式

编程中常见的另外一种模式如下：不将通道作为参数传递给协程，而用函数来生成一个通道并返回（工厂角色）；函数内有个匿名函数被协程调用。

在 [channel\_block2.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/channel_block2.go) 加入这种模式便有了示例 14.5-[channel\_idiom.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/channel_idiom.go)：

```text
package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
```

### 14.2.10 给通道使用 for 循环

`for` 循环的 `range` 语句可以用在通道 `ch` 上，便可以从通道中获取值，像这样：

```text
for v := range ch {
	fmt.Printf("The value is %v\n", v)
}
```

它从指定通道中读取数据直到通道关闭，才继续执行下边的代码。很明显，另外一个协程必须写入 `ch`（不然代码就阻塞在 for 循环了），而且必须在写入完成后才关闭。`suck` 函数可以这样写，且在协程中调用这个动作，程序变成了这样：

示例 14.6-[channel\_idiom2.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/channel_idiom2.go)：

```text
package main

import (
	"fmt"
	"time"
)

func main() {
	suck(pump())
	time.Sleep(1e9)
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}
```

习惯用法：通道迭代模式

这个模式用到了后边14.6章示例 [producer\_consumer.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/exercises/chapter_14/producer_consumer.go) 的生产者-消费者模式，通常，需要从包含了地址索引字段 items 的容器给通道填入元素。为容器的类型定义一个方法 `Iter()`，返回一个只读的通道（参见第 [14.2.11](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.2.md#14211-%E9%80%9A%E9%81%93%E7%9A%84%E6%96%B9%E5%90%91) 节）items，如下：

```text
func (c *container) Iter () <- chan item {
	ch := make(chan item)
	go func () {
		for i:= 0; i < c.Len(); i++{	// or use a for-range loop
			ch <- c.items[i]
		}
	} ()
	return ch
}
```

在协程里，一个 for 循环迭代容器 c 中的元素（对于树或图的算法，这种简单的 for 循环可以替换为深度优先搜索）。

调用这个方法的代码可以这样迭代容器：

```text
for x := range container.Iter() { ... }
```

其运行在自己启动的协程中，所以上边的迭代用到了一个通道和两个协程（可能运行在不同的线程上）。 这样我们就有了一个典型的生产者-消费者模式。如果在程序结束之前，向通道写值的协程未完成工作，则这个协程不会被垃圾回收；这是设计使然。这种看起来并不符合预期的行为正是由通道这种线程安全的通信方式所导致的。如此一来，一个协程为了写入一个永远无人读取的通道而被挂起就成了一个bug，而并非你预想中的那样被悄悄回收掉（garbage-collected）了。

习惯用法：生产者消费者模式

假设你有 `Produce()` 函数来产生 `Consume` 函数需要的值。它们都可以运行在独立的协程中，生产者在通道中放入给消费者读取的值。整个处理过程可以替换为无限循环：

```text
for {
	Consume(Produce())
}
```

### 14.2.11 通道的方向

通道类型可以用注解来表示它只发送或者只接收：

```text
var send_only chan<- int 		// channel can only receive data
var recv_only <-chan int		// channel can only send data
```

只接收的通道（&lt;-chan T）无法关闭，因为关闭通道是发送者用来表示不再给通道发送值了，所以对只接收通道是没有意义的。通道创建的时候都是双向的，但也可以分配有方向的通道变量，就像以下代码：

```text
var c = make(chan int) // bidirectional
go source(c)
go sink(c)

func source(ch chan<- int){
	for { ch <- 1 }
}

func sink(ch <-chan int) {
	for { <-ch }
}
```

习惯用法：管道和选择器模式

更具体的例子还有协程处理它从通道接收的数据并发送给输出通道：

```text
sendChan := make(chan int)
receiveChan := make(chan string)
go processChannel(sendChan, receiveChan)

func processChannel(in <-chan int, out chan<- string) {
	for inValue := range in {
		result := ... /// processing inValue
		out <- result
	}
}
```

通过使用方向注解来限制协程对通道的操作。

这里有一个来自 Go 指导的很赞的例子，打印了输出的素数，使用选择器（‘筛’）作为它的算法。每个 prime 都有一个选择器，如下图：

[![](https://github.com/unknwon/the-way-to-go_ZH_CN/raw/master/eBook/images/14.2_fig14.2.png?raw=true)](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/images/14.2_fig14.2.png?raw=true)

版本1：示例 14.7-[sieve1.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/sieve1.go)

```text
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}

// The prime sieve: Daisy-chain filter processes together.
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for {
		prime := <-ch
		fmt.Print(prime, " ")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
```

协程 `filter(in, out chan int, prime int)` 拷贝整数到输出通道，丢弃掉可以被 prime 整除的数字。然后每个 prime 又开启了一个新的协程，生成器和选择器并发请求。

输出：

```text
2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97 101
103 107 109 113 127 131 137 139 149 151 157 163 167 173 179 181 191 193 197 199 211 223
227 229 233 239 241 251 257 263 269 271 277 281 283 293 307 311 313 317 331 337 347 349
353 359 367 373 379 383 389 397 401 409 419 421 431 433 439 443 449 457 461 463 467 479
487 491 499 503 509 521 523 541 547 557 563 569 571 577 587 593 599 601 607 613 617 619
631 641 643 647 653 659 661 673 677 683 691 701 709 719 727 733 739 743 751 757 761 769
773 787 797 809 811 821 823 827 829 839 853 857 859 863 877 881 883 887 907 911 919 929
937 941 947 953 967 971 977 983 991 997 1009 1013...
```

第二个版本引入了上边的习惯用法：函数 `sieve`、`generate` 和 `filter` 都是工厂；它们创建通道并返回，而且使用了协程的 lambda 函数。`main` 函数现在短小清晰：它调用 `sieve()` 返回了包含素数的通道，然后通过 `fmt.Println(<-primes)` 打印出来。

版本2：示例 14.8-[sieve2.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/sieve2.go)

```text
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
)

// Send the sequence 2, 3, 4, ... to returned channel
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// Filter out input values divisible by 'prime', send rest to returned channel
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}

func main() {
	primes := sieve()
	for {
		fmt.Println(<-primes)
	}
}
```

## 14.3 协程的同步：关闭通道-测试阻塞的通道

通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。

继续看示例 [goroutine2.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/goroutine2.go)（示例 14.2）：我们如何在通道的 `sendData()` 完成的时候发送一个信号，`getData()` 又如何检测到通道是否关闭或阻塞？

第一个可以通过函数 `close(ch)` 来完成：这个将通道标记为无法通过发送操作 `<-` 接受更多的值；给已经关闭的通道发送或者再次关闭都会导致运行时的 panic。在创建一个通道后使用 defer 语句是个不错的办法（类似这种情况）：

```text
ch := make(chan float64)
defer close(ch)
```

第二个问题可以使用逗号，ok 操作符：用来检测通道是否被关闭。

如何来检测可以收到没有被阻塞（或者通道没有被关闭）？

```text
v, ok := <-ch   // ok is true if v received value
```

通常和 if 语句一起使用：

```text
if v, ok := <-ch; ok {
  process(v)
}
```

或者在 for 循环中接收的时候，当关闭或者阻塞的时候使用 break：

```text
v, ok := <-ch
if !ok {
  break
}
process(v)
```

在示例程序 14.2 中使用这些可以改进为版本 goroutine3.go，输出相同。

实现非阻塞通道的读取，需要使用 select（参见第 [14.4](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/14.4.md) 节）。

示例 14.9-[goroutine3.go](https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/examples/chapter_14/goroutine3.go)：

```text
package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}
```

改变了以下代码：

* 现在只有 `sendData()` 是协程，`getData()` 和 `main()` 在同一个线程中：

```text
go sendData(ch)
getData(ch)
```

* 在 `sendData()` 函数的最后，关闭了通道：

```text
func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	close(ch)
}
```

* 在 for 循环的 `getData()` 中，在每次接收通道的数据之前都使用 `if !open` 来检测：

```text
for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
```

使用 for-range 语句来读取通道是更好的办法，因为这会自动检测通道是否关闭：

```text
for input := range ch {
  	process(input)
}
```

阻塞和生产者-消费者模式：

在第 14.2.10 节的通道迭代器中，两个协程经常是一个阻塞另外一个。如果程序工作在多核心的机器上，大部分时间只用到了一个处理器。可以通过使用带缓冲（缓冲空间大于 0）的通道来改善。比如，缓冲大小为 100，迭代器在阻塞之前，至少可以从容器获得 100 个元素。如果消费者协程在独立的内核运行，就有可能让协程不会出现阻塞。

由于容器中元素的数量通常是已知的，需要让通道有足够的容量放置所有的元素。这样，迭代器就不会阻塞（尽管消费者协程仍然可能阻塞）。然而，这实际上加倍了迭代容器所需要的内存使用量，所以通道的容量需要限制一下最大值。记录运行时间和性能测试可以帮助你找到最小的缓存容量带来最好的性能。

