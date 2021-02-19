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

别忘了我们在哪儿：

```text
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    if parent == nil {
        panic("cannot create context from nil parent")
    } // √
    c := newCancelCtx(parent) // √
    propagateCancel(parent, &c) // <-- 我们在这里
    return &c, func() { c.cancel(true, Canceled) } // <--未到达
}
```

`propagateCancel`函数眼见地复杂，但其最终目的很简单：把当前`child canceler`挂载到`parent Context`的`children map`里，以`child`指针本身作为`map`的`key`值。而实际上`canceler`也必然是四大`context`结构体中的某一种

最终要形成的效果如下:

![](https://pic3.zhimg.com/80/v2-9f7f80ebfbc57934f3c6b881df67d98a_1440w.jpg)

解读这部分代码，我们可以将其拆解为三部分：

```text
// propagateCancel函数用以确保parent被取消时，child同样被取消
func propagateCancel(parent Context, child canceler) {
    // --------- 第一部分 --------------------
    done := parent.Done()
    if done == nil {
        return // parent is never canceled
    }
    // --------- 第二部分 --------------------
    select {
    case <-done:
        // parent is already canceled
        child.cancel(false, parent.Err())
        return
    default:
    }
    // --------- 第三部分 --------------------
    if p, ok := parentCancelCtx(parent); ok {
        p.mu.Lock()
        if p.err != nil {
            // parent has already been canceled
            child.cancel(false, p.err)
        } else {
            if p.children == nil {
                p.children = make(map[canceler]struct{})
            }
            p.children[child] = struct{}{}
        }
        p.mu.Unlock()
    } else {
        atomic.AddInt32(&goroutines, +1)
        go func() {
            select {
            case <-parent.Done():
                child.cancel(false, parent.Err())
            case <-child.Done():
            }
        }()
    }
}
```

1 获取`parent.Done()`，赋值给**本地变量**`done`，判断**本地变量**`done`是否为`nil`

\`\`

```text

// --------- 第一部分 --------------------  
done := parent.Done()
if done == nil {
    return // parent is never canceled
}
```

* 只有当`parent`的实际类型为`emptyCtx`的时候，调用`emptyContext.Done()`会无条件返回`nil`。按照我们写的示例代码：`ctx, cancel := context.WithCancel(context.Background())`这里只传递了一个`context.Background()`，也就是`emptyCtx`，那么`propagateCancel`代码到这里就结束了

```text
// 源码如下
func (*emptyCtx) Done() <-chan struct{} {
    return nil
}
```

* 如果是`cancelCtx`类型，那么会先判断是否当前实例中的`done`属性是否为空，如果为空就初始化，不为空直接返回`cancelCtx.done`

```text
// cancelCtx对接口函数Done()的实现源码如下
func (c *cancelCtx) Done() <-chan struct{} {
    c.mu.Lock() // 上锁
    if c.done == nil {
        c.done = make(chan struct{})
    }
    d := c.done
    c.mu.Unlock()
    return d
}
```

2`select+channel`组合。如果接收到`done`通道传来的信息，就说明当前`parent context`已经被取消了，那么作为准备挂载到其身上的`child`变量（函数参数）也要被取消，调用了子节点的`cancel`函数，然后返回。如果没有接收到`done`传递来的消息，那么继续往下运行

```text
// --------- 第二部分 --------------------
select {
    case <-done:
    // parent is already canceled
    child.cancel(false, parent.Err())
    return
    default:
}
```

3 第三部分本质上是一个`if-else`的判断，所以最关键的在于第一条语句:

`if p, ok := parentCancelCtx(parent); ok`

```text
// --------- 第三部分 --------------------
if p, ok := parentCancelCtx(parent); ok {
    p.mu.Lock()
    if p.err != nil {
        // parent has already been canceled
        child.cancel(false, p.err)
    } else {
        if p.children == nil {
            p.children = make(map[canceler]struct{})
        }
        p.children[child] = struct{}{}
    }
    p.mu.Unlock()
} else {
    atomic.AddInt32(&goroutines, +1)
    go func() {
        select {
            case <-parent.Done():
            child.cancel(false, parent.Err())
            case <-child.Done():
            }
    }()
}
```

这里`parentCancelCtx(parent)` 函数返回两个值：一个`cancelCtx`指针和一个布尔值，其实该函数的作用是一个校验，并将获取`parent`指针并将其强转为`cancelCtx`类型然后赋值给本地变量`p`。用布尔值`ok`判断这个过程是否成功



为什么要这么做？

这里我们回看一下`propagateCancel`函数原型： `propagateCancel(parent Context, child canceler)`

要注意一点，从形参传递过来的`parent`变量类型是`Context`，而`child`是`canceler`

`Context`和`canceler`都是接口，他们可以实际上指向不同的`Ctx`类型，比如`parent`有可能是`timerCtx`，而`child`是`cancelCtx`类型`parentCancelCtx(parent)`类型强制转化为`cancelCtx`指针：

比如，`timerCtx`结构体内容为：

```text
// timerCtx结构体组合了一个cancelCtx结构体作为成员属性
// 因此当我们采用.(*cancelCtx)强转时，原本指向整个timerCtx的指针指向了timer.cancelCtx
// 从作用而言timerCtx和cancelCtx本质是相同的，只是timerCtx附带了两个额外属性
type timerCtx struct {
    cancelCtx
    timer *time.Timer
    deadline time.Time
}
```

`parentCancelCtx(parent)`源码为：

```text
func parentCancelCtx(parent Context) (*cancelCtx, bool) {
    // 先判断parent.Done()的管道是否存在，如果不存在或关闭，则返回nil和false
    done := parent.Done()
    if done == closedchan || done == nil {
        return nil, false
    }
    // 如果管道已经初始化，则进一步根据key校对value
    // 这里cancelCtxKey是一个int类型的全局变量，专门用来根据其获得parent指向的对象指针
    p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)
    // 如果获取失败，则返回nil和false
    if !ok {
        return nil, false
    }
    // 如果获取成功，先加锁，再判断重新获取的对象当前的done属性和原先的done属性是否相同
    // 因为有可能在这段期间内，done发送了取消请求
    p.mu.Lock()
    ok = p.done == done // 判断done的状态是否相同
    p.mu.Unlock()
    // 如果状态不一，则返回nil和false
    if !ok {
        return nil, false
    }
    // 状态相同返回cancelCtx指针和true
    return p, true
}
```

从上面这段代码看，我们其实相当于传进去`parent`对象返回一个指向`parent`对象的`cancelCtx`指针，似乎多此一举，然而并没有那么简单，整段代码的核心在于：

`p, ok := parent.Value(&cancelCtxKey).(*cancelCtx)`

我们可以看一下`cancelCtx`实现`Value`方法的源代码:

```text
// 注意Value方法是Context接口中的四大方法之一
var cancelCtxKey int // 全局变量作为cancelCtx的key，这样保持接口调用一致性
// 从这段代码中我们可以看出如果我们调用的parent实际指向的类型是cancelCtx或者组合了cancelCtx的timerCtx
// 则直接返回指向parent的指针
// 如果此时是valueCtx类型，那么valueCtx结构体中自带的key与传递过来的cancelCtxKey先进行对比
// 如果刚好对的上，则返回valueCtx.val，也就是挂载的数据
// 然而一般而言是对不上的，进而返回valueCtx.Context.Value(key)，相当于继续向上查找
// 向当于顺着context链不断匹配知道搜索到符合key值得context节点
func (c *cancelCtx) Value(key interface{}) interface{} {
    if key == &cancelCtxKey {
        return c
    }
    return c.Context.Value(key)
}
// valueCtx的结构如下，封装了一个Context接口作为自己的成员属性
type valueCtx struct {
    Context
    key, val interface{}
}
func (c *valueCtx) Value(key interface{}) interface{} {
    if c.key == key {
        return c.val
    }
    return c.Context.Value(key)
}
```

讲了这么多终于把第三部分第一行代码讲完了：

`if p, ok := parentCancelCtx(parent); ok // <-- √`

