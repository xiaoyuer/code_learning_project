# Go泛型

{% embed url="https://s3.amazonaws.com/bizzabo.file.upload/p1oqhNDwQq2kb6d78TDA\_R%20Griesemer%20-%20Typing%20-Generic-%20Go.pdf" %}

{% embed url="https://www.youtube.com/watch?v=TborQFPY2IM&ab\_channel=GopherAcademy" %}

\

Go语言技术大会[GopherCon 2020](https://www.gophercon.com/)上，Go语言之父之一的[Robert Griesemer](https://github.com/griesemer)为全世界Gopher们带来了本次大会最重量级的演讲**“Typing \[Generic\] Go”**

Robert Griesemer向Gopher们介绍了自从今年中旬在Go官网发表文章[“The Next Step for Generics”](https://blog.golang.org/generics-next-step)以来[Go泛型\(Go Generics\)技术草案](https://tonybai.com/2020/06/18/the-go-generics-is-coming-and-supported-in-go-1-17-at-the-earliest/)的最新变化，并详细介绍了类型参数\(type parameter\)是如何满足Go现有的类型系统的，以及Go编译器是如何对Go泛型代码进行类型检查的。



#### 一. 预备知识

为了更好地理解Robert Griesemer的讲解，这里先带着大家回顾一下Go generics技术草案演化史。

  
图：Go泛型技术草案演化时间线

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-evolution-timeline.png)

* 2017年7月，Go核心团队领军人物[Russ Cox](https://github.com/rsc/)在Gophercon 2017大会上发表演讲“[Toward Go 2](https://blog.golang.org/toward-go2)”，正式吹响Go向下一个阶段演化的号角；
* 2018年8月，在Gophercon 2018大会结束后不久，Go核心团队发布了Go2 draft proposal，这里面涵盖了由Ian Lance Taylor和Robert Griesemer操刀主写的Go泛型的[第一版draft proposal](https://github.com/golang/proposal/blob/00fd2f65291738699cd265243559718f1fb7d8c5/design/go2draft-contracts.md)。这版草案引入了**contract关键字**来定义泛型类型参数\(type parameter\)的约束、类型参数放在普通函数参数列表前面的**小括号**中，并用type关键字声明：

```text
// 第一版泛型技术草案中的典型泛型语法

contract stringer(x T) {
    var s string = x.String()
}

func Stringify(type T stringer)(s []T) (ret []string) {

}
```

* 2019年7月，Ian Lance Taylor在GopherCon 2019大会上发表演讲[“Why Generics?”](https://blog.golang.org/why-generics)，并更新了[泛型的技术草案](https://github.com/golang/proposal/blob/4a54a00950b56dd0096482d0edae46969d7432a6/design/go2draft-contracts.md)，简化了contract的语法设计：

```text
// 简化后的contract语法如下：

contract stringer(T) {
    T String() string
}
```

* 2020年6月，《Featherweight Go》论文发表在arxiv.org上，该论文缘于Rob Pike向著名计算机科学家、函数语言专家、[Haskell语言](https://tonybai.com/tag/haskell)的设计者之一、Java泛型的设计者PHILIP WADLER发出的一次邀请，希望PHILIP WADLER帮助Go核心团队解决Go语言的泛型扩展问题：

  
图：Rob Pike向PHILIP WADLER发出的邀请

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-4.png)

而这篇论文则是对这次邀请的回应。这篇论文为Go语言的一个最小语法子集设计了泛型语法Featherweight Generic Go\(FGG\)，并成功地给出了FGG到Feighterweight Go\(FG\)的可行性实现的形式化证明。

该篇论文采用monomorphisation\(单态\)的实现，而非Java使用的擦触法\(Erasure\)，这样的好处之一是如果代码中没有使用任何泛型抽象，程序的运行时不会因支持泛型而承担额外的消耗。

该论文的形式化证明给Go团队带来了信心，也是的Go团队在一些语法问题上达成更广泛的一致。

  
图：Robert Griesemer表达了对该论文团队的感谢

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-5.png)

* 2020.6月末，Ian Lance Taylor和Robert Griesemer在Go官方博客发表了文章[《The Next Step for Generics》](http://blog.golang.org/generics-next-step)，介绍了Go泛型工作的最新进展。Go团队放弃了之前的技术草案，并重新编写了[一个新草案](https://github.com/golang/proposal/blob/d44c4ded9c1a13dcf715ac641ce760170fbbcf64/design/go2draft-type-parameters.md)。在这份新技术方案中，Go团队放弃了引入contract关键字作为泛型类型参数的约束，而采用扩展后的interface来替代contract。这样上面的Stringify函数就可以写成如下形式：

```text
type Stringer interface {
    String() string
}

func Stringify(type T Stringer)(s []T) (ret []string) {
    ... ...
}
```

同时，Go团队还推出了可以[在线试验Go泛型语法的playground](https://go2goplay.golang.org/)：https://go2goplay.golang.org，这样gopher们可以直观体验新语法，并给出自己的意见反馈。

* 2020年11月的GopherCon 2020大会，Griesemer与全世界Gopher同步了Go泛型的最新进展和roadmap，在最新的技术草案版本中，小括号被方括号取代，类型参数前面的type关键字也不再需要了：

```text
func Stringify[T Stringer](s []T) (ret []string) {
    ... ...
}
```

go2goplay.golang.org也支持了方括号语法，gopher可以在线体验。

**下面我们就来看看Griesemer对最新Go泛型技术草案的详细讲解**。

#### 二. 类型参数\(Type parameters\)技术草案详解

这版草案与2019年中旬发布的草案的最大变动就是**使用interface而不是contract来表达对类型参数的约束**。

该版设计的主要特性：

* 类型参数\(Type parameters\) – 一种将类型或函数进行参数化的机制
* 约束\(Constraints\) – 一种表达对类型参数的约束的机制
* 类型推导\(Type inference，可选\)

**普通函数参数列表 vs. 泛型函数的类型参数列表**

我们知道，普通函数的参数列表是这样的：

```text
(x, y aType, z anotherType)
```

* x, y, z是形参\(parameter\)的名字，即变量；
* aType，anotherType是形参的类型，即类型。

我们再来看一下类型参数\(type parameter\)列表：

```text
[P, Q aConstraint, R anotherConstraint]
```

* P，Q，R是类型形参的名字，即类型；
* aConstraint，anotherConstraint代表类型参数的约束\(constraint\)，可以理解为一种元类型\(meta-type，即修饰类型的类型\)。

> 注：按惯例，类型参数\(type parameter\)的名字都是头母大写的。

**为什么需要类型参数\(type parameter\)**

我们先来看一下当前Go语言标准库中提供的排序方案：

```text
// $GOROOT/src/sort/sort.go
type Interface interface {
        Len() int
        Less(i, j int) bool
        Swap(i, j int)
}

func Sort(data Interface) {
    ... ...
}
```

为了应用这个排序函数Sort，我们需要让被排序的类型实现sort.Interface接口，就像下面例子中这样：

```text
type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
        sl := IntSlice([]int{89, 14, 8, 9, 17, 56, 95, 3})
        fmt.Println(sl)
        sort.Sort(sl)
        fmt.Println(sl)
}
```

这真是我们想要的实现方式吗？我们真正需要的是这样的：

```text
func Sort(list []Elem)

// 使用
var myList = []Elem{...}
Sort(myList)
```

解决办法：**使用type parameter**\(类型参数或叫做参数化的类型，将类型作为参数传递\)：

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-6.png)  
图：使用类型参数的Sort

**约束\(constraints\)**

约束\(constraint\)规定了一个类型实参\(type argument\)必须满足的条件要求。而在泛型Go中，**我们使用interface来定义约束**。

如果某个类型实现了某个约束\(规定的所有条件要求\)，那么它就是一个合法的类型实参。

下面是一个泛型版本的Sort函数：

```text
func Sort[Elem interface{ Less(y Elem) bool }](list []Elem)
```

我们看到上面函数Sort的类型形参\(type parameter\)Elem的约束是一个interface，这样传入的类型实参\(type argument\)只要实现了该接口即可。

约束的定义中也可以引用类型形参，比如下面这个泛型函数：

  
图：约束的定义中引用类型形参

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-7.png)

**类型形参的声明与作用域**

  
图：类型参数的声明与作用域

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-8.png)

类型参数的作用域始于**\[**，终于泛型函数的函数体结尾或泛型类型的声明结尾。

**泛型的类型具化与类型检查**

下面是一个使用泛型版本Sort函数的例子：

```text
func Sort[Elem interface{ Less(y Elem) bool }](list []Elem)

type book struct{…}
func (x book) Less(y book) bool {…}
var bookshelf []book
…
Sort[book](bookshelf) // 泛型函数调用
```

上面的泛型函数调用`Sort[book](bookshelf)`将分成两个阶段：

1. 具化\(instantiation\)

形象点说，具化\(instantiation\)就**好比一家生产“排序机器”的工厂根据要排序的对象的类型将这样的机器生产出来的过程**。以上面的例子来说，整个具化过程如下：

* 工厂接单：**Sort\[book\]**，发现要排序的对象类型为book；
* 模具检查与匹配：检查book类型是否满足模具的约束要求\(即是否实现了Less方法\)，如满足，则将其作为类型实参替换Sort函数中的类型形参，结果为**Sort\[book interface{ Less\(y book\) bool }\]**；
* 生产机器：将泛型函数Sort具化为一个**新函数**，这里将其起名为**booksort**，其函数原型为**func\(\[\]book\)**。本质上**booksort := Sort\[book\]**。

1. 调用\(invocation\)

一旦“排序机器”被生产出来，那么它就可以对目标对象进行排序了，这和普通的函数调用没有区别。这里就相当于调用booksort\(bookshelf\)，整个过程只需检查传入的函数实参\(bookshelf\)的类型与booksort函数原型中的形参类型\(\[\]book\)是否匹配即可。

用伪代码来表述上面两个过程如下：

```text
Sort[book](bookshelf)

<=>

具化：booksort := Sort[book]
调用：booksort(bookshelf)
```

**泛型类型**

除了函数可以携带类型参数变身为“泛型函数”外，类型也可以拥有类型参数而化身为“泛型类型”：

```text
type Lesser[T any] interface{
   Less(y T) bool
}
```

上面代码中的**any**代表没有任何约束，等价于interface{}。

**泛型类型的类型参数的声明与作用域范围**

泛型类型的类型参数的声明方式如下，类型参数的作用域范围也同见下图：

  
图：泛型类型的类型参数的声明与作用域

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-9.png)

**用泛型类型改造Sort**

用泛型类型定义一个具名的约束条件- Lesser接口类型：

```text
type Lesser[T any] interface{
   Less(y T) bool
}
```

使用Lesser\[T\]作为约束的Sort函数可以这样写：

```text
func Sort[Elem Lesser[Elem]](list []Elem)
```

> 注意：任何泛型函数或泛型类型在使用前都必须先“具化\(instantiation\)”。

我们再来看看Sort函数的内部实现：

```text
func Sort[Elem Lesser[Elem]](list []Elem) {
    ...
    var i, j int
    ...
    if list[i].Less(List[j]) {
        ...
    }
    ...
}
```

* 这里的list\[i\]和list\[j\]的类型是Elem；
* Elem不是一个接口类型，它是泛型函数\(Sort\)的类型参数，Lesser\[Elem\]是作为类型参数的约束而存在的，不要与函数常规参数列表混淆。

再次强调：**类型参数是一个真实的类型，不是一个接口类型\(变量\)，当然我们可以使用一个接口类型作为类型实参来具化一个泛型函数或泛型类型**。

**实参类型自动推导\(Argument type inference\)**

我们是想要：

```text
Sort[book](bookshelf)
```

还是：

```text
Sort(bookshelf)
```

显然是后者。我们希望Go编译器能够根据传入的变量自动推导出类型参数的实参类型。

  
图：实参类型的自动推导

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-10.png)

这样，在具化之前，如果泛型函数调用没有显式提供实参类型，那么Go编译器将进行自动实参类型推导。**有了是实参类型的自动推导，大多数泛型调用的方式与常规函数调用一致**。

**类型列表\(type lists\)**

到这里，约束仅限于描述方法要求。下面的函数调用仍然无法工作：

```text
Sort([]int{1, 2, 3})
```

因为原生的int类型不满足Elem的约束，没有实现Less方法。虽然我们可以用下面替代方法实现整型切片的排序：

```text
type myInt int
func (x myInt) Less(y myInt) bool { return x < y }
```

但这还是太麻烦了。

Go泛型扩展了interface语法，除了让interface拥有自己的方法列表外，还支持在interface中定义类型列表\(type list\)：

```text
type Float interface {
   type float32, float64
}

// float32和float64都可以作为类型实参传递给Sin
func Sin[T Float](x T) T
```

现在，一个类型实参要想满足约束，要么它实现了约束中的所有方法，要么它或它的底层类型\(underlying type\)在约束的类型列表中。

下面是一个泛型函数min的声明与约束定义：

```text
func min[T Ordered](x, y T) T ...

type Ordered interface {
    type int, int8, int16, ..., uint, uint8, uint16, ..., float32, float64, string
}
```

函数min的实现如下：

```text
func min[T Ordered](x, y T) T {
    if x < y {
        return x
    }
    return y
}
```

* x和y的类型都是T，T类型要满足约束Ordered；
* x &lt; y是合法的，因为在Ordered的类型列表中的每个类型都支持"&lt;"比较。

但不同类型参数代表的却是不同类型：

```text
func invalid[Tx, Ty Ordered](x Tx, y Ty) Tx {
    ...
    if x < y { // 不合法
        ...
    }
}
```

* x的类型是Tx，y的类型是Ty；
* Tx和Ty是不同类型；
* "&lt;"需要两个操作数拥有相同的类型。

**类型列表应用的典型示例**

* 将\[\]byte和string的操作整合在一起

我们知道目前标准库中有一个bytes包和一个strings包，这两个包一个用于处理\[\]byte，一个则用于处理string。但使用过这两个包的gopher会发现，这两个包中大部分函数和方法是一样的，甚至处理逻辑都是一样的。有了泛型后，我们可以将对两种类型的大部分操作整合在一起，以Index函数为例：

```text
type Bytes interface {
   type []byte, string
}

// Index returns the index of the first instance of sep
// in s, or -1 if sep is not present in s.
func Index[bytes Bytes](s, sep bytes) int
```

* 类型参数\(type parameter\)之间的关系

```text
type Pointer[T any] interface {
    type *T
}

func f[T any, PT Pointer[T]](x T)

或

func foo[T any, PT interface{type *T}](x T)
```

上面是基于类型列表表述“一个类型的指针类型”约束的方案。PT的实参的类型必须是T的实参类型的指针类型。

下面这几个函数和接口很大可能会加入到标准库：

```text
func BasicSort[Elem Ordered](list []Elem)

func Sort[Elem Lesser[Elem]](list []Elem)

type Lesser[Elem any] interface {
    Less(Elem) Elem
}
```

**小结**

关于泛型声明：

* 类型参数列表和普通参数列表相似，只是使用"\[ \]"括起；
* 函数和类型都可以拥有类型参数列表；
* 使用interface表达对类型参数的约束。

关于泛型使用：

* 泛型函数和类型在使用之前必须先“具化\(instantiated\)”；
* 类型自动推导可实现函数隐式具化；
* 如果类型实参满足约束，那么具化才会合法。

截至2020.10月份的泛型设计草案版本，我们对以下特性设计的满意度为：

![img{512x368}](https://tonybai.com/wp-content/uploads/go-generics-at-gophercon-2020-11.png)

#### 三. 结束语

**“能力越大，责任越大”**

* 类型参数\(泛型\)是Go工具集中的新成员；
* 它与语言的其他部分正交；
* 其正交性也打开了编码风格的一个新维度。

**泛型引入了抽象，无用的抽象带来复杂性。请三思而后行！**

**示例1**

```text
func ReadAll(r io.Reader) ([]byte, error)

对比：

func ReadAll[reader io.Reader](r reader) ([]byte, error)
```

=&gt; 引入泛型的版本并未解决任何实际问题\(还带来了复杂难以理解的抽象\)

**示例2**

```text
// Drain drains any elements remaining on the channel.
func Drain[T any](c <-chan T)

// Merge merges two channels of some element type into
// a single channel.
func Merge[T any](c1, c2 <-chan T) <-chan T
```

=&gt; 类型参数让以往无法实现的逻辑成为现实。

**何时使用泛型**

* 增强静态类型安全性
* 更高效的内存使用
* \(显著的\)更好的性能

**泛型是带有类型检查的宏\(macro\)。使用宏之前请三思！**

**接下来的工作**

Go核心团队正在着手做出一个完整的泛型实现，以便我们解决所有未解决的问题。我们继续欢迎大家的反馈！

如何抢先体验泛型：

* playground: https://go2goplay.golang.org/
* go2go命令工具：git checkout dev.go2go

> 注：2020.11.21日，Go开发团队技术负责人Russ Cox在golang-dev上的mail确认了Go泛型\(type parameter\)[将在Go 1.18版本落地，即2022.2月份](https://groups.google.com/g/golang-dev/c/U7eW9i0cqmo/m/ffs0tyIYBAAJ?pli=1)。

