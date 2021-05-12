# concurrent

## 按序打印

```go
package main

import (
	"fmt"
)

type Foo struct {
	startTwo   chan bool
	startThree chan bool
	over       chan bool
}

func (f *Foo) one() {
	fmt.Println("one")
	f.startTwo <- true
}

func (f *Foo) two() {
	<-f.startTwo
	fmt.Println("two")
	f.startThree <- true
}

func (f *Foo) three() {
	<-f.startThree
	fmt.Println("three")
	f.over <- true
}

func main() {
	f := new(Foo)
	f.startTwo = make(chan bool)
	f.startThree = make(chan bool)
	f.over = make(chan bool)
	go f.one()
	go f.two()
	go f.three()
	<-f.over
}



```





```go
//文件名为print-in-order_test.go
package test
import (
	"fmt"
	"sync"
	"testing"
)

type Foo struct{}

func (*Foo) first()  { fmt.Println("first") }
func (*Foo) second() { fmt.Println("second") }
func (*Foo) third()  { fmt.Println("third") }
//根据题意声明结构体方法

func TestPrintInOrder(t *testing.T) {
	f := new(Foo)
	var wg sync.WaitGroup
	wg.Add(len(order))
	syncChan := make(chan int)
	flag := 1
	order := [...]int{3, 1, 2}
    //通过order、switch模拟题意中3个协程按不同顺序调用方法
	for i := 0; i < len(order); i++ {
		switch order[i] {
		case 1:
			go func() {
				defer wg.Done()
				f.first()
				flag++
				syncChan <- flag
			}()
		case 2:
			go func() {
				defer wg.Done()
			LOOP1:
				for {
					if i := <-syncChan; i == 2 {
						f.second()
						flag++
						syncChan <- flag
						break LOOP1
					} else {
						syncChan <- i
					}
				}
			}()
		case 3:
			go func() {
				defer wg.Done()
			LOOP2:
				for {
					if i := <-syncChan; i == 3 {
						f.third()
						break LOOP2
					} else {
						syncChan <- i
					}
				}
			}()
		}
	}
	wg.Wait()
}


```

```go
var secondChan = make(chan int)
var thirdChan = make(chan int)
var mainChan = make(chan int)

func first() {
	fmt.Print("first")
	secondChan <- 1
}

func second() {
	signal := <-secondChan
	fmt.Print("second")
	thirdChan <- signal
	close(secondChan)
}

func third() {
	signal := <-thirdChan
	fmt.Print("third")
	mainChan <- signal
	close(thirdChan)
}

func main() {
	funcMap := map[int]func(){1: first, 2: second, 3: third}
	inputList := [3]int{1, 2, 3}

	for _, num := range inputList {
		go funcMap[num]()
	}

	_ = <-mainChan
	close(mainChan)
}

```





python 

第一种基础方法超时，中间五种方法在线上测试的时间上基本没有区别都可以做到最快72ms, 95-97%这样。

最后几种数据结构的方法没有多线程模块threading的阻塞，分别用的是多线程专用的阻塞队列数据结构queue的两种方法，和本身线程安全的字典，队列最快可以达到68ms, 98%这样，字典最快可以达到64ms, 99%这样，对于不懂多线程的同学来说也算是比较好理解的方法了。（11月更新一下，字典法现在被判不在规定的线程输出，已经不能用了）

另外sleep大法已经不能过了，包括threading模块里的Timer也一样，主要是每个函数的执行时间不固定，休眠时间短未必能做到正解，休眠时间太久又会超时。

方法一，while循环法（超时）：

可能是不懂多线程的同学最能够接受的基础解法，可以大体理解多线程的阻塞是什么意思。

就相当于先用某些方法卡住执行顺序，然后不断监控目标，直到目标符合条件时才跳出当前断点继续执行后续语句。

输出是正确的，只是因为没法像threading模块那样很好的监控线程，所以大概率会超时，其他语言或许可以用这种方法AC，但python相对较慢，大约只能过30/37的数据。

对于单次阻塞来说，运行时间大约是threading模块时间的10-14倍这样，整个程序平均时间差距就会在15-25倍这样。

python

class Foo: def **init**\(self\): self.t = 0

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.t = 1

def second(self, printSecond: 'Callable[[], None]') -> None:
    while self.t != 1: 
        pass
    printSecond()
    self.t = 2

def third(self, printThird: 'Callable[[], None]') -> None:
    while self.t != 2: 
        pass
    printThird()
```

来自评论区，改进后可以通过。 while循环跑满CPU，会影响GIL线程上下文切换的判定，可能是导致超时的重要原因之一，time.sleep会把CPU交还给GIL，可以让GIL及时切换线程，即使sleep很短的时间（比如1e-9或更小），也可以通过线上测试。

python

from time import sleep class Foo: def **init**\(self\): self.t = 0

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.t = 1

def second(self, printSecond: 'Callable[[], None]') -> None:
    while self.t != 1: 
        sleep(1e-3)
    printSecond()
    self.t = 2

def third(self, printThird: 'Callable[[], None]') -> None:
    while self.t != 2: 
        sleep(1e-3)
    printThird()
```

