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