但是其余代码比较简单，总体而言遵循一个`if - else`结构

```text
if ok == true { // p获取成功的情况
    // ...
} else { // p获取不成功的情况
    // ...
}
```

接下来我们再来看如果`p`获取成功的情况，`ok == true`：

```text
p.mu.Lock() // 加锁
if p.err != nil { // 如果指针存在err则说明，parent指针指向得cancelCtx已经被取消
    // 其子节点一并取消
    child.cancel(false, p.err)
} else { // 如果p没有取消
    // 如果p指向的cancelCtx还没有挂载过chirdren，则初始化children
    if p.children == nil {
        p.children = make(map[canceler]struct{})
    }
    // 添加当前child到p.chirdren(一个map对象)
    p.children[child] = struct{}{}
}
p.mu.Unlock() // 解锁结束
```

如果`p`获取不成功，`ok == false`

```text
atomic.AddInt32(&goroutines, +1) // 原子操作goroutines是一个int类型的全局变量，对其+1
// 生成一个goroutine，同时监控parent和child是否已经接收到了Done发出的取消信号
// 如果parent有的话，child需要条用cancel函数取消，并发送一个parent.Err消息
go func() { 
    select {
        case <-parent.Done():
        child.cancel(false, parent.Err())
        case <-child.Done():
        }
}()
```

