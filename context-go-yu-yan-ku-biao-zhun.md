---
description: '转自https://zhuanlan.zhihu.com/p/293666788'
---

# context Go语言库标准

##  什么是context

用来控制goroutine的一种模式：在复杂goroutine应用场景中，往往需要在api边界和过程之间传递截止时间、取消信号或其他相关的数据。

## `context`的应用场景

### goroutine的使用一般有三种方式：

* WaitGroup 希望把一件工作拆成多个job去运行，然后需要等全部工作完成拿到最终结果
* Channel 某个goroutine跑太久了，我们需要发送一个信息让他停止下来，这种情况下可以使用Channel+Select的模式
* Context

我们会一个场景，就是有多个goroutine的情况下，每个中可能又创建了其他goroutine，类似某个施工项目总包后层层拆解分包，分包又分包的情况，这种情况下我们会有一个需求，就是当某个goroutine结束之后，其创建出来的子goroutine任务也应该结束，这种向下传递的模式，需要我们使用Context完成，

_**Context更像一个通信器，可以在主goroutine创建，然后在子goroutine中使用，通知子goroutine结束信息，因此可以直接取代Channel。**_

```go
func CtxWaitGroup() {
    var wg sync.WaitGroup
    wg.Add(2) //在waitgroup中添加job数量
    go func() {
        time.Sleep(2 * time.Second)
        fmt.Println("老财做账")
        wg.Done() // 通知waitgroup本job完成
    }()

    go func() {
        time.Sleep(1 * time.Second)
        fmt.Println("老财审单")
        wg.Done()
    }()
    wg.Wait() //等待waitgroup中的job完成
    fmt.Println("这就是老财们的日常工作")
}
// 运行结果：
// 老财审单
// 老财做账
// 这就是老财们的日常工作
```

```go
// 如何主动通知停止
func CtxStopInitiative() {
    stop := make(chan bool) // 定义一个channel，传递true/false
    go func() { // 创建一个goroutine
        for {
            select {
            case <-stop: // 如果channel接收到停止请求
                fmt.Println("You are fired！")
                return
            default: // 未接收到停止请求前
                fmt.Println("老财工作中")
                time.Sleep(1 * time.Second)
            }
        }
    }()

    time.Sleep(5 * time.Second) // 等待五秒
    fmt.Println("那个老财动作太慢了！开除！")
    stop <- true // 等不下去了，向channel发送一个停止请求
    time.Sleep(5 * time.Second)
}
// 打印结果：
// 老财工作中
// 老财工作中
// 老财工作中
// 老财工作中
// 老财工作中
// 那个老财动作太慢了！开除！
// You are fired！
// 老财滚蛋了！
```

（大包工头准备停工，就要通知分包出去的小包工头先停工，小包工头要停工前要通知接单的施工队长先停工）

#### 改造一下Channel+Select的例子：

```go
// 使用Context
func CtxContext() {
  // context 宣告，context.Background()就是当前函数创建context应用的root goroutine(类似总包开始)
    // WithCancel()函数返回两个东西：
    // 1 一个Context对象，内部携带了一个channel
    // 2 cancel函数，用以发送结束请求
    ctx, cancel := context.WithCancel(context.Background())
    go func() {
        for {
            select {
       // 如果接收到ctx.Done()的反馈表示root goroutine要结束了
       case <-ctx.Done(): 
                fmt.Println("You are fired！")
                return
            default:
                fmt.Println("老财工作中")
                time.Sleep(1 * time.Second)
            }
        }
    }()
    time.Sleep(5 * time.Second)
    fmt.Println("那个老财动作太慢了！开除！")
  // 同样等不下去，但无需通过管道直接调用cancel函数
  // 通过<-ctx.Done()所有子goroutine发送结束信号
    cancel() 
    time.Sleep(1 * time.Second)
    fmt.Println("老财滚蛋了！")
}
// 打印结果：
// 老财工作中
// 老财工作中
// 老财工作中
// 老财工作中
// 老财工作中
// 那个老财动作太慢了！开除！
// You are fired！
// 老财滚蛋了！
```

#### 如果有多个goroutine或goroutine内又有goroutine：

```go
func CtxContextManyGoroutine() {
    // 父goroutine其实创建了三个子goroutine：worker；
    // 而每个worker又创建了自己的goroutine;
    // 仍然在父goroutine创建一个context对象
  // 并将其通过函数参数，分发给所有worker，当父goroutine需要停止时
  // 调用cancel()函数，所有子goroutine会接收到<-ctx.Done()结束消息，作出相应处理
    ctx, cancel := context.WithCancel(context.Background())
    go worker(ctx, "老财1")
    go worker(ctx, "老财2")
    go worker(ctx, "老财3")
    time.Sleep(1 * time.Second) // 主goroutine阻塞1秒，观察三个worker-goroutine运行情况
    fmt.Println("建立财务共享中心，老财全部优化！")
    cancel() // ctx发出了结束信号，代表主goroutine即将结束
    time.Sleep(1 * time.Second)
    fmt.Println("老财们都滚蛋了！")
}

func worker(ctx context.Context, str string) {
    go func() {
        for {
            select {
            case <-ctx.Done(): // worker-goroutine接收到结束信号，打印消息后直接返回结束
                fmt.Println(str, " 你被优化了！")
                return
            default:
                fmt.Println(str, " 工作中")
                time.Sleep(1 * time.Second)
            }
        }
    }()
}
// 运行结果：
// 老财3  工作中
// 老财2  工作中
// 老财1  工作中
// 老财1  工作中
// 老财3  工作中
// 老财2  工作中
// 建立财务共享中心，老财全部优化！
// 老财3  你被优化了！
// 老财1  你被优化了！
// 老财2  你被优化了！
// 老财们都滚蛋了！
```

