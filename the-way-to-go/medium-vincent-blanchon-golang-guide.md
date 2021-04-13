# Medium Vincent Blanchon Golang Guide

## Go: How to Reduce Lock Contention with the Atomic Package <a id="c2b4"></a>

{% embed url="https://medium.com/a-journey-with-go/go-how-to-reduce-lock-contention-with-the-atomic-package-ba3b2664b549" %}

ℹ️ _This article is based on Go 1.14._

Go provides memory synchronization mechanisms such as channel or mutex that help to solve different issues. In the case of shared memory, mutex protects the memory against data races. However, although two mutexes exist, Go also provides atomic memory primitives via the `atomic` package to improve performance. Let’s first go back to the data races before diving into the solutions.

## Data Race <a id="0744"></a>

A data race can occur when two or more goroutines access the same memory location concurrently, and at least one of them is writing. While the maps have a native mechanism to protect against data races, a simple structure does not have any, making it vulnerable to data races.

To illustrate a data race, I will take an example of a configuration that is continuously updated by a goroutine. Here is the code:![](https://miro.medium.com/max/60/1*w5b8O7ijqoqGkiVEOI3Q6A.png?q=20)![](https://miro.medium.com/max/2816/1*w5b8O7ijqoqGkiVEOI3Q6A.png)

Running this code clearly shows that the result is non-deterministic due to the data race:

```text
[...]
&{[79167 79170 79173 79176 79179 79181]}
&{[79216 79219 79220 79221 79222 79223]}
&{[79265 79268 79271 79274 79278 79281]}
```

Each line was expecting to be a continuous sequence of integers when the result is quite random. Running the same program with the flag `-race` points out the data races:

```text
WARNING: DATA RACE
Read at 0x00c0003aa028 by goroutine 9:
  [...]
  fmt.Printf()
      /usr/local/go/src/fmt/print.go:213 +0xb5
  main.main.func2()
      main.go:30 +0x3b

Previous write at 0x00c0003aa028 by goroutine 7:
  main.main.func1()
      main.go:20 +0xfe
```

Protecting our reads and writes from data races can be done by a mutex — probably the most common one — or by the `atomic` package.

## Mutex vs Atomic <a id="775f"></a>

The standard library provides two kinds of mutex with the `sync` package: `sync.Mutex` and `sync.RWMutex`; the latter is optimized when your program deals with multiples readers and very few writers. Here is one solution:![](https://miro.medium.com/max/60/1*PDqI7-YpPCdpvLnLkytxdQ.png?q=20)![](https://miro.medium.com/max/2720/1*PDqI7-YpPCdpvLnLkytxdQ.png)

The program now prints out the expected result; the numbers are properly incremented:

```text
[...]
&{[213 214 215 216 217 218]}
&{[214 215 216 217 218 219]}
&{[215 216 217 218 219 220]}
```

The second solution can be done thanks to the `atomic` package. Here is the code:![](https://miro.medium.com/max/60/1*J4gwjLHj7V7M1RhYrQ6q5Q.png?q=20)![](https://miro.medium.com/max/2744/1*J4gwjLHj7V7M1RhYrQ6q5Q.png)

The result is also the expected one:

```text
[...]
&{[32724 32725 32726 32727 32728 32729]}
&{[32733 32734 32735 32736 32737 32738]}
&{[32753 32754 32755 32756 32757 32758]}
```

Regarding the generated output, it looks like the solution using the `atomic` package is much faster since it can generate a higher sequence of numbers. Benchmarking both of the programs would help to figure out which one is the most efficient.

## Performance <a id="24c5"></a>

A benchmark should be interpreted according to what is measured. In this case, I will measure the previous program where it has a writer that constantly stores a new config along with multiple readers that constantly read it. To cover more potential cases, I will also include benchmarks for a program that only has readers, assuming the config does not change often. Here is an example of this new case:![](https://miro.medium.com/max/60/1*Yt9DjmoSMnh_wo93WkD8zQ.png?q=20)![](https://miro.medium.com/max/2684/1*Yt9DjmoSMnh_wo93WkD8zQ.png)

Running the benchmark ten times give the following results:

```text
name                              time/op
AtomicOneWriterMultipleReaders-4  72.2ns ± 2%
AtomicMultipleReaders-4           65.8ns ± 2%MutexOneWriterMultipleReaders-4    717ns ± 3%
MutexMultipleReaders-4             176ns ± 2%
```

The benchmark confirms what we have seen before in terms of performance. To understand where exactly the bottleneck is with the mutex, we can rerun the program with the tracer enabled.

_For more information about the `trace` package, I suggest you read my article ”_[_Go: Discovery of the Trace Package_](https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c)_.”_

Here is the profile with the program using the `atomic` package:![](https://miro.medium.com/max/60/1*Wr8Xx10Y07ED-fJ-TqStMg.png?q=20)![](https://miro.medium.com/max/2960/1*Wr8Xx10Y07ED-fJ-TqStMg.png)

The goroutines run with no interruption and are able to complete their tasks. Regarding the profile of the program with the mutex, that is quite different:![](https://miro.medium.com/max/60/1*hccsICnf6hF8ORwcU4IONg.png?q=20)![](https://miro.medium.com/max/3016/1*hccsICnf6hF8ORwcU4IONg.png)

The running time is now quite fragmented, and this is due to the mutex that parks the goroutine. This is confirmed from the goroutine’s overview, where it shows the time spent blocked on synchronization:![](https://miro.medium.com/max/60/1*koHTP_rP80pM7QbZo39evQ.png?q=20)![](https://miro.medium.com/max/3096/1*koHTP_rP80pM7QbZo39evQ.png)

The blocking time accounts for roughly a third of the time. It can be detailed from the blocking profile:![](https://miro.medium.com/max/60/1*OLbWMA1mllWjcpcWqr1P5Q.png?q=20)![](https://miro.medium.com/max/2984/1*OLbWMA1mllWjcpcWqr1P5Q.png)

The `atomic` package definitely brings an advantage in that case. However, performance could be degraded in some. For instance, if you would have to store a large map, you would have to copy it every single time the map is updated, making it inefficient.

_For more information about the mutex, I suggest you read my article ”_[_Go: Mutex and Starvation_](https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50)_.”_

\_\_

{% embed url="https://medium.com/a-journey-with-go/go-discovery-of-the-trace-package-e5a821743c3c" %}

{% embed url="https://medium.com/a-journey-with-go/go-mutex-and-starvation-3f4f4e75ad50" %}

\_\_