方法二，Condition条件对象法：

threading模块里的Condition方法，后面五种的方法也都是调用这个模块和使用不同的方法了，方法就是启动wait\_for来阻塞每个函数，直到指示self.t为目标值的时候才释放线程，with是配合Condition方法常用的语法糖，主要是替代try语句的。

python

import threading

class Foo: def **init**\(self\): self.c = threading.Condition\(\) self.t = 0

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    self.res(0, printFirst)

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.res(1, printSecond)

def third(self, printThird: 'Callable[[], None]') -> None:
    self.res(2, printThird)

def res(self, val: int, func: 'Callable[[], None]') -> None:
    with self.c:
        self.c.wait_for(lambda: val == self.t) #参数是函数对象，返回值是bool类型
        func()
        self.t += 1
        self.c.notify_all()
```

方法三，Lock锁对象法：

在这题里面功能都是类似的，就是添加阻塞，然后释放线程，只是类初始化的时候不能包含有参数，所以要写一句acquire进行阻塞，调用其他函数的时候按顺序release释放。

python

import threading

class Foo: def **init**\(self\): self.l1 = threading.Lock\(\) self.l1.acquire\(\) self.l2 = threading.Lock\(\) self.l2.acquire\(\)

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.l1.release()

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.l1.acquire()
    printSecond()
    self.l2.release()

def third(self, printThird: 'Callable[[], None]') -> None:
    self.l2.acquire()
    printThird()
```

方法四，Semaphore信号量对象法：

和方法三是类似的，不过在类赋值的时候可以带有参数自带阻塞。

import threading

class Foo: def **init**\(self\): self.s1 = threading.Semaphore\(0\) self.s2 = threading.Semaphore\(0\)

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.s1.release()

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.s1.acquire()
    printSecond()
    self.s2.release()

def third(self, printThird: 'Callable[[], None]') -> None:
    self.s2.acquire()
    printThird()
```

方法五，Event事件对象法：

原理同上，用wait方法作为阻塞，用set来释放线程，默认类赋值就是阻塞的。

python

import threading

class Foo: def **init**\(self\): self.e1 = threading.Event\(\) self.e2 = threading.Event\(\)

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.e1.set()

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.e1.wait()
    printSecond()
    self.e2.set()

def third(self, printThird: 'Callable[[], None]') -> None:
    self.e2.wait()
    printThird()
```

方法六，Barrier栅栏对象法：

Barrier初始化的时候定义了parties = 2个等待线程，调用完了parties个wait就会释放线程。

python

import threading

class Foo: def **init**\(self\): self.b1 = threading.Barrier\(2\) self.b2 = threading.Barrier\(2\)

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.b1.wait()

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.b1.wait()
    printSecond()
    self.b2.wait()

def third(self, printThird: 'Callable[[], None]') -> None:
    self.b2.wait()
    printThird()
```

方法七，Queue队列法1：

直接使用多线程专用的阻塞队列，对于队列为空时，get方法就会自动阻塞，直到put使之非空才会释放进程。

python

import queue

class Foo: def **init**\(self\): self.q1 = queue.Queue\(\) self.q2 = queue.Queue\(\)

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.q1.put(0)

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.q1.get()
    printSecond()
    self.q2.put(0)

def third(self, printThird: 'Callable[[], None]') -> None:
    self.q2.get()
    printThird()
```

方法八，Queue队列法2：

反过来，对于定容队列来说，如果队列满了，put方法也是阻塞。

python

import queue

class Foo: def **init**\(self\): self.q1 = queue.Queue\(1\) self.q1.put\(0\) self.q2 = queue.Queue\(1\) self.q2.put\(0\)

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    printFirst()
    self.q1.get()

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.q1.put(0)
    printSecond()
    self.q2.get()

def third(self, printThird: 'Callable[[], None]') -> None:
    self.q2.put(0)
    printThird()
```

方法九，dict字典法（WA）：

把三个函数指针，按指定键值存入线程安全的字典，当字典长度为3时，按序输出字典。

11月更新一下，字典法现在被判不在规定的线程输出，已经不能用了。

python

class Foo: def **init**\(self\): self.d = {}

```text
def first(self, printFirst: 'Callable[[], None]') -> None:
    self.d[0] = printFirst
    self.res()

def second(self, printSecond: 'Callable[[], None]') -> None:
    self.d[1] = printSecond
    self.res()

def third(self, printThird: 'Callable[[], None]') -> None:
    self.d[2] = printThird
    self.res()

def res(self) -> None:
    if len(self.d) == 3:
        self.d[0]()
        self.d[1]()
        self.d[2]()
```

线上测试最快时间来自于字典法：

线下测试运行十次全排列的时间对比，反复运行下来，大体上除了第一种，后面几种方法阻塞效率区分不大，谁快谁慢偶然性很高：

![](../.gitbook/assets/image%20%2827%29.png)