## `context`的源码解读



一个`Context`接口（包括4中基本方法）

```go
type Context interface{
    Deadline() (deadline time.Time, ok bool)
    // a Done channel for cancellation.
    Done() <-chan struct{}
    Err() error
    // Value returns the value associated with this context for key
    Value(key interface{}) interface{}
}
```

四个结构体分别实现了`Context`接口：_**`emptyCtx`, `cancelCtx`, `timerCtx`, `valueCtx`**_

其中`emptyCtx`, `cancelCtx`分别有对context接口函数的具体实现

而`timerCtx`, `valueCtx`，结构分别为组合了`cancelCtx`结构体和`Context`接口本身

六个具体方法：`Background`, `WithCancel`, `WithDeadLine`, `WithTimeout`, `WithValue`, `TODO`  


### emptyCtx用于初始化

```text
type emptyCtx int // emptyCtx的原型为int
// emtpyCtx对Context接口的实现
func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
    return
}
func (*emptyCtx) Done() <-chan struct{} {
    return nil
}
func (*emptyCtx) Err() error {
    return nil
}
func (*emptyCtx) Value(key interface{}) interface{} {
    return nil
}
```



写Context应用时第一个调用的函数时`Background()`：

```text
var (
    // new是用来分配内存的内建函数，区别于其他语言中new会初始化内存，golang中的new只会将内存置零
    background = new(emptyCtx)
)
func Background() Context {
    return background // 返回初始化的内存地址
}
```

### cancelCtx

`WithCancel(context.Background())`：

```text
// 我们将Background()创建的Context类型传递到WithCancel函数中来，也就是这里的parent形参
// WithCancel返回一个Context类型和一个CancelFunc类型
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    if parent == nil {
        panic("cannot create context from nil parent")
    }
    c := newCancelCtx(parent)
    propagateCancel(parent, &c)
    return &c, func() { c.cancel(true, Canceled) }
}
```

首先判断传进来的`parent`是不是存在，

* 不存在`panic`报错；
* 如果存在，调用`newCancelCtx(parent)`函数，将结果返回赋值给变量`c`，
* 然后调用`propagateCancel(parent, &c)`, 
* 最终返回`&c, func() { c.cancel(true, Canceled) }`



从这里衍生出了几个新的类型和函数，他们到底有什么功能实现了什么作用呢，我们罗列一下：

* 作为函数返回值类型之一的`CancelFunc`类型
* `newCancelCtx`函数
* `propagateCancel`函数
* `return`语句中回调函数里的：`cancel`函数和`Canceled`变量。

首先来看看`CancelFunc`的定义：

```text
// CancelFunc定义的是一个函数类型，相当于传递一个“传递参数为空+返回值为空”的函数指针
type CancelFunc func()
```

再来看看`newCancelCtx`函数：其主要作用为创建了另一个名为`cancelCtx`的对象，并将一个实现了`Context`接口的对象包裹其中，作为它的`父Context节点指针`，这样一来形成了一条`Context链`

```text
// newCancelCtx函数返回一个初始化的cancelCtx类型.
// 作用是传入一个实现Context接口的类型，然后将其包裹为cancelCtx对象并返回
func newCancelCtx(parent Context) cancelCtx {
    return cancelCtx{Context: parent}
}
```

因此我们又引申出了`cancelCtx`类型，我们知道了`WithCancel()`函数中`c := newCancelCtx(parent)`中的变量`c`最终赋值的类型是`cancelCtx`，是重新打包后的`parent`。

  
`cancelCtx`类型的在源码中的定义：

```text
// cancelCtx结构体是可以被取消的，一旦本身取消
// 其children字段中所有实现了canceler接口的Context类型都会被取消
type cancelCtx struct {
    Context
    mu       sync.Mutex            // 互斥锁用以保护以下属性字段
    done     chan struct{}         // done用于获取取消通知
    children map[canceler]struct{} // 一个map类型，存储以当前节点为root的所有可取消的Context
    err      error                 // 存储取消时指定的错误信息
}

// canceler接口定义如下，而cancelCtx类型同样实现了canceler接口
type canceler interface {
    // cancel取消函数，传入两个参数，一是bool类型，明确是否要从父节点删除，二是一个error参数
    // cancel函数的主要作用是关闭cancelCtx.done的通道
    // 如果removeFromParent为false，将遍历cancelCtx.children，将map中的每个元素都递归调用cancel函数
    // 如果removeFromParent为true，除了false时的操作外，
    // 还会直接调用delete函数删除cancelCtx.children中的数据
    cancel(removeFromParent bool, err error)
    // Done取消通知，返回一个通道类型
    Done() <-chan struct{}
}
```

变量`c`的类型明确之后我们就可以来看看`propagateCancel`函数，源码中`propagateCancel(parent, &c)`的调用。在我们的代码中，实际上传递了两个不同的`Context`类型。`parent`是一个`emptyContext`，而`c`是`cancelCtx`