好的，我们终于到达`WithCancel`函数的最后一条代码了，别忘了我们现在在哪儿：

```text
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    if parent == nil {
        panic("cannot create context from nil parent")
    } // √
    c := newCancelCtx(parent) // √
    propagateCancel(parent, &c) // √
    return &c, func() { c.cancel(true, Canceled) } // <-- 我们在这里
}
```

现在困扰我们的只有最后两个问题：`Canceled`变量和`cancel`函数

先看`Canceled`类型：

```text
// Canceled变量很简单就是一个全局变量存储一个定义好的错误提示消息
// 当我们取消时随同cancel函数发送
var Canceled = errors.New("context canceled")
```

至于`cancel`函数则复杂得多：

我们可以将其拆解为三部分：

* 第一部分：两个`err`的校验：传递进来的形参`err`和当前`context`对象中`err`属性的校验，前者必须不为`nil`，后者必须等于`nil`
* 第二部分：将传递进来的形参赋值`context`对象中的`err`属性，同时`close(context.done)`，关闭所有子`context`
* 第三部分，根据`removeFromParent`形参判断是否要在当前`context`对象的`父context`中删除自己

```text
// closedchan是一个全局变量，初始化了一个通道类型
// 但在整个context.go的init()函数中被直接关闭
// closedchan在cancel函数中令cancelCtx.done保持关闭
var closedchan = make(chan struct{})
func init() {
    close(closedchan)
}
// 整个cancel函数的核心作用：close(cancelCtx.done)，这样通过管道就发出了cancel信号
// 那些在select中监控 <-context.done的goroutine就会接收到信息，作出相应的行为
// cancel函数传递两个参数一个是布尔类型的removeFromParent，用于控制是否从父context对象中删除本子节点
// 还有一个是error对象，用以传递error信息
func (c *cancelCtx) cancel(removeFromParent bool, err error) {
  // ---------------第一部分------------------------------------
  // 1.1 对传递进来的形参err的校验
  // 当调用cancel时，error是必须传递的参数，所以error = nil 会引发panic
    if err == nil {
        panic("context: internal error: missing cancel error")
    }
  // 1.2 对当前context对象上的err属性进行校验
    c.mu.Lock() 
  // 如果当前context对象中本身的err不为nil
  // 代表了该对象已经被取消，其err标记已经更新了，直接退出
  // 如果当前context对象中本身的err等于nil，则继续运行
    if c.err != nil { 
        c.mu.Unlock()
        return 
    }
  // ---------------第二部分------------------------------------
  // 如果当前context对象中本身的err==nil
  // 则代表当前context，并没有发出过取消信号，需要我们手动取消
  // 2.1 重制context.err属性，将其打上err标记
    c.err = err
  // 2.2 先判断当前context.done是否等于nil，也就是没有初始化过
  // 如果是的话就直接将关闭的closedchan赋值给c.done
  // 否则的话，直接调用close函数进行关闭
    if c.done == nil { 
        c.done = closedchan
    } else {
        close(c.done)
    }
  // 2.3 依次遍历当前context.children，对map所有child都调用cancel函数
  // 父context关闭所有子context必须也关闭
    for child := range c.children {
        // NOTE: acquiring the child's lock while holding parent's lock.
        child.cancel(false, err)
    }
    c.children = nil
    c.mu.Unlock()
    // ---------------第三部分------------------------------------
  // 如果removeFromParent == true 则要到父context中删除挂载的自己
    if removeFromParent {
        removeChild(c.Context, c)
    }
}
```

