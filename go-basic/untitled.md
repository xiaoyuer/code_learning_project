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





