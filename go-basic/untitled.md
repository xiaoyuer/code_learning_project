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