`removeChild`函数：顺着`parent`指针到`父context`的`children属性`中调用`delete(map,key)`函数删除自己

```text
// 仍然要先调用parentCancelCtx(parent) 确定parent指针
// 同时删除过程中要加锁
func removeChild(parent Context, child canceler) {
    p, ok := parentCancelCtx(parent)
    if !ok {
        return
    }
    p.mu.Lock()
    if p.children != nil {
        delete(p.children, child)
    }
    p.mu.Unlock()
}
```

最后一条代码我们也讲完了

```text
func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
    if parent == nil {
        panic("cannot create context from nil parent")
    } // √
    c := newCancelCtx(parent) // √
    propagateCancel(parent, &c) // √
    return &c, func() { c.cancel(true, Canceled) } // √
}
```

尽管我们只讲了`WithCancel`函数把整个context.go的作用基本已经缕了一遍。

于是我们再回到最早的应用函数那里，来想想具体发生了什么：

```text
// 更进一步 如果有多个goroutine或goroutine内又有goroutine
func CtxContextManyGoroutine() {
    // WithCancel函数创建了一个withContext对象，以及一个cancel函数句柄
  // ctx被分发给了每个worker-goroutine
  // cancel被留在了主goroutine随时发送结束请求
  // 一旦cancel()被调用，则所有worker-goroutine
  // 会通过context对象内部的channel(context.done)接收到结束信号
  // 这种模式听起来有点像给每个worker配了个只能接收信息的BB机，由工头统一发送结束信息
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
```

我们现在再改修改一下程序，假设我们创建一个根`ctx0`，在他下面挂载三个子`ctx1~3`，分别分发给3个子goruntine

我们先单独调用`cancel1`和`cancel3`函数分别终止`ctx1`和`ctx3`，然后调用`cancel0`直接发出根`ctx0`的终止请求，这样的话按照源代码中`cancel`函数的定义，会遍历`ctx0`挂载的子context对象，统一进行终止处理。

```text
func CtxManyContexts() {
    // 分别定义4个context对象，一个根ctx0，三个子ctx1-3
    ctx0, cancel0 := context.WithCancel(context.Background())
    ctx1, cancel1 := context.WithCancel(ctx0)
    ctx2, _ := context.WithCancel(ctx0)
    ctx3, cancel3 := context.WithCancel(ctx0)
  // 分发给三个任务
    go bookeeper(ctx1, "老财1")
    go bookeeper(ctx2, "老财2")
    go bookeeper(ctx3, "老财3")
    time.Sleep(1 * time.Second) // 主goroutine阻塞1秒，观察三个bookeeper-goroutine运行情况
  // 发出指令：结束ctx1和ctx3
    fmt.Println("老财1和老财3优化！")
    cancel1() // 调用ctx1的cancel函数，结束ctx1
    cancel3() // 调用ctx3的cancel函数，结束ctx3
    time.Sleep(2 * time.Second)
    fmt.Println("建立财务共享中心，剩下的老财全部优化！")
  // 由于到这里ctx2还未结束，但ctx0发出了结束信号，代表主goroutine即将结束
  // ctx2的cancel被自动调用，完成了整个context树的结束请求
    cancel0() 
    time.Sleep(2 * time.Second)
    fmt.Println("老财们都滚蛋了！")
}
func bookeeper(ctx context.Context, str string) {
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
// 输出结果：
// 老财1  工作中
// 老财3  工作中
// 老财2  工作中
// 老财2  工作中
// 老财3  工作中
// 老财1和老财3优化！
// 老财1  你被优化了！
// 老财3  你被优化了！
// 老财2  工作中
// 建立财务共享中心，剩下的老财全部优化！
// 老财2  你被优化了！
// 老财们都滚蛋了！
```

好，到此为止我们大致上已经对`cancelCtx`有了基本了解，剩下的`valueCtx`和`timerCtx`只是它的扩充和延伸，就会简单许多

#### 3.3 valueCxt

这里只讲一下，valueCtx的主要功能：

基本定义如下，我们看到就是一个`Context`接口和`key:val`形式的键值对，主要用于实现`withValue`函数，重载了`Context接口`中的`Value函数`，用于根据`key`查找对应的`value`

```text
type valueCtx struct {
    Context
    key, val interface{}
}

func WithValue(parent Context, key, val interface{}) Context {
  // parent不能为nil，相当于valueCtx必须挂载载其他context之下
    if parent == nil {
        panic("cannot create context from nil parent")
    }
  // key不能为空，这是必须的
    if key == nil {
        panic("nil key")
    }
  // key同时必须具有可比较性
    if !reflectlite.TypeOf(key).Comparable() {
        panic("key is not comparable")
    }
  // 最终返回创建的valueCtx对象的地址
    return &valueCtx{parent, key, val}
}

// 根据key查找value
func (c *valueCtx) Value(key interface{}) interface{} {
    if c.key == key {
        return c.val
    }
    return c.Context.Value(key)
}
```

#### 3.4 timerCrx

这里只讲一下，`timerCrx`的主要功能：

`timerCtx`结构体是在`cancelCtx`结构体的基础上封装了一个定时器`timer`和截止时间`deadline`，由于cancelCtx存在，所以本质上同样受`cancelCtx.mu锁`的保护

```text
type timerCtx struct {
    cancelCtx
    timer *time.Timer 
    deadline time.Time
}
// Deadline函数，可以针对当前timerCtx设定一个截止时间
func (c *timerCtx) Deadline() (deadline time.Time, ok bool) {
    return c.deadline, true
}
// timerCtx.cancel函数
// 前半部分是对cancelCtx.cancel函数的调用
// 后半部分是停止当前timerCtx对象的计数器
func (c *timerCtx) cancel(removeFromParent bool, err error) {
    c.cancelCtx.cancel(false, err)
    if removeFromParent {
        // Remove this timerCtx from its parent cancelCtx's children.
        removeChild(c.cancelCtx.Context, c)
    }
    c.mu.Lock() // 上锁
  // 停止计数器
    if c.timer != nil {
        c.timer.Stop()
        c.timer = nil
    }
    c.mu.Unlock()
}
```

由`WithDeadline`和`WithTimeout`都可以创建`timerCtx`，基本逻辑同`WithCancel`，但`WithDeadline`需要额外传递一个`time.Time`类型的`deadline`，这是我们设定的具体截止时间。

`WithTimeout`函数其实就是调用`WithDeadline`，不同的是`WithDeadline`需要我们传入的是具体的结束时间，而`WithTimeout`需要我们传入的是具体多少时间后结束。

`WithDeadline`函数的核心在于：

`c.timer = time.AfterFunc(dur, func() { c.cancel(true, DeadlineExceeded)})`

`dur`是到达我们设定的`deadline`需要多少时间，`timerCtx`中的定时器`timerCtx.timer`会运行一个`AfterFunc`函数，当我们时间到达`deadline`后自动触发`cancel`程序并发送一个`DeadlineExceeded`信息

```text
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc) {
    if parent == nil {
        panic("cannot create context from nil parent")
    }
  // 父context的截止时间早于了我们对当前context设定的deadline，相当于这个timerCtx没必要创建
  // 直接返回父context
    if cur, ok := parent.Deadline(); ok && cur.Before(d) {
        // The current deadline is already sooner than the new one.
        return WithCancel(parent)
    }
  // 创建timerCtx，我们设定的截止时间用作timerCtx中的deadline属性
    c := &timerCtx{
        cancelCtx: newCancelCtx(parent),
        deadline:  d,
    }
    propagateCancel(parent, c)
  // 时间校验，当前的时间是不是早就过了我们设定的截止时间
  // 直接调用cancel函数，并发送一个DeadlineExceeded错误信息
    dur := time.Until(d) // 计算到达deadline需要多少时间
    if dur <= 0 {
        c.cancel(true, DeadlineExceeded) 
        return c, func() { c.cancel(false, Canceled) }
    }
  
    c.mu.Lock() // 加锁
    defer c.mu.Unlock()
  // timerCtx.timer 设定一个AfterFunc函数，延迟多少dur时间之后自动运行cancel函数
    if c.err == nil {
        c.timer = time.AfterFunc(dur, func() {
            c.cancel(true, DeadlineExceeded)
        })
    }
    return c, func() { c.cancel(true, Canceled) }
}
// WithTimeout函数其实就是调用WithDeadline
// 不同的是WithDeadline需要我们传入的是具体的结束时间，
// 而WithTimeout需要我们传入的是具体多少时间后结束
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
    return WithDeadline(parent, time.Now().Add(timeout))
}

// DeadlineExceeded是一个error类型的全局变量，定义了截止时间超过的错误信息
var DeadlineExceeded error = deadlineExceededError{}
type deadlineExceededError struct{}
func (deadlineExceededError) Error() string   { return "context deadline exceeded" }
func (deadlineExceededError) Timeout() bool   { return true }
func (deadlineExceededError) Temporary() bool { return true }
```

#### 3.5 改造最终的案例

我们最后大致梳理一下：![](https://pic2.zhimg.com/80/v2-1a2fc1928a675f6493a7dfd44c659591_1440w.jpg)

然后再改造一下我们的最终案例：

我们仍然同前创建一个`ctx0`作为`根context`，`ctx1-3`挂载到ctx0下面，不同的是ctx1是cancelCtx类型，而ctx2我们设定为`valueCtx`类型，给它赋予一个键值对，`ctx3`是一个`timerCtx`类型，给予一个自动结束时间，我们主动调用`ctx1`的`cancel1`函数主动结束`ctx1`，然后坐等`ctx3`到达截止时间自动结束，最后调用`ctx0`的`cancel0`函数结束整个`context树`。

```text
func CtxFinalContexts() {
    ctx0, cancel0 := context.WithCancel(context.Background())
    ctx1, cancel1 := context.WithCancel(ctx0)
    ctx2 := context.WithValue(ctx0, "老财2", "老财2：十年老会计 信奉越老越吃香 没经验的才被优化！")
    ctx3, _ := context.WithTimeout(ctx0, 3*time.Second)
    go bookeeper(ctx1, "老财1")
    go bookeeper(ctx2, "老财2")
    go bookeeper(ctx3, "老财3")
    time.Sleep(2 * time.Second) // 主goroutine阻塞1秒，观察三个bookeeper-goroutine运行情况
    fmt.Println("老板：老财1，优化！")
    fmt.Println(ctx2.Value("老财2"))
    cancel1() // 调用ctx1的cancel函数
    time.Sleep(1 * time.Second)
    fmt.Println("老板：老财3，合同到期自动清退！") // 3秒后ctx3结束
    time.Sleep(2 * time.Second)
    fmt.Println("老板：建立财务共享中心，剩下的老财全部优化！")
    cancel0() // ctx0发出了结束信号，代表主goroutine即将结束，ctx2自动结束
    time.Sleep(3 * time.Second)
    fmt.Println("老板：老财们都滚蛋了！")
}
// 运行结果：
// 老财3  工作中
// 老财1  工作中
// 老财2  工作中
// 老财2  工作中
// 老财1  工作中
// 老财3  工作中
// 老板：老财1，优化！
// 老财2：十年老会计 信奉越老越吃香 没经验的才被优化！
// 老财3  工作中
// 老财1  你被优化了！
// 老财2  工作中
// 老板：老财3，合同到期自动清退！
// 老财3  你的临时工合同到期了！
// 老财2  工作中
// 老财2  工作中
// 老板：建立财务共享中心，剩下的老财全部优化！
// 老财2  你被优化了！
// 老板：老财们都滚蛋了！
```









