---
description: >-
  type/interface-struct-loop-map-switch-package-27Gotchas-tutorial-cheetsheet-commontasks
---

# Why Go-Key advantages

Go makes it easier \(than Java or Python\) to write correct, clear and efficient code.

## **Minimalism**

When you add new features to a language, the complexity doesn’t just add up, it often multiplies: language features can interact in many ways. This is a significant problem – language complexity affects **all developers** \(not just the ones writing the spec and implementing the compiler\).

Here are some core Go features:

* The **built-in** frameworks for [testing](https://golang.org/doc/code.html#Testing) and [profiling](https://blog.golang.org/profiling-go-programs) are small and easy to learn, but still fully functional. There are plenty of third-party add-ons, but chances are you won’t need them.
* It’s possible to [**debug**](https://blog.golang.org/debugging-what-you-deploy) and [**profile**](https://blog.golang.org/profiling-go-programs) an optimized binary running in production through an HTTP server.
* Go has [automatically generated documentation](https://blog.golang.org/godoc-documenting-go-code) with [testable examples](https://blog.golang.org/examples). Once again, the **interface is minimal**, and there is very little to learn.
* Go is **strongly** and **statically** typed with [**no implicit conversions**](https://yourbasic.org/golang/conversions/), but the syntactic overhead is still surprisingly small. This is achieved by simple [type inference in assign­ments](https://tour.golang.org/basics/14) together with [untyped numeric constants](https://yourbasic.org/golang/untyped-constants/). This gives Go stronger type safety than Java \(which has implicit conversions\), but the code reads more like Python \(which has untyped variables\).
* Programs are constructed from [packages](https://yourbasic.org/golang/packages-explained/) that offer clear [**code separation**](https://yourbasic.org/golang/public-private/) and allow efficient management of dependencies. The package mechanism is perhaps the single most well-designed feature of the language, and certainly one of the most overlooked.
* Structurally typed [interfaces](https://yourbasic.org/golang/interfaces-explained/) provide runtime **polymorphism** through [dynamic dispatch](https://en.wikipedia.org/wiki/Runtime_polymorphism).
* [Concurrency](https://yourbasic.org/golang/concurrent-programming/) is an **integral part** of Go, supported by [goroutines](https://yourbasic.org/golang/goroutines-explained/), [channels](https://yourbasic.org/golang/channels-explained/) and the [select statement](https://yourbasic.org/golang/select-explained/).

> See [Go vs. Java: 15 main differences](https://yourbasic.org/golang/go-vs-java/) for a small code example, and a sample of basic data types, methods and control structures.

#### Features for the future <a id="features-for-the-future"></a>



Go **omits several features** found in other modern languages.

Here’s the [language designers’ answer](https://golang.org/doc/faq#Design) to _Why does Go not have feature X?_

> Every language contains novel features and omits someone’s favorite feature. Go was designed with an eye on felicity of programming, speed of compilation, orthogonality of concepts, and the need to support features such as concurrency and garbage collection. Your favorite feature may be missing because it doesn’t fit, because it affects compilation speed or clarity of design, or because it would make the fundamental system model too difficult.

New features are considered only if there is a pressing need demon­strated by [experience reports](https://github.com/golang/go/wiki/ExperienceReports) from real-world projects.

There are a few likely major additions in the pipeline:

* Package management through [modules](https://blog.golang.org/modules2019) were [preliminary introduced in Go 1.11](https://golang.org/doc/go1.11#modules).
* There is a [generics draft design](https://go.googlesource.com/proposal/+/master/design/go2draft-contracts.md), which may be implemented in Go 2. For now, have a look at [Generics \(alternatives and workarounds\)](https://yourbasic.org/golang/generics/).
* Similarly, there is an [error handling draft design](https://github.com/golang/proposal/blob/master/design/go2draft-error-handling.md) that extends the current minimalist [error handling](https://yourbasic.org/golang/errors-explained/).

Currently, some [minor additions](https://blog.golang.org/go2-here-we-come) are considered to help create and test a [community-driven process](https://blog.golang.org/toward-go2) for developing Go.

However, features such as [optional parameters, default parameter values and method overloading](https://yourbasic.org/golang/overload-overwrite-optional-parameter/) probably won’t be part of a future [Go 2](https://blog.golang.org/go2-here-we-come).



#### Java comparison <a id="java-comparison"></a>

[The Java® Language Specification](https://docs.oracle.com/javase/specs/jls/se12/jls12.pdf) is currently 750 pages. Much of the complexity is due to [feature creep](https://en.wikipedia.org/wiki/Feature_creep). Here are but three examples.

Java [inner classes](https://en.wikipedia.org/wiki/Inner_class) suddenly appeared in 1997; it took more than a year to update the specification, and it became almost twice as big as a result. That’s a high price to pay for a non-essential feature.

[Generics in Java](https://en.wikipedia.org/wiki/Generics_in_Java), implemented by [type erasure](https://en.wikipedia.org/wiki/Type_erasure), make some code cleaner and allow for additional runtime checks, but quickly become complex when you move beyond basic examples: [generic arrays](http://www.tothenew.com/blog/why-is-generic-array-creation-not-allowed-in-java/) aren’t supported and [type wildcards](https://docs.oracle.com/javase/tutorial/extra/generics/wildcards.html) with upper and lower bounds are quite complicated. The string “generic” appears 280 times in the specification. It’s not clear to me if this feature is worth its cost.

A Java [enum](https://docs.oracle.com/javase/tutorial/java/javaOO/enum.html), introduced in 2004, is a special type of class that represents a group of constants. It’s certainly nice to have, but offers little that couldn’t be done with ordinary classes. The string “enum” appears 241 times in the specification.  


### Code transparency <a id="code-transparency"></a>

Your project is doomed if you can’t understand your code.

* You **always** need to know **exactly what** your coding is doing,
* and **sometimes** need to **estimate the resources** \(time and memory\) it uses.

Go tries to meet both of these goals.

The syntax is designed to be transparent and there is **one** [standard code format](https://golang.org/doc/effective_go.html#formatting), automatically generated by the [fmt tool](https://blog.golang.org/go-fmt-your-code).

Another example is that Go programs with [unused package imports do not compile](https://yourbasic.org/golang/unused-imports/). This improves code clarity and long-term performance.

I suspect that the Go designers made some things **difficult on purpose**. You need to jump through hoops to [catch a panic](https://yourbasic.org/golang/recover-from-panic/) \(exception\) and you are forced to sprinkle your code with the word “unsafe” to step around type safety.

#### Python comparison <a id="python-comparison"></a>

The Python code snippet `del a[i]` deletes the element at index `i` from a list `a`. This code certainly is quite **readable**, but not so **transparent**: it’s easy to miss that the [time complexity](https://yourbasic.org/algorithms/time-complexity-explained/) is O\(n\), where n is the number of elements in the list.

Go doesn’t have a similar utility function. This is definitely **less convenient**, but also **more transparent**. Your code needs to explicitly state if it copies a section of the list. See [2 ways to delete an element from a slice](https://yourbasic.org/golang/delete-element-slice/) for a Go code example.  


Go doesn’t have a similar utility function. This is definitely **less convenient**, but also **more transparent**. Your code needs to explicitly state if it copies a section of the list. See [2 ways to delete an element from a slice](https://yourbasic.org/golang/delete-element-slice/) for a Go code example.  


### Fast version \(changes order\) <a id="fast-version-changes-order"></a>

```text
a := []string{"A", "B", "C", "D", "E"}
i := 2

// Remove the element at index i from a.
a[i] = a[len(a)-1] // Copy last element to index i.
a[len(a)-1] = ""   // Erase last element (write zero value).
a = a[:len(a)-1]   // Truncate slice.

fmt.Println(a) // [A B E D]
```

The code copies a single element and runs in **constant time**.

### Slow version \(maintains order\) <a id="slow-version-maintains-order"></a>

```text
a := []string{"A", "B", "C", "D", "E"}
i := 2

// Remove the element at index i from a.
copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
a[len(a)-1] = ""     // Erase last element (write zero value).
a = a[:len(a)-1]     // Truncate slice.

fmt.Println(a) // [A B D E]
```

The code copies len\(a\) - i - 1 elements and runs in **linear time**.  


#### Java comparison <a id="java-comparison-1"></a>

Code transparency is not just a syntactic issue. Here are two examples where the rules for [Go package initialization and program execution order](https://yourbasic.org/golang/package-init-function-main-execution-order/) make it easier to reason about and maintain a project.

* [Circular dependencies](https://en.wikipedia.org/wiki/Circular_dependency) can cause many unwanted effects. As opposed to Java, a Go program with initialization cycles will not compile.
* When the main function of a Go program returns, the program exits. A Java programs exits when all user non-daemon threads finish.

This means that you may need to study large parts of a Java program to understand some of its behavior. This may even be impossible if you use third-party libraries.

### Compatibility <a id="compatibility"></a>

A language that **changes abruptly**, or becomes **unavailable**, can end your project.

Go 1 has succinct and strict [**compatibility guarantees**](https://golang.org/doc/go1compat) for the core language and standard packages – Go programs that work today should continue to work with future releases of Go 1. Backward compatibility has been excellent so far.

Go is an **open source** project and comes with a [BSD-style license](https://golang.org/LICENSE) that permits commercial use, modification, distribution, and private use. Copyright belongs to [The Go Authors](https://golang.org/AUTHORS), those of us who [contributed](https://golang.org/doc/contribute.html) to the project. There is also a [patent grant](https://golang.org/PATENTS) by Google.

#### Python comparison <a id="python-comparison-1"></a>

If you’re a Python developer, you know the pain of having to deal with the [differ­ences between Python 2.7.x and Python 3.x](https://sebastianraschka.com/Articles/2014_python_2_3_key_diff.html). There are [strong reasons](https://wiki.python.org/moin/Python2orPython3) for choosing Python 3, but if you depend on libraries that are only available for an older version, you may not be able.

#### Java comparison <a id="java-comparison-2"></a>

Java has a very good history of backward compatibility and the [Compatibility Guide for JDK 8](https://www.oracle.com/technetwork/java/javase/8-compatibility-guide-2156366.html) is extensive. Also, Java has been freely available to developers for a long time.

Unfortunately, there are some dark clouds on the horizon with the [Oracle America, Inc. v. Google, Inc.](https://en.wikipedia.org/wiki/Oracle_America,_Inc._v._Google,_Inc.) legal case about the nature of computer code and copyright law, and Oracle’s new [Java licensing model](https://www.infoworld.com/article/3284164/oracle-now-requires-a-subscription-to-use-java-se.html).  


### Performance <a id="performance"></a>

The exterior of Go is far from flashy, but there is a **fine-tuned engine** underneath.

It makes little sense to discuss performance issues out of context. Running time and memory use is heavily influenced by factors such as algorithms, data structures, input, coding skill, operating systems, and hardware.

Still, **language**, **runtime** and **standard libraries** can have a large effect on perfor­mance. This discussion is limited to high-level issues and design decisions. See the [Go FAQ](https://golang.org/doc/faq) for a more detailed look at the [implementation](https://golang.org/doc/faq#Implementation) and its [performance](https://golang.org/doc/faq#Performance).

First, Go is a **compiled** language. An executable Go program typically consists of a **single standalone binary**, with no separate dynamic libraries or virtual machines, which can be **directly deployed**.

**Size and speed of generated code** will vary depending on target architecture. Go code generation is [fairly mature](https://about.sourcegraph.com/go/generating-better-machine-code-with-ssa/) and the major OSes \(Linux, macOS, Windows\) and architectures \(Intel x86/x86-64, ARM64, WebAssembly, ARM\), as well as many others, are supported. You can expect performance to be on a similar level to that of C++ or Java. Compared to interpreted Python code, the improvement can be huge.

Go is **garbage collected**, protecting against memory leaks. The collection has [**very low latency**](https://blog.golang.org/ismmkeynote). In fact, you may never notice that the GC thread is there.

The **standard libraries** are typically of high quality, with optimized code using efficient algorithms. As an example, [regular expressions](https://yourbasic.org/golang/regexp-cheat-sheet/) are very efficient in Go with running time linear in the size of the input. Unfortunately, this is [not true for Java and Python](https://swtch.com/~rsc/regexp/regexp1.html).

**Build speeds**, in absolute terms, are currently fairly good. More importantly, Go is **designed** to make compilation and dependency analysis easy, making it possible to create programming tools that **scales well** with growing projects.

## Type, value and equality of interfaces

## Interface type

实现了接口的方法，就是实现了接口

> An interface type consists of a set of method signatures. A variable of interface type can hold any value that implements these methods.

In this example both `Temp` and `*Point` implement the `MyStringer` interface.

```text
type MyStringer interface {
	String() string
}
```

```text
type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
```

Actually, `*Temp` also implements `MyStringer`, since the method set of a pointer type `*T` is the set of all methods with receiver `*T` or `T`.

When you call a method on an interface value, the method of its underlying type is executed.

```text
var x MyStringer

x = Temp(24)
fmt.Println(x.String()) // 24 °C

x = &Point{1, 2}
fmt.Println(x.String()) // (1,2)

```

### Structural typing <a id="structural-typing"></a>

> A type implements an interface by implementing its methods. No explicit declaration is required.

In fact, the `Temp`, `*Temp` and `*Point` types also implement the standard library [`fmt.Stringer`](https://golang.org/pkg/fmt/#Stringer) interface. The `String` method in this interface is used to print values passed as an operand to functions such as [`fmt.Println`](https://golang.org/pkg/fmt/#Println).

```text
var x MyStringer

x = Temp(24)
fmt.Println(x) // 24 °C

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

### The empty interface <a id="the-empty-interface"></a>

The interface type that specifies no methods is known as the empty interface.

```text
interface{}
```

An empty interface can hold values of any type since every type implements at least zero methods.

```text
var x interface{}

x = 2.4
fmt.Println(x) // 2.4

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

The [`fmt.Println`](https://golang.org/pkg/fmt/#Println) function is a chief example. It takes any number of arguments of any type.

```text
func Println(a ...interface{}) (n int, err error)
```

### Interface values <a id="interface-values"></a>

> An **interface value** consists of a **concrete value** and a **dynamic type**: `[Value, Type]`

In a call to [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf), you can use `%v` to print the concrete value and `%T` to print the dynamic type.

```text
var x MyStringer
fmt.Printf("%v %T\n", x, x) // <nil> <nil>

x = Temp(24)
fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

x = &Point{1, 2}
fmt.Printf("%v %T\n", x, x) // (1,2) *main.Point

x = (*Point)(nil)
fmt.Printf("%v %T\n", x, x) // <nil> *main.Point
```

The **zero value** of an interface type is nil, which is represented as `[nil, nil]`.

Calling a method on a nil interface is a run-time error. However, it’s quite common to write methods that can handle a receiver value `[nil, Type]`, where `Type` isn’t nil.

You can use [**type assertions**](https://yourbasic.org/golang/type-assertion-switch/) or [**type switches**](https://yourbasic.org/golang/type-assertion-switch/) to access the dynamic type of an interface value. See [Find the type of an object](https://yourbasic.org/golang/find-type-of-object/) for more details.

### Equality <a id="equality"></a>

Two interface values are equal

* if they have equal concrete values **and** identical dynamic types,
* or if both are nil.

A value `t` of interface type `T` and a value `x` of non-interface type `X` are equal if

* `t`’s concrete value is equal to `x`
* **and** `t`’s dynamic type is identical to `X`.

```text
var x MyStringer
fmt.Println(x == nil) // true

x = (*Point)(nil)
fmt.Println(x == nil) // false
```

In the second print statement, the concrete value of `x` equals `nil`, but its dynamic type is `*Point`, which is not `nil`.

> **Warning:** See [Nil is not nil](https://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/) for a real-world example where this definition of equality leads to puzzling results.

#### Further reading <a id="further-reading"></a>

[Generics \(alternatives and workarounds\)](https://yourbasic.org/golang/generics/) discusses how interfaces, multiple functions, type assertions, reflection and code generation can be use in place of parametric polymorphism in Go.  


## Create, initialize and compare structs

### Struct types <a id="struct-types"></a>

A struct is a typed collection of fields, useful for grouping data into records.

```text
type Student struct {
    Name string
    Age  int
}

var a Student    // a == Student{"", 0}
a.Name = "Alice" // a == Student{"Alice", 0}
```

* To define a new **struct type**, you list the names and types of each field.
* The default **zero value** of a struct has all its fields zeroed.
* You can access individual fields with **dot notation**.

### 2 ways to create and initialize a new struct <a id="2-ways-to-create-and-initialize-a-new-struct"></a>

The **`new`** keyword can be used to create a new struct. It returns a [pointer](https://yourbasic.org/golang/pointers-explained/) to the newly created struct.

```text
var pa *Student   // pa == nil
pa = new(Student) // pa == &Student{"", 0}
pa.Name = "Alice" // pa == &Student{"Alice", 0}
```

You can also create and initialize a struct with a **struct literal**.

```text
b := Student{ // b == Student{"Bob", 0}
    Name: "Bob",
}
    
pb := &Student{ // pb == &Student{"Bob", 8}
    Name: "Bob",
    Age:  8,
}

c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
d := Student{}             // d == Student{"", 0}
```

* An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
* An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
* A literal may omit the element list; such a literal evaluates to the zero value for its type.

For further details, see [The Go Language Specification: Composite literals](https://golang.org/ref/spec#Composite_literals).

### Compare structs <a id="compare-structs"></a>

You can compare struct values with the comparison operators `==` and `!=`. Two values are equal if their corresponding fields are equal.

```text
d1 := Student{"David", 1}
d2 := Student{"David", 2}
fmt.Println(d1 == d2) // false

```

## Maps explained: create, add, get, delete



### Create a new map <a id="create-a-new-map"></a>

```text
var m map[string]int                // nil map of string-int pairs

m1 := make(map[string]float64)      // Empty map of string-float64 pairs
m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

m3 := map[string]float64{           // Map literal
    "e":  2.71828,
    "pi": 3.1416,
}
fmt.Println(len(m3))                // Size of map: 2
```

* A map \(or dictionary\) is an **unordered** collection of **key-value** pairs, where each key is **unique**.
* You create a new map with a [**make**](https://golang.org/pkg/builtin/#make) statement or a **map literal**.
* The default **zero value** of a map is `nil`. A nil map is equivalent to an empty map except that **elements can’t be added**.
* The [**`len`**](https://golang.org/pkg/builtin/#len) function returns the **size** of a map, which is the number of key-value pairs.

> **Warning:** If you try to add an element to an uninitialized map you get the mysterious run-time error [_Assignment to entry in nil map_](https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/).

### Add, update, get and delete keys/values <a id="add-update-get-and-delete-keys-values"></a>

```text
m := make(map[string]float64)

m["pi"] = 3.14             // Add a new key-value pair
m["pi"] = 3.1416           // Update value
fmt.Println(m)             // Print map: "map[pi:3.1416]"

v := m["pi"]               // Get value: v == 3.1416
v = m["pie"]               // Not found: v == 0 (zero value)

_, found := m["pi"]        // found == true
_, found = m["pie"]        // found == false

if x, found := m["pi"]; found {
    fmt.Println(x)
}                           // Prints "3.1416"

delete(m, "pi")             // Delete a key-value pair
fmt.Println(m)              // Print map: "map[]"
```

* When you index a map you get two return values; the second one \(which is optional\) is a boolean that indicates if the key exists.
* If the key doesn’t exist, the first value will be the default [zero value](https://yourbasic.org/golang/default-zero-value/).

### For-each range loop <a id="for-each-range-loop"></a>

```text
m := map[string]float64{
    "pi": 3.1416,
    "e":  2.71828,
}
fmt.Println(m) // "map[e:2.71828 pi:3.1416]"

for key, value := range m { // Order not specified 
    fmt.Println(key, value)
}
```

* Iteration order is not specified and may vary from iteration to iteration.
* If an entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced.（删除不迟到）
* If an entry is created during iteration, that entry may or may not be produced during the iteration.（add 可能会迟到）

> Starting with [Go 1.12](https://tip.golang.org/doc/go1.12), the [fmt package](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) prints maps in key-sorted order to ease testing.

### Performance and implementation <a id="performance-and-implementation"></a>

* Maps are backed by [hash tables](https://yourbasic.org/algorithms/hash-tables-explained/).
* Add, get and delete operations run in **constant** expected time. The time complexity for the add operation is [amortized](https://yourbasic.org/algorithms/amortized-time-complexity-analysis/).
* The comparison operators `==` and `!=` must be defined for the key type.

## 5 basic for loop patterns

### Three-component loop <a id="three-component-loop"></a>

This version of the Go for loop works just as in C or Java.

```text
sum := 0
for i := 1; i < 5; i++ {
    sum += i
}
fmt.Println(sum) // 10 (1+2+3+4)
```

1. The init statement, `i := 1`, runs.
2. The condition, `i < 5`, is computed.
   * If true, the loop body runs,
   * otherwise the loop is done.
3. The post statement, `i++`, runs.
4. Back to step 2.

The scope of `i` is limited to the loop.

### While loop <a id="while-loop"></a>

If you skip the init and post statements, you get a while loop.

```text
n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n) // 8 (1*2*2*2)
```

1. The condition, `n < 5`, is computed.
   * If true, the loop body runs,
   * otherwise the loop is done.
2. Back to step 1.

### Infinite loop <a id="infinite-loop"></a>

If you skip the condition as well, you get an infinite loop.

```text
sum := 0
for {
    sum++ // repeated forever
}
fmt.Println(sum) // never reached
```

### For-each range loop <a id="for-each-range-loop"></a>

Looping over elements in _slices_, _arrays_, _maps_, _channels_ or _strings_ is often better done with a range loop.

```text
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
```

```text
0 hello
1 world
```

See [4 basic range loop patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for a complete set of examples.

### Exit a loop <a id="exit-a-loop"></a>

The break and continue keywords work just as they do in C and Java.

```text
sum := 0
for i := 1; i < 5; i++ {
    if i%2 != 0 { // skip odd numbers
        continue
    }
    sum += i
}
fmt.Println(sum) // 6 (2+4)
```

* A **continue** statement begins the next iteration of the innermost **for** loop at its post statement \(`i++`\).
* A **break** statement leaves the innermost **for**, [**switch**](https://yourbasic.org/golang/switch-statement/) or [**select**](https://yourbasic.org/golang/select-explained/) statement.

#### Further reading <a id="further-reading"></a>

See [4 basic range loop \(for-each\) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for a detailed description of how to loop over slices, arrays, strings, maps and channels in Go.

## 5 switch statement patterns

### Basic switch with default <a id="basic-switch-with-default"></a>

* A switch statement runs the first case equal to the condition expression.
* The cases are evaluated from top to bottom, stopping when a case succeeds.
* If no case matches and there is a default case, its statements are executed.

```text
switch time.Now().Weekday() {
case time.Saturday:
    fmt.Println("Today is Saturday.")
case time.Sunday:
    fmt.Println("Today is Sunday.")
default:
    fmt.Println("Today is a weekday.")
}
```

> Unlike C and Java, the case expressions do not need to be constants.

### No condition <a id="no-condition"></a>

A switch without a condition is the same as switch true.

```text
switch hour := time.Now().Hour(); { // missing expression means "true"
case hour < 12:
    fmt.Println("Good morning!")
case hour < 17:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}
```

### Case list <a id="case-list"></a>

```text
func WhiteSpace(c rune) bool {
    switch c {
    case ' ', '\t', '\n', '\f', '\r':
        return true
    }
    return false
}
```

### Fallthrough <a id="fallthrough"></a>

* A `fallthrough` statement transfers control to the next case.
* It may be used only as the final statement in a clause.

```text
switch 2 {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")
    fallthrough
case 3:
    fmt.Println("3")
}
```

```text
2
3
```

### Exit with break <a id="exit-with-break"></a>

A `break` statement terminates execution of the **innermost** `for`, `switch`, or `select` statement.

If you need to break out of a surrounding loop, not the switch, you can put a **label** on the loop and break to that label. This example shows both uses.

```text
Loop:
    for _, ch := range "a b\nc" {
        switch ch {
        case ' ': // skip space
            break
        case '\n': // break at newline
            break Loop
        default:
            fmt.Printf("%c\n", ch)
        }
    }
```

```text
a
b
```

### Execution order <a id="execution-order"></a>

* First the switch expression is evaluated once.
* Then case expressions are evaluated left-to-right and top-to-bottom:
  * the first one that equals the switch expression triggers execution of the statements of the associated case,
  * the other cases are skipped.

```text
// Foo prints and returns n.
func Foo(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Foo(2) {
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
        fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}
```

```text
2
1
2
First case
Second case
```

## Packages explained: declare, import, download, document

### Basics <a id="basics"></a>

Every Go program is made up of packages and each package has an **import path**:

* `"fmt"`
* `"math/rand"`
* `"github.com/yourbasic/graph"`

Packages in the standard library have short import paths, such as `"fmt"` and `"math/rand"`. Third-party packages, such as `"github.com/yourbasic/graph"`, typically have an import path that includes a hosting service \(`github.com`\) and an organization name \(`yourbasic`\).

By convention, the **package name** is the same as the last element of the import path:

* `fmt`
* `rand`
* `graph`

References to other packages’ definitions must always be prefixed with their package names, and only the capitalized names from other packages are accessible.

```text
package main

import (
    "fmt"
    "math/rand"

    "github.com/yourbasic/graph"
)

func main() {
    n := rand.Intn(100)
    g := graph.New(n)
    fmt.Println(g)
}
```

### Declare a package <a id="declare-a-package"></a>

Every Go source file starts with a package declaration, which contains only the package name.

For example, the file [`src/math/rand/exp.go`](https://golang.org/src/math/rand/exp.go), which is part of the implementation of the [`math/rand`](https://golang.org/pkg/math/rand/) package, contains the following code.

```text
package rand
  
import "math"
  
const re = 7.69711747013104972
…
```

You don’t need to worry about package name collisions, only the import path of a package must be unique. [How to Write Go Code](https://golang.org/doc/code.html) shows how to organize your code and its packages in a file structure.

### Package name conflicts <a id="package-name-conflicts"></a>

You can customize the name under which you refer to an imported package.

```text
package main

import (
    csprng "crypto/rand"
    prng "math/rand"

    "fmt"
)

func main() {
    n := prng.Int() // pseudorandom number
    b := make([]byte, 8)
    csprng.Read(b) // cryptographically secure pseudorandom number
    fmt.Println(n, b)
}
```

### Dot imports <a id="dot-imports"></a>

If a period `.` appears instead of a name in an import statement, all the package’s exported identifiers can be accessed without a qualifier.

```text
package main

import (
    "fmt"
    . "math"
)

func main() {
    fmt.Println(Sin(Pi/2)*Sin(Pi/2) + Cos(Pi)/2) // 0.5
}
```

Dot imports can make programs hard to read and **generally should be avoided**.

### Package download <a id="package-download"></a>

The [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) command downloads packages named by import paths, along with their dependencies, and then installs the packages.

```text
$ go get github.com/yourbasic/graph
```

The import path corresponds to the repository hosting the code. This reduces the likelihood of future name collisions.

The [Go Wiki](https://github.com/golang/go/wiki/Projects) and [Awesome Go](https://github.com/avelino/awesome-go) provide lists of high-quality Go packages and resources.

For more information on using remote repositories with the go tool, see [Command go: Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths).

### Package documentation <a id="package-documentation"></a>

The [GoDoc](https://godoc.org/) web site hosts documentation for all public Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad:

* [`https://godoc.org/fmt`](https://godoc.org/fmt)
* [`https://godoc.org/math/rand`](https://godoc.org/math/rand)
* [`https://godoc.org/github.com/yourbasic/graph`](https://godoc.org/github.com/yourbasic/graph)

The [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command extracts and generates documentation for all locally installed Go programs. The following command starts a web server that presents the documentation at `http://localhost:6060/`.

```text
$ godoc -http=:6060 
```

For more on how to access and create documentation, see the [Package documentation](https://yourbasic.org/golang/package-documentation/) article.

## 27 Go Gotcha Ninja Pitfalls

## 1.Assignment to entry in nil map

Why does this program panic?

```text
var m map[string]float64
m["pi"] = 3.1416
```

```text
panic: assignment to entry in nil map
```

### Answer <a id="answer"></a>

You have to initialize the map using the make function \(or a map literal\) before you can add any elements:

```text
m := make(map[string]float64)
m["pi"] = 3.1416
```



## 2.Invalid memory address or nil pointer dereference

Why does this program panic?

```text
type Point struct {
    X, Y float64
}

func (p *Point) Abs() float64 {
    return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func main() {
    var p *Point
    fmt.Println(p.Abs())
}
```

```text
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xd2c5a]

goroutine 1 [running]:
main.(*Point).Abs(...)
	../main.go:6
main.main()
	../main.go:11 +0x1a
```

### Answer <a id="answer"></a>

The uninitialized pointer `p` in the `main` function is `nil`, and you can’t follow the nil pointer.

> If `x` is nil, an attempt to evaluate `*x` will cause a run-time panic.[The Go Programming Language Specification: Address operators](https://golang.org/ref/spec#Address_operators)

You need to create a `Point`:

```text
func main() {
    var p *Point = new(Point)
    fmt.Println(p.Abs())
}
```

Since methods with pointer receivers take either a value or a pointer, you could also skip the pointer altogether:

```text
func main() {
    var p Point // has zero value Point{X:0, Y:0}
    fmt.Println(p.Abs())
}
```

See [Pointers](https://yourbasic.org/golang/pointers-explained/) for more about pointers in Go.



## 3.Multiple-value in single-value context



Why does this code give a compile error?

```text
t := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
fmt.Println(t)
```

```text
../main.go:9:17: multiple-value time.Parse() in single-value context
```

### Answer <a id="answer"></a>

The [`time.Parse`](https://golang.org/pkg/time/#Parse) function returns two values, a [`time.Time`](https://golang.org/pkg/time/#Time) and an [`error`](https://yourbasic.org/golang/errors-explained/), and you must use both.

```text
t, err := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
if err != nil {
    // TODO: Handle error.
}
fmt.Println(t)
```

```text
2018-04-06 10:49:05 +0000 UTC
```

#### Blank identifier \(underscore\) <a id="blank-identifier-underscore"></a>

You can use the [blank identifier](https://yourbasic.org/golang/underscore/) to ignore unwanted return values.

```text
m := map[string]float64{"pi": 3.1416}
_, exists := m["pi"] // exists == true
```



## 4.Array won’t change

Why does the array value stick?

```text
func Foo(a [2]int) {
    a[0] = 8
}

func main() {
    a := [2]int{1, 2}
    Foo(a)         // Try to change a[0].
    fmt.Println(a) // Output: [1 2]
}
```

### Answer <a id="answer"></a>

* Arrays in Go are **values**.
* When you pass an array to a function, the array is copied.

If you want `Foo` to update the elements of `a`, _use a slice instead_.

```text
func Foo(a []int) {
    if len(a) > 0 {
        a[0] = 8
    }
}

func main() {
    a := []int{1, 2}
    Foo(a)         // Change a[0].
    fmt.Println(a) // Output: [8 2]
}
```

A slice does not store any data, it just describes a section of an underlying array.

When you change an element of a slice, you modify the corresponding element of its underlying array, and other slices that share the same underlying array will see the change.

See [Slices and arrays in 6 easy steps](https://yourbasic.org/golang/slices-explained/) for all about slices in Go.

## 5.Shadowed variables

Why doesn’t `n` change?

```text
func main() {
    n := 0
    if true {
        n := 1
        n++
    }
    fmt.Println(n) // 0
}
```

### Answer <a id="answer"></a>

The statement `n := 1` declares a new variable which **shadows** the original `n` throughout the scope of the if statement.

To reuse `n` from the outer block, write `n = 1` instead.

```text
func main() {
    n := 0
    if true {
        n = 1
        n++
    }
    fmt.Println(n) // 2
}
```

### Detecting shadowed variables <a id="detecting-shadowed-variables"></a>

To help detect shadowed variables, you may use the experimental `-shadow` feature provided by the [vet](https://golang.org/cmd/vet/) tool. It flags variables that **may have been** unintentionally shadowed. Passing the original version of the code to `vet` gives a warning message.

```text
$ go vet -shadow main.go
main.go:4: declaration of "n" shadows declaration at main.go:2
```

[Go 1.12](https://tip.golang.org/doc/go1.12) no longer supports this. Instead you may do

```text
go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow
go vet -vettool=$(which shadow)
```

Additionally, the Go compiler detects and disallows some cases of shadowing.

```text
func Foo() (n int, err error) {
    if true {
        err := fmt.Errorf("Invalid")
        return
    }
    return
}
```

```text
../main.go:4:3: err is shadowed during return
```

## 6.Unexpected newline

Why doesn’t this program compile?

```text
func main() {
    fruit := []string{
        "apple",
        "banana",
        "cherry"
    }
    fmt.Println(fruit)
}
```

```text
../main.go:5:11: syntax error: unexpected newline, expecting comma or }
```

### Answer <a id="answer"></a>

In a multi-line slice, array or map literal, every line **must end with a comma**.

```text
func main() {
    fruit := []string{
        "apple",
        "banana",
        "cherry", // comma added
    }
    fmt.Println(fruit) // "[apple banana cherry]"
}
```

This behavior is a consequence of the Go [semicolon insertion rules](https://golang.org/ref/spec#Semicolons).

As a result, you can add and remove lines without modifying the surrounding code.

## 7.Immutable strings

Why doesn’t this code compile?

```text
s := "hello"
s[0] = 'H'
fmt.Println(s)
```

```text
../main.go:3:7: cannot assign to s[0]
```

### Answer <a id="answer"></a>

Go strings are immutable and behave like read-only byte slices \(with a few extra properties\).

To update the data, use a rune slice instead.

```text
buf := []rune("hello")
buf[0] = 'H'
s := string(buf)
fmt.Println(s)  // "Hello"
```

If the string only contains ASCII characters, you could also use a byte slice.

See [String functions cheat sheet](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/) for an overview of strings in Go.

## 8.How does characters add up?

Why doesn’t these print statements give the same result?

```text
fmt.Println("H" + "i")
fmt.Println('H' + 'i')
```

```text
Hi
177
```

### Answer <a id="answer"></a>

The rune literals `'H'` and `'i'` are integer values identifying Unicode code points: `'H'` is 72 and `'i'` is 105.

You can turn a code point into a string with a conversion.

```text
fmt.Println(string(72) + string('i')) // "Hi"
```

You can also use the [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function.

```text
s := fmt.Sprintf("%c%c, world!", 72, 'i')
fmt.Println(s)// "Hi, world!"
```

This [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) lists the most common formatting verbs and flags.

## 9.What happened to ABBA?

What’s up with [strings.TrimRight](https://golang.org/pkg/strings/#TrimRight)?

```text
fmt.Println(strings.TrimRight("ABBA", "BA")) // Output: ""
```

### Answer <a id="answer"></a>

The `Trim`, `TrimLeft` and `TrimRight` functions strip all Unicode code points contained in a **cutset**. In this case, all trailing A:s and B:s are stripped from the string, leaving the empty string.

To strip a trailing **string**, use [`strings.TrimSuffix`](https://golang.org/pkg/strings/#TrimSuffix).

```text
fmt.Println(strings.TrimSuffix("ABBA", "BA")) // Output: "AB"
```

See [String functions cheat sheet](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/) for more about strings in Go.

## 10.Where is my copy?

Why does the copy disappear?

```text
var src, dst []int
src = []int{1, 2, 3}
copy(dst, src) // Copy elements to dst from src.
fmt.Println("dst:", dst)
```

```text
dst: []
```

### Answer <a id="answer"></a>

The number of elements copied by the `copy` function is the minimum of `len(dst)` and `len(src)`. To make a full copy, you must allocate a big enough destination slice.

```text
var src, dst []int
src = []int{1, 2, 3}
dst = make([]int, len(src))
n := copy(dst, src)
fmt.Println("dst:", dst, "(copied", n, "numbers)")
```

```text
dst: [1 2 3] (copied 3 numbers)
```

The return value of the `copy` function is the number of elements copied. See [Copy function](https://yourbasic.org/golang/copy-explained/) for more about the built-in `copy` function in Go.

#### Using append

You could also use the `append` function to make a copy by appending to a nil slice.

```text
var src, dst []int
src = []int{1, 2, 3}
dst = append(dst, src...)
fmt.Println("dst:", dst)
```

```text
dst: [1 2 3]
```

Note that the capacity of the slice allocated by `append` may be a bit larger than `len(src)`.

## 11.Why doesn’t append work every time? \[scary bug\]



```text
a := []byte("ba")

a1 := append(a, 'd')
a2 := append(a, 'g')

fmt.Println(string(a1)) // bag
fmt.Println(string(a2)) // bag
```

### Answer <a id="answer"></a>

If there is room for more elements, `append` reuses the underlying array. Let's take a look:

```text
a := []byte("ba")
fmt.Println(len(a), cap(a)) // 2 32
```

This means that the slices `a`, `a1` and `a2` will refer to the same underlying array in our example.

To avoid this, we need to use two separate byte arrays.

```text
const prefix = "ba"

a1 := append([]byte(prefix), 'd')
a2 := append([]byte(prefix), 'g')

fmt.Println(string(a1)) // bad
fmt.Println(string(a2)) // bag
```

#### The scary case: It “worked” for me <a id="the-scary-case-it-worked-for-me"></a>

> In some Go implementations `[]byte("ba")` only allocates two bytes, and then the code seems to work: the first string is `"bad"` and the second one `"bag"`.
>
> Unfortunately the code is **still wrong**, even though it seems to work. The program may behave differently when you run it in another environment.

See [How to append anything](https://yourbasic.org/golang/append-explained/) for more about the built-in `append` function in Go.

## 12.Constant overflows int

Why doesn’t this code compile?

```text
const n = 9876543210 * 9876543210
fmt.Println(n)
```

```text
../main.go:2:13: constant 97546105778997104100 overflows int
```

### Answer <a id="answer"></a>

The untyped constant `n` must be converted to a type before it can be assigned to the `interface{}` parameter in the call to `fmt.Println`.

```text
fmt.Println(a ...interface{})
```

When the type can’t be inferred from the context, an untyped constant is converted to a `bool`, `int`, `float64`, `complex128`, `string` or `rune` depending of the format of the constant.

In this case the constant is an integer, but `n` is larger than the maximum value of an `int`.

However, `n` can be represented as a `float64`.

```text
const n = 9876543210 * 9876543210
fmt.Println(float64(n))
```

```text
9.75461057789971e+19
```

For exact representation of big numbers, the [math/big](https://golang.org/pkg/math/big/) package implements arbitrary-precision arithmetic. It supports signed integers, rational numbers and floating-point numbers.



## 13.Unexpected ++, expecting expression



Why doesn’t these lines compile?

```text
i := 0
fmt.Println(++i)
fmt.Println(i++)
```

```text
main.go:9:14: syntax error: unexpected ++, expecting expression
main.go:10:15: syntax error: unexpected ++, expecting comma or )
```

### Answer <a id="answer"></a>

In Go increment and decrement operations can’t be used as expressions, only as **statements**. Also, only the **postfix notation** is allowed.

The above snippet needs to be written as:

```text
i := 0
i++
fmt.Println(i)
fmt.Println(i)
i++
```

> Without pointer arithmetic, the convenience value of pre- and postfix increment operators drops. By removing them from the expression hierarchy altogether, expression syntax is simplified and the messy issues around order of evaluation of `++` and `--` \(consider `f(i++)` and `p[i] = q[++i]`\) are eliminated as well. The simplification is significant.[Go FAQ: Why are ++ and – statements and not expressions?](https://golang.org/doc/faq#inc_dec)

## 14.Get your priorities right

Why doesn’t this code compute the number of hours and seconds?

```text
n := 43210 // time in seconds
fmt.Println(n/60*60, "hours and", n%60*60, "seconds")
```

```text
43200 hours and 600 seconds
```

### Answer <a id="answer"></a>

The `*`, `/`, and `%` operators have the same precedence and are evaluated left to right: `n/60*60` is the same as `(n/60)*60`.

Insert a pair of parantheses to force the correct evaluation order.

```text
fmt.Println(n/(60*60), "hours and", n%(60*60), "seconds")
```

```text
12 hours and 10 seconds
```

Or better yet, use a constant.

```text
const SecPerHour = 60 * 60
fmt.Println(n/SecPerHour, "hours and", n%SecPerHour, "seconds")
```

```text
12 hours and 10 seconds
```

See [Operators: complete list](https://yourbasic.org/golang/operators/) for a full explanation of the evalutation order of operations in Go expressions.  


## 15.Go and Pythagoras

Pythagorean triples are integer solutions to the Pythagorean Theorem, a2 + b2 = c2.

A well-known example is \(3, 4, 5\):

```text
fmt.Println(3^2+4^2 == 5^2) // true
```

The triple \(6, 8, 10\) is another example, but Go doesn't seem to agree.

```text
fmt.Println(6^2+8^2 == 10^2) // false
```

### Answer <a id="answer"></a>

The circumflex `^` denotes bitwise XOR in Go. The computation written in base 2 looks like this:

```text
0011 ^ 0010 == 0001   (3^2 == 1)
0100 ^ 0010 == 0110   (4^2 == 6)
0101 ^ 0010 == 0111   (5^2 == 7)
```

Of course, `1 + 6 == 7`; Go and Pythagoras agree on that. See [Bitwise operators cheat sheet](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) for more about bitwise calculations in Go.

To raise an integer to the power 2, use multiplication.

```text
fmt.Println(6*6 + 8*8 == 10*10) // true
```

Go has no built-in support for integer power computations, but there is a [`math.Pow`](https://golang.org/pkg/math/#Pow) function for floating-point numbers.

## 16.No end in sight

Why does this loop run forever?

```text
var b byte
for b = 250; b <= 255; b++ {
    fmt.Printf("%d %c\n", b, b)
}
```

### Answer <a id="answer"></a>

byte溢出 重置为0

After the `b == 255` iteration, `b++` is executed. This overflows \(since the maximum value for a byte is 255\) and results in `b == 0`. Therefore `b <= 255` still holds and the loop restarts from 0.

> For unsigned integer values, the operations +, -, \*, and &lt;&lt; are computed modulo 2n, where n is the bit width of the unsigned integer type.
>
> For signed integers, the operations +, -, \*, and &lt;&lt; may legally overflow and the resulting value exists and is deterministically defined by the signed integer representation, the operation, and its operands. No exception is raised as a result of overflow.[The Go Programming Language Specification: Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators)

If we use the standard loop idiom with a strict inequality, the compiler will catch the bug.

```text
var b byte
for b = 250; b < 256; b++ {
    fmt.Printf("%d %c\n", b, b)
}
```

```text
../main.go:2:17: constant 256 overflows byte
```

One solution is to use a wider data type, such as an `int`.

```text
for i := 250; i < 256; i++ {
    fmt.Printf("%d %c\n", i, i)
}
```

```text
250 ú
251 û
252 ü
253 ý
254 þ
255 ÿ
```

Another option is to put the termination test at the end of the loop.

```text
for b := byte(250); ; b++ {
    fmt.Printf("%d %c\n", b, b)
    if b == 255 {
        break
    }
}
```

```text
250 ú
251 û
252 ü
253 ý
254 þ
255 ÿ
```

## 17.Numbers that start with zero

What’s up with the counting in this example?

```text
const (
    Century = 100
    Decade  = 010
    Year    = 001
)
// The world's oldest person, Emma Morano, lived for a century,
// two decades and two years.
fmt.Println("She was", Century+2*Decade+2*Year, "years old.")
```

```text
She was 118 years old.
```

### Answer <a id="answer"></a>

`010` is a number in **base 8**, therefore it means 8, not 10.

Integer literals in Go are specified in octal, decimal or hexadecimal. The number 16 can be written as `020`, `16` or `0x10`.

| Literal | Base | Note |
| :--- | :--- | :--- |
| `020` | 8 | Starts with `0` |
| `16` | 10 | Never starts with `0` |
| `0x10` | 16 | Starts with `0x` |

This [bitwise operators cheat sheet](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) covers all bitwise operators and functions in Go.

#### Zero knowledge \(trivia\) <a id="zero-knowledge-trivia"></a>

There are many ways to write zero in base 8 in Go, including `0`, `00` and `000`. If you prefer hexadecimal notation, you also have a smörgåsbord of options: such as `0x0`, `0x00` and `0x000` \(as well as `0X0`, `0X00` and `0X000`\). However, there is no decimal zero integer literal in Go.

In fact, Go doesn’t have any negative decimal literals either: `-1` is the unary negation operator followed by the decimal literal `1`.

## 18.Whatever remains

Why isn’t -1 odd?

```text
func Odd(n int) bool {
    return n%2 == 1
}

func main() {
    fmt.Println(Odd(-1)) // false
}
```

### Answer <a id="answer"></a>

The remainder operator can give negative answers if the dividend is negative: if `n` is an odd negative number, `n % 2` equals `-1`.

The quotient `q = x / y` and remainder `r = x % y` satisfy the relationships

```text
x = q*y + r  and  |r| < |y|
```

where `x / y` is truncated towards zero.

```text
 x     y     x / y     x % y
 5     3       1         2
-5     3      -1        -2
 5    -3      -1         2
-5    -3       1        -2
```

\(There is one exception: if `x` is the most negative value of its type, the quotient `q = x / -1` is equal to `x`. See [Compute absolute values](https://yourbasic.org/golang/absolute-value-int-float/) for more on this anomaly.\)

One solution is to write the function like this:

```text
// Odd tells whether n is an odd number.
func Odd(n int) bool {
    return n%2 != 0
}
```

You can also use the bitwise AND operator `&`.

```text
// Odd tells whether n is an odd number.
func Odd(n int) bool {
    return n&1 == 1
}
```

## 19.Time is not a number

Why doesn’t this compile?

```text
n := 100
time.Sleep(n * time.Millisecond)
```

```text
invalid operation: n * time.Millisecond (mismatched types int and time.Duration)
```

### Answer <a id="answer"></a>

There is no mixing of numeric types in Go. You can only multiply a `time.Duration` with

* another `time.Duration`, or
* an **untyped integer** constant.

Here are three correct examples.

```text
var n time.Duration = 100
time.Sleep(n * time.Millisecond)
```

```text
const n = 100
time.Sleep(n * time.Millisecond)
```

```text
time.Sleep(100 * time.Millisecond)
```

See [Untyped numeric constants with no limits](https://yourbasic.org/golang/untyped-constants/) for details about typed and untyped integer and floating point constants and their limits.

## 20.Index out of range

Why does this program crash?

```text
a := []int{1, 2, 3}
for i := 1; i <= len(a); i++ {
    fmt.Println(a[i])
}
```

```text
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
	../main.go:3 +0xe0
```

### Answer <a id="answer"></a>

In the last iteration, `i` equals `len(a)` which is outside the bounds of `a`.

Arrays, slices and strings are indexed **starting from zero** so the values of `a` are found at `a[0]`, `a[1]`, `a[2]`, …, `a[len(a)-1]`.

Loop from `0` to `len(a)-1` instead.

```text
for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
}
```

Or, better yet, use a range loop.

```text
for _, n := range a {
    fmt.Println(n)
}
```

## 21.Unexpected values in range loop



```text
primes := []int{2, 3, 5, 7}
for p := range primes {
    fmt.Println(p)
}
```

print

```text
0
1
2
3
```

### Answer <a id="answer"></a>

For arrays and slices, the range loop generates **two values**:

* first the index,
* then the data at this position.

If you omit the second value, you get only the indices.

To print the data, use the second value instead:

```text
primes := []int{2, 3, 5, 7}
for _, p := range primes {
    fmt.Println(p)
}
```

In this case, the blank identifier \(underscore\) is used for the return value you're not interested in.

See [4 basic range loop \(for-each\) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for all about range loops in Go.

## 22.Can’t change entries in range loop

Why isn’t the slice updated in this example?

```text
s := []int{1, 1, 1}
for _, n := range s {
    n += 1
}
fmt.Println(s)
// Output: [1 1 1]
```

### Answer <a id="answer"></a>

The range loop copies the values from the slice to a **local variable** `n`; updating `n` will not affect the slice.

Update the slice entries like this:

```text
s := []int{1, 1, 1}
for i := range s {
    s[i] += 1
}
fmt.Println(s)
// Output: [2 2 2]
```

See [4 basic range loop \(for-each\) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for all about range loops in Go.

## 23.Iteration variable doesn’t see change in range loop

Why doesn’t the iteration variable `x` notice that `a[1]` has been updated?

```text
var a [2]int
for _, x := range a {
    fmt.Println("x =", x)
    a[1] = 8
}
fmt.Println("a =", a)
```

```text
x = 0
x = 0        <- Why isn't this 8?
a = [0 8]
```

### Answer <a id="answer"></a>

The range expression `a` is evaluated once before beginning the loop and a copy of the array is used to generate the iteration values.

To avoid copying the array, iterate over a slice instead.

```text
var a [2]int
for _, x := range a[:] {
    fmt.Println("x =", x)
    a[1] = 8
}
fmt.Println("a =", a)
```

```text
x = 0
x = 8
a = [0 8]
```

See [4 basic range loop \(for-each\) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for all about range loops in Go.

## 24._**\(data race\)**_ Iteration variables and closures

Why does this program

```text
func main() {
    var wg sync.WaitGroup
    wg.Add(5)
    for i := 0; i < 5; i++ {
        go func() {
            fmt.Print(i)
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println()
}
```

print

```text
55555
```

\(A [WaitGroup](https://yourbasic.org/golang/wait-for-goroutines-waitgroup/) waits for a collection of goroutines to finish.\)

### Answer <a id="answer"></a>

There is a **data race**: the variable `i` is shared by six \(6\) goroutines.

> A data race occurs when two goroutines access the same variable concurrently and at least one of the accesses is a write.

To avoid this, use a local variable and pass the number as a parameter when starting the goroutine.

```text
func main() {
    var wg sync.WaitGroup
    wg.Add(5)
    for i := 0; i < 5; i++ {
        go func(n int) { // Use a local variable.
            fmt.Print(n)
            wg.Done()
        }(i)
    }
    wg.Wait()
    fmt.Println()
}
```

Example output:

```text
40123
```

It’s also possible to avoid this data race while still using a closure, but then we must take care to use a unique variable for each goroutine.

```text
func main() {
    var wg sync.WaitGroup
    wg.Add(5)
    for i := 0; i < 5; i++ {
        n := i // Create a unique variable for each closure.
        go func() {
            fmt.Print(n)
            wg.Done()
        }()
    }
    wg.Wait()
    fmt.Println()
}
```

See [Data races](https://yourbasic.org/golang/data-races-explained) for a detailed explanation of data races in Go.

## 25.No JSON in sight

Why does [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) produce empty structs in the JSON text output?

```text
type Person struct {
    name string
    age  int
}

p := Person{"Alice", 22}
jsonData, _ := json.Marshal(p)
fmt.Println(string(jsonData))
```

```text
{}
```

### Answer <a id="answer"></a>

Only **exported** fields of a Go struct will be present in the JSON output.

```text
type Person struct {
    Name string // Changed to capital N
    Age  int    // Changed to capital A
}

p := Person{"Alice", 22}

jsonData, _ := json.Marshal(p)
fmt.Println(string(jsonData))
```

```text
{"Name":"Alice","Age":22}
```

You can specify the JSON field name explicitly with a `json:` tag.

```text
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

p := Person{"Alice", 22}

jsonData, _ := json.Marshal(p)
fmt.Println(string(jsonData))
```

```text
{"name":"Alice","age":22}
```

See [JSON by example](https://yourbasic.org/golang/json-example/) for an extensive guide to the Go JSON library.

## 26.Is "three" a digit?

Why does the regular expression `[0-9]*`, which is supposed to match a string with zero or more digits, match a string with characters in it?

```text
matched, err := regexp.MatchString(`[0-9]*`, "12three45")
fmt.Println(matched) // true
fmt.Println(err)     // nil (regexp is valid)
```

### Answer <a id="answer"></a>

The function [`regexp.MatchString`](https://golang.org/pkg/regexp/#MatchString) \(as well as most functions in the [`regexp`](https://golang.org/pkg/regexp/) package\) does **substring** matching.

To check if a full string matches `[0-9]*`, anchor the start and the end of the regular expression:

* the caret ^ matches the beginning of a text or line,
* the dollar sign $ matches the end of a text.

```text
matched, err := regexp.MatchString(`^[0-9]*$`, "12three45")
fmt.Println(matched) // false
fmt.Println(err)     // nil (regexp is valid)
```

See this [Regexp in-depth tutorial](https://yourbasic.org/golang/regexp-cheat-sheet/) for a cheat sheet and plenty of code examples.

## 27.Nil is not nil

Why is nil not equal to nil in this example?

```text
func Foo() error {
    var err *os.PathError = nil
    // …
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)        // <nil>
    fmt.Println(err == nil) // false
}
```

### Answer <a id="answer"></a>

An interface value is equal to `nil` only if both its value and dynamic type are `nil`. In the example above, `Foo()` returns `[nil, *os.PathError]` and we compare it with `[nil, nil]`.  


You can think of the interface value `nil` as typed, and `nil` _without type_ doesn’t equal `nil` _with type_. If we convert `nil` to the correct type, the values are indeed equal.

```text
…
fmt.Println(err == (*os.PathError)(nil)) // true
…
```

#### A better approach <a id="a-better-approach"></a>

To avoid this problem use a variable of type `error` instead, for example a [named return value](https://yourbasic.org/golang/named-return-values-parameters/).

```text
func Foo() (err error) {
    // …
    return // err is unassigned and has zero value [nil, nil]
}

func main() {
    err := Foo()
    fmt.Println(err)        // <nil>
    fmt.Println(err == nil) // true
}
```

See [Interfaces in 5 easy steps](https://yourbasic.org/golang/interfaces-explained/) for an extensive guide to interfaces in Go.

## Go Tutorial

## Go beginner’s guide: top 4 resources to get you started\[already known\]

## How to use JSON with Go \[best practices\]

### Default types <a id="default-types"></a>

The default Go types for decoding and encoding JSON are

* `bool` for JSON booleans,
* `float64` for JSON numbers,
* `string` for JSON strings, and
* `nil` for JSON null.

Additionally, [`time.Time`](https://golang.org/pkg/time/#Time) and the numeric types in the [`math/big`](https://golang.org/pkg/math/big/) package can be automatically encoded as JSON strings.

Note that JSON doesn’t support basic integer types. They can often be approximated by floating-point numbers.

> Since software that implements IEEE 754-2008 binary64 \(double precision\) numbers is generally available and widely used, good interoperability can be achieved by implementations that expect no more precision or range than these provide \[…\]
>
> Note that when such software is used, numbers that are integers and are in the range \[-253 + 1, 253 - 1\] are interoperable in the sense that implementations will agree exactly on their numeric values.[RFC 7159: The JSON Data Interchange Format](https://tools.ietf.org/html/rfc7159#section-6)

### Encode \(marshal\) struct to JSON <a id="encode-marshal-struct-to-json"></a>

The [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) function in package [`encoding/json`](https://golang.org/pkg/encoding/json/) generates JSON data.

```text
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

basket := FruitBasket{
    Name:    "Standard",
    Fruit:   []string{"Apple", "Banana", "Orange"},
    Id:      999,
    private: "Second-rate",
    Created: time.Now(),
}

var jsonData []byte
jsonData, err := json.Marshal(basket)
if err != nil {
    log.Println(err)
}
fmt.Println(string(jsonData))
```

Output:

```text
{"Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999,"Created":"2018-04-09T23:00:00Z"}
```

Only data that can be represented as JSON will be encoded; see [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) for the complete rules.

* Only the exported \(public\) fields of a struct will be present in the JSON output. **Other fields are ignored**.
* A field with a `json:` **struct tag** is stored with its tag name instead of its variable name.
* Pointers will be encoded as the values they point to, or `null` if the pointer is `nil`.

### Pretty print <a id="pretty-print"></a>

Replace `json.Marshal` with [`json.MarshalIndent`](https://golang.org/pkg/encoding/json/#MarshalIndent) in the example above to indent the JSON output.

```text
jsonData, err := json.MarshalIndent(basket, "", "    ")
```

Output:

```text
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}
```

### Decode \(unmarshal\) JSON to struct <a id="decode-unmarshal-json-to-struct"></a>

The [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) function in package [`encoding/json`](https://golang.org/pkg/encoding/json/) parses JSON data.

```text
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64 `json:"ref"`
    Created time.Time
}

jsonData := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

var basket FruitBasket
err := json.Unmarshal(jsonData, &basket)
if err != nil {
    log.Println(err)
}
fmt.Println(basket.Name, basket.Fruit, basket.Id)
fmt.Println(basket.Created)
```

Output:

```text
Standard [Apple Banana Orange] 999
2018-04-09 23:00:00 +0000 UTC
```

Note that `Unmarshal` allocated a new slice all by itself. This is how unmarshaling works for slices, maps and pointers.

For a given JSON key `Foo`, `Unmarshal` will attempt to match the struct fields in this order:

1. an exported \(public\) field with a struct tag `json:"Foo"`,
2. an exported field named `Foo`, or
3. an exported field named `FOO`, `FoO`, or some other case-insensitive match.

Only fields thar are found in the destination type will be decoded:

* This is useful when you wish to pick only a few specific fields.
* In particular, any unexported fields in the destination struct will be unaffected.

### Arbitrary objects and arrays <a id="arbitrary-objects-and-arrays"></a>

The [encoding/json](https://golang.org/pkg/encoding/json/) package uses

* `map[string]interface{}` to store arbitrary JSON objects, and
* `[]interface{}` to store arbitrary JSON arrays.

It will unmarshal any valid JSON data into a plain `interface{}` value.

Consider this JSON data:

```text
{
    "Name": "Eve",
    "Age": 6,
    "Parents": [
        "Alice",
        "Bob"
    ]
}
```

The [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) function will parse it into a map whose keys are strings, and whose values are themselves stored as empty interface values:

```text
map[string]interface{}{
    "Name": "Eve",
    "Age":  6.0,
    "Parents": []interface{}{
        "Alice",
        "Bob",
    },
}
```

We can iterate through the map with a range statement and use a type switch to access its values.

```text
jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

var v interface{}
json.Unmarshal(jsonData, &v)
data := v.(map[string]interface{})

for k, v := range data {
    switch v := v.(type) {
    case string:
        fmt.Println(k, v, "(string)")
    case float64:
        fmt.Println(k, v, "(float64)")
    case []interface{}:
        fmt.Println(k, "(array):")
        for i, u := range v {
            fmt.Println("    ", i, u)
        }
    default:
        fmt.Println(k, v, "(unknown)")
    }
}
```

Output:

```text
Name Eve (string)
Age 6 (float64)
Parents (array):
     0 Alice
     1 Bob
```

### JSON file example <a id="json-file-example"></a>

The [`json.Decoder`](https://golang.org/pkg/encoding/json/#Decoder) and [`json.Encoder`](https://golang.org/pkg/encoding/json/#Encoder) types in package [`encoding/json`](https://golang.org/pkg/encoding/json/) offer support for reading and writing streams, e.g. files, of JSON data.

The code in this example

* reads a stream of JSON objects from a [Reader](https://yourbasic.org/golang/io-reader-interface-explained/) \([`strings.Reader`](https://golang.org/pkg/strings/#Reader)\),
* removes the `Age` field from each object,
* and then writes the objects to a [Writer](https://yourbasic.org/golang/io-writer-interface-explained/) \([`os.Stdout`](https://golang.org/pkg/os/#pkg-variables)\).

```text
const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
`
reader := strings.NewReader(jsonData)
writer := os.Stdout

dec := json.NewDecoder(reader)
enc := json.NewEncoder(writer)

for {
    // Read one JSON object and store it in a map.
    var m map[string]interface{}
    if err := dec.Decode(&m); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }

    // Remove all key-value pairs with key == "Age" from the map.
    for k := range m {
        if k == "Age" {
            delete(m, k)
        }
    }

    // Write the map as a JSON object.
    if err := enc.Encode(&m); err != nil {
        log.Println(err)
    }
}
```

Output:

```text
{"Name":"Alice"}
{"Name":"Bob"}
```

#### Further reading

[Tutorials](https://yourbasic.org/golang/tutorials/) for beginners and experienced developers alike: best practices and production-quality code examples.

## Regexp tutorial and cheat sheet

### Basics <a id="basics"></a>

The regular expression `a.b` matches any string that starts with an `a`, ends with a `b`, and has a single character in between \(the period matches any character\).

To check if there is a **substring** matching `a.b`, use the [regexp.MatchString](https://golang.org/pkg/regexp/#MatchString) function.

```text
matched, err := regexp.MatchString(`a.b`, "aaxbb")
fmt.Println(matched) // true
fmt.Println(err)     // nil (regexp is valid)
```

To check if a **full string** matches `a.b`, anchor the start and the end of the regexp:

* the caret `^` matches the beginning of a text or line,
* the dollar sign `$` matches the end of a text.

```text
matched, _ := regexp.MatchString(`^a.b$`, "aaxbb")
fmt.Println(matched) // false
```

Similarly, we can check if a string **starts with** or **ends with** a pattern by using only the start or end anchor.

#### Compile <a id="compile"></a>

For more complicated queries, you should compile a regular expression to create a [`Regexp`](https://golang.org/pkg/regexp/#Regexp) object. There are two options:

```text
re1, err := regexp.Compile(`regexp`) // error if regexp invalid
re2 := regexp.MustCompile(`regexp`)  // panic if regexp invalid
```

#### Raw strings <a id="raw-strings"></a>

It’s convenient to use ```raw strings``` when writing regular expressions, since both ordinary string literals and regular expressions use backslashes for special characters.

A [raw string](https://yourbasic.org/golang/multiline-string/#raw-string-literals), delimited by backticks, is interpreted literally and backslashes have no special meaning.

### Cheat sheet <a id="cheat-sheet"></a>

#### Choice and grouping <a id="choice-and-grouping"></a>

| Regexp | Meaning |
| :--- | :--- |
| `xy` | `x` followed by `y` |
| `x|y` | `x` or `y`, prefer `x` |
| `xy|z` | same as `(xy)|z` |
| `xy*` | same as `x(y*)` |

#### Repetition \(greedy and non-greedy\) <a id="repetition-greedy-and-non-greedy"></a>

| Regexp | Meaning |
| :--- | :--- |
| `x*` | zero or more x, prefer more |
| `x*?` | prefer fewer \(non-greedy\) |
| `x+` | one or more x, prefer more |
| `x+?` | prefer fewer \(non-greedy\) |
| `x?` | zero or one x, prefer one |
| `x??` | prefer zero |
| `x{n}` | exactly n x |

#### Character classes <a id="character-classes"></a>

| Expression | Meaning |
| :--- | :--- |
| `.` | any character |
| `[ab]` | the character a or b |
| `[^ab]` | any character except a or b |
| `[a-z]` | any character from a to z |
| `[a-z0-9]` | any character from a to z or 0 to 9 |
| `\d` | a digit: `[0-9]` |
| `\D` | a non-digit: `[^0-9]` |
| `\s` | a whitespace character: `[\t\n\f\r ]` |
| `\S` | a non-whitespace character: `[^\t\n\f\r ]` |
| `\w` | a word character: `[0-9A-Za-z_]` |
| `\W` | a non-word character: `[^0-9A-Za-z_]` |
| `\p{Greek}` | Unicode character class\* |
| `\pN` | one-letter name |
| `\P{Greek}` | negated Unicode character class\* |
| `\PN` | one-letter name |

\* [RE2: Unicode character class names](https://github.com/google/re2/wiki/Syntax)

#### Special characters <a id="special-characters"></a>

To match a **special character** `\^$.|?*+-[]{}()` literally, escape it with a backslash. For example `\{` matches an opening brace symbol.

Other escape sequences are:

| Symbol | Meaning |
| :--- | :--- |
| `\t` | horizontal tab = `\011` |
| `\n` | newline = `\012` |
| `\f` | form feed = `\014` |
| `\r` | carriage return = `\015` |
| `\v` | vertical tab = `\013` |
| `\123` | octal character code \(up to three digits\) |
| `\x7F` | hex character code \(exactly two digits\) |

#### Text boundary anchors <a id="text-boundary-anchors"></a>

| Symbol | Matches |
| :--- | :--- |
| `\A` | at beginning of text |
| `^` | at beginning of text or line |
| `$` | at end of text |
| `\z` |  |
| `\b` | at ASCII word boundary |
| `\B` | not at ASCII word boundary |

#### Case-insensitive and multiline matches <a id="case-insensitive-and-multiline-matches"></a>

To change the default matching behavior, you can add a set of flags to the beginning of a regular expression.

For example, the prefix `"(?is)"` makes the matching case-insensitive and lets `.` match `\n`. \(The default matching is case-sensitive and `.` doesn’t match `\n`.\)

| Flag | Meaning |
| :--- | :--- |
| `i` | case-insensitive |
| `m` | let `^` and `$` match begin/end line in addition to begin/end text \(multi-line mode\) |
| `s` | let `.` match `\n` \(single-line mode\) |

### Code examples <a id="code-examples"></a>

#### First match <a id="first-match"></a>

Use the [`FindString`](https://golang.org/pkg/regexp/#Regexp.FindString) method to find the **text of the first match**. If there is no match, the return value is an empty string.

```text
re := regexp.MustCompile(`foo.?`)
fmt.Printf("%q\n", re.FindString("seafood fool")) // "food"
fmt.Printf("%q\n", re.FindString("meat"))         // ""
```

#### Location <a id="location"></a>

Use the [`FindStringIndex`](https://golang.org/pkg/regexp/#Regexp.FindStringIndex) method to find `loc`, the **location of the first match**, in a string `s`. The match is at `s[loc[0]:loc[1]]`. A return value of nil indicates no match.

```text
re := regexp.MustCompile(`ab?`)
fmt.Println(re.FindStringIndex("tablett"))    // [1 3]
fmt.Println(re.FindStringIndex("foo") == nil) // true
```

#### All matches <a id="all-matches"></a>

Use the [`FindAllString`](https://golang.org/pkg/regexp/#Regexp.FindAllString) method to find the **text of all matches**. A return value of nil indicates no match.

The method takes an integer argument `n`; if `n >= 0`, the function returns at most `n` matches.

```text
re := regexp.MustCompile(`a.`)
fmt.Printf("%q\n", re.FindAllString("paranormal", -1)) // ["ar" "an" "al"]
fmt.Printf("%q\n", re.FindAllString("paranormal", 2))  // ["ar" "an"]
fmt.Printf("%q\n", re.FindAllString("graal", -1))      // ["aa"]
fmt.Printf("%q\n", re.FindAllString("none", -1))       // [] (nil slice)
```

#### Replace <a id="replace"></a>

Use the [`ReplaceAllString`](https://golang.org/pkg/regexp/#Regexp.ReplaceAllString) method to **replace the text of all matches**. It returns a copy, replacing all matches of the regexp with a replacement string.

```text
re := regexp.MustCompile(`ab*`)
fmt.Printf("%q\n", re.ReplaceAllString("-a-abb-", "T")) // "-T-T-"
```

#### Split <a id="split"></a>

Use the [`Split`](https://golang.org/pkg/regexp/#Regexp.Split) method to **slice a string into substrings** separated by the regexp. It returns a slice of the substrings between those expression matches. A return value of nil indicates no match.

The method takes an integer argument `n`; if `n >= 0`, the function returns at most `n` matches.

```text
a := regexp.MustCompile(`a`)
fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

zp := regexp.MustCompile(`z+`)
fmt.Printf("%q\n", zp.Split("pizza", -1)) // ["pi" "a"]
fmt.Printf("%q\n", zp.Split("pizza", 0))  // [] (nil slice)
fmt.Printf("%q\n", zp.Split("pizza", 1))  // ["pizza"]
fmt.Printf("%q\n", zp.Split("pizza", 2))  // ["pi" "a"]
```

**More functions**

There are 16 functions following the naming pattern

```text
Find(All)?(String)?(Submatch)?(Index)?
```

For example: `Find`, `FindAllString`, `FindStringIndex`, …

* If `All` is present, the function matches successive non-overlapping matches.
* `String` indicates that the argument is a string; otherwise it’s a byte slice.
* If `Submatch` is present, the return value is a slice of successive submatches. Submatches are matches of parenthesized subexpressions within the regular expression. See [`FindSubmatch`](https://golang.org/pkg/regexp/#Regexp.FindSubmatch) for an example.
* If `Index` is present, matches and submatches are identified by byte index pairs.

### Implementation <a id="implementation"></a>

* The [`regexp`](https://golang.org/pkg/regexp/) package implements regular expressions with [RE2](https://golang.org/s/re2syntax) syntax.
* It supports UTF-8 encoded strings and Unicode character classes.
* The implementation is very efficient: the running time is linear in the size of the input.
* Backreferences are not supported since they cannot be efficiently implemented.

#### Further reading <a id="further-reading"></a>

[Regular expression matching can be simple and fast \(but is slow in Java, Perl, PHP, Python, Ruby, …\)](https://swtch.com/~rsc/regexp/regexp1.html).

## CheetSheet

### String literals \(escape characters\) <a id="string-literals-escape-characters"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `""` |  | [Default zero value](https://yourbasic.org/golang/default-zero-value/) for type `string` |
| `"Japan 日本"` | Japan 日本 | Go code is [Unicode text encoded in UTF‑8](https://yourbasic.org/golang/rune/) |
| `"\xe6\x97\xa5"` | 日 | `\xNN` specifies a byte |
| `"\u65E5"` | 日 | `\uNNNN` specifies a Unicode value |
| `"\\"` | \ | Backslash |
| `"\""` | " | Double quote |
| `"\n"` |  | Newline |
| `"\t"` |  | Tab |
| ```\xe6``` | \xe6 | Raw string literal\* |
| `html.EscapeString("<>")` | &lt;&gt; | HTML escape for &lt;, &gt;, &, ' and " |
| `url.PathEscape("A B")` | A%20B | URL percent-encoding net/url |

\* In ```````` string literals, text is interpreted literally and backslashes have no special meaning. See [Escapes and multiline strings](https://yourbasic.org/golang/multiline-string/) for more on raw strings, escape characters and string encodings.

### Concatenate <a id="concatenate"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `"Ja" + "pan"` | Japan | Concatenation |

> **Performance tips**  
> See [3 tips for efficient string concatenation](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/) for how to best use a string builder to concatenate strings without redundant copying.

### Equal and compare \(ignore case\) <a id="equal-and-compare-ignore-case"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `"Japan" == "Japan"` | true | Equality |
| `strings.EqualFold("Japan", "JAPAN")` | true | Unicode case folding |
| `"Japan" < "japan"` | true | Lexicographic order |

### Length in bytes or runes <a id="length-in-bytes-or-runes"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `len("日")` | 3 | Length in bytes |
| `utf8.RuneCountInString("日")` | 1 | in runes unicode/utf8 |
| `utf8.ValidString("日")` | true | UTF-8? unicode/utf8 |

### Index, substring, iterate <a id="index-substring-iterate"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `"Japan"[2]` | 'p' | Byte at position 2 |
| `"Japan"[1:3]` | ap | Byte indexing |
| `"Japan"[:2]` | Ja |  |
| `"Japan"[2:]` | pan |  |

A Go [range loop](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) iterates over UTF-8 encoded characters \([runes](https://yourbasic.org/golang/rune/)\):

```text
for i, ch := range "Japan 日本" {
    fmt.Printf("%d:%q ", i, ch)
}
// Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'
```

Iterating over bytes produces nonsense characters for non-ASCII text:

```text
s := "Japan 日本"
for i := 0; i < len(s); i++ {
    fmt.Printf("%q ", s[i])
}
// Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'
```

### Search \(contains, prefix/suffix, index\) <a id="search-contains-prefix-suffix-index"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Contains("Japan", "abc")` | false | Is abc in Japan? |
| `strings.ContainsAny("Japan", "abc")` | true | Is a, b or c in Japan? |
| `strings.Count("Banana", "ana")` | 1 | Non-overlapping instances of ana |
| `strings.HasPrefix("Japan", "Ja")` | true | Does Japan start with Ja? |
| `strings.HasSuffix("Japan", "pan")` | true | Does Japan end with pan? |
| `strings.Index("Japan", "abc")` | -1 | Index of first abc |
| `strings.IndexAny("Japan", "abc")` | 1 | a, b or c |
| `strings.LastIndex("Japan", "abc")` | -1 | Index of last abc |
| `strings.LastIndexAny("Japan", "abc")` | 3 | a, b or c |

### Replace \(uppercase/lowercase, trim\) <a id="replace-uppercase-lowercase-trim"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Replace("foo", "o", ".", 2)` | f.. | Replace first two “o” with “.” Use -1 to replace all |
| `f := func(r rune) rune {     return r + 1 } strings.Map(f, "ab")` | bc | Apply function to each character |
| `strings.ToUpper("Japan")` | JAPAN | Uppercase |
| `strings.ToLower("Japan")` | japan | Lowercase |
| `strings.Title("ja pan")` | Ja Pan | Initial letters to uppercase |
| `strings.TrimSpace(" foo\n")` | foo | Strip leading and trailing white space |
| `strings.Trim("foo", "fo")` |  | Strip _leading and trailing_ f:s and o:s |
| `strings.TrimLeft("foo", "f")` | oo | _only leading_ |
| `strings.TrimRight("foo", "o")` | f | _only trailing_ |
| `strings.TrimPrefix("foo", "fo")` | o |  |
| `strings.TrimSuffix("foo", "o")` | fo |  |

### Split by space or comma <a id="split-by-space-or-comma"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Fields(" a\t b\n")` | `["a" "b"]` | Remove white space |
| `strings.Split("a,b", ",")` | `["a" "b"]` | Remove separator |
| `strings.SplitAfter("a,b", ",")` | `["a," "b"]` | Keep separator |

### Join strings with separator <a id="join-strings-with-separator"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Join([]string{"a", "b"}, ":")` | a:b | Add separator |
| `strings.Repeat("da", 2)` | dada | 2 copies of “da” |

### Format and convert <a id="format-and-convert"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strconv.Itoa(-42)` | `"-42"` | Int to string |
| `strconv.FormatInt(255, 16)` | `"ff"` | Base 16 |

#### Sprintf <a id="sprintf"></a>

The [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function is often your best friend when formatting data:

```text
s := fmt.Sprintf("%.4f", math.Pi) // s == "3.1416
```

This [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) covers the most common formatting flags.

### Regular expression <a id="regular-expressions"></a>

For more advanced string handling, see this [Regular expressions tutorial](https://yourbasic.org/golang/regexp-cheat-sheet/), a gentle introduction to the `regexp` package with cheat sheet and plenty of examples.

## Conversions \[complete list\]



### Basics <a id="basics"></a>

The expression `T(x)` converts the value `x` to the type `T`.

```text
x := 5.1
n := int(x) // convert float to int
```

The conversion rules are extensive but predictable:

* all conversions between typed expressions must be explicitly stated,
* illegal conversions are caught by the compiler.

Conversions to and from numbers and strings may **change the representation** and have a **run-time cost**. All other conversions only change the type but not the representation of `x`.

### Interfaces <a id="interfaces"></a>

> To “convert” an [interface](https://yourbasic.org/golang/interfaces-explained/) to a string, struct or map you should use a **type assertion** or a **type switch**. A type assertion doesn’t really convert an interface to another data type, but it provides access to an interface’s concrete value, which is typically what you want.[Type assertions and type switches](https://yourbasic.org/golang/type-assertion-switch/)

### Integers <a id="integers"></a>

* When converting to a shorter integer type, the value is **truncated** to fit in the result type’s size.
* When converting to a longer integer type,
  * if the value is a signed integer, it is [**sign extended**](https://en.wikipedia.org/wiki/Sign_extension);
  * otherwise it is **zero extended**.

```text
a := uint16(0x10fe) // 0001 0000 1111 1110
b := int8(a)        //           1111 1110 (truncated to -2)
c := uint16(b)      // 1111 1111 1111 1110 (sign extended to 0xfffe)
```

### Floats <a id="floats"></a>

* When converting a floating-point number to an integer, the **fraction is discarded** \(truncation towards zero\).
* When converting an integer or floating-point number to a floating-point type, the result value is **rounded to the precision** specified by the destination type.

```text
var x float64 = 1.9
n := int64(x) // 1
n = int64(-x) // -1

n = 1234567890
y := float32(n) // 1.234568e+09
```

> **Warning:** In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.[The Go Programming Language Specification: Conversions](https://golang.org/ref/spec#Conversions)

### Integer to string <a id="integer-to-string"></a>

* When converting an integer to a string, the value is interpreted as a Unicode code point, and the resulting string will contain the character represented by that code point, encoded in UTF-8.
* If the value does not represent a valid code point \(for instance if it’s negative\), the result will be `"\ufffd"`, the Unicode replacement character �.

```text
string(97) // "a"
string(-1) // "\ufffd" == "\xef\xbf\xbd"
```

> Use [`strconv.Itoa`](https://golang.org/pkg/strconv/#Itoa) to get the decimal string representation of an integer.
>
> ```text
> strconv.Itoa(97) // "97"
> ```

### Strings and byte slices <a id="strings-and-byte-slices"></a>

* Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.
* Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string.

```text
string([]byte{97, 230, 151, 165}) // "a日"
[]byte("a日")                     // []byte{97, 230, 151, 165}
```

### Strings and rune slices <a id="strings-and-rune-slices"></a>

* Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.
* Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.

```text
string([]rune{97, 26085}) // "a日"
[]rune("a日")             // []rune{97, 26085}
```

### Underlying type <a id="underlying-type"></a>

A non-constant value can be converted to type `T` if it has the same underlying type as `T`.

In this example, the underlying type of `int64`, `T1`, and `T2` is `int64`.

```text
type (
	T1 int64
	T2 T1
)
```

It’s idiomatic in Go to convert the type of an expression to access a specific method.

```text
var n int64 = 12345
fmt.Println(n)                // 12345
fmt.Println(time.Duration(n)) // 12.345µs
```

\(The underlying type of [`time.Duration`](https://golang.org/pkg/time/#Duration) is `int64`, and the `time.Duration` type has a [`String`](https://golang.org/pkg/time/#Duration.String) method that returns the duration formatted as a time.\)

### Implicit conversions <a id="implicit-conversions"></a>

The only implicit conversion in Go is when an untyped constant is used in a situation where a type is required.

In this example the untyped literals `1` and `2` are implicitly converted.

```text
var x float64
x = 1 // Same as x = float64(1)

t := 2 * time.Second // Same as t := time.Duration(2) * time.Second
```

\(The implicit conversions are necessary since there is no mixing of numeric types in Go. You can only multiply a `time.Duration` with another `time.Duration`.\)

When the type can’t be inferred from the context, an untyped constant is converted to a `bool`, `int`, `float64`, `complex128`, `string` or `rune` depending on the syntactical format of the constant.

```text
n := 1   // Same as n := int(1)
x := 1.0 // Same as x := float64(1.0)
s := "A" // Same as s := string("A")
c := 'A' // Same as c := rune('A')
```

Illegal implicit conversions are caught by the compiler.

```text
var b byte = 256 // Same as var b byte = byte(256)
```

```text
../main.go:2:6: constant 256 overflows byte
```

### Pointers <a id="pointers"></a>

The Go compiler does not allow conversions between pointers and integers. Package [`unsafe`](https://golang.org/pkg/unsafe/) implements this functionality under restricted circumstances.

> **Warning:** The built-in package unsafe, known to the compiler, provides facilities for low-level programming including operations that violate the type system. A package using unsafe must be vetted manually for type safety and may not be portable.[The Go Programming Language Specification: Package unsafe](https://golang.org/ref/spec#Package_unsafe)



## fmt.Printf formatting tutorial and cheat sheet

### Basics <a id="basics"></a>

With the Go [`fmt`](https://golang.org/pkg/fmt) package you can format numbers and strings padded with spaces or zeroes, in different bases, and with optional quotes.

You submit a **template string** that contains the text you want to format plus some **annotation verbs** that tell the `fmt` functions how to format the trailing arguments.

#### Printf <a id="printf"></a>

In this example, [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf) formats and writes to standard output:

```text
fmt.Printf("Binary: %b\\%b", 4, 5) // Prints `Binary: 100\101`
```

* the **template string** is `"Binary: %b\\%b"`,
* the **annotation verb** `%b` formats a number in binary, and
* the **special value** `\\` is a backslash.

As a special case, the verb `%%`, which consumes no argument, produces a percent sign:

```text
fmt.Printf("%d %%", 50) // Prints `50 %`
```

#### Sprintf \(format without printing\) <a id="sprintf-format-without-printing"></a>

Use [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) to format a string without printing it:

```text
s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
```

#### Find fmt errors with vet <a id="find-fmt-errors-with-vet"></a>

If you try to compile and run this incorrect line of code

```text
fmt.Printf("Binary: %b\\%b", 4) // An argument to Printf is missing.
```

you’ll find that the program will compile, and then print

```text
Binary: 100\%!b(MISSING)
```

To catch this type of errors early, you can use the [vet command](https://golang.org/cmd/vet/) – it can find calls whose arguments do not align with the format string.

```text
$ go vet example.go
example.go:8: missing argument for Printf("%b"): format reads arg 2, have only 1 args
```

### Cheat sheet <a id="cheat-sheet"></a>

#### Default formats and type <a id="default"></a>

* **Value:** `[]int64{0, 1}`

| Format | Verb | Description |
| :--- | :--- | :--- |
| \[0 1\] | `%v` | Default format |
| \[\]int64{0, 1} | `%#v` | Go-syntax format |
| \[\]int64 | `%T` | The type of the value |

#### Integer \(indent, base, sign\) <a id="integer"></a>

* **Value:** `15`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 15 | `%d` | Base 10 |
| +15 | `%+d` | Always show sign |
| ␣␣15 | `%4d` | Pad with spaces \(width 4, right justified\) |
| 15␣␣ | `%-4d` | Pad with spaces \(width 4, left justified\) |
| 0015 | `%04d` | Pad with zeroes \(width 4\) |
| 1111 | `%b` | Base 2 |
| 17 | `%o` | Base 8 |
| f | `%x` | Base 16, lowercase |
| F | `%X` | Base 16, uppercase |
| 0xf | `%#x` | Base 16, with leading 0x |

#### Character \(quoted, Unicode\) <a id="character"></a>

* **Value:** `65`   \(Unicode letter A\)

| Format | Verb | Description |
| :--- | :--- | :--- |
| A | `%c` | Character |
| 'A' | `%q` | Quoted character |
| U+0041 | `%U` | Unicode |
| U+0041 'A' | `%#U` | Unicode with character |

#### Boolean \(true/false\) <a id="boolean"></a>

Use `%t` to format a boolean as `true` or `false`.

#### Pointer \(hex\) <a id="pointer"></a>

Use `%p` to format a pointer in base 16 notation with leading `0x`.

#### Float \(indent, precision, scientific notation\) <a id="float"></a>

* **Value:** `123.456`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 1.234560e+02 | `%e` | Scientific notation |
| 123.456000 | `%f` | Decimal point, no exponent |
| 123.46 | `%.2f` | Default width, precision 2 |
| ␣␣123.46 | `%8.2f` | Width 8, precision 2 |
| 123.456 | `%g` | Exponent as needed, necessary digits only |

#### String or byte slice \(quote, indent, hex\) <a id="string-or-byte-slice"></a>

* **Value:** `"café"`

| Format | Verb | Description |
| :--- | :--- | :--- |
| café | `%s` | Plain string |
| ␣␣café | `%6s` | Width 6, right justify |
| café␣␣ | `%-6s` | Width 6, left justify |
| "café" | `%q` | Quoted string |
| 636166c3a9 | `%x` | Hex dump of byte values |
| 63 61 66 c3 a9 | `% x` | Hex dump with spaces |

#### Special values <a id="special-values"></a>

| Value | Description |
| :--- | :--- |
| `\a` | U+0007 alert or bell |
| `\b` | U+0008 backspace |
| `\\` | U+005c backslash |
| `\t` | U+0009 horizontal tab |
| `\n` | U+000A line feed or newline |
| `\f` | U+000C form feed |
| `\r` | U+000D carriage return |
| `\v` | U+000b vertical tab |

Arbitrary values can be encoded with backslash escapes and can be used in any `""` string literal.

There are four different formats:

* `\x` followed by exactly two hexadecimal digits,
* `\` followed by exactly three octal digits,
* `\u` followed by exactly four hexadecimal digits,
* `\U` followed by exactly eight hexadecimal digits.

The escapes `\u` and `\U` represent Unicode code points.

```text
fmt.Println("\\caf\u00e9") // Prints \café
```

### Further readin <a id="further-reading"></a>

[40+ practical string tips \[cheat sheet\]](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)

## Format a time or date \[complete guide\]

### Basic example <a id="basic-example"></a>

Go doesn’t use yyyy-mm-dd layout to format a time. Instead, you format a special **layout parameter**

`Mon Jan 2 15:04:05 MST 2006`

the same way as the time or date should be formatted. \(This date is easier to remember when written as `01/02 03:04:05PM ‘06 -0700`.\)

```text
const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
)
date := "1999-12-31"
t, _ := time.Parse(layoutISO, date)
fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
fmt.Println(t.Format(layoutUS)) // December 31, 1999
```

The function

* [`time.Parse`](https://golang.org/pkg/time/#Parse) parses a date string, and
* [`Format`](https://golang.org/pkg/time/#Time.Format) formats a [`time.Time`](https://golang.org/pkg/time/#Time).

They have the following signatures:

```text
func Parse(layout, value string) (Time, error)
func (t Time) Format(layout string) string
```

### Standard time and date formats <a id="standard-time-and-date-formats"></a>

| Go layout | Note |
| :--- | :--- |
| `January 2, 2006` | Date |
| `01/02/06` |  |
| `Jan-02-06` |  |
| `15:04:05` | Time |
| `3:04:05 PM` |  |
| `Jan _2 15:04:05` | Timestamp |
| `Jan _2 15:04:05.000000` | with microseconds |
| `2006-01-02T15:04:05-0700` | [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) \([RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)\) |
| `2006-01-02` |  |
| `15:04:05` |  |
| `02 Jan 06 15:04 MST` | [RFC 822](https://www.ietf.org/rfc/rfc822.txt) |
| `02 Jan 06 15:04 -0700` | with numeric zone |
| `Mon, 02 Jan 2006 15:04:05 MST` | [RFC 1123](https://www.ietf.org/rfc/rfc1123.txt) |
| `Mon, 02 Jan 2006 15:04:05 -0700` | with numeric zone |

The following predefined date and timestamp [format constants](https://golang.org/pkg/time/#pkg-constants) are also available.

```text
ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
```

### Layout options <a id="layout-options"></a>

| Type | Options |
| :--- | :--- |
| Year | `06`   `2006` |
| Month | `01`   `1`   `Jan`   `January` |
| Day | `02`   `2`   `_2`   \(width two, right justified\) |
| Weekday | `Mon`   `Monday` |
| Hours | `03`   `3`   `15` |
| Minutes | `04`   `4` |
| Seconds | `05`   `5` |
| ms μs ns | `.000`   `.000000`   `.000000000` |
| ms μs ns | `.999`   `.999999`   `.999999999`   \(trailing zeros removed\) |
| am/pm | `PM`   `pm` |
| Timezone | `MST` |
| Offset | `-0700`   `-07`   `-07:00`   `Z0700`   `Z07:00` |

### Corner cases <a id="corner-cases"></a>

It’s not possible to specify that an hour should be rendered without a leading zero in a 24-hour time format.

It’s not possible to specify midnight as `24:00` instead of `00:00`. A typical usage for this would be giving opening hours ending at midnight, such as `07:00-24:00`.

It’s not possible to specify a time containing a leap second: `23:59:60`. In fact, the time package assumes a Gregorian calendar without leap seconds.

## Regexp tutorial and cheat sheet

A regular expression is a sequence of characters that define a search pattern.

### Basics <a id="basics"></a>

The regular expression `a.b` matches any string that starts with an `a`, ends with a `b`, and has a single character in between \(the period matches any character\).

To check if there is a **substring** matching `a.b`, use the [regexp.MatchString](https://golang.org/pkg/regexp/#MatchString) function.

```text
matched, err := regexp.MatchString(`a.b`, "aaxbb")
fmt.Println(matched) // true
fmt.Println(err)     // nil (regexp is valid)
```

To check if a **full string** matches `a.b`, anchor the start and the end of the regexp:

* the caret `^` matches the beginning of a text or line,
* the dollar sign `$` matches the end of a text.

```text
matched, _ := regexp.MatchString(`^a.b$`, "aaxbb")
fmt.Println(matched) // false
```

Similarly, we can check if a string **starts with** or **ends with** a pattern by using only the start or end anchor.

#### Compile <a id="compile"></a>

For more complicated queries, you should compile a regular expression to create a [`Regexp`](https://golang.org/pkg/regexp/#Regexp) object. There are two options:

```text
re1, err := regexp.Compile(`regexp`) // error if regexp invalid
re2 := regexp.MustCompile(`regexp`)  // panic if regexp invalid
```

#### Raw strings <a id="raw-strings"></a>

It’s convenient to use ```raw strings``` when writing regular expressions, since both ordinary string literals and regular expressions use backslashes for special characters.

A [raw string](https://yourbasic.org/golang/multiline-string/#raw-string-literals), delimited by backticks, is interpreted literally and backslashes have no special meaning.

### Cheat sheet <a id="cheat-sheet"></a>

#### Choice and grouping <a id="choice-and-grouping"></a>

| Regexp | Meaning |
| :--- | :--- |
| `xy` | `x` followed by `y` |
| `x|y` | `x` or `y`, prefer `x` |
| `xy|z` | same as `(xy)|z` |
| `xy*` | same as `x(y*)` |

#### Repetition \(greedy and non-greedy\) <a id="repetition-greedy-and-non-greedy"></a>

| Regexp | Meaning |
| :--- | :--- |
| `x*` | zero or more x, prefer more |
| `x*?` | prefer fewer \(non-greedy\) |
| `x+` | one or more x, prefer more |
| `x+?` | prefer fewer \(non-greedy\) |
| `x?` | zero or one x, prefer one |
| `x??` | prefer zero |
| `x{n}` | exactly n x |

#### Character classes <a id="character-classes"></a>

| Expression | Meaning |
| :--- | :--- |
| `.` | any character |
| `[ab]` | the character a or b |
| `[^ab]` | any character except a or b |
| `[a-z]` | any character from a to z |
| `[a-z0-9]` | any character from a to z or 0 to 9 |
| `\d` | a digit: `[0-9]` |
| `\D` | a non-digit: `[^0-9]` |
| `\s` | a whitespace character: `[\t\n\f\r ]` |
| `\S` | a non-whitespace character: `[^\t\n\f\r ]` |
| `\w` | a word character: `[0-9A-Za-z_]` |
| `\W` | a non-word character: `[^0-9A-Za-z_]` |
| `\p{Greek}` | Unicode character class\* |
| `\pN` | one-letter name |
| `\P{Greek}` | negated Unicode character class\* |
| `\PN` | one-letter name |

\* [RE2: Unicode character class names](https://github.com/google/re2/wiki/Syntax)

#### Special characters <a id="special-characters"></a>

To match a **special character** `\^$.|?*+-[]{}()` literally, escape it with a backslash. For example `\{` matches an opening brace symbol.

Other escape sequences are:

| Symbol | Meaning |
| :--- | :--- |
| `\t` | horizontal tab = `\011` |
| `\n` | newline = `\012` |
| `\f` | form feed = `\014` |
| `\r` | carriage return = `\015` |
| `\v` | vertical tab = `\013` |
| `\123` | octal character code \(up to three digits\) |
| `\x7F` | hex character code \(exactly two digits\) |

#### Text boundary anchors <a id="text-boundary-anchors"></a>

| Symbol | Matches |
| :--- | :--- |
| `\A` | at beginning of text |
| `^` | at beginning of text or line |
| `$` | at end of text |
| `\z` |  |
| `\b` | at ASCII word boundary |
| `\B` | not at ASCII word boundary |

#### Case-insensitive and multiline matches <a id="case-insensitive-and-multiline-matches"></a>

To change the default matching behavior, you can add a set of flags to the beginning of a regular expression.

For example, the prefix `"(?is)"` makes the matching case-insensitive and lets `.` match `\n`. \(The default matching is case-sensitive and `.` doesn’t match `\n`.\)

| Flag | Meaning |
| :--- | :--- |
| `i` | case-insensitive |
| `m` | let `^` and `$` match begin/end line in addition to begin/end text \(multi-line mode\) |
| `s` | let `.` match `\n` \(single-line mode\) |

### Code examples <a id="code-examples"></a>

#### First match <a id="first-match"></a>

Use the [`FindString`](https://golang.org/pkg/regexp/#Regexp.FindString) method to find the **text of the first match**. If there is no match, the return value is an empty string.

```text
re := regexp.MustCompile(`foo.?`)
fmt.Printf("%q\n", re.FindString("seafood fool")) // "food"
fmt.Printf("%q\n", re.FindString("meat"))         // ""
```

#### Location <a id="location"></a>

Use the [`FindStringIndex`](https://golang.org/pkg/regexp/#Regexp.FindStringIndex) method to find `loc`, the **location of the first match**, in a string `s`. The match is at `s[loc[0]:loc[1]]`. A return value of nil indicates no match.

```text
re := regexp.MustCompile(`ab?`)
fmt.Println(re.FindStringIndex("tablett"))    // [1 3]
fmt.Println(re.FindStringIndex("foo") == nil) // true
```

#### All matches <a id="all-matches"></a>

Use the [`FindAllString`](https://golang.org/pkg/regexp/#Regexp.FindAllString) method to find the **text of all matches**. A return value of nil indicates no match.

The method takes an integer argument `n`; if `n >= 0`, the function returns at most `n` matches.

```text
re := regexp.MustCompile(`a.`)
fmt.Printf("%q\n", re.FindAllString("paranormal", -1)) // ["ar" "an" "al"]
fmt.Printf("%q\n", re.FindAllString("paranormal", 2))  // ["ar" "an"]
fmt.Printf("%q\n", re.FindAllString("graal", -1))      // ["aa"]
fmt.Printf("%q\n", re.FindAllString("none", -1))       // [] (nil slice)
```

#### Replace <a id="replace"></a>

Use the [`ReplaceAllString`](https://golang.org/pkg/regexp/#Regexp.ReplaceAllString) method to **replace the text of all matches**. It returns a copy, replacing all matches of the regexp with a replacement string.

```text
re := regexp.MustCompile(`ab*`)
fmt.Printf("%q\n", re.ReplaceAllString("-a-abb-", "T")) // "-T-T-"
```

#### Split <a id="split"></a>

Use the [`Split`](https://golang.org/pkg/regexp/#Regexp.Split) method to **slice a string into substrings** separated by the regexp. It returns a slice of the substrings between those expression matches. A return value of nil indicates no match.

The method takes an integer argument `n`; if `n >= 0`, the function returns at most `n` matches.

```text
a := regexp.MustCompile(`a`)
fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

zp := regexp.MustCompile(`z+`)
fmt.Printf("%q\n", zp.Split("pizza", -1)) // ["pi" "a"]
fmt.Printf("%q\n", zp.Split("pizza", 0))  // [] (nil slice)
fmt.Printf("%q\n", zp.Split("pizza", 1))  // ["pizza"]
fmt.Printf("%q\n", zp.Split("pizza", 2))  // ["pi" "a"]
```

**More functions**

There are 16 functions following the naming pattern

```text
Find(All)?(String)?(Submatch)?(Index)?
```

For example: `Find`, `FindAllString`, `FindStringIndex`, …

* If `All` is present, the function matches successive non-overlapping matches.
* `String` indicates that the argument is a string; otherwise it’s a byte slice.
* If `Submatch` is present, the return value is a slice of successive submatches. Submatches are matches of parenthesized subexpressions within the regular expression. See [`FindSubmatch`](https://golang.org/pkg/regexp/#Regexp.FindSubmatch) for an example.
* If `Index` is present, matches and submatches are identified by byte index pairs.

### Implementation <a id="implementation"></a>

* The [`regexp`](https://golang.org/pkg/regexp/) package implements regular expressions with [RE2](https://golang.org/s/re2syntax) syntax.
* It supports UTF-8 encoded strings and Unicode character classes.
* The implementation is very efficient: the running time is linear in the size of the input.
* Backreferences are not supported since they cannot be efficiently implemented.

#### Further reading <a id="further-reading"></a>

[Regular expression matching can be simple and fast \(but is slow in Java, Perl, PHP, Python, Ruby, …\)](https://swtch.com/~rsc/regexp/regexp1.html).

##  Bitwise operators \[cheat sheet\]

### Number literals <a id="number-literals"></a>

The binary number 100002 can be written as `020`, `16` or `0x10` in Go.

| Literal | Base | Note |
| :--- | :--- | :--- |
| `020` | 8 | Starts with `0` |
| `16` | 10 | Never starts with `0` |
| `0x10` | 16 | Starts with `0x` |

### Built-in operators <a id="built-in-operators"></a>

| Operation | Result | Description |
| :--- | :--- | :--- |
| `0011 & 0101` | 0001 | Bitwise AND |
| `0011 | 0101` | 0111 | Bitwise OR |
| `0011 ^ 0101` | 0110 | Bitwise XOR |
| `^0101` | 1010 | Bitwise NOT \(same as `1111 ^ 0101`\) |
| `0011 &^ 0101` | 0010 | Bitclear \(AND NOT\) |
| `00110101<<2` | 11010100 | Left shift |
| `00110101<<100` | 00000000 | No upper limit on shift count |
| `00110101>>2` | 00001101 | Right shift |

* The binary numbers in the examples are for explanation only. Integer literals in Go must be specified in octal, decimal or hexadecimal.
* The bitwise operators take both signed and unsigned integers as input. The right-hand side of a shift operator, however, must be an unsigned integer.
* Shift operators implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer.

### Package [`math/bits`](https://golang.org/pkg/math/bits/) <a id="package-math-bits"></a>

| Operation | Result | Description |
| :--- | :--- | :--- |
| `bits.UintSize` | 32 or 64 | Size of a uint in bits |
| `bits.OnesCount8(00101110)` | 4 | Number of one bits \(population count\) |
| `bits.Len8(00101110)` | 6 | Bits required to represent number |
| `bits.Len8(00000000)` | 0 |  |
| `bits.LeadingZeros8(00101110)` | 2 | Number of leading zero bits |
| `bits.LeadingZeros8(00000000)` | 8 |  |
| `bits.TrailingZeros8(00101110)` | 1 | Number of trailing zero bits |
| `bits.TrailingZeros8(00000000)` | 8 |  |
| `bits.RotateLeft8(00101110, 3)` | 01110001 | The value rotated left by 3 bits |
| `bits.RotateLeft8(00101110, -3)` | 11000101 | The value rotated **right** by 3 bits |
| `bits.Reverse8(00101110)` | 01110100 | Bits in reversed order |
| `bits.ReverseBytes16(0x00ff)` | `0xff00` | Bytes in reversed order |

* The functions operate on **unsigned integers**.
* They come in different forms that take arguments of different sizes. For example, `Len`, `Len8`, `Len16`, `Len32`, and `Len64` apply to the types `uint`, `uint8`, `uint16`, `uint32`, and `uint64`, respectively.
* The functions are recognized by the compiler and on most architectures they are treated as [intrinsics](https://dave.cheney.net/2019/08/20/go-compiler-intrinsics) for additional performance.

### Bit manipulation code exampl <a id="bit-manipulation-code-example"></a>

[Bitmasks, bitsets and flags](https://yourbasic.org/golang/bitmask-flag-set-clear/) shows how to implement a bitmask, a small set of booleans, often called flags, represented by the bits in a single number.

## Start a new Go project \[standard layout\]

The repository at [github.com/yourbasic/fenwick](https://github.com/yourbasic/fenwick) is a small but complete Go library [package](https://yourbasic.org/golang/packages-explained/). It shows the structure of a basic project and can be used as a template.

In addition to source code and resources, it includes

* a [README](https://github.com/yourbasic/fenwick/blob/master/README.md) file with sections on
  * installation,
  * documentation and
  * compatibility policy,
* unit tests,
* benchmarks,
* godoc links,
* a testable doc example and
* a licence.

#### Further readin <a id="further-reading"></a>

[Your basic API](https://yourbasic.org/algorithms/your-basic-api/) is an introduction to API design with examples in Go and Java.  


## Learn to love your compiler

The Go compiler sometimes confuses and annoys developers who are new to the language.

This is a list of short articles with strategies and workarounds for common compiler error messages that tend to confuse fresh Go programmers.

* [`imported and not used`](https://yourbasic.org/golang/unused-imports/)Programs with unused imports won't compile.
* [`declared and not used`](https://yourbasic.org/golang/unused-local-variables/)You must use all local variables.
* [`multiple-value in single-value context`](https://yourbasic.org/golang/gotcha-multiple-value-single-value-context/)When a function returns multiple values, you must use all of them.
* [`syntax error: unexpected newline, expecting comma or }`](https://yourbasic.org/golang/gotcha-missing-comma-slice-array-map-literal/)In a multi-line slice, array or map literal, every line must end with a comma.
* [`cannot assign to …`](https://yourbasic.org/golang/gotcha-strings-are-immutable/)Go strings are immutable and behave like read-only byte slices.
* [`constant overflows int`](https://yourbasic.org/golang/gotcha-constant-overflows-int/)An untyped constant is converted before it is assigned to a variable.
* [`syntax error: unexpected ++, expecting expression comma or )`](https://yourbasic.org/golang/gotcha-increment-decrement-statement/)Increment and decrement operations can’t be used as expressions, only as statements.
* [`syntax error: non-declaration statement outside function body`](https://yourbasic.org/golang/short-variable-declaration-outside-function/)Short variable declarations can only be used inside functions.
* [`missing function body for …`](https://yourbasic.org/golang/opening-brace-separate-line/)An opening brace cannot appear on a line by itself.

#### Further reading

[Tutorials](https://yourbasic.org/golang/tutorials/) for beginners and experienced developers alike: best practices and production-quality code examples.

## Go Code For Common Tasks

## 2 basic FIFO queue implementations

A simple way to implement a temporary queue data structure in Go is to use a [slice](https://yourbasic.org/golang/slices-explained/):

* to enqueue you use the built-in `append` function, and
* to dequeue you slice off the first element.

```text
var queue []string

queue = append(queue, "Hello ") // Enqueue
queue = append(queue, "world!")

for len(queue) > 0 {
    fmt.Print(queue[0]) // First element
    queue = queue[1:]   // Dequeue
}
```

```text
Hello world!
```

### Watch out for memory leaks <a id="watch-out-for-memory-leaks"></a>

You may want to remove the first element before dequeuing.

```text
// Dequeue
queue[0] = "" // Erase element (write zero value)
queue = queue[1:]
```

> **Warning:** The memory allocated for the array is never returned. For a long-living queue you should probably use a dynamic data structure, such as a linked list.

### Linked list <a id="linked-list"></a>

The [`container/list`](https://golang.org/pkg/container/list/) package implements a doubly linked list which can be used as a queue.

```text
queue := list.New()

queue.PushBack("Hello ") // Enqueue
queue.PushBack("world!")

for queue.Len() > 0 {
    e := queue.Front() // First element
    fmt.Print(e.Value)

    queue.Remove(e) // Dequeue
}
```

```text
Hello world!
```

#### More code example

#### [Go blueprints: code for com­mon tasks](https://yourbasic.org/golang/blueprint/) is a collection of handy code examples.

## 2 basic set implementations

### Map implementation <a id="map-implementation"></a>

The idiomatic way to implement a set in Go is to use a [map](https://yourbasic.org/golang/maps-explained/).

```text
set := make(map[string]bool) // New empty set
set["Foo"] = true            // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")    // Delete
size := len(set)      // Size
exists := set["Foo"]  // Membership
```

#### Alternative <a id="alternative"></a>

If the memory used by the booleans is an issue, which seems unlikely, you could replace them with empty structs. In Go, an empty struct typically doesn’t use any memory.

```text
type void struct{}
var member void

set := make(map[string]void) // New empty set
set["Foo"] = member          // Add
for k := range set {         // Loop
    fmt.Println(k)
}
delete(set, "Foo")      // Delete
size := len(set)        // Size
_, exists := set["Foo"] // Membership
```

### Bitset implementation <a id="bitset-implementation"></a>

For small sets of integers, you may want to consider a **bitset**, a small set of booleans, often called flags, represented by the bits in a single number.

See [Bitmasks and flags](https://yourbasic.org/golang/bitmask-flag-set-clear/) for a complete example.

## A basic stack \(LIFO\) data structure

The idiomatic way to implement a stack data structure in Go is to use a [slice](https://yourbasic.org/golang/slices-explained/):

* to push you use the built-in [append function](https://yourbasic.org/golang/append-explained/), and
* to pop you slice off the top element.

```text
var stack []string

stack = append(stack, "world!") // Push
stack = append(stack, "Hello ")

for len(stack) > 0 {
    n := len(stack) - 1 // Top element
    fmt.Print(stack[n])

    stack = stack[:n] // Pop
}
```

```text
Hello world!
```

### Performanc <a id="performance"></a>

Appending a single element to a slice takes **constant amortized time**. See [Amortized time complexity](https://yourbasic.org/algorithms/amortized-time-complexity-analysis/) for a detailed explanation.

If the stack is permanent and the elements temporary, you may want to remove the top element before popping the stack to avoid memory leaks.

```text
// Pop
stack[n] = "" // Erase element (write zero value)
stack = stack[:n]
```

## Access environment variables

Use the [`os.Setenv`](https://golang.org/pkg/os/#Setenv), [`os.Getenv`](https://golang.org/pkg/os/#Getenv) and [`os.Unsetenv`](https://golang.org/pkg/os/#Unsetenv) functions to access environment variables.

```text
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/bash"

os.Unsetenv("SHELL")
fmt.Printf("%q\n", os.Getenv("SHELL")) // ""

os.Setenv("SHELL", "/bin/dash")
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/dash"
```

The [`os.Environ`](https://golang.org/pkg/os/#Environ) function returns a slice of `"key=value"` strings listing all environment variables.

```text
for _, s := range os.Environ() {
    kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
    fmt.Printf("key:%q value:%q\n", kv[0], kv[1])
}
```

```text
key:"SHELL" value:"/bin/bash"
key:"SESSION" value:"ubuntu"
key:"TERM" value:"xterm-256color"
key:"LANG" value:"en_US.UTF-8"
key:"XMODIFIERS" value:"@im=ibus"
…
```

## Access private fields with reflection

With reflection it's possible to read, _but not write_, unexported fields of a struct defined in another package.

In this example, we access the unexported field `len` in the `List` struct in [package](https://yourbasic.org/golang/packages-explained/) [`container/list`](https://golang.org/pkg/container/list/):

```text
package list

type List struct {
    root Element
    len  int
}
```

This code reads the value of `len` with reflection.

```text
package main

import (
    "container/list"
    "fmt"
    "reflect"
)

func main() {
    l := list.New()
    l.PushFront("foo")
    l.PushFront("bar")

    // Get a reflect.Value fv for the unexported field len.
    fv := reflect.ValueOf(l).Elem().FieldByName("len")
    fmt.Println(fv.Int()) // 2

    // Try to set the value of len.
    fv.Set(reflect.ValueOf(3)) // ILLEGAL
}
```

```text
2
panic: reflect: reflect.Value.Set using value obtained using unexported field

goroutine 1 [running]:
reflect.flag.mustBeAssignable(0x1a2, 0x285a)
	/usr/local/go/src/reflect/value.go:225 +0x280
reflect.Value.Set(0xee2c0, 0x10444254, 0x1a2, 0xee2c0, 0x1280c0, 0x82)
	/usr/local/go/src/reflect/value.go:1345 +0x40
main.main()
	../main.go:18 +0x280
```

## Bitmasks, bitsets and flags

### Bitmask <a id="bitmask"></a>

A bitmask is a small set of booleans, often called flags, represented by the bits in a single number.

```text
type Bits uint8

const (
    F0 Bits = 1 << iota
    F1
    F2
)

func Set(b, flag Bits) Bits    { return b | flag }
func Clear(b, flag Bits) Bits  { return b &^ flag }
func Toggle(b, flag Bits) Bits { return b ^ flag }
func Has(b, flag Bits) bool    { return b&flag != 0 }

func main() {
    var b Bits
    b = Set(b, F0)
    b = Toggle(b, F2)
    for i, flag := range []Bits{F0, F1, F2} {
        fmt.Println(i, Has(b, flag))
    }
}
```

```text
0 true
1 false
2 true
```

### Larger bitsets <a id="larger-bitsets"></a>

To represent larger sets of bits, you may want to use a custom data structure. The package [`github.com/yourbasic/bit`](https://github.com/yourbasic/bit) provides a bit array implementation and some utility bit functions.

Because a bit array uses bit-level parallelism, limits memory access, and efficiently uses the data cache, it often outperforms other data structures. Here is an example that shows how to create the set of all primes less than n in O\(n log log n\) time using the [`bit.Set`](https://godoc.org/github.com/yourbasic/bit#Set) data structure from package [`bit`](https://godoc.org/github.com/yourbasic/bit). Try the code with n equal to a few hundred millions and be pleasantly surprised.

```text
// Sieve of Eratosthenes
const n = 50
sieve := bit.New().AddRange(2, n)
sqrtN := int(math.Sqrt(n))
for p := 2; p <= sqrtN; p = sieve.Next(p) {
    for k := p * p; k < n; k += p {
        sieve.Delete(k)
    }
}
fmt.Println(sieve)
```

```text
{2 3 5 7 11 13 17 19 23 29 31 37 41 43 47}
```

## Check if a number is prime

### Ints <a id="ints"></a>

For integer types, use [`ProbablyPrime(0)`](https://golang.org/pkg/math/big/#Int.ProbablyPrime) from package [`math/big`](https://golang.org/pkg/math/big/). This primality test is 100% accurate for inputs less than 264.

```text
const n = 1212121
if big.NewInt(n).ProbablyPrime(0) {
    fmt.Println(n, "is prime")
} else {
    fmt.Println(n, "is not prime")
}
```

```text
1212121 is prime
```

### Larger numbers <a id="larger-numbers"></a>

For larger numbers, you need to provide the desired number of tests to [ProbablyPrime\(n\)](https://golang.org/pkg/math/big/#Int.ProbablyPrime). For n tests, the probability of returning true for a randomly chosen non-prime is at most \(1/4\)n. A common choice is to use n = 20; this gives a false positive rate 0.000,000,000,001.

```text
z := new(big.Int)
fmt.Sscan("170141183460469231731687303715884105727", z)
if z.ProbablyPrime(20) {
    fmt.Println(z, "is probably prime")
} else {
    fmt.Println(z, "is not prime")
}
```

```text
170141183460469231731687303715884105727 is probably prime
```

## Go as a scripting language: lightweight, safe and fast

This example is a simplified version of the Unix `grep` command. The program searches the input file for lines containing the given pattern and prints these lines.

```text
func main() {
    log.SetPrefix("grep: ")
    log.SetFlags(0) // no extra info in log messages

    if len(os.Args) != 3 {
        fmt.Printf("Usage: %v PATTERN FILE\n", os.Args[0])
        return
    }

    pattern, err := regexp.Compile(os.Args[1])
    if err != nil {
        log.Fatalln(err)
    }

    file, err := os.Open(os.Args[2])
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if pattern.MatchString(line) {
            fmt.Println(line)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
```

## Command-line arguments and flags

The [`os.Args`](https://golang.org/pkg/os/#pkg-variables) variable holds the command-line arguments – starting with the program name – which are passed to a Go program.

```text
func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage:", os.Args[0], "PATTERN", "FILE")
        return
    }
    pattern := os.Args[1]
    file := os.Args[2]
    // ...
}
```

```text
$ go build grep.go
$ ./grep
Usage: ./grep PATTERN FILE
```

#### Flag parsing <a id="flag-parsing"></a>

The [flag](https://golang.org/pkg/flag/) package implements basic command-line flag parsing.

## Compute absolute value of an int/float

### Integers <a id="integers"></a>

There is no built-in abs function for integers, but it’s simple to write your own.

```text
// Abs returns the absolute value of x.
func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}
```

> **Warning:** The smallest value of a signed integer doesn’t have a matching positive value.
>
> * [`math.MinInt64`](https://golang.org/pkg/math/#pkg-constants) is -922337203685477580**8**, but
> * [`math.MaxInt64`](https://golang.org/pkg/math/#pkg-constants) is 922337203685477580**7**.
>
> Unfortunately, our `Abs` function returns a **negative** value in this case.
>
> ```text
> fmt.Println(Abs(math.MinInt64))
> // Output: -9223372036854775808
> ```
>
> \(The Java and C libraries behave like this as well.\)

### Floats <a id="floats"></a>

The [`math.Abs`](https://golang.org/pkg/math/#Abs) function returns the absolute value of `x`.

```text
func Abs(x float64) float64
```

Special cases:

```text
Abs(±Inf) = +Inf
Abs(NaN) = NaN
```

## Compute max of two ints/floats

### Integers <a id="integers"></a>

There is no built-in max or min function for integers, but it’s simple to write your own.

```text
// Max returns the larger of x or y.
func Max(x, y int64) int64 {
    if x < y {
        return y
    }
    return x
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
    if x > y {
        return y
    }
    return x
}
```

### Floats <a id="floats"></a>

For floats, use [`math.Max`](https://golang.org/pkg/math/#Max) and [`math.Min`](https://golang.org/pkg/math/#Min).

```text
// Max returns the larger of x or y. 
func Max(x, y float64) float64

// Min returns the smaller of x or y.
func Min(x, y float64) float64
```

Special cases:

```text
Max(x, +Inf) = Max(+Inf, x) = +Inf
Max(x, NaN) = Max(NaN, x) = NaN
Max(+0, ±0) = Max(±0, +0) = +0
Max(-0, -0) = -0

Min(x, -Inf) = Min(-Inf, x) = -Inf
Min(x, NaN) = Min(NaN, x) = NaN
Min(-0, ±0) = Min(±0, -0) = -0
```

## Format byte size as kilobytes, megabytes, gigabytes, ...

These utility functions convert a size in bytes to a human-readable string in either SI \(decimal\) or IEC \(binary\) format.

| Input | ByteCountSI | ByteCountIEC |
| :--- | :--- | :--- |
| 999 | `"999 B"` | `"999 B"` |
| 1000 | `"1.0 kB"` | `"1000 B"` |
| 1023 | `"1.0 kB"` | `"1023 B"` |
| 1024 | `"1.0 kB"` | `"1.0 KiB"` |
| 987,654,321 | `"987.7 MB"` | `"941.9 MiB"` |
| `math.MaxInt64` | `"9.2 EB"` | `"8.0 EiB"` |

```text
func ByteCountSI(b int64) string {
    const unit = 1000
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %cB",
        float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountIEC(b int64) string {
    const unit = 1024
    if b < unit {
        return fmt.Sprintf("%d B", b)
    }
    div, exp := int64(unit), 0
    for n := b / unit; n >= unit; n /= unit {
        div *= unit
        exp++
    }
    return fmt.Sprintf("%.1f %ciB",
        float64(b)/float64(div), "KMGTPE"[exp])
}
```

{% embed url="https://byte-count.go" %}

## 3 simple ways to create an error

### String-based errors <a id="string-based-errors"></a>

The standard library offers two out-of-the-box options.

```text
// simple string-based error
err1 := errors.New("math: square root of negative number")

// with formatting
err2 := fmt.Errorf("math: square root of negative number %g", x)
```

### Custom errors with data <a id="custom-errors-with-data"></a>

To define a custom error type, you must satisfy the predeclared `error` [interface](https://yourbasic.org/golang/interfaces-explained/).

```text
type error interface {
    Error() string
}
```

Here are two examples.

```text
type SyntaxError struct {
    Line int
    Col  int
}

func (e *SyntaxError) Error() string {
    return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}
```

```text
type InternalError struct {
    Path string
}

func (e *InternalError) Error() string {
    return fmt.Sprintf("parse %v: internal error", e.Path)
}
```

If `Foo` is a function that can return a `SyntaxError` or an `InternalError`, you may handle the two cases like this.

```text
if err := Foo(); err != nil {
    switch e := err.(type) {
    case *SyntaxError:
        // Do something interesting with e.Line and e.Col.
    case *InternalError:
        // Abort and file an issue.
    default:
        log.Println(e)
    }
}
```

## Create a new image

Use the [`image`](https://golang.org/pkg/image/), [`image/color`](https://golang.org/pkg/image/color/), and [`image/png`](https://golang.org/pkg/image/png/) packages to create a PNG image.

```text
width := 200
height := 100

upLeft := image.Point{0, 0}
lowRight := image.Point{width, height}

img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// Colors are defined by Red, Green, Blue, Alpha uint8 values.
cyan := color.RGBA{100, 200, 200, 0xff}

// Set color for each pixel.
for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
        switch {
        case x < width/2 && y < height/2: // upper left quadrant
            img.Set(x, y, cyan)
        case x >= width/2 && y >= height/2: // lower right quadrant
            img.Set(x, y, color.White)
        default:
            // Use zero value.
        }
    }
}

// Encode as PNG.
f, _ := os.Create("image.png")
png.Encode(f, img)
```

Output \([image.png](https://yourbasic.org/golang/image.png)\):

```text

```

**Note:** The upper right and lower left quadrants of the image are transparent \(the alpha value is 0\) and will be the same color as the background.

### Go image support <a id="go-image-support"></a>

The [`image`](https://golang.org/pkg/image/) package implements a basic 2-D image library without painting or drawing functionality. The article [The Go image package](https://blog.golang.org/go-image-package) has a nice introduction to images, color models, and image formats in Go.

Additionally, the [`image/draw`](https://golang.org/pkg/image/draw) package provides image composition functions that can be used to perform a number of common image manipulation tasks. The article [The Go image/draw package](https://blog.golang.org/go-imagedraw-package) has plenty of examples.

## Generate all permutations

```text
// Perm calls f with each permutation of a.
func Perm(a []rune, f func([]rune)) {
    perm(a, f, 0)
}

// Permute the values at index i to len(a)-1.
func perm(a []rune, f func([]rune), i int) {
    if i > len(a) {
        f(a)
        return
    }
    perm(a, f, i+1)
    for j := i + 1; j < len(a); j++ {
        a[i], a[j] = a[j], a[i]
        perm(a, f, i+1)
        a[i], a[j] = a[j], a[i]
    }
}
```

Example usage:

```text
Perm([]rune("abc"), func(a []rune) {
	fmt.Println(string(a))
})
```

Output:

```text
abc
acb
bac
bca
cba
cab
```

## Hash checksums: MD5, SHA-1, SHA-256

### String checksum <a id="string-checksum"></a>

To compute the hash value of a **string** or **byte slice**, use the `Sum` function from a crypto package such as [`crypto/md5`](https://golang.org/pkg/crypto/md5/), [`crypto/sha1`](https://golang.org/pkg/crypto/sha1/), or [`crypto/sha256`](https://golang.org/pkg/crypto/sha256/).

```text
s := "Foo"

md5 := md5.Sum([]byte(s))
sha1 := sha1.Sum([]byte(s))
sha256 := sha256.Sum256([]byte(s))

fmt.Printf("%x\n", md5)
fmt.Printf("%x\n", sha1)
fmt.Printf("%x\n", sha256)
```

```text
1356c67d7ad1638d816bfb822dd2c25d
201a6b3053cc1422d2c3670b62616221d2290929
1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa
```

### File checksum <a id="file-checksum"></a>

To compute the hash value of a **file** or other **input stream**:

* create a new [`hash.Hash`](https://golang.org/pkg/hash/#Hash) from a crypto package such as [`crypto/md5`](https://golang.org/pkg/crypto/md5/), [`crypto/sha1`](https://golang.org/pkg/crypto/sha1/), or [`crypto/sha256`](https://golang.org/pkg/crypto/sha256/),
* add data by writing to its `io.Writer` function,
* extract the checksum by calling its `Sum` function.

```text
input := strings.NewReader("Foo")

hash := sha256.New()
if _, err := io.Copy(hash, input); err != nil {
    log.Fatal(err)
}
sum := hash.Sum(nil)

fmt.Printf("%x\n", sum)
```

```text
1cbec737f863e4922cee63cc2ebbfaafcd1cff8b790d8cfd2e6a5d550b648afa
```

## Hello world HTTP server example

###  <a id="a-basic-web-server"></a>

If you access the URL `http://localhost:8080/world` on a machine where the program below is running, you will be greeted by this page.

![Web browser localhost:8080](https://yourbasic.org/golang/hello-server.png)

```text
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
```

* The call to [`http.HandleFunc`](https://golang.org/pkg/net/http/#HandleFunc) tells the [`net.http`](https://golang.org/pkg/net/http/) package to handle all requests to the web root with the `HelloServer` function.
* The call to [`http.ListenAndServe`](https://golang.org/pkg/net/http/#ListenAndServe) tells the server to listen on the TCP network address `:8080`. This function blocks until the program is terminated.
* Writing to an [`http.ResponseWriter`](https://golang.org/pkg/net/http/#ResponseWriter) sends data to the HTTP client.
* An [`http.Request`](https://golang.org/pkg/net/http/#Request) is a data structure that represents a client HTTP request.
* `r.URL.Path` is the path component of the requested URL. In this case, `"/world"` is the path component of `"http://localhost:8080/world"`.

### Further reading: a complete wiki <a id="further-reading-a-complete-wiki"></a>

The [Writing Web Applications](https://golang.org/doc/articles/wiki/) tutorial shows how to extend this small example into a complete wiki.

The tutorial covers how to

* create a data structure with load and save methods,
* use the [`net/http`](https://golang.org/pkg/net/http/) package to build web applications,
* use the [`html/template`](https://golang.org/pkg/html/template/) package to process HTML templates,
* use the [`regexp`](https://yourbasic.org/golang/regexp-cheat-sheet/) package to validate user input.

## How to best implement an iterator

Go has a built-in range loop for iterating over slices, arrays, strings, maps and channels. See [4 basic range loop \(for-each\) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/).

To iterate over other types of data, an iterator function with callbacks is a clean and fairly efficient abstraction.

### Basic iterator pattern <a id="basic-iterator-pattern"></a>

```text
// Iterate calls the f function with n = 1, 2, and 3.
func Iterate(f func(n int)) {
    for i := 1; i <= 3; i++ {
        f(i)
    }
}
```

In use:

```text
Iterate(func(n int) { fmt.Println(n) })
```

```text
1
2
3
```

### Iterator with break <a id="iterator-with-break"></a>

```text
// Iterate calls the f function with n = 1, 2, and 3.
// If f returns true, Iterate returns immediately
// skipping any remaining values.
func Iterate(f func(n int) (skip bool)) {
    for i := 1; i <= 3; i++ {
        if f(i) {
            return
        }
    }
}
```

In use:

```text
Iterate(func(n int) (skip bool) {
	fmt.Println(n)
	return n == 2
})
```

```text
1
2
```

## 4 iota enum examples



### Iota basic example <a id="iota-basic-example"></a>

* The [`iota`](https://yourbasic.org/golang/iota/) keyword represents successive integer constants 0, 1, 2,…
* It resets to 0 whenever the word `const` appears in the source code,
* and increments after each const specification.

```text
const (
    C0 = iota
    C1 = iota
    C2 = iota
)
fmt.Println(C0, C1, C2) // "0 1 2"
```

This can be simplified to

```text
const (
	C0 = iota
	C1
	C2
)
```

Here we rely on the fact that expressions are implicitly repeated in a paren­thesized const declaration – this indicates a repetition of the preceding expression and its type.

#### Start from one <a id="start-from-one"></a>

To start a list of constants at 1 instead of 0, you can use `iota` in an arithmetic expression.

```text
const (
    C1 = iota + 1
    C2
    C3
)
fmt.Println(C1, C2, C3) // "1 2 3"
```

#### Skip value <a id="skip-value"></a>

You can use the blank identifier to skip a value in a list of constants.

```text
const (
    C1 = iota + 1
    _
    C3
    C4
)
fmt.Println(C1, C3, C4) // "1 3 4"
```

### Complete enum type with strings \[best practice\] <a id="complete-enum-type-with-strings-best-practice"></a>

Here’s an idiomatic way to implement an enumerated type:

* create a new integer type,
* list its values using `iota`,
* give the type a `String` function.

```text
type Direction int

const (
    North Direction = iota
    East
    South
    West
)

func (d Direction) String() string {
    return [...]string{"North", "East", "South", "West"}[d]
}
```

In use:

```text
var d Direction = North
fmt.Print(d)
switch d {
case North:
    fmt.Println(" goes up.")
case South:
    fmt.Println(" goes down.")
default:
    fmt.Println(" stays put.")
}
// Output: North goes up.
```

#### Naming convention <a id="naming-convention"></a>

> The standard naming convention is to use mixed caps also for constants. For example, an exported constant is `NorthWest`, not `NORTH_WEST`.

## Maximum value of an int

Go has two predeclared integer types with implementation-specific sizes:

* a `uint` \(unsigned integer\) has either 32 or 64 bits,
* an `int` \(signed integer\) has the same size as a `uint`.

This code computes the limit values as **untyped constants**.

```text
const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

const (
    MaxInt  = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
    MinInt  = -MaxInt - 1         // -1 << 31 or -1 << 63
    MaxUint = 1<<UintSize - 1     // 1<<32 - 1 or 1<<64 - 1
)
```

> The [`UintSize`](https://golang.org/pkg/math/bits/#pkg-constants) constant is also available in package [`math/bits`](https://golang.org/pkg/math/bits/).

## Round float to 2 decimal places

### Float to string <a id="float-to-string"></a>

To display the value as a string, use the [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) method.

```text
s := fmt.Sprintf("%.2f", 12.3456) // s == "12.35"
```

The [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) lists the most common formatting verbs and flags.

### Float to float <a id="float-to-float"></a>

To round to a floating-point value, use one of these techniques.

```text
x := 12.3456
fmt.Println(math.Floor(x*100)/100) // 12.34 (round down)
fmt.Println(math.Round(x*100)/100) // 12.35 (round to nearest)
fmt.Println(math.Ceil(x*100)/100)  // 12.35 (round up)
```

Due to the quirks of floating point representation, these rounded values may be slightly off.

### Float to integer value <a id="float-to-integer-value"></a>

[Round float to integer value](https://yourbasic.org/golang/round-float-to-int/) has further details on how to round a float64 to an integer \(away from zero, to even number, converted to an int type\).

### Before Go 1.10 <a id="before-go-1-10"></a>

The [`math.Round`](https://golang.org/pkg/math/#Round) function was introduced in Go 1.10. See [Round float to integer value](https://yourbasic.org/golang/round-float-to-int/) for equivalent code.

## Round float to integer value

### Round away from zero[Go 1.10](https://golang.org/doc/go1.10) <a id="round-away-from-zero"></a>

Use [`math.Round`](https://golang.org/pkg/math/#Round) to return the nearest integer, as a `float64`, rounding ties away from zero.

```text
fmt.Println(math.Round(-0.6)) // -1
fmt.Println(math.Round(-0.4)) // -0
fmt.Println(math.Round(0.4))  // 0
fmt.Println(math.Round(0.6))  // 1
```

Note the special cases.

```text
Round(±0) = ±0
Round(±Inf) = ±Inf
Round(NaN) = NaN
```

### Round to even number[Go 1.10](https://golang.org/doc/go1.10) <a id="round-to-even-number"></a>

Use [`math.RoundToEven`](https://golang.org/pkg/math/#RoundToEven) to return the nearest integer, as a `float64`, rounding ties to an even number.

```text
fmt.Println(math.RoundToEven(0.5)) // 0
fmt.Println(math.RoundToEven(1.5)) // 2
```

### Convert to an int type <a id="convert-to-an-int-type"></a>

Note that when converting a floating-point number to an `int` type, the fraction is discarded \(truncation towards zero\).

```text
fmt.Println(int64(1.9))  //  1
fmt.Println(int64(-1.9)) // -1
```

> **Warning:** If the result type cannot represent the value the conversion succeeds but the result is implementation-dependent.

### Before Go 1.10 <a id="before-go-1-10"></a>

The following implementations are equivalent to [`math.Round`](https://golang.org/pkg/math/#Round) and [`math.RoundToEven`](https://golang.org/pkg/math/#RoundToEven), but less efficient.

```text
// Round returns the nearest integer, rounding ties away from zero.
func Round(x float64) float64 {
    t := math.Trunc(x)
    if math.Abs(x-t) >= 0.5 {
        return t + math.Copysign(1, x)
    }
    return t
}
```

```text
// RoundToEven returns the nearest integer, rounding ties to an even number.
func RoundToEven(x float64) float64 {
    t := math.Trunc(x)
    odd := math.Remainder(t, 2) != 0
    if d := math.Abs(x - t); d > 0.5 || (d == 0.5 && odd) {
        return t + math.Copysign(1, x)
    }
    return t
}
```

## Table-driven unit tests

Here is the code we want to test.

```text
package search

// Find returns the smallest index i at which x <= a[i].
// If there is no such index, it returns len(a).
// The slice must be sorted in ascending order.
func Find(a []int, x int) int {
    switch len(a) {
    case 0:
        return 0
    case 1:
        if x <= a[0] {
            return 0
        }
        return 1
    }
    mid := len(a) / 2
    if x <= a[mid-1] {
        return Find(a[:mid], x)
    }
    return mid + Find(a[mid:], x)
}
```

* Put the test code in a file whose name ends with **\_test.go**.
* Write a function `TestXXX` with a single argument of type [`*testing.T`](https://golang.org/pkg/testing/#T). The test framework runs each such function.
* To indicate a failed test, call a failure function such as [`t.Errorf`](https://golang.org/pkg/testing/#T.Errorf).

```text
package search

import "testing"

var tests = []struct {
    a   []int
    x   int
    exp int
}{
    {[]int{}, 1, 0},
    {[]int{1, 2, 3, 3}, 0, 0},
    {[]int{1, 2, 3, 3}, 1, 0},
    {[]int{1, 2, 3, 3}, 2, 1},
    {[]int{1, 2, 3, 3}, 3, 3}, // incorrect test case
    {[]int{1, 2, 3, 3}, 4, 4},
}

func TestFind(t *testing.T) {
    for _, e := range tests {
        res := Find(e.a, e.x)
        if res != e.exp {
            t.Errorf("Find(%v, %d) = %d, expected %d",
                e.a, e.x, res, e.exp)
        }
    }
}
```

Run the tests with [`go test`](https://golang.org/cmd/go/#hdr-Test_packages).

```text
$ go test
--- FAIL: TestFind (0.00s)
    search_test.go:22: Find([1 2 3 3], 3) = 2, expected 3
FAIL
exit status 1
FAIL    .../search  0.001s
```

#### Further reading <a id="further-reading"></a>

The [Induction and recursive functions](https://yourbasic.org/algorithms/induction-recursive-functions/) article has a correctness proof for the `Find` function.

## The 3 ways to sort in Go

### Sort a slice of ints, float64s or strings <a id="sort-a-slice-of-ints-float64s-or-strings"></a>

Use one of the functions

* [`sort.Ints`](https://golang.org/pkg/sort/#Ints)
* [`sort.Float64s`](https://golang.org/pkg/sort/#Float64s)
* [`sort.Strings`](https://golang.org/pkg/sort/#Strings)

```text
s := []int{4, 2, 3, 1}
sort.Ints(s)
fmt.Println(s) // [1 2 3 4]
```

> Package [radix](https://godoc.org/github.com/yourbasic/radix) contains a drop-in replacement for sort.Strings, which can be more than twice as fast in some settings.

### Sort with custom comparator <a id="sort-with-custom-comparator"></a>

* Use the function [`sort.Slice`](https://golang.org/pkg/sort/#Slice). It sorts a slice using a provided function `less(i, j int) bool`.
* To sort the slice while keeping the original order of equal elements, use [`sort.SliceStable`](https://golang.org/pkg/sort/#SliceStable) instead.

```text
family := []struct {
    Name string
    Age  int
}{
    {"Alice", 23},
    {"David", 2},
    {"Eve", 2},
    {"Bob", 25},
}

// Sort by age, keeping original order or equal elements.
sort.SliceStable(family, func(i, j int) bool {
    return family[i].Age < family[j].Age
})
fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
```

### Sort custom data structures <a id="sort-custom-data-structures"></a>

* Use the generic [`sort.Sort`](https://golang.org/pkg/sort/#Sort) and [`sort.Stable`](https://golang.org/pkg/sort/#Stable) functions.
* They sort any collection that implements the [`sort.Interface`](https://golang.org/pkg/sort/#Interface) [interface](https://yourbasic.org/golang/interfaces-explained/).

```text
type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
```

Here’s an example.

```text
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    family := []Person{
        {"Alice", 23},
        {"Eve", 2},
        {"Bob", 25},
    }
    sort.Sort(ByAge(family))
    fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}
```

### Bonus: Sort a map by key or value <a id="bonus-sort-a-map-by-key-or-value"></a>

A [map](https://yourbasic.org/golang/maps-explained/) is an **unordered** collection of key-value pairs. If you need a stable iteration order, you must maintain a separate data structure.

This code example uses a slice of keys to sort a map in key order.

```text
m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
// Output:
// Alice 2
// Bob 3
// Cecil 1
```

> Also, starting with [Go 1.12](https://tip.golang.org/doc/go1.12), the [`fmt`](https://golang.org/pkg/fmt/) package prints maps in key-sorted order to ease testing.

### Performance and implementation <a id="performance-and-implementation"></a>

All algorithms in the Go sort package make _O_\(_n_ log _n_\) comparisons in the worst case, where _n_ is the number of elements to be sorted.

Most of the functions are implemented using an [optimized version of quicksort](https://yourbasic.org/golang/quicksort-optimizations/).

## Object-oriented programming without inheritance



Go doesn’t have inheritance – instead composition, embed­ding and inter­faces support code reuse and poly­morphism.

### Object-oriented programming with inheritance <a id="object-oriented-programming-with-inheritance"></a>

### Object-oriented programming with inheritance <a id="object-oriented-programming-with-inheritance"></a>

Inheritance in traditional object-oriented languages offers three features in one. When a `Dog` inherits from an `Animal`

1. the `Dog` class reuses code from the `Animal` class,
2. a variable `x` of type `Animal` can refer to either a `Dog` or an `Animal`,
3. `x.Eat()` will choose an `Eat` method based on what type of object `x` refers to.

In object-oriented lingo, these features are known as **code reuse**, **poly­mor­phism** and **dynamic dispatch**.

All of these are available in Go, using separate constructs:

* **composition** and **embedding** provide code reuse,
* [**interfaces**](https://yourbasic.org/golang/interfaces-explained/) take care of polymorphism and dynamic dispatch.

### Code reuse by composition <a id="code-reuse-by-composition"></a>

> Don't worry about type hierarchies when starting a new Go project –  
> it's easy to introduce polymorphism and dynamic dispatch later on.

If a `Dog` needs some or all of the functionality of an `Animal`, simply use **composition**.

```text
type Animal struct {
	// …
}

type Dog struct {
	beast Animal
	// …
}
```

This gives you full freedom to use the `Animal` part of your `Dog` as needed. Yes, it’s that simple.

### Code reuse by embedding <a id="code-reuse-by-embedding"></a>

If the `Dog` class inherits **the exact behavior** of an `Animal`, this approach can result in some tedious coding.

```text
type Animal struct {
	// …
}

func (a *Animal) Eat()   { … }
func (a *Animal) Sleep() { … }
func (a *Animal) Breed() { … }

type Dog struct {
	beast Animal
	// …
}

func (a *Dog) Eat()   { a.beast.Eat() }
func (a *Dog) Sleep() { a.beast.Sleep() }
func (a *Dog) Breed() { a.beast.Breed() }
```

This code pattern is known as **delegation**.

Go uses **embedding** for situations like this. The declaration of the `Dog` struct and it’s three methods can be reduced to:

```text
type Dog struct {
	Animal
	// …
}
```

### Polymorphism and dynamic dispatch with interfaces <a id="polymorphism-and-dynamic-dispatch-with-interfaces"></a>

> Keep your interfaces short, and introduce them only when needed.

Further down the road your project might have grown to include more animals. At this point you can introduce polymorphism and dynamic dispatch using [**interfaces**](https://yourbasic.org/golang/interfaces-explained/).

If you need to put all your pets to sleep, you can define a `Sleeper` interface.

```text
type Sleeper interface {
	Sleep()
}

func main() {
	pets := []Sleeper{new(Cat), new(Dog)}
	for _, x := range pets {
		x.Sleep()
	}
}
```

No explicit declaration is required by the `Cat` and `Dog` types. Any type that provides the methods named in an inter­face may be treated as an imple­mentation.

_When I see a bird that walks like a duck and swims like a duck and quacks like a duck, I call that bird a duck._  
–James Whitcomb Riley

### What about constructors? <a id="what-about-constructors"></a>

See [Constructors deconstructed](https://yourbasic.org/golang/constructor-best-practice/) for best practices on how to set up new data structures in Go.

## Constructors deconstructed \[best practice\]

Go doesn't have explicit constructors. The idiomatic way to set up new data structures is to use proper **zero values** coupled with **factory** functions.

### Zero value <a id="zero-value"></a>

Try to make the default [zero value](https://yourbasic.org/golang/default-zero-value/) useful and document its behavior. Sometimes this is all that’s needed.

```text
// A StopWatch is a simple clock utility.
// Its zero value is an idle clock with 0 total time.
type StopWatch struct {
    start   time.Time
    total   time.Duration
    running bool
}

var clock StopWatch // Ready to use, no initialization needed.
```

* `StopWatch` takes advantage of the useful zero values of `time.Time`, `time.Duration` and `bool`.
* In turn, users of `StopWatch` can benefit from _its_ useful zero value.

### Factory <a id="factory"></a>

If the zero value doesn’t suffice, use factory functions named `NewFoo` or just `New`.

```text
scanner := bufio.NewScanner(os.Stdin)
err := errors.New("Houston, we have a problem")
```

## Methods explained

Go doesn't have classes, but you can define methods on types.



* A method is a function with an extra **receiver** argument.
* The receiver sits between the `func` keyword and the method name.

In this example, the `HasGarage` method is associated with the `House` type. The method receiver is called `p`.

```text
type House struct {
    garage bool
}

func (p *House) HasGarage() bool { return p.garage }

func main() {
    house := new(House)
    fmt.Println(house.HasGarage()) // Prints "false" (zero value)
}
```

#### Conversions and methods <a id="conversions-and-methods"></a>

If you [convert](https://yourbasic.org/golang/conversions/) a value to a different type, the new value will have the methods of the new type, but not the old.

```text
type MyInt int

func (m MyInt) Positive() bool { return m > 0 }

func main() {
    var m MyInt = 2
    m = m * m // The operators of the underlying type still apply.

    fmt.Println(m.Positive())        // Prints "true"
    fmt.Println(MyInt(3).Positive()) // Prints "true"

    var n int
    n = int(m) // The conversion is required.
    n = m      // ILLEGAL
}
```

```text
../main.go:14:4: cannot use m (type MyInt) as type int in assignment
```

It’s idiomatic in Go to convert the type of an expression to access a specific method.

```text
var n int64 = 12345
fmt.Println(n)                // 12345
fmt.Println(time.Duration(n)) // 12.345µs
```

\(The underlying type of [`time.Duration`](https://golang.org/pkg/time/#Duration) is `int64`, and the `time.Duration` type has a [`String`](https://golang.org/pkg/time/#Duration.String) method that returns the duration formatted as a time.\)

#### Further reading <a id="further-reading"></a>

[Object-oriented programming without inheritance](https://yourbasic.org/golang/inheritance-object-oriented/) explains how composition, embedding and interfaces support code reuse and polymorphism in Go.

## Public vs. private

A package is the smallest unit of private encap­sulation in Go.

* All identifiers defined within a [package](https://yourbasic.org/golang/packages-explained/) are visible throughout that package.
* When importing a package you can access only its **exported** identifiers.
* An identifier is exported if it begins with a **capital letter**.

Exported and unexported identifiers are used to describe the public interface of a package and to guard against certain programming errors.

> **Warning:** Unexported identifiers is not a security measure and it does not hide or protect any information.

### Example <a id="example"></a>

In this package, the only exported identifiers are `StopWatch` and `Start`.

```text
package timer

import "time"

// A StopWatch is a simple clock utility.
// Its zero value is an idle clock with 0 total time.
type StopWatch struct {
    start   time.Time
    total   time.Duration
    running bool
}

// Start turns the clock on.
func (s *StopWatch) Start() {
    if !s.running {
        s.start = time.Now()
        s.running = true
    }
}
```

The `StopWatch` and its exported methods can be imported and used in a different package.

```text
package main

import "timer"

func main() {
    clock := new(timer.StopWatch)
    clock.Start()
    if clock.running { // ILLEGAL
        // …
    }
}
```

```text
../main.go:8:15: clock.running undefined (cannot refer to unexported field or method clock.running)
```

## Pointer vs. value receiver

### Basic guidelines <a id="basic-guidelines"></a>

* For a given type, _don’t mix_ value and pointer receivers.
* If in doubt, _use pointer receivers_ \(they are safe and extendable\).

### Pointer receivers <a id="pointer-receivers"></a>

You _must_ use pointer receivers

* if any method needs to mutate the receiver,
* for structs that contain a `sync.Mutex` or similar synchronizing field \(they musn’t be copied\).

You _probably want_ to use pointer receivers

* for large structs or arrays \(it can be more efficient\),
* in all other cases.

### Value receivers <a id="value-receivers"></a>

You _probably want_ to use value receivers

* for `map`, `func` and `chan` types,
* for simple basic types such as `int` or `string`,
* for small arrays or structs that are value types, with no mutable fields and no pointers.

You _may want_ to use value receivers

* for slices with methods that do not reslice or reallocate the slice.

## Functional Programming

## Functional programming in Go \[case study\]

A graph implementation based entirely on functions

### Introduction <a id="introduction"></a>

This text is about the implementation of a Go tool based entirely on functions – the API contains only immutable data types, and the code is built on top of a `struct` with five `func` fields.

It’s a tool for building **virtual graphs**. In a virtual graph no vertices or edges are stored in memory, they are instead computed as needed. The tool is part of a larger library of generic graph algorithms:

* the package [graph](https://github.com/yourbasic/graph) contains the basic graph library and
* the subpackage [graph/build](https://github.com/yourbasic/graph/tree/master/build) is the tool for building virtual graphs.

There is an online reference for the build tool at [godoc.org](https://godoc.org/github.com/yourbasic/graph/build).

## Function types and values

Function types and function values can be used and passed around just like other values:

```text
type Operator func(x float64) float64

// Map applies op to each element of a.
func Map(op Operator, a []float64) []float64 {
    res := make([]float64, len(a))
    for i, x := range a {
        res[i] = op(x)
    }
    return res
}

func main() {
    op := math.Abs
    a := []float64{1, -2}
    b := Map(op, a)
    fmt.Println(b) // [1 2]

    c := Map(func(x float64) float64 { return 10 * x }, b)
    fmt.Println(c) // [10, 20]
}
```

The second call to `Map` uses a **function literal** \(or **lambda**\). See [Anonymous functions and closures](https://yourbasic.org/golang/anonymous-function-literal-lambda-closure/) for more about lambdas in Go.

### Details <a id="details"></a>

A function type describes the set of all functions with the same parameter and result types.

* The value of an uninitialized variable of function type is `nil`.
* The parameter names are optional.

The following two function types are identical.

```text
func(x, y int) int
func(int, int) int
```

#### Further reading <a id="further-reading"></a>

[Anonymous functions and closures](https://yourbasic.org/golang/anonymous-function-literal-lambda-closure/)

## Anonymous functions and closures

A function literal \(or lambda\) is a function without a name.

In this example a **function literal** is passed as the `less` argument to the [`sort.Slice`](https://golang.org/pkg/sort/#Slice) function.

```text
func Slice(slice interface{}, less func(i, j int) bool)
```

```text
people := []string{"Alice", "Bob", "Dave"}
sort.Slice(people, func(i, j int) bool {
    return len(people[i]) < len(people[j])
})
fmt.Println(people)
// Output: [Bob Dave Alice]
```

You can also use an intermediate variable.

```text
people := []string{"Alice", "Bob", "Dave"}
less := func(i, j int) bool {
    return len(people[i]) < len(people[j])
}
sort.Slice(people, less)
```

Note that the `less` function is a **closure**: it references the `people` variable, which is declared outside the function.

### Closures <a id="closures"></a>

Function literals in Go are **closures**: they may refer to variables defined in an enclosing function. Such variables

* are shared between the surrounding function and the function literal,
* survive as long as they are accessible.

In this example, the function literal uses the local variable `n` from the enclosing scope to count the number of times it has been invoked.

```text
// New returns a function Count.
// Count prints the number of times it has been invoked.
func New() (Count func()) {
    n := 0
    return func() {
        n++
        fmt.Println(n)
    }
}

func main() {
    f1, f2 := New(), New()
    f1() // 1
    f2() // 1 (different n)
    f1() // 2
    f2() // 2
}
```

## How to best implement an iterator

Go has a built-in range loop for iterating over slices, arrays, strings, maps and channels. See [4 basic range loop \(for-each\) patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/).

To iterate over other types of data, an iterator function with callbacks is a clean and fairly efficient abstraction.

### Basic iterator pattern <a id="basic-iterator-pattern"></a>

```text
// Iterate calls the f function with n = 1, 2, and 3.
func Iterate(f func(n int)) {
    for i := 1; i <= 3; i++ {
        f(i)
    }
}
```

In use:

```text
Iterate(func(n int) { fmt.Println(n) })
```

```text
1
2
3
```

### Iterator with break <a id="iterator-with-break"></a>

```text
// Iterate calls the f function with n = 1, 2, and 3.
// If f returns true, Iterate returns immediately
// skipping any remaining values.
func Iterate(f func(n int) (skip bool)) {
    for i := 1; i <= 3; i++ {
        if f(i) {
            return
        }
    }
}
```

In use:

```text
Iterate(func(n int) (skip bool) {
	fmt.Println(n)
	return n == 2
})
```

```text
1
2
```

## Scripting

## Go as a scripting language: lightweight, safe and fast

This example is a simplified version of the Unix `grep` command. The program searches the input file for lines containing the given pattern and prints these lines.

```text
func main() {
    log.SetPrefix("grep: ")
    log.SetFlags(0) // no extra info in log messages

    if len(os.Args) != 3 {
        fmt.Printf("Usage: %v PATTERN FILE\n", os.Args[0])
        return
    }

    pattern, err := regexp.Compile(os.Args[1])
    if err != nil {
        log.Fatalln(err)
    }

    file, err := os.Open(os.Args[2])
    if err != nil {
        log.Fatalln(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if pattern.MatchString(line) {
            fmt.Println(line)
        }
    }
    if err := scanner.Err(); err != nil {
        log.Println(err)
    }
}
```

## Read a file \(stdin\) line by line

Use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to read a file line by line.

```text
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

### Read from stdin <a id="read-from-stdin"></a>

Use [`os.Stdin`](https://golang.org/pkg/os/#pkg-variables) to read from the standard input stream.

```text
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	log.Println(err)
}
```

### Read from any stream <a id="read-from-any-stream"></a>

A bufio.Scanner can read from any stream of bytes, as long as it implements the [`io.Reader`](https://golang.org/pkg/io/#Reader) interface. See [How to use the io.Reader interface](https://yourbasic.org/golang/io-reader-interface-explained/).

#### Further reading <a id="further-reading"></a>

For more advanced scanning, see the examples in the [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) documentation.  


## How to use the io.Reader interface

### Basics <a id="basics"></a>

The [`io.Reader`](https://golang.org/pkg/io/#Reader) [interface](https://yourbasic.org/golang/interfaces-explained/) represents an entity from which you can read a stream of bytes.

```text
type Reader interface {
        Read(buf []byte) (n int, err error)
}
```

`Read` reads up to `len(buf)` bytes into `buf` and returns the number of bytes read – it returns an [`io.EOF`](https://golang.org/pkg/io/#pkg-variables) error when the stream ends.

The standard library provides numerous Reader [implementations](https://golang.org/search?q=Read#Global) \(including in-memory byte buffers, files and network connections\), and Readers are accepted as input by many utilities \(including the HTTP client and server implementations\).

### Use a built-in reader <a id="use-a-built-in-reader"></a>

As an example, you can create a Reader from a string using the [`strings.Reader`](https://golang.org/pkg/strings/#Reader) function and then pass the Reader directly to the [`http.Post`](https://golang.org/pkg/net/http/#Post) function in package [`net/http`](https://golang.org/pkg/net/http/). The Reader is then used as the source for the data to be posted.

```text
r := strings.NewReader("my request")
resp, err := http.Post("http://foo.bar",
	"application/x-www-form-urlencoded", r)
```

Since `http.Post` uses a Reader instead of a `[]byte` it’s trivial to, for instance, use the contents of a file instead.

### Read directly from a byte stream <a id="read-directly-from-a-byte-stream"></a>

You can use the `Read` function directly \(this is the least common use case\).

```text
r := strings.NewReader("abcde")

buf := make([]byte, 4)
for {
	n, err := r.Read(buf)
	fmt.Println(n, err, buf[:n])
	if err == io.EOF {
		break
	}
}
```

```text
4 <nil> [97 98 99 100]
1 <nil> [101]
0 EOF []

```

Use [`io.ReadFull`](https://golang.org/pkg/io/#ReadFull) to read exactly `len(buf)` bytes into `buf`:

```text
r := strings.NewReader("abcde")

buf := make([]byte, 4)
if _, err := io.ReadFull(r, buf); err != nil {
	log.Fatal(err)
}
fmt.Println(buf)

if _, err := io.ReadFull(r, buf); err != nil {
	fmt.Println(err)
}
```

```text
[97 98 99 100]
unexpected EOF
```

Use [`ioutil.ReadAll`](https://golang.org/pkg/io/ioutil/#ReadAll) to read everything:

```text
r := strings.NewReader("abcde")

buf, err := ioutil.ReadAll(r)
if err != nil {
	log.Fatal(err)
}
fmt.Println(buf)
```

```text
[97 98 99 100 101]
```

### Buffered reading and scanning <a id="buffered-reading-and-scanning"></a>

The [`bufio.Reader`](https://golang.org/pkg/bufio/#Reader) and [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) types wrap a Reader creating another Reader that also implements the interface but provides buffering and some help for textual input.

In this example we use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to count the number of words in a text.

```text
const input = `Beware of bugs in the above code;
I have only proved it correct, not tried it.`

scanner := bufio.NewScanner(strings.NewReader(input))
scanner.Split(bufio.ScanWords) // Set up the split function.

count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Println(err)
}
fmt.Println(count)
```

```text
16
```

## Command-line arguments and flags

## The [`os.Args`](https://golang.org/pkg/os/#pkg-variables) variable holds the command-line arguments – starting with the program name – which are passed to a Go program.

```text
func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage:", os.Args[0], "PATTERN", "FILE")
        return
    }
    pattern := os.Args[1]
    file := os.Args[2]
    // ...
}
```

```text
$ go build grep.go
$ ./grep
Usage: ./grep PATTERN FILE
```

#### Flag parsing <a id="flag-parsing"></a>

The [flag](https://golang.org/pkg/flag/) package implements basic command-line flag parsing.

## Access environment variables

```text
key:"SHELL" value:"/bin/bash"
key:"SESSION" value:"ubuntu"
key:"TERM" value:"xterm-256color"
key:"LANG" value:"en_US.UTF-8"
key:"XMODIFIERS" value:"@im=ibus"
…
```

```text
for _, s := range os.Environ() {
    kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
    fmt.Printf("key:%q value:%q\n", kv[0], kv[1])
}
```

The [`os.Environ`](https://golang.org/pkg/os/#Environ) function returns a slice of `"key=value"` strings listing all environment variables.

```text
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/bash"

os.Unsetenv("SHELL")
fmt.Printf("%q\n", os.Getenv("SHELL")) // ""

os.Setenv("SHELL", "/bin/dash")
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/dash"
```

Use the [`os.Setenv`](https://golang.org/pkg/os/#Setenv), [`os.Getenv`](https://golang.org/pkg/os/#Getenv) and [`os.Unsetenv`](https://golang.org/pkg/os/#Unsetenv) functions to access environment variables.

## Strings

### fmt.Printf formatting tutorial and cheat sheet

### Basics

With the Go [`fmt`](https://golang.org/pkg/fmt) package you can format numbers and strings padded with spaces or zeroes, in different bases, and with optional quotes.

You submit a **template string** that contains the text you want to format plus some **annotation verbs** that tell the `fmt` functions how to format the trailing arguments.

#### Printf <a id="printf"></a>

In this example, [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf) formats and writes to standard output:

```text
fmt.Printf("Binary: %b\\%b", 4, 5) // Prints `Binary: 100\101`
```

* the **template string** is `"Binary: %b\\%b"`,
* the **annotation verb** `%b` formats a number in binary, and
* the **special value** `\\` is a backslash.

As a special case, the verb `%%`, which consumes no argument, produces a percent sign:

```text
fmt.Printf("%d %%", 50) // Prints `50 %`
```

#### Sprintf \(format without printing\) <a id="sprintf-format-without-printing"></a>

Use [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) to format a string without printing it:

```text
s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
```

#### Find fmt errors with vet <a id="find-fmt-errors-with-vet"></a>

If you try to compile and run this incorrect line of code

```text
fmt.Printf("Binary: %b\\%b", 4) // An argument to Printf is missing.
```

you’ll find that the program will compile, and then print

```text
Binary: 100\%!b(MISSING)
```

To catch this type of errors early, you can use the [vet command](https://golang.org/cmd/vet/) – it can find calls whose arguments do not align with the format string.

```text
$ go vet example.go
example.go:8: missing argument for Printf("%b"): format reads arg 2, have only 1 args
```

### Cheat sheet <a id="cheat-sheet"></a>

#### Default formats and type <a id="default"></a>

* **Value:** `[]int64{0, 1}`

| Format | Verb | Description |
| :--- | :--- | :--- |
| \[0 1\] | `%v` | Default format |
| \[\]int64{0, 1} | `%#v` | Go-syntax format |
| \[\]int64 | `%T` | The type of the value |

#### Integer \(indent, base, sign\) <a id="integer"></a>

* **Value:** `15`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 15 | `%d` | Base 10 |
| +15 | `%+d` | Always show sign |
| ␣␣15 | `%4d` | Pad with spaces \(width 4, right justified\) |
| 15␣␣ | `%-4d` | Pad with spaces \(width 4, left justified\) |
| 0015 | `%04d` | Pad with zeroes \(width 4\) |
| 1111 | `%b` | Base 2 |
| 17 | `%o` | Base 8 |
| f | `%x` | Base 16, lowercase |
| F | `%X` | Base 16, uppercase |
| 0xf | `%#x` | Base 16, with leading 0x |

#### Character \(quoted, Unicode\) <a id="character"></a>

* **Value:** `65`   \(Unicode letter A\)

| Format | Verb | Description |
| :--- | :--- | :--- |
| A | `%c` | Character |
| 'A' | `%q` | Quoted character |
| U+0041 | `%U` | Unicode |
| U+0041 'A' | `%#U` | Unicode with character |

#### Boolean \(true/false\) <a id="boolean"></a>

Use `%t` to format a boolean as `true` or `false`.

#### Pointer \(hex\) <a id="pointer"></a>

Use `%p` to format a pointer in base 16 notation with leading `0x`.

#### Float \(indent, precision, scientific notation\) <a id="float"></a>

* **Value:** `123.456`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 1.234560e+02 | `%e` | Scientific notation |
| 123.456000 | `%f` | Decimal point, no exponent |
| 123.46 | `%.2f` | Default width, precision 2 |
| ␣␣123.46 | `%8.2f` | Width 8, precision 2 |
| 123.456 | `%g` | Exponent as needed, necessary digits only |

#### String or byte slice \(quote, indent, hex\) <a id="string-or-byte-slice"></a>

* **Value:** `"café"`

| Format | Verb | Description |
| :--- | :--- | :--- |
| café | `%s` | Plain string |
| ␣␣café | `%6s` | Width 6, right justify |
| café␣␣ | `%-6s` | Width 6, left justify |
| "café" | `%q` | Quoted string |
| 636166c3a9 | `%x` | Hex dump of byte values |
| 63 61 66 c3 a9 | `% x` | Hex dump with spaces |

#### Special values <a id="special-values"></a>

| Value | Description |
| :--- | :--- |
| `\a` | U+0007 alert or bell |
| `\b` | U+0008 backspace |
| `\\` | U+005c backslash |
| `\t` | U+0009 horizontal tab |
| `\n` | U+000A line feed or newline |
| `\f` | U+000C form feed |
| `\r` | U+000D carriage return |
| `\v` | U+000b vertical tab |

Arbitrary values can be encoded with backslash escapes and can be used in any `""` string literal.

There are four different formats:

* `\x` followed by exactly two hexadecimal digits,
* `\` followed by exactly three octal digits,
* `\u` followed by exactly four hexadecimal digits,
* `\U` followed by exactly eight hexadecimal digits.

The escapes `\u` and `\U` represent Unicode code points.

```text
fmt.Println("\\caf\u00e9") // Prints \café
```

### Further reading <a id="further-reading"></a>

[40+ practical string tips \[cheat sheet\]](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)

## Runes and character encoding

### Characters, ASCII and Unicode <a id="characters-ascii-and-unicode"></a>

> The `rune` type is an alias for `int32`, and is used to emphasize than an integer represents a code point.

**ASCII** defines 128 characters, identified by the **code points** 0–127. It covers English letters, Latin numbers, and a few other characters.

**Unicode**, which is a superset of ASCII, defines a codespace of 1,114,112 code points. Unicode version 10.0 covers 139 modern and historic scripts \(including the runic alphabet, but not Klingon\) as well as multiple symbol sets.

### Strings and UTF-8 encoding <a id="strings-and-utf-8-encoding"></a>

> A `string` is a sequence of bytes, not runes.

However, strings often contain Unicode text encoded in [UTF-8](https://research.swtch.com/utf8), which encodes all Unicode code points using one to four bytes. \(ASCII characters are encoded with one byte, while other code points use more.\)

Since Go source code itself is encoded as UTF-8, string literals will automatically get this encoding.

For example, in the string `"café"` the character `é` \(code point 233\) is encoded using two bytes, while the ASCII characters `c`, `a` and `f` \(code points 99, 97 and 102\) only use one:

```text
fmt.Println([]byte("café")) // [99 97 102 195 169]
fmt.Println([]rune("café")) // [99 97 102 233]
```

See [Convert between byte array/slice and string](https://yourbasic.org/golang/convert-string-to-byte-slice/) and [Convert between rune array/slice and string](https://yourbasic.org/golang/convert-string-to-rune-slice/).

#### Further reading <a id="further-reading"></a>

[Escapes and multiline strings](https://yourbasic.org/golang/multiline-string/)

### Clean and simple string building <a id="clean-and-simple-string-building"></a>

For simple cases where performance is a non-issue, [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend. It’s clean, simple and fairly efficient.

```text
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

The [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) lists the most common formatting verbs and flags.

## Efficient string concatenation \[full guide\]



### Clean and simple string building <a id="clean-and-simple-string-building"></a>

For simple cases where performance is a non-issue, [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend. It’s clean, simple and fairly efficient.

```text
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

The [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) lists the most common formatting verbs and flags.

### High-performance string concatenation[Go 1.10](https://golang.org/doc/go1.10) <a id="high-performance-string-concatenation"></a>

A [`strings.Builder`](https://golang.org/pkg/strings/#Builder) is used to efficiently append strings using write methods.

* It offers a subset of the [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) methods that allows it to safely avoid extra copying when converting a builder to a string.
* You can use the [`fmt`](https://golang.org/pkg/fmt/) package for formatting since the builder implements the [`io.Writer`](https://yourbasic.org/golang/io-writer-interface-explained/) interface.
* The [`Grow`](https://golang.org/pkg/strings/#Builder.Grow) method can be used to preallocate memory when the maximum size of the string is known.

```text
var b strings.Builder
b.Grow(32)
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&b, "%d:%d, ", i+1, p)
}
s := b.String()   // no copying
s = s[:b.Len()-2] // no copying (removes trailing ", ")
fmt.Println(s)
```

```text
1:2, 2:3, 3:5, 4:7, 5:11, 6:13
```

### Before Go 1.10 <a id="before-go-1-10"></a>

Use [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) to print into a [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer).

```text
var buf bytes.Buffer
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&buf, "%d:%d, ", i+1, p)
}
buf.Truncate(buf.Len() - 2) // Remove trailing ", "
s := buf.String()           // Copy into a new string
fmt.Println(s)
```

```text
1:2, 2:3, 3:5, 4:7, 5:11, 6:13
```

This solution is pretty efficient but may generate some excess garbage. For higher performance, you can try to use the append functions in package [`strconv`](https://golang.org/pkg/strconv/).

```text
buf := []byte("Size: ")
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```

If the expected maximum length of the string is known, you may want to preallocate the slice.

```text
buf := make([]byte, 0, 16)
buf = append(buf, "Size: "...)
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```

## Escapes and multiline strings

### Raw string literals <a id="raw-string-literals"></a>

Raw [string literals](https://golang.org/ref/spec#String_literals), delimited by **backticks** \(back quotes\), are interpreted literally. They can contain line breaks, and backslashes have no special meaning.

```text
const s = `First line
Second line`
fmt.Println(s)
```

```text
First line
Second line
```

#### Backtick escape <a id="backtick-escape"></a>

It’s [not possible](https://github.com/golang/go/issues/24475) to include a backtick in a raw string literal, but you can do

```text
fmt.Println("`" + "foo" + "`") // Prints: `foo`
```

### Interpreted string literals <a id="interpreted-string-literals"></a>

To insert escape characters, use interpreted string literals delimited by **double quotes**.

```text
const s = "\tFirst line\n" +
"Second line"
fmt.Println(s)
```

```text
   First line
Second line
```

The escape character `\t` denotes a horizontal tab and `\n` is a line feed or newline.

#### Double quote escape <a id="double-quote-escape"></a>

Use `\"` to insert a double quote in an interpreted string literal:

```text
fmt.Println("\"foo\"") // Prints: "foo"
```

### Escape HTML <a id="escape-html"></a>

Use [`html.EscpapeString`](https://golang.org/pkg/html/#EscapeString) to encode a string so that it can be safely placed inside HTML text. The function escapes the five characters `<`, `>`, `&`, `'` and `"`.

```text
const s = `"Foo's Bar" <foobar@example.com>`
fmt.Println(html.EscapeString(s))
```

```text
&#34;Foo&#39;s Bar&#34; &lt;foobar@example.com&gt;
```

[`html.UnescapeString`](https://golang.org/pkg/html/#UnescapeString) does the inverse transformation.

### Escape URL <a id="escape-url"></a>

Use [`url.PathEscape`](https://golang.org/pkg/net/url/#PathEscape) in package [`net/url`](https://golang.org/pkg/net/url/) to encode a string so it can be safely placed inside a URL. The function uses [percent-encoding](https://en.wikipedia.org/wiki/Percent-encoding).

```text
const s = `Foo's Bar?`
fmt.Println(url.PathEscape(s))
```

```text
Foo%27s%20Bar%3F
```

[`url.PathUnescape`](https://golang.org/pkg/net/url/#PathUnescape) does the inverse transformation.

### All escape characters <a id="all-escape-characters"></a>

Arbitrary character values can be encoded with backslash escapes and used in string or [rune literals](https://golang.org/ref/spec#Rune_literals). There are four different formats:

* `\x` followed by exactly two hexadecimal digits,
* `\` followed by exactly three octal digits,
* `\u` followed by exactly four hexadecimal digits,
* `\U` followed by exactly eight hexadecimal digits,

where the escapes `\u` and `\U` represent Unicode code points.

The following special escape values are also available.

| Value | Description |
| :--- | :--- |
| `\a` | Alert or bell |
| `\b` | Backspace |
| `\\` | Backslash |
| `\t` | Horizontal tab |
| `\n` | Line feed or newline |
| `\f` | Form feed |
| `\r` | Carriage return |
| `\v` | Vertical tab |
| `\'` | Single quote \(only in rune literals\) |
| `\"` | Double quote \(only in string literals\) |

```text
fmt.Println("\\caf\u00e9") // Prints string: \café
fmt.Printf("%c", '\u00e9') // Prints character: é
```

## 3 ways to split a string into a slice

### Split on comma or other substring <a id="split-on-comma-or-other-substring"></a>

Use the [`strings.Split`](https://golang.org/pkg/strings/#Split) function to split a string into its comma separated values.

```text
s := strings.Split("a,b,c", ",")
fmt.Println(s)
// Output: [a b c]
```

To include the separators, use [`strings.SplitAfter`](https://golang.org/pkg/strings/#SplitAfter). To split only the first n values, use [`strings.SplitN`](https://golang.org/pkg/strings/#SplitN) and [`strings.SplitAfterN`](https://golang.org/pkg/strings/#SplitAfterN).

You can use [`strings.TrimSpace`](https://yourbasic.org/golang/trim-whitespace-from-string/) to strip leading and trailing whitespace from the resulting strings.

### Split by whitespace and newline <a id="split-by-whitespace-and-newline"></a>

Use the [`strings.Fields`](https://golang.org/pkg/strings/#Fields) function to split a string into substrings removing any space characters, including newlines.

```text
s := strings.Fields(" a \t b \n")
fmt.Println(s)
// Output: [a b]
```

### Split on regular expression <a id="split-on-regular-expression"></a>

In more complicated situations, the regexp [`Split`](https://yourbasic.org/golang/regexp-cheat-sheet/#split) method might do the trick.

It splits a string into substrings separated by a regular expression. The method takes an integer argument `n`; if `n >= 0`, it returns at most `n` substrings.

```text
a := regexp.MustCompile(`a`)              // a single `a`
fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

zp := regexp.MustCompile(` *, *`)             // spaces and one comma
fmt.Printf("%q\n", zp.Split("a,b ,  c ", -1)) // ["a" "b" "c "]
```

See this [Regexp tutorial and cheat sheet](https://yourbasic.org/golang/regexp-cheat-sheet/) for a gentle introduction to the Go regexp package with plenty of examples.

## Convert between rune array/slice and string

### Convert string to runes <a id="convert-string-to-runes"></a>

* When you convert a string to a rune slice, you get a new slice that contains the [Unicode code points](https://yourbasic.org/golang/rune/) \(runes\) of the string.
* For an invalid UTF-8 sequence, the rune value will be `0xFFFD` for each invalid byte.

```text
r := []rune("ABC€")
fmt.Println(r)        // [65 66 67 8364]
fmt.Printf("%U\n", r) // [U+0041 U+0042 U+0043 U+20AC]
```

> You can also use a [range loop](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) to access the code points of a string.

### Convert runes to string <a id="convert-runes-to-string"></a>

* When you convert a slice of runes to a string, you get a new string that is the concatenation of the runes converted to UTF-8 encoded strings.
* Values outside the range of valid Unicode code points are converted to `\uFFFD`, the Unicode replacement character `�`.

```text
s := string([]rune{'\u0041', '\u0042', '\u0043', '\u20AC', -1})
fmt.Println(s) // ABC€�
```

### Performance <a id="performance"></a>

These conversions create a new slice or string, and therefore have [time complexity](https://yourbasic.org/algorithms/time-complexity-explained/) proportional to the number of bytes that are processed.

#### More efficient alternative <a id="more-efficient-alternative"></a>

In some cases, you might be able to use a [string builder](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/), which can concatenate strings without redundant copying:

[Efficient string concatenation \[full guide\]](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/)

## Convert between float and string

### String to float <a id="string-to-float"></a>

Use the [`strconv.ParseFloat`](https://golang.org/pkg/strconv/#ParseFloat) function to parse a string as a floating-point number with the precision specified by `bitSize`: 32 for `float32`, or 64 for `float64`.

```text
func ParseFloat(s string, bitSize int) (float64, error)
```

When `bitSize` is 32, the result still has type `float64`, but it will be convertible to `float32` without changing its value.

```text
f := "3.14159265"
if s, err := strconv.ParseFloat(f, 32); err == nil {
    fmt.Println(s) // 3.1415927410125732
}
if s, err := strconv.ParseFloat(f, 64); err == nil {
    fmt.Println(s) // 3.14159265
}
```

### Float to string <a id="float-to-string"></a>

Use the [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) method to format a floating-point number as a string.

```text
s := fmt.Sprintf("%f", 123.456) // s == "123.456000"
```

| Formatting | Description | Verb |
| :--- | :--- | :--- |
| 1.234560e+02 | Scientific notation | `%e` |
| 123.456000 | Decimal point, no exponent | `%f` |
| 123.46 | Default width, precision 2 | `%.2f` |
| ␣␣123.46 | Width 8, precision 2 | `%8.2f` |
| 123.456 | Exponent as needed, necessary digits only | `%g` |

## Convert between int, int64 and string

### int/int64 to string <a id="int-int64-to-string"></a>

Use [`strconv.Itoa`](https://golang.org/pkg/strconv/#Itoa) to convert an int to a decimal string.

```text
s := strconv.Itoa(97) // s == "97"
```

> **Warning:** In a plain [conversion](https://yourbasic.org/golang/conversions/) the value is interpreted as a Unicode code point, and the resulting string will contain the character represented by that code point, encoded in UTF-8.
>
> ```text
> s := string(97) // s == "a"
> ```

Use [`strconv.FormatInt`](https://golang.org/pkg/strconv/#FormatInt) to format an int64 in a given base.

```text
var n int64 = 97
s := strconv.FormatInt(n, 10) // s == "97" (decimal)
```

```text
var n int64 = 97
s := strconv.FormatInt(n, 16) // s == "61" (hexadecimal)
```

### string to int/int64 <a id="string-to-int-int64"></a>

Use [`strconv.Atoi`](https://golang.org/pkg/strconv/#Atoi) to parse a decimal string to an int.

```text
s := "97"
if n, err := strconv.Atoi(s); err == nil {
    fmt.Println(n+1)
} else {
    fmt.Println(s, "is not an integer.")
}
// Output: 98
```

Use [`strconv.ParseInt`](https://golang.org/pkg/strconv/#ParseInt) to parse a decimal string \(base `10`\) and check if it fits into an int64.

```text
s := "97"
n, err := strconv.ParseInt(s, 10, 64)
if err == nil {
    fmt.Printf("%d of type %T", n, n)
}
// Output: 97 of type int64
```

The two numeric arguments represent a base \(0, 2 to 36\) and a bit size \(0 to 64\).

If the first argument is 0, the base is implied by the string’s prefix: base 16 for `"0x"`, base 8 for `"0"`, and base 10 otherwise.

The second argument specifies the integer type that the result must fit into. Bit sizes 0, 8, 16, 32, and 64 correspond to `int`, `int8`, `int16`, `int32`, and `int64`.

### int to int64 \(and back\) <a id="int-to-int64-and-back"></a>

The size of an `int` is implementation-specific, it’s either 32 or 64 bits, and hence you won’t lose any information when converting from int to  int64.

```text
var n int = 97
m := int64(n) // safe
```

> However, when converting to a shorter integer type, the value is **truncated** to fit in the result type's size.

```text
var m int64 = 2 << 32
n := int(m)    // truncated on machines with 32-bit ints
fmt.Println(n) // either 0 or 4,294,967,296
```

* See [Maximum value of an int](https://yourbasic.org/golang/max-min-int-uint/) for code to compute the size of an `int`.
* See [Pick the right one: int vs. int64](https://yourbasic.org/golang/int-vs-int64/) for best practices.

### General formatting \(width, indent, sign\) <a id="general-formatting-width-indent-sign"></a>

The [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function is a useful general tool for converting data to string:

```text
s := fmt.Sprintf("%+8d", 97)
// s == "     +97" (width 8, right justify, always show sign)
```

## Convert interface to string

Use [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) to convert an [interface value](https://yourbasic.org/golang/interfaces-explained/) to a string.

```text
var x interface{} = "abc"
str := fmt.Sprintf("%v", x)
```

In fact, the same technique can be used to get a string representation of any data structure.

```text
var x interface{} = []int{1, 2, 3}
str := fmt.Sprintf("%v", x)
fmt.Println(str) // "[1 2 3]"

```

## Remove all duplicate whitespace



```text
space := regexp.MustCompile(`\s+`)
s := space.ReplaceAllString("Hello  \t \n world!", " ")
fmt.Printf("%q", s) // "Hello world!"
```

`\s+` is a [regular expression](https://yourbasic.org/golang/regexp-cheat-sheet/):

* the character class `\s` matches a space, tab, new line, carriage return or form feed,
* and `+` says “one or more of those”.

In other words, the code will replace each whitespace substring with a single space character.

#### Trim leading and trailing space <a id="trim-leading-and-trailing-space"></a>

To [trim leading and trailing whitespace](https://yourbasic.org/golang/trim-whitespace-from-string/), use the `strings.TrimSpace` function.

## 3 ways to trim whitespace \(or other characters\) from a string

Use the [`strings.TrimSpace`](https://golang.org/pkg/strings/#TrimSpace) function to remove leading and trailing whitespace as defined by Unicode.

```text
s := strings.TrimSpace("\t Goodbye hair!\n ")
fmt.Printf("%q", s) // "Goodbye hair!"
```

* To remove other leading and trailing characters, use [`strings.Trim`](https://golang.org/pkg/strings/#Trim).
* To remove only the leading or the trailing characters, use [`strings.TrimLeft`](https://golang.org/pkg/strings/#TrimLeft) or [`strings.TrimRight`](https://golang.org/pkg/strings/#TrimRight).

## How to reverse a string by byte or rune

### Byte by byte <a id="byte-by-byte"></a>

It’s pretty straightforward to reverse a string one byte at a time.

```text
// Reverse returns a string with the bytes of s in reverse order.
func Reverse(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}
```

### Rune by rune <a id="rune-by-rune"></a>

To reverse a string by UTF-8 encoded characters is a bit trickier.

```text
// ReverseRune returns a string with the runes of s in reverse order.
// Invalid UTF-8 sequences, if any, will be reversed byte by byte.
func ReverseRune(s string) string {
    res := make([]byte, len(s))
    prevPos, resPos := 0, len(s)
    for pos := range s {
        resPos -= pos - prevPos
        copy(res[resPos:], s[prevPos:pos])
        prevPos = pos
    }
    copy(res[0:], s[prevPos:])
    return string(res)
}
```

#### Example usage <a id="example-usage"></a>

```text
for _, s := range []string{
	"Ångström",
	"Hello, 世界",
	"\xff\xfe\xfd", // invalid UTF-8
} {
	fmt.Printf("%q\n", ReverseRune(s))
}
```

```text
"mörtsgnÅ"
"界世 ,olleH"
"\xfd\xfe\xff"
```

## Maps



## Maps explained: create, add, get, delete

Go maps are implemented by hash tables and have efficient add, get and delete operations.



### Create a new map <a id="create-a-new-map"></a>

```text
var m map[string]int                // nil map of string-int pairs

m1 := make(map[string]float64)      // Empty map of string-float64 pairs
m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

m3 := map[string]float64{           // Map literal
    "e":  2.71828,
    "pi": 3.1416,
}
fmt.Println(len(m3))                // Size of map: 2
```

* A map \(or dictionary\) is an **unordered** collection of **key-value** pairs, where each key is **unique**.
* You create a new map with a [**make**](https://golang.org/pkg/builtin/#make) statement or a **map literal**.
* The default **zero value** of a map is `nil`. A nil map is equivalent to an empty map except that **elements can’t be added**.
* The [**`len`**](https://golang.org/pkg/builtin/#len) function returns the **size** of a map, which is the number of key-value pairs.

> **Warning:** If you try to add an element to an uninitialized map you get the mysterious run-time error [_Assignment to entry in nil map_](https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/).

### Add, update, get and delete keys/values <a id="add-update-get-and-delete-keys-values"></a>

```text
m := make(map[string]float64)

m["pi"] = 3.14             // Add a new key-value pair
m["pi"] = 3.1416           // Update value
fmt.Println(m)             // Print map: "map[pi:3.1416]"

v := m["pi"]               // Get value: v == 3.1416
v = m["pie"]               // Not found: v == 0 (zero value)

_, found := m["pi"]        // found == true
_, found = m["pie"]        // found == false

if x, found := m["pi"]; found {
    fmt.Println(x)
}                           // Prints "3.1416"

delete(m, "pi")             // Delete a key-value pair
fmt.Println(m)              // Print map: "map[]"
```

* When you index a map you get two return values; the second one \(which is optional\) is a boolean that indicates if the key exists.
* If the key doesn’t exist, the first value will be the default [zero value](https://yourbasic.org/golang/default-zero-value/).

### For-each range loop <a id="for-each-range-loop"></a>

```text
m := map[string]float64{
    "pi": 3.1416,
    "e":  2.71828,
}
fmt.Println(m) // "map[e:2.71828 pi:3.1416]"

for key, value := range m { // Order not specified 
    fmt.Println(key, value)
}
```

* Iteration order is not specified and may vary from iteration to iteration.
* If an entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced.
* If an entry is created during iteration, that entry may or may not be produced during the iteration.

> Starting with [Go 1.12](https://tip.golang.org/doc/go1.12), the [fmt package](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) prints maps in key-sorted order to ease testing.

### Performance and implementation <a id="performance-and-implementation"></a>

* Maps are backed by [hash tables](https://yourbasic.org/algorithms/hash-tables-explained/).
* Add, get and delete operations run in **constant** expected time. The time complexity for the add operation is [amortized](https://yourbasic.org/algorithms/amortized-time-complexity-analysis/).
* The comparison operators `==` and `!=` must be defined for the key type.

## 3 ways to find a key in a map



### Basics <a id="basics"></a>

When you index a [map](https://yourbasic.org/golang/maps-explained/) in Go you get two return values; the second one \(which is optional\) is a boolean that indicates if the key exists.

If the key doesn’t exist, the first value will be the default [zero value](https://yourbasic.org/golang/default-zero-value/).

### Check second return value <a id="check-second-return-value"></a>

```text
m := map[string]float64{"pi": 3.14}
v, found := m["pi"] // v == 3.14  found == true
v, found = m["pie"] // v == 0.0   found == false
_, found = m["pi"]  // found == true
```

### Use second return value directly in an if statement <a id="use-second-return-value-directly-in-an-if-statement"></a>

```text
m := map[string]float64{"pi": 3.14}
if v, found := m["pi"]; found {
    fmt.Println(v)
}
// Output: 3.14
```

### Check for zero value <a id="check-for-zero-value"></a>

```text
m := map[string]float64{"pi": 3.14}

v := m["pi"] // v == 3.14
v = m["pie"] // v == 0.0 (zero value)
```

> **Warning:** This approach doesn't work if the zero value is a possible key.



## Get slices of keys and values from a map

You can use a range statement to extract slices of keys and values from a [map](https://yourbasic.org/golang/maps-explained/).

```text
keys := make([]keyType, 0, len(myMap))
values := make([]valueType, 0, len(myMap))

for k, v := range myMap {
	keys = append(keys, k)
	values = append(values, v)
}
```

## Sort a map by key or value



* A [map](https://yourbasic.org/golang/maps-explained/) is an **unordered** collection of key-value pairs.
* If you need a stable iteration order, you must maintain a separate data structure.

This example uses a sorted slice of keys to print a `map[string]int` in key order.

```text
m := map[string]int{"Alice": 23, "Eve": 2, "Bob": 25}

keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
	fmt.Println(k, m[k])
}
```

Output:

```text
Alice 23
Bob 25
Eve 2
```

> Also, starting with [Go 1.12](https://tip.golang.org/doc/go1.12), the [`fmt`](https://golang.org/pkg/fmt/) package prints maps in key-sorted order to ease testing.

## Slice and Arrays

## Slices/arrays explained: create, index, slice, iterate



### Basics <a id="basics"></a>

A slice doesn’t store any data, it just describes a section of an underlying [array](https://yourbasic.org/algorithms/time-complexity-arrays/).

* When you change an element of a slice, you modify the corresponding element of its underlying array, and other slices that share the same underlying array will see the change.
* A slice can grow and shrink within the bounds of the underlying array.
* Slices are indexed in the usual way: `s[i]` accesses the `i`th element, starting from zero. 

### Construction <a id="construction"></a>

```text
var s []int                   // a nil slice
s1 := []string{"foo", "bar"}
s2 := make([]int, 2)          // same as []int{0, 0}
s3 := make([]int, 2, 4)       // same as new([4]int)[:2]
fmt.Println(len(s3), cap(s3)) // 2 4
```

* The default **zero value** of a slice is `nil`. The functions `len`, `cap` and `append` all regard `nil` as an empty slice with 0 capacity.
* You create a slice either by a **slice literal** or a call to the [`make`](https://golang.org/pkg/builtin/#make) function, which takes the **length** and an optional **capacity** as arguments.
* The built-in [`len`](https://golang.org/pkg/builtin/#len) and [`cap`](https://golang.org/pkg/builtin/#cap) functions retrieve the length and capacity.

### Slicing <a id="slicing"></a>

```text
a := [...]int{0, 1, 2, 3} // an array
s := a[1:3]               // s == []int{1, 2}        cap(s) == 3
s = a[:2]                 // s == []int{0, 1}        cap(s) == 4
s = a[2:]                 // s == []int{2, 3}        cap(s) == 2
s = a[:]                  // s == []int{0, 1, 2, 3}  cap(s) == 4
```

You can also create a slice by slicing an existing array or slice.

* A slice is formed by specifying a low bound and a high bound: `a[low:high]`. This selects a half-open range which includes the first element, but excludes the last.
* You may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

```text
s := []int{0, 1, 2, 3, 4} // a slice
s = s[1:4]                // s == []int{1, 2, 3}
s = s[1:2]                // s == []int{2} (index relative to slice)
s = s[:3]                 // s == []int{2, 3, 4} (extend length)
```

When you slice a slice, the indexes are relative to the slice itself, not to the backing array.

* The high bound is not bound by the slice’s length, but by it’s capacity, which means you can extend the length of the slice.
* Trying to extend beyond the capacity causes a panic.

### Iteration <a id="iteration"></a>

```text
s := []string{"Foo", "Bar"}
for i, v := range s {
    fmt.Println(i, v)
}
```

```text
0 Foo
1 Bar
```

* The range expression, `s`, is **evaluated once** before beginning the loop.
* The iteration values are assigned to the respective iteration variables, `i` and `v`, **as in an assignment statement**.
* The second iteration variable is optional.
* If the slice is `nil`, the number of iterations is 0.

### Append and copy <a id="append-and-copy"></a>

* The `append` function appends elements to a slice. It will **automatically allocate** a larger backing array if the capacity is exceeded. See [Append function](https://yourbasic.org/golang/append-explained/).
* The `copy` function copies elements into a destination slice `dst` from a source slice `src`. The number of elements copied is the **minimum** of `len(dst)` and `len(src)`. See [Copy function](https://yourbasic.org/golang/copy-explained/).

### Stacks and queues <a id="stacks-and-queues"></a>

The idiomatic way to implement a stack or queue in Go is to use a slice directly. For code examples, see

* [Implement a stack \(LIFO\)](https://yourbasic.org/golang/implement-stack/)
* [Implement a FIFO queue](https://yourbasic.org/golang/implement-fifo-queue/)

## 3 ways to compare slices \(arrays\)

### Basic case <a id="basic-case"></a>

In most cases, you will want to write your own code to compare the elements of two [**slices**](https://yourbasic.org/golang/slices-explained/).

```text
// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
```

For [**arrays**](https://yourbasic.org/golang/slices-explained/), however, you can use the comparison operators `==` and `!=`.

```text
a := [2]int{1, 2}
b := [2]int{1, 3}
fmt.Println(a == b) // false
```

> Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.[The Go Programming Language Specification: Comparison operators](https://golang.org/ref/spec#Comparison_operators)

### Optimized code for byte slices <a id="optimized-code-for-byte-slices"></a>

To compare byte slices, use the optimized [`bytes.Equal`](https://golang.org/pkg/bytes/#Equal). This function also treats nil arguments as equivalent to empty slices.

### General-purpose code for recursive comparison <a id="general-purpose-code-for-recursive-comparison"></a>

For testing purposes, you may want to use [`reflect.DeepEqual`](https://golang.org/pkg/reflect/#DeepEqual). It compares two elements of any type recursively.

```text
var a []int = nil
var b []int = make([]int, 0)
fmt.Println(reflect.DeepEqual(a, b)) // false
```

The performance of this function is much worse than for the code above, but it’s useful in test cases where simplicity and correctness are crucial. The semantics, however, are quite complicated.

## How to best clear a slice: empty vs. nil

### Remove all elements <a id="remove-all-elements"></a>

To remove all elements, simply set the slice to `nil`.

```text
a := []string{"A", "B", "C", "D", "E"}
a = nil
fmt.Println(a, len(a), cap(a)) // [] 0 0
```

This will release the underlying array to the garbage collector \(assuming there are no other references\).

### Keep allocated memory <a id="keep-allocated-memory"></a>

To keep the underlying array, slice the slice to zero length.

```text
a := []string{"A", "B", "C", "D", "E"}
a = a[:0]
fmt.Println(a, len(a), cap(a)) // [] 0 5
```

If the slice is extended again, the original data reappears.

```text
fmt.Println(a[:2]) // [A B]
```

### Empty slice vs. nil slice <a id="empty-slice-vs-nil-slice"></a>

In practice, **nil slices** and **empty slices** can often be treated in the same way:

* they have zero length and capacity,
* they can be used with the same effect in [for loops](https://yourbasic.org/golang/for-loop/) and [append functions](https://yourbasic.org/golang/append-explained/),
* and they even look the same when printed.

```text
var a []int = nil
fmt.Println(len(a)) // 0
fmt.Println(cap(a)) // 0
fmt.Println(a)      // []
```

However, if needed, you can tell the difference.

```text
var a []int = nil
var a0 []int = make([]int, 0)

fmt.Println(a == nil)  // true
fmt.Println(a0 == nil) // false

fmt.Printf("%#v\n", a)  // []int(nil)
fmt.Printf("%#v\n", a0) // []int{}
```

The official Go wiki recommends using nil slices over empty slices.

> \[…\] the nil slice is the preferred style.
>
> Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects \(a nil slice encodes to null, while \[\]string{} encodes to the JSON array \[\]\).
>
> When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.[The Go wiki: Declaring empty slices](https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices)

#### Further reading <a id="further-reading"></a>

[Slices and arrays in 6 easy steps](https://yourbasic.org/golang/slices-explained/)

## 2 ways to delete an element from a slice

### Fast version \(changes order\) <a id="fast-version-changes-order"></a>

```text
a := []string{"A", "B", "C", "D", "E"}
i := 2

// Remove the element at index i from a.
a[i] = a[len(a)-1] // Copy last element to index i.
a[len(a)-1] = ""   // Erase last element (write zero value).
a = a[:len(a)-1]   // Truncate slice.

fmt.Println(a) // [A B E D]
```

The code copies a single element and runs in **constant time**.

### Slow version \(maintains order\) <a id="slow-version-maintains-order"></a>

```text
a := []string{"A", "B", "C", "D", "E"}
i := 2

// Remove the element at index i from a.
copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
a[len(a)-1] = ""     // Erase last element (write zero value).
a = a[:len(a)-1]     // Truncate slice.

fmt.Println(a) // [A B D E]
```

The code copies len\(a\) - i - 1 elements and runs in **linear time**.

[Slices and arrays in 6 easy steps](https://yourbasic.org/golang/slices-explained/)

## Find element in slice/array with linear or binary search

### Linear search <a id="linear-search"></a>

Go doesn’t have an out-of-the-box linear search function for [slices and arrays](https://yourbasic.org/golang/slices-explained/). Here are two example **linear search** implemen­tations, which you can use as templates.

```text
// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}
```

### Binary search <a id="binary-search"></a>

Binary search is faster than linear search, but only works if your data is in order. It's a sortcut. – Dan Bentley

If the array is sorted, you can use a binary search instead. This will be much more efficient, since binary search runs in worst-case logarithmic time, making _**O**_**\(log** _**n**_**\)** comparisons, where _n_ is the size of the slice.

There are the three custom binary search functions: [`sort.SearchInts`](https://golang.org/pkg/sort/#SearchInts), [`sort.SearchStrings`](https://golang.org/pkg/sort/#SearchStrings) or [`sort.SearchFloat64s`](https://golang.org/pkg/sort/#SearchFloat64s).

They all have the signature

```text
func SearchType(a []Type, x Type) int
```

and return

* the smallest index `i` at which `x <= a[i]`
* or `len(a)` if there is no such index.

The slice must be sorted in **ascending order**.

```text
a := []string{"A", "C", "C"}

fmt.Println(sort.SearchStrings(a, "A")) // 0
fmt.Println(sort.SearchStrings(a, "B")) // 1
fmt.Println(sort.SearchStrings(a, "C")) // 1
fmt.Println(sort.SearchStrings(a, "D")) // 3
```

#### Generic binary search <a id="generic-binary-search"></a>

There is also a **generic binary search** function [`sort.Search`](https://golang.org/pkg/sort/#Search).

```text
func Search(n int, f func(int) bool) int
```

It returns

* the smallest index `i` at which `f(i)` is true,
* or `n` if there is no such index.

It requires that `f` is false for some \(possibly empty\) prefix of the input range and then true for the remainder.

This example mirrors the one above, but uses the generic [`sort.Search`](https://golang.org/pkg/sort/#Search) instead of [`sort.SearchInts`](https://golang.org/pkg/sort/#SearchInts).

```text
a := []string{"A", "C", "C"}
x := "C"

i := sort.Search(len(a), func(i int) bool { return x <= a[i] })
if i < len(a) && a[i] == x {
    fmt.Printf("Found %s at index %d in %v.\n", x, i, a)
} else {
    fmt.Printf("Did not find %s in %v.\n", x, a)
}
// Output: Found C at index 1 in [A C C].
```

### The map option <a id="the-map-option"></a>

If you are doing repeated searches and updates, you may want to use a [map](https://yourbasic.org/golang/maps-explained/) instead of a slice. A map provides lookup, insert, and delete operations in _**O**_**\(1\)** expected [amortized time](https://yourbasic.org/algorithms/amortized-time-complexity-analysis/).

## Last item in a slice/array

Use the index `len(a)-1` to access the last element of a slice or array `a`.

```text
a := []string{"A", "B", "C"}
s := a[len(a)-1] // C
```

> Go doesn't have negative indexing like Python does. This is a deliberate design decision — keeping the language simple can help save you from [subtle bugs](https://github.com/golang/go/issues/11245).

### Remove last element <a id="remove-last-element"></a>

```text
a = a[:len(a)-1] // [A B]
```

#### Watch out for memory leaks <a id="watch-out-for-memory-leaks"></a>

> **Warning:** If the slice is permanent and the element temporary, you may want to remove the reference to the element before slicing it off.
>
> ```text
> a[len(a)-1] = "" // Erase element (write zero value)
> a = a[:len(a)-1] // [A B]
> ```

## Files

## Read a file \(stdin\) line by line

### Read from file <a id="read-from-file"></a>

Use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to read a file line by line.

```text
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

### Read from stdin <a id="read-from-stdin"></a>

Use [`os.Stdin`](https://golang.org/pkg/os/#pkg-variables) to read from the standard input stream.

```text
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	log.Println(err)
}
```

### Read from any stream <a id="read-from-any-stream"></a>

A bufio.Scanner can read from any stream of bytes, as long as it implements the [`io.Reader`](https://golang.org/pkg/io/#Reader) interface. See [How to use the io.Reader interface](https://yourbasic.org/golang/io-reader-interface-explained/).

## Append text to a file

This code appends a line of text to the file `text.log`. It creates the file if it doesn’t already exist.

```text
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()
if _, err := f.WriteString("text to append\n"); err != nil {
	log.Println(err)
}

```

## Find current working directory



### Current directory <a id="current-directory"></a>

Use [`os.Getwd`](https://golang.org/pkg/os/#Getwd) to find the path name for the current directory.

```text
path, err := os.Getwd()
if err != nil {
    log.Println(err)
}
fmt.Println(path)  // for example /home/user
```

> **Warning:** If the current directory can be reached via multiple paths \(due to symbolic links\), Getwd may return any one of them.

### Current executable <a id="current-executable"></a>

Use [`os.Executable`](https://golang.org/pkg/os/#Executable) to find the path name for the executable that started the current process.

```text
path, err := os.Executable()
if err != nil {
    log.Println(err)
}
fmt.Println(path) // for example /tmp/go-build872132473/b001/exe/main
```

> **Warning:** There is no guarantee that the path is still pointing to the correct executable. If a symlink was used to start the process, depending on the operating system, the result might be the symlink or the path it pointed to. If a stable result is needed, [`path/filepath.EvalSymlinks`](https://golang.org/pkg/path/filepath/#EvalSymlinks) might help.

## List all files \(recursively\) in a directory

### Directory listing <a id="directory-listing"></a>

Use the [`ioutil.ReadDir`](https://golang.org/pkg/io/ioutil/#ReadDir) function in package [`io/ioutil`](https://golang.org/pkg/io/ioutil/). It returns a sorted slice containing elements of type [`os.FileInfo`](https://golang.org/pkg/os/#FileInfo).

The code in this example prints a sorted list of all file names in the current directory.

```text
files, err := ioutil.ReadDir(".")
if err != nil {
    log.Fatal(err)
}
for _, f := range files {
    fmt.Println(f.Name())
}
```

Example output:

```text
dev
etc
tmp
usr
```

### Visit all files and folders in a directory tree <a id="visit-all-files-and-folders-in-a-directory-tree"></a>

Use the [`filepath.Walk`](https://golang.org/pkg/path/filepath/#Walk) function in package [`path/filepath`](https://golang.org/pkg/path/filepath/).

* It walks a file tree calling a function of type [`filepath.WalkFunc`](https://golang.org/pkg/path/filepath/#WalkFunc) for each file or directory in the tree, including the root.
* The files are walked in lexical order.
* Symbolic links are not followed.

The code in this example lists the paths and sizes of all files and directories in the file tree rooted at the current directory.

```text
err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.Size())
    return nil
})
if err != nil {
    log.Println(err)
}
```

Example output:

```text
. 1644
dev 1644
dev/null 0
dev/random 0
dev/urandom 0
dev/zero 0
etc 1644
etc/group 116
etc/hosts 20
etc/passwd 0
etc/resolv.conf 0
tmp 548
usr 822
usr/local 822
usr/local/go 822
usr/local/go/lib 822
usr/local/go/lib/time 822
usr/local/go/lib/time/zoneinfo.zip 366776
```

## Create a temporary file or directory

### File <a id="file"></a>

Use [`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile) in package [`io/ioutil`](https://golang.org/pkg/io/ioutil/) to create a **globally unique temporary file**. It’s your own job to remove the file when it’s no longer needed.

```text
file, err := ioutil.TempFile("dir", "prefix")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(file.Name())

fmt.Println(file.Name()) // For example "dir/prefix054003078"
```

The call to [`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile)

* creates a new file with a name starting with `"prefix"` in the directory `"dir"`,
* opens the file for reading and writing,
* and returns the new [`*os.File`](https://golang.org/pkg/os/#File).

To put the new file in [`os.TempDir()`](https://golang.org/pkg/os/#TempDir), the default directory for temporary files, call [`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile) with an empty directory string.

#### Add a suffix to the temporary file name[Go 1.11](https://tip.golang.org/doc/go1.11) <a id="add-suffix"></a>

Starting with Go 1.11, if the second string given to `TempFile` includes a `"*"`, the random string replaces this `"*"`.

```text
file, err := ioutil.TempFile("dir", "myname.*.bat")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(file.Name())

fmt.Println(file.Name()) // For example "dir/myname.054003078.bat"
```

If no `"*"` is included the old behavior is retained, and the random digits are appended to the end.

### Directory <a id="directory"></a>

Use [`ioutil.TempDir`](https://golang.org/pkg/io/ioutil/#TempDir) in package [`io/ioutil`](https://golang.org/pkg/io/ioutil/) to create a **globally unique temporary directory**.

```text
dir, err := ioutil.TempDir("dir", "prefix")
if err != nil {
	log.Fatal(err)
}
defer os.RemoveAll(dir)
```

The call to [`ioutil.TempDir`](https://golang.org/pkg/io/ioutil/#TempDir)

* creates a new directory with a name starting with `"prefix"` in the directory `"dir"`
* and returns the path of the new directory.

To put the new directory in [`os.TempDir()`](https://golang.org/pkg/os/#TempDir), the default directory for temporary files, call [`ioutil.TempDir`](https://golang.org/pkg/io/ioutil/#TempDir) with an empty directory string.

## Time and Date

## Format a time or date \[complete guide\]

### Basic example <a id="basic-example"></a>

Go doesn’t use yyyy-mm-dd layout to format a time. Instead, you format a special **layout parameter**

`Mon Jan 2 15:04:05 MST 2006`

the same way as the time or date should be formatted. \(This date is easier to remember when written as `01/02 03:04:05PM ‘06 -0700`.\)

```text
const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
)
date := "1999-12-31"
t, _ := time.Parse(layoutISO, date)
fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
fmt.Println(t.Format(layoutUS)) // December 31, 1999
```

The function

* [`time.Parse`](https://golang.org/pkg/time/#Parse) parses a date string, and
* [`Format`](https://golang.org/pkg/time/#Time.Format) formats a [`time.Time`](https://golang.org/pkg/time/#Time).

They have the following signatures:

```text
func Parse(layout, value string) (Time, error)
func (t Time) Format(layout string) string
```

### Standard time and date formats <a id="standard-time-and-date-formats"></a>

| Go layout | Note |
| :--- | :--- |
| `January 2, 2006` | Date |
| `01/02/06` |  |
| `Jan-02-06` |  |
| `15:04:05` | Time |
| `3:04:05 PM` |  |
| `Jan _2 15:04:05` | Timestamp |
| `Jan _2 15:04:05.000000` | with microseconds |
| `2006-01-02T15:04:05-0700` | [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) \([RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)\) |
| `2006-01-02` |  |
| `15:04:05` |  |
| `02 Jan 06 15:04 MST` | [RFC 822](https://www.ietf.org/rfc/rfc822.txt) |
| `02 Jan 06 15:04 -0700` | with numeric zone |
| `Mon, 02 Jan 2006 15:04:05 MST` | [RFC 1123](https://www.ietf.org/rfc/rfc1123.txt) |
| `Mon, 02 Jan 2006 15:04:05 -0700` | with numeric zone |

The following predefined date and timestamp [format constants](https://golang.org/pkg/time/#pkg-constants) are also available.

```text
ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
```

### Layout options <a id="layout-options"></a>

| Type | Options |
| :--- | :--- |
| Year | `06`   `2006` |
| Month | `01`   `1`   `Jan`   `January` |
| Day | `02`   `2`   `_2`   \(width two, right justified\) |
| Weekday | `Mon`   `Monday` |
| Hours | `03`   `3`   `15` |
| Minutes | `04`   `4` |
| Seconds | `05`   `5` |
| ms μs ns | `.000`   `.000000`   `.000000000` |
| ms μs ns | `.999`   `.999999`   `.999999999`   \(trailing zeros removed\) |
| am/pm | `PM`   `pm` |
| Timezone | `MST` |
| Offset | `-0700`   `-07`   `-07:00`   `Z0700`   `Z07:00` |

### Corner cases <a id="corner-cases"></a>

It’s not possible to specify that an hour should be rendered without a leading zero in a 24-hour time format.

It’s not possible to specify midnight as `24:00` instead of `00:00`. A typical usage for this would be giving opening hours ending at midnight, such as `07:00-24:00`.

It’s not possible to specify a time containing a leap second: `23:59:60`. In fact, the time package assumes a Gregorian calendar without leap seconds.

## Time zones



Each [`Time`](https://golang.org/pkg/time/#Time) has an associated [`Location`](https://golang.org/pkg/time/#Location), which is used for display purposes.

The method [`In`](https://golang.org/pkg/time/#Time.In) returns a time with a specific location. Changing the location in this way changes only the presentation; it does not change the instant in time.

Here is a convenience function that changes the location associated with a time.

```text
// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
    loc, err := time.LoadLocation(name)
    if err == nil {
        t = t.In(loc)
    }
    return t, err
}
```

In use:

```text
for _, name := range []string{
	"",
	"Local",
	"Asia/Shanghai",
	"America/Metropolis",
} {
	t, err := TimeIn(time.Now(), name)
	if err == nil {
		fmt.Println(t.Location(), t.Format("15:04"))
	} else {
		fmt.Println(name, "<time unknown>")
	}
}
```

```text
UTC 19:32
Local 20:32
Asia/Shanghai 03:32
America/Metropolis <time unknown>
```

> **Warning:** A daylight savings time transition skips or repeats times. For example, in the United States, March 13, 2011 2:15am never occurred, while November 6, 2011 1:15am occurred twice. In such cases, the choice of time zone, and therefore the time, is not well-defined. Date returns a time that is correct in one of the two zones involved in the transition, but it does not guarantee which.[Package time: Date](https://golang.org/pkg/time/#Date)

## How to get current timestamp

Use [`time.Now`](https://golang.org/pkg/time/#Now) and one of [`time.Unix`](https://golang.org/pkg/time/#Time.Unix) or [`time.UnixNano`](https://golang.org/pkg/time/#Time.UnixNano) to get a timestamp.

```text
now := time.Now()      // current local time
sec := now.Unix()      // number of seconds since January 1, 1970 UTC
nsec := now.UnixNano() // number of nanoseconds since January 1, 1970 UTC

fmt.Println(now)  // time.Time
fmt.Println(sec)  // int64
fmt.Println(nsec) // int64
```

```text
2009-11-10 23:00:00 +0000 UTC m=+0.000000000
1257894000
1257894000000000000
```

## Get year, month, day from time

he [`Date`](https://golang.org/pkg/time/#Time.Date) function returns the year, month and day of a [`time.Time`](https://golang.org/pkg/time/#Time).

```text
func (t Time) Date() (year int, month Month, day int)
```

In use:

```text
year, month, day := time.Now().Date()
fmt.Println(year, month, day)      // For example 2009 November 10
fmt.Println(year, int(month), day) // For example 2009 11 10
```

You can also extract the information with seperate calls:

```text
t := time.Now()
year := t.Year()   // type int
month := t.Month() // type time.Month
day := t.Day()     // type int
```

The [`time.Month`](https://golang.org/pkg/time/#Month) type specifies a month of the year \(January = 1, …\).

```text
type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)
```

## How to find the day of week

The [`Weekday`](https://golang.org/pkg/time/#Time.Weekday) function returns returns the day of the week of a [`time.Time`](https://golang.org/pkg/time/#Time).

```text
func (t Time) Weekday() Weekday
```

In use:

```text
weekday := time.Now().Weekday()
fmt.Println(weekday)      // "Tuesday"
fmt.Println(int(weekday)) // "2"
```

### Type Weekday <a id="type-weekday"></a>

The [`time.Weekday`](https://golang.org/pkg/time/#Weekday) type specifies a day of the week \(Sunday = 0, …\).

```text
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
```

## Days between two dates

```text
func main() {
    // The leap year 2016 had 366 days.
    t1 := Date(2016, 1, 1)
    t2 := Date(2017, 1, 1)
    days := t2.Sub(t1).Hours() / 24
    fmt.Println(days) // 366
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
```

## Days in a month

To compute the last day of a month, you can use the fact that [`time.Date`](https://golang.org/pkg/time/#Date) accepts values outside their usual ranges – the values are normalized during the conversion.

To compute the number of days in February, look at the day before March 1.

```text
func main() {
    t := Date(2000, 3, 0) // the day before 2000-03-01
    fmt.Println(t)        // 2000-02-29 00:00:00 +0000 UTC
    fmt.Println(t.Day())  // 29
}

func Date(year, month, day int) time.Time {
    return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
```

[`AddDate`](https://golang.org/pkg/time/#Time.AddDate) normalizes its result in the same way. For example, adding one month to October 31 yields December 1, the normalized form of November 31.

```text
t = Date(2000, 10, 31).AddDate(0, 1, 0) // a month after October 31
fmt.Println(t)                          // 2000-12-01 00:00:00 +0000 UTC
```

## Measure execution time

### Measure a piece of code <a id="measure-a-piece-of-code"></a>

```text
start := time.Now()
// Code to measure
duration := time.Since(start)

// Formatted string, such as "2h3m0.5s" or "4.503μs"
fmt.Println(duration)

// Nanoseconds as int64
fmt.Println(duration.Nanoseconds())
```

### Measure a function call <a id="measure-a-function-call"></a>

You can track the execution time of a complete function call with this one-liner, which logs the result to the standard error stream.

```text
func foo() {
    defer duration(track("foo"))
    // Code to measure
}
```

```text
func track(msg string) (string, time.Time) {
    return msg, time.Now()
}

func duration(msg string, start time.Time) {
    log.Printf("%v: %v\n", msg, time.Since(start))
}
```

### Benchmarks <a id="benchmarks"></a>

The [`testing`](https://golang.org/pkg/testing/) package has support for benchmarking that can be used to examine the performance of your code.

## Random Numbers



## Generate random numbers, characters and slice elements

### Go pseudo-random number basics <a id="go-pseudo-random-number-basics"></a>

Use the [`rand.Seed`](https://golang.org/pkg/math/rand/#Seed) and [`rand.Int63`](https://golang.org/pkg/math/rand/#Int63) functions in package [`math/rand`](https://golang.org/pkg/math/rand/) to generate a non-negative pseudo-random number of type `int64`:

```text
rand.Seed(time.Now().UnixNano())
n := rand.Int63() // for example 4601851300195147788
```

Similarly, [`rand.Float64`](https://golang.org/pkg/math/rand/#Float64) generates a pseudo-random float x, where 0 ≤ x &lt; 1:

```text
x := rand.Float64() // for example 0.49893371771268225
```

> **Warning:** Without an initial call to `rand.Seed`, you will get the same sequence of numbers each time you run the program.

See [What’s a seed in a random number generator?](https://yourbasic.org/algorithms/random-number-generator-seed/) for an explanation of pseuodo-random number generators.

#### Several random sources <a id="several-random-sources"></a>

The functions in the [`math/rand`](https://golang.org/pkg/math/rand/) package all use a single random source.

If needed, you can create a new random generator of type [`Rand`](https://golang.org/pkg/math/rand/#Rand) with its own source, and then use its methods to generate random numbers:

```text
generator := rand.New(rand.NewSource(time.Now().UnixNano()))
n := generator.Int63()
x := generator.Float64()
```

### Integers and characters in a given range <a id="integers-and-characters-in-a-given-range"></a>

#### Number between a and b <a id="number-between-a-and-b"></a>

Use [`rand.Intn(m)`](https://golang.org/pkg/math/rand/#Intn), which returns a pseudo-random number n, where 0 ≤ n &lt; m.

```text
n := a + rand.Intn(b-a+1) // a ≤ n ≤ b
```

#### Character between 'a' and 'z' <a id="character-between-39-a-39-and-39-z-39"></a>

```text
c := 'a' + rune(rand.Intn('z'-'a'+1)) // 'a' ≤ c ≤ 'z'
```

### Random element from slice <a id="random-element-from-slice"></a>

To generate a character from an arbitrary set, choose a random index from a slice of characters:

```text
chars := []rune("AB⌘")
c := chars[rand.Intn(len(chars))] // for example '⌘'

```

[Runes and character encoding](https://yourbasic.org/golang/rune/)

## Generate a random string \(password\)

### Random string <a id="random-string"></a>

This code generates a random string of numbers and characters from the Swedish alphabet \(which includes the non-ASCII characters å, ä and ö\).

```text
rand.Seed(time.Now().UnixNano())
chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö" +
    "0123456789")
length := 8
var b strings.Builder
for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
}
str := b.String() // E.g. "ExcbsVQs"
```

> **Warning:** To generate a password, you should use cryptographically secure pseudorandom numbers. See [User-friendly access to crypto/rand](https://yourbasic.org/golang/crypto-rand-int/).

### Random string with restrictions <a id="random-string-with-restrictions"></a>

This code generates a random ASCII string with at least one digit and one special character.

```text
rand.Seed(time.Now().UnixNano())
digits := "0123456789"
specials := "~=+%^*/()[]{}/!@#$?|"
all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
    "abcdefghijklmnopqrstuvwxyz" +
    digits + specials
length := 8
buf := make([]byte, length)
buf[0] = digits[rand.Intn(len(digits))]
buf[1] = specials[rand.Intn(len(specials))]
for i := 2; i < length; i++ {
    buf[i] = all[rand.Intn(len(all))]
}
rand.Shuffle(len(buf), func(i, j int) {
    buf[i], buf[j] = buf[j], buf[i]
})
str := string(buf) // E.g. "3i[g0|)z"
```

#### Before Go 1.10 <a id="before-go-1-10"></a>

In code before Go 1.10, replace the call to [rand.Shuffle](https://golang.org/pkg/math/rand/#Shuffle) with this code:

```text
for i := len(buf) - 1; i > 0; i-- { // Fisher–Yates shuffle
    j := rand.Intn(i + 1)
    buf[i], buf[j] = buf[j], buf[i]
}
```

## Generate a unique string \(UUID, GUID\)

A [universally unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) \(UUID\), or globally unique identifier \(GUID\), is a 128-bit number used to identify information.

* A UUID is for practical purposes unique: the probability that it will be duplicated is very close to zero.
* UUIDs don’t depend on a central authority or on coordination between those generating them.

The string representation of a UUID consists of 32 hexadecimal digits displayed in 5 groups separated by hyphens. For example:

```text
123e4567-e89b-12d3-a456-426655440000
```

### UUID generator example <a id="uuid-generator-example"></a>

You can use the [`rand.Read`](https://golang.org/pkg/crypto/rand/#Read) function from package [`crypto/rand`](https://golang.org/pkg/crypto/rand/) to generate a basic UUID.

```text
b := make([]byte, 16)
_, err := rand.Read(b)
if err != nil {
    log.Fatal(err)
}
uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
    b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
fmt.Println(uuid)
```

```text
9438167c-9493-4993-fd48-950b27aad7c9
```

#### Limitations <a id="limitations"></a>

This UUID doesn’t conform to [RFC 4122](https://tools.ietf.org/html/rfc4122). In particular, it doesn’t contain any version or variant numbers.

> **Warning:** The `rand.Read` call returns an error if the underlying system call fails. For instance if it can't read `/dev/urandom` on a Unix system, or if [`CryptAcquireContext`](https://msdn.microsoft.com/en-us/library/windows/desktop/aa379886%28v=vs.85%29.aspx) fails on a Windows system.

## Shuffle a slice or array

The [`rand.Shuffle`](https://golang.org/pkg/math/rand/#Shuffle) function in package [`math/rand`](https://golang.org/pkg/math/rand/) shuffles an input sequence using a given swap function.

```text
a := []int{1, 2, 3, 4, 5, 6, 7, 8}
rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
```

```text
[5 8 6 4 3 7 2 1]
```

> **Warning:** Without the call to `rand.Seed` you will get the same sequence of pseudo­random numbers each time you run the program.

#### Before Go 1.10 <a id="before-go-1-10"></a>

Use the [`rand.Seed`](https://golang.org/pkg/math/rand/#Seed) and [`rand.Intn`](https://golang.org/pkg/math/rand/#Intn) functions in package [`math/rand`](https://golang.org/pkg/math/rand/).

```text
a := []int{1, 2, 3, 4, 5, 6, 7, 8}
rand.Seed(time.Now().UnixNano())
for i := len(a) - 1; i > 0; i-- { // Fisher–Yates shuffle
    j := rand.Intn(i + 1)
    a[i], a[j] = a[j], a[i]
}
```

## User-friendly access to crypto/rand

Go has two packages for random numbers:

* [`math/rand`](https://golang.org/pkg/math/rand/) implements a large selection of pseudo-random number generators.
* [`crypto/rand`](https://golang.org/pkg/crypto/rand/) implements a cryptographically secure pseudo-random number generator with a limited interface.

The two packages can be combined by calling [`rand.New`](https://golang.org/pkg/math/rand/#New) in package `math/rand` with a source that gets its data from `crypto/rand`.

```text
import (
    crand "crypto/rand"
    rand "math/rand"

    "encoding/binary"
    "fmt"
    "log"
)

func main() {
    var src cryptoSource
    rnd := rand.New(src)
    fmt.Println(rnd.Intn(1000)) // a truly random number 0 to 999
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
    return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
    err := binary.Read(crand.Reader, binary.BigEndian, &v)
    if err != nil {
        log.Fatal(err)
    }
    return v
}
```

> **Warning:** The `crand.Reader` returns an error if the underlying system call fails. For instance if it can't read `/dev/urandom` on a Unix system, or if [`CryptAcquireContext`](https://msdn.microsoft.com/en-us/library/windows/desktop/aa379886%28v=vs.85%29.aspx) fails on a Windows system.

## Language Basics

## Packages explained: declare, import, download, document

### Basics <a id="basics"></a>

Every Go program is made up of packages and each package has an **import path**:

* `"fmt"`
* `"math/rand"`
* `"github.com/yourbasic/graph"`

Packages in the standard library have short import paths, such as `"fmt"` and `"math/rand"`. Third-party packages, such as `"github.com/yourbasic/graph"`, typically have an import path that includes a hosting service \(`github.com`\) and an organization name \(`yourbasic`\).

By convention, the **package name** is the same as the last element of the import path:

* `fmt`
* `rand`
* `graph`

References to other packages’ definitions must always be prefixed with their package names, and only the capitalized names from other packages are accessible.

```text
package main

import (
    "fmt"
    "math/rand"

    "github.com/yourbasic/graph"
)

func main() {
    n := rand.Intn(100)
    g := graph.New(n)
    fmt.Println(g)
}
```

### Declare a package <a id="declare-a-package"></a>

Every Go source file starts with a package declaration, which contains only the package name.

For example, the file [`src/math/rand/exp.go`](https://golang.org/src/math/rand/exp.go), which is part of the implementation of the [`math/rand`](https://golang.org/pkg/math/rand/) package, contains the following code.

```text
package rand
  
import "math"
  
const re = 7.69711747013104972
…
```

You don’t need to worry about package name collisions, only the import path of a package must be unique. [How to Write Go Code](https://golang.org/doc/code.html) shows how to organize your code and its packages in a file structure.

### Package name conflicts <a id="package-name-conflicts"></a>

You can customize the name under which you refer to an imported package.

```text
package main

import (
    csprng "crypto/rand"
    prng "math/rand"

    "fmt"
)

func main() {
    n := prng.Int() // pseudorandom number
    b := make([]byte, 8)
    csprng.Read(b) // cryptographically secure pseudorandom number
    fmt.Println(n, b)
}
```

### Dot imports <a id="dot-imports"></a>

If a period `.` appears instead of a name in an import statement, all the package’s exported identifiers can be accessed without a qualifier.

```text
package main

import (
    "fmt"
    . "math"
)

func main() {
    fmt.Println(Sin(Pi/2)*Sin(Pi/2) + Cos(Pi)/2) // 0.5
}
```

Dot imports can make programs hard to read and **generally should be avoided**.

### Package download <a id="package-download"></a>

The [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) command downloads packages named by import paths, along with their dependencies, and then installs the packages.

```text
$ go get github.com/yourbasic/graph
```

The import path corresponds to the repository hosting the code. This reduces the likelihood of future name collisions.

The [Go Wiki](https://github.com/golang/go/wiki/Projects) and [Awesome Go](https://github.com/avelino/awesome-go) provide lists of high-quality Go packages and resources.

For more information on using remote repositories with the go tool, see [Command go: Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths).

### Package documentation <a id="package-documentation"></a>

The [GoDoc](https://godoc.org/) web site hosts documentation for all public Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad:

* [`https://godoc.org/fmt`](https://godoc.org/fmt)
* [`https://godoc.org/math/rand`](https://godoc.org/math/rand)
* [`https://godoc.org/github.com/yourbasic/graph`](https://godoc.org/github.com/yourbasic/graph)

The [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command extracts and generates documentation for all locally installed Go programs. The following command starts a web server that presents the documentation at `http://localhost:6060/`.

```text
$ godoc -http=:6060 &
```

## Package documentation

### godoc.org website <a id="godoc-org-website"></a>

The [GoDoc](https://godoc.org/) website hosts docu­men­tation for all public Go [packages](https://yourbasic.org/golang/packages-explained/) on Bitbucket, GitHub, Google Project Hosting and Launchpad.

### Local godoc server <a id="local-godoc-server"></a>

The [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command extracts and generates documentation for all locally installed Go programs, both your own code and the standard libraries.

The following command starts a web server that presents the documentation at `http://localhost:6060/`.

```text
$ godoc -http=:6060 &
```

![Web browser localhost:6060](https://yourbasic.org/golang/localhost-6060.png)

The documentation is tightly coupled with the code. For example, you can navigate from a function’s documentation to its implementation with a single click.

### go doc command-line tool <a id="go-doc-command-line-tool"></a>

The [go doc](https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol) command prints plain text documentation to standard output:

```text
$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```

### Create documentation <a id="create-documentation"></a>

To document a function, type, constant, variable, or even a complete package, write a regular comment directly preceding its declaration, with no blank line in between. For example, this is the documentation for the [`fmt.Println`](https://golang.org/src/fmt/print.go?s=7388:7437#L246) function:

```text
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
…
```

For best practices on how to document Go code, see [Effective Go: Commentary](https://golang.org/doc/effective_go.html#commentary).

### Runnable documentation examples <a id="runnable-documentation-examples"></a>

You can add example code snippets to the package documentation; this code is verified by running it as a test. For more information on how to create such testable examples, see [The Go Blog: Testable Examples in Go](https://blog.golang.org/examples).

## Package initialization and program execution order



### Basics <a id="basics"></a>

* First the `main` [package](https://yourbasic.org/golang/packages-explained/) is initialized.
  * Imported packages are initialized before the package itself.
  * Packages are initialized one at a time:
  * first package-level variables are initialized in declaration order,
  * then the `init` functions are run.
* Finally the `main` function is called.

### Program execution <a id="program-execution"></a>

Program execution begins by initializing the `main` package and then calling the function `main`. When `main` returns, the program exits. It **does not wait** for other goroutines to complete.

### Package initialization <a id="package-initialization"></a>

* Package-level variables are initialized in **declaration order**, but after any of the variables they **depend** on.
* Initialization of variables declared in multiple files is done in **lexical file name order**. Variables declared in the first file are declared before any of the variables declared in the second file.
* Initialization cycles are **not allowed**.
* Dependency analysis is performed **per package**; only references referring to variables, functions, and methods declared in the current package are considered.

#### Example <a id="example"></a>

In this example, taken directly from the [Go language specification](https://golang.org/ref/spec#Package_initialization), the initialization order is d, b, c, a.

```text
var (
    a = c + b
    b = f()
    c = f()
    d = 3
)

func f() int {
    d++
    return d
}
```

### Init function <a id="init-function"></a>

Variables may also be initialized using `init` functions.

```text
func init() { … }
```

Multiple such functions may be defined. They cannot be called from inside a program.

* A package with **no imports** is initialized
  * by assigning initial values to all its package-level variables,
  * followed by calling all `init` functions in the order they appear in the source.
* Imported packages are initialized before the package itself.
* Each package is initialized **once**, regardless if it’s imported by multiple other packages.

It follows that there can be **no cyclic dependencies**.

Package initialization happens in a single goroutine, sequentially, one package at a time.

### Warning <a id="warning"></a>

Lexical ordering according to file names is not part of the formal language specification.

> To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.  
> [The Go Programming Language Specification: Package initialization](https://golang.org/ref/spec#Package_initialization)

{% embed url="https://tutorialedge.net/golang/the-go-init-function/" %}

## Statements

## 4 basic if-else statement patterns

### Basic syntax <a id="basic-syntax"></a>

```text
if x > max {
    x = max
}
```

```text
if x <= y {
    min = x
} else {
    min = y
}
```

An **if statement** executes one of two branches according to a boolean expression.

* If the expression evaluates to true, the **if** branch is executed,
* otherwise, if present, the **else** branch is executed.

### With init statement <a id="with-init-statement"></a>

```text
if x := f(); x <= y {
    return x
}
```

The expression may be preceded by a **simple statement**, which executes before the expression is evaluated. The **scope** of `x` is limited to the if statement.

### Nested if statements <a id="nested-if-statements"></a>

```text
if x := f(); x < y {
    return x
} else if x > z {
    return z
} else {
    return y
}
```

Complicated conditionals are often best expressed in Go with a **switch statement**. See [5 switch statement patterns](https://yourbasic.org/golang/switch-statement/) for details.

### Ternary ? operator alternatives <a id="ternary-operator-alternatives"></a>



You can’t write a short one-line conditional in Go; there is no ternary conditional operator. Instead of

```text
res = expr ? x : y
```

you write

```text
if expr {
    res = x
} else {
    res = y
}
```

In some cases, you may want to create a dedicated function.

```text
func Min(x, y int) int {
    if x <= y {
        return x
    }
    return y
}
```

## 5 switch statement patterns

### Basic switch with default <a id="basic-switch-with-default"></a>

* A switch statement runs the first case equal to the condition expression.
* The cases are evaluated from top to bottom, stopping when a case succeeds.
* If no case matches and there is a default case, its statements are executed.

```text
switch time.Now().Weekday() {
case time.Saturday:
    fmt.Println("Today is Saturday.")
case time.Sunday:
    fmt.Println("Today is Sunday.")
default:
    fmt.Println("Today is a weekday.")
}
```

> Unlike C and Java, the case expressions do not need to be constants.

### No condition <a id="no-condition"></a>

A switch without a condition is the same as switch true.

```text
switch hour := time.Now().Hour(); { // missing expression means "true"
case hour < 12:
    fmt.Println("Good morning!")
case hour < 17:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}
```

### Case list <a id="case-list"></a>

```text
func WhiteSpace(c rune) bool {
    switch c {
    case ' ', '\t', '\n', '\f', '\r':
        return true
    }
    return false
}
```

### Fallthrough <a id="fallthrough"></a>

* A `fallthrough` statement transfers control to the next case.
* It may be used only as the final statement in a clause.

```text
switch 2 {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")
    fallthrough
case 3:
    fmt.Println("3")
}
```

```text
2
3
```

### Exit with break <a id="exit-with-break"></a>

A `break` statement terminates execution of the **innermost** `for`, `switch`, or `select` statement.

If you need to break out of a surrounding loop, not the switch, you can put a **label** on the loop and break to that label. This example shows both uses.

```text
Loop:
    for _, ch := range "a b\nc" {
        switch ch {
        case ' ': // skip space
            break
        case '\n': // break at newline
            break Loop
        default:
            fmt.Printf("%c\n", ch)
        }
    }
```

```text
a
b
```

### Execution order <a id="execution-order"></a>

* First the switch expression is evaluated once.
* Then case expressions are evaluated left-to-right and top-to-bottom:
  * the first one that equals the switch expression triggers execution of the statements of the associated case,
  * the other cases are skipped.

```text
// Foo prints and returns n.
func Foo(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Foo(2) {
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
        fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}
```

```text
2
1
2
First case
Second case
```

## 5 basic for loop patterns



### Three-component loop <a id="three-component-loop"></a>

This version of the Go for loop works just as in C or Java.

```text
sum := 0
for i := 1; i < 5; i++ {
    sum += i
}
fmt.Println(sum) // 10 (1+2+3+4)
```

1. The init statement, `i := 1`, runs.
2. The condition, `i < 5`, is computed.
   * If true, the loop body runs,
   * otherwise the loop is done.
3. The post statement, `i++`, runs.
4. Back to step 2.

The scope of `i` is limited to the loop.

### While loop <a id="while-loop"></a>

If you skip the init and post statements, you get a while loop.

```text
n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n) // 8 (1*2*2*2)
```

1. The condition, `n < 5`, is computed.
   * If true, the loop body runs,
   * otherwise the loop is done.
2. Back to step 1.

### Infinite loop <a id="infinite-loop"></a>

If you skip the condition as well, you get an infinite loop.

```text
sum := 0
for {
    sum++ // repeated forever
}
fmt.Println(sum) // never reached
```

### For-each range loop <a id="for-each-range-loop"></a>

Looping over elements in _slices_, _arrays_, _maps_, _channels_ or _strings_ is often better done with a range loop.

```text
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
```

```text
0 hello
1 world
```

See [4 basic range loop patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for a complete set of examples.

### Exit a loop <a id="exit-a-loop"></a>

The break and continue keywords work just as they do in C and Java.

```text
sum := 0
for i := 1; i < 5; i++ {
    if i%2 != 0 { // skip odd numbers
        continue
    }
    sum += i
}
fmt.Println(sum) // 6 (2+4)
```

* A **continue** statement begins the next iteration of the innermost **for** loop at its post statement \(`i++`\).
* A **break** statement leaves the innermost **for**, [**switch**](https://yourbasic.org/golang/switch-statement/) or [**select**](https://yourbasic.org/golang/select-explained/) statement.

## 2 patterns for a do-while loop in Go

There is no **do-while loop** in Go. To emulate the C/Java code

```text
do {
    work();
} while (condition);
```

you may use a [for loop](https://yourbasic.org/golang/for-loop/) in one of these two ways:

```text
for ok := true; ok; ok = condition {
    work()
}
```

```text
for {
    work()
    if !condition {
        break
    }
}
```

#### Repeat-until loop <a id="repeat-until-loop"></a>

To write a **repeat-until loop**

```text
repeat
    work();
until condition;
```

simply change the condition in the code above to its complement:

```text
for ok := true; ok; ok = !condition {
    work()
}
```

```text
for {
    work()
    if condition {
        break
    }
}
```

## 4 basic range loop \(for-each\) patterns



### Basic for-each loop \(slice or array\) <a id="basic-for-each-loop-slice-or-array"></a>

```text
a := []string{"Foo", "Bar"}
for i, s := range a {
    fmt.Println(i, s)
}
```

```text
0 Foo
1 Bar
```

* The range expression, `a`, is **evaluated once** before beginning the loop.
* The iteration values are assigned to the respective iteration variables, `i` and `s`, **as in an assignment statement**.
* The second iteration variable is optional.
* For a nil slice, the number of iterations is 0.

### String iteration: runes or bytes <a id="string-iteration-runes-or-bytes"></a>

For strings, the range loop iterates over [Unicode code points](https://yourbasic.org/golang/rune/).

```text
for i, ch := range "日本語" {
    fmt.Printf("%#U starts at byte position %d\n", ch, i)
}
```

```text
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

* The index is the first byte of a UTF-8-encoded code point; the second value, of type `rune`, is the value of the code point.
* For an invalid UTF-8 sequence, the second value will be 0xFFFD, and the iteration will advance a single byte.

> To loop over individual bytes, simply use a [normal for loop](https://yourbasic.org/golang/for-loop/) and string indexing:
>
> ```text
> const s = "日本語"
> for i := 0; i < len(s); i++ {
>     fmt.Printf("%x ", s[i])
> }
> ```
>
> ```text
> e6 97 a5 e6 9c ac e8 aa 9e
> ```

### Map iteration: keys and values <a id="map-iteration-keys-and-values"></a>

The iteration order over [maps](https://yourbasic.org/golang/maps-explained/) is not specified and is not guaranteed to be the same from one iteration to the next.

```text
m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
for k, v := range m {
    fmt.Println(k, v)
}
```

```text
two 2
three 3
one 1
```

* If a map entry that has not yet been reached is removed during iteration, this value will not be produced.
* If a map entry is created during iteration, that entry may or may not be produced.
* For a nil map, the number of iterations is 0.

### Channel iteration <a id="channel-iteration"></a>

For [channels](https://yourbasic.org/golang/channels-explained/), the iteration values are the successive values sent on the channel until closed.

```text
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()
for n := range ch {
    fmt.Println(n)
}
```

```text
1
2
3
```

* For a nil channel, the range loop blocks forever.

### Gotchas <a id="gotchas"></a>

Here are two traps that you want to avoid when using range loops:

* [Unexpected values in range loop](https://yourbasic.org/golang/gotcha-unexpected-values-range/)
* [Can’t change entries in range loop](https://yourbasic.org/golang/gotcha-change-value-range/)

## Defer a function call \(with return value\)

### Defer statement basics <a id="defer-statement-basics"></a>

A `defer` statement postpones the execution of a function until the surrounding function returns, either normally or through a panic.

```text
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

```text
Hello
World
```

Deferred calls are executed even when the function panics:

```text
func main() {
    defer fmt.Println("World")
    panic("Stop")
    fmt.Println("Hello")
}
```

```text
World
panic: Stop

goroutine 1 [running]:
main.main()
    ../main.go:3 +0xa0
```

#### Order of execution <a id="order-of-execution"></a>

The deferred call’s **arguments are evaluated immediately**, even though the function call is not executed until the surrounding function returns.

If there are several deferred function calls, they are executed in last-in-first-out order.

```text
func main() {
    fmt.Println("Hello")
    for i := 1; i <= 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("World")
}
```

```text
Hello
World
3
2
1
```

#### Use func to return a value <a id="use-func-to-return-a-value"></a>

Deferred anonymous functions may access and modify the surrounding function’s named return parameters.

In this example, the `foo` function returns “Change World”.

```text
func foo() (result string) {
    defer func() {
        result = "Change World" // change value at the very last moment
    }()
    return "Hello World"
}
```

### Common applications <a id="common-applications"></a>

Defer is often used to perform clean-up actions, such as closing a file or unlocking a mutex. Such actions should be performed both when the function returns normally and when it panics.

#### Close a file <a id="close-a-file"></a>

In this example, defer statements are used to ensure that all files are closed before leaving the `CopyFile` function, whichever way that happens.

```text
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

#### Error handling: catch a panic <a id="error-handling-catch-a-panic"></a>

The [Recover from a panic](https://yourbasic.org/golang/recover-from-panic/#recover-and-catch-a-panic) code example shows how to use a defer statement to recover from a panic and update the return value.

## Type assertions and type switches

### Type assertions <a id="type-assertions"></a>

A **type assertion** doesn’t really convert an [interface](https://yourbasic.org/golang/interfaces-explained/) to another data type, but it provides access to an interface’s concrete value, which is typically what you want.

The type assertion `x.(T)` asserts that the concrete value stored in `x` is of type `T`, and that `x` is not nil.

* If `T` is not an interface, it asserts that the dynamic type of `x` is identical to `T`.
* If `T` is an interface, it asserts that the dynamic type of `x` implements `T`.

```text
var x interface{} = "foo"

var s string = x.(string)
fmt.Println(s)     // "foo"

s, ok := x.(string)
fmt.Println(s, ok) // "foo true"

n, ok := x.(int)
fmt.Println(n, ok) // "0 false"

n = x.(int)        // ILLEGAL
```

```text
panic: interface conversion: interface {} is string, not int
```

### Type switches <a id="type-switches"></a>

A **type switch** performs several type assertions in series and runs the first case with a matching type.

```text
var x interface{} = "foo"

switch v := x.(type) {
case nil:
    fmt.Println("x is nil")            // here v has type interface{}
case int: 
    fmt.Println("x is", v)             // here v has type int
case bool, string:
    fmt.Println("x is bool or string") // here v has type interface{}
default:
    fmt.Println("type unknown")        // here v has type interface{}
}
```

```text
x is bool or string
```

## Type alias explained

An **alias declaration** has the form

```text
type T1 = T2
```

as opposed to a standard **type definition**

```text
type T1 T2
```

An alias declaration doesn’t create a new distinct type different from the type it’s created from. It just introduces an alias name `T1`, an alternate spelling, for the type denoted by `T2`.

Type aliases are not meant for everyday use. They were introduced to support gradual code repair while moving a type between packages during large-scale refactoring. [Codebase Refactoring \(with help from Go\)](https://talks.golang.org/2016/refactor.article) covers this in detail.

## Expressions



## Create, initialize and compare structs

### Struct types <a id="struct-types"></a>

A struct is a typed collection of fields, useful for grouping data into records.

```text
type Student struct {
    Name string
    Age  int
}

var a Student    // a == Student{"", 0}
a.Name = "Alice" // a == Student{"Alice", 0}
```

* To define a new **struct type**, you list the names and types of each field.
* The default **zero value** of a struct has all its fields zeroed.
* You can access individual fields with **dot notation**.

### 2 ways to create and initialize a new struct <a id="2-ways-to-create-and-initialize-a-new-struct"></a>

The **`new`** keyword can be used to create a new struct. It returns a [pointer](https://yourbasic.org/golang/pointers-explained/) to the newly created struct.

```text
var pa *Student   // pa == nil
pa = new(Student) // pa == &Student{"", 0}
pa.Name = "Alice" // pa == &Student{"Alice", 0}
```

You can also create and initialize a struct with a **struct literal**.

```text
b := Student{ // b == Student{"Bob", 0}
    Name: "Bob",
}
    
pb := &Student{ // pb == &Student{"Bob", 8}
    Name: "Bob",
    Age:  8,
}

c := Student{"Cecilia", 5} // c == Student{"Cecilia", 5}
d := Student{}             // d == Student{"", 0}
```

* An element list that contains keys does not need to have an element for each struct field. Omitted fields get the zero value for that field.
* An element list that does not contain any keys must list an element for each struct field in the order in which the fields are declared.
* A literal may omit the element list; such a literal evaluates to the zero value for its type.

For further details, see [The Go Language Specification: Composite literals](https://golang.org/ref/spec#Composite_literals).

### Compare structs <a id="compare-structs"></a>

You can compare struct values with the comparison operators `==` and `!=`. Two values are equal if their corresponding fields are equal.

```text
d1 := Student{"David", 1}
d2 := Student{"David", 2}
fmt.Println(d1 == d2) // false
```

## Pointers explained

A pointer is a vari­able that con­tains the address of an object.

### Basics <a id="basics"></a>

Structs and arrays are **copied** when used in assignments and passed as arguments to functions. With pointers this can be avoided.

Pointers store **addresses** of objects. The addresses can be passed around more efficiently than the actual objects.

A pointer has type `*T`. The keyword `new` allocates a new object and returns its address.

```text
type Student struct {
    Name string
}

var ps *Student = new(Student) // ps holds the address of the new struct
```

The variable declaration can be written more compactly.

```text
ps := new(Student)
```

### Address operator <a id="address-operator"></a>

The `&` operator returns the address of an object.

```text
s := Student{"Alice"} // s holds the actual struct 
ps := &s              // ps holds the address of the struct 
```

The `&` operator can also be used with **composite literals**. The two lines above can be written as

```text
ps := &Student{"Alice"}
```

### Pointer indirection <a id="pointer-indirection"></a>

For a pointer `x`, the **pointer indirection** `*x` denotes the value which `x` points to. Pointer indirection is rarely used, since Go can automatically take the address of a variable.

```text
ps := new(Student)
ps.Name = "Alice" // same as (*ps).Name = "Alice"
```

### Pointers as parameters <a id="pointers-as-parameters"></a>

When using a pointer to modify an object, you’re affecting all code that uses the object.

```text
// Bob is a function that has no effect.
func Bob(s Student) {
    s.Name = "Bob" // changes only the local copy
}

// Charlie sets pp.Name to "Charlie".
func Charlie(ps *Student) {
    ps.Name = "Charlie"
}

func main() {
    s := Student{"Alice"}

    Bob(s)
    fmt.Println(s) // prints {Alice}

    Charlie(&s)
    fmt.Println(s) // prints {Charlie}
}
```

## Untyped numeric constants with no limits

Constants may be **typed** or **untyped**.

```text
const a uint = 17
const b = 55
```

An untyped constant has **no limits**. When it’s used in a context that requires a type, a type will be inferred and a limit applied.

```text
const big = 10000000000  // Ok, even though it's too big for an int.
const bigger = big * 100 // Still ok.
var i int = big / 100    // No problem: the new result fits in an int.

// Compile time error: "constant 10000000000 overflows int"
var j int = big
```

The inferred type is determined by the syntax of the value:

* `123` gets type `int`, and
* `123.4` becomes a `float64`.

The other possibilities are `rune` \(alias for `int32`\) and `complex128`.

### Enumerations <a id="enumerations"></a>

Go does not have enumerated types. Instead, you can use the special name `iota` in a single `const` declaration to get a series of increasing values. When an initialization expression is omitted for a `const`, it reuses the preceding expression.

```text
const (
    red = iota // red == 0
    blue       // blue == 1
    green      // green == 2
)
```

See [4 iota enum examples](https://yourbasic.org/golang/iota/) for further examples.

## Make slices, maps and channels

[Slices](https://yourbasic.org/golang/slices-explained/), [maps](https://yourbasic.org/golang/maps-explained/) and [channels](https://yourbasic.org/golang/channels-explained/) can be created with the built-in `make` function. The memory is initialized with [zero values](https://yourbasic.org/golang/default-zero-value/).

| Call | Type | Description |
| :--- | :--- | :--- |
| `make(T, n)` | slice | slice of type T with length n |
| `make(T, n, c)` |  | capacity c |
| `make(T)` | map | map of type T |
| `make(T, n)` |  | initial room for approximately n elements |
| `make(T)` | channel | unbuffered channel of type T |
| `make(T, n)` |  | buffered channel with buffer size n |

```text
s := make([]int, 10, 100)      // slice with len(s) == 10, cap(s) == 100
m := make(map[string]int, 100) // map with initial room for ~100 elements
c := make(chan int, 10)        // channel with a buffer size of 10
```

Slices, arrays and maps can also be created with [composite literals](https://golang.org/ref/spec#Composite_literals).

```text
s := []string{"f", "o", "o"} // slice with len(s) == 3, cap(s) == 3
a := [...]int{1, 2}          // array with len(a) == 2
m := map[string]float64{     // map with two key-value elements
    "e":  2.71828,
    "pi": 3.1416,
}
```

## How to append anything \(element, slice or string\) to a slice

### Append function basics <a id="append-function-basics"></a>

With the built-in [append function](https://golang.org/ref/spec#Appending_and_copying_slices) you can use a slice as a [dynamic array](https://yourbasic.org/algorithms/time-complexity-arrays/). The function appends any number of elements to the end of a [slice](https://yourbasic.org/golang/slices-explained/):

* if there is enough capacity, the underlying array is reused;
* if not, a new underlying array is allocated and the data is copied over.

Append **returns the updated slice**. Therefore you need to store the result of an append, often in the variable holding the slice itself:

```text
a := []int{1, 2}
a = append(a, 3, 4) // a == [1 2 3 4]
```

In particular, it’s perfectly fine to **append to an empty slice**:

```text
a := []int{}
a = append(a, 3, 4) // a == [3 4]
```

> **Warning:** See [Why doesn’t append work every time?](https://yourbasic.org/golang/gotcha-append/) for an example of what can happen if you forget that `append` may reuse the underlying array.

### Append one slice to another <a id="append-one-slice-to-another"></a>

You can **concatenate two slices** using the [three dots notation](https://yourbasic.org/golang/variadic-function/):

```text
a := []int{1, 2}
b := []int{11, 22}
a = append(a, b...) // a == [1 2 11 22]
```

The `...` unpacks `b`. Without the dots, the code would attempt to append the slice as a whole, which is invalid.

The result does not depend on whether the **arguments overlap**:

```text
a := []int{1, 2}
a = append(a, a...) // a == [1 2 1 2]
```

### Append string to byte slice <a id="append-string-to-byte-slice"></a>

As a special case, it’s legal to append a string to a byte slice:

```text
slice := append([]byte("Hello "), "world!"...)
```

### Performance <a id="performance"></a>

Appending a single element takes **constant amortized time**. See [Amortized time complexity](https://yourbasic.org/algorithms/amortized-time-complexity-analysis/) for a detailed explanation.

## How to use the copy function

The built-in [copy function](https://golang.org/ref/spec#Appending_and_copying_slices) copies elements into a destination slice `dst` from a source slice `src`.

```text
func copy(dst, src []Type) int
```

It returns the number of elements copied, which will be the **minimum** of `len(dst)` and `len(src)`. The result does not depend on whether the arguments overlap.

As a **special case**, it’s legal to copy bytes from a string to a slice of bytes.

```text
copy(dst []byte, src string) int
```

### Examples <a id="examples"></a>

#### Copy from one slice to another <a id="copy-from-one-slice-to-another"></a>

```text
var s = make([]int, 3)
n := copy(s, []int{0, 1, 2, 3}) // n == 3, s == []int{0, 1, 2}
```

#### Copy from a slice to itself <a id="copy-from-a-slice-to-itself"></a>

```text
s := []int{0, 1, 2}
n := copy(s, s[1:]) // n == 2, s == []int{1, 2, 2}
```

#### Copy from a string to a byte slice \(special case\) <a id="copy-from-a-string-to-a-byte-slice-special-nbsp-case"></a>

```text
var b = make([]byte, 5)
copy(b, "Hello, world!") // b == []byte("Hello")
```

## Default zero values for all Go types

Variables declared without an initial value are set to their [zero values](https://golang.org/ref/spec#The_zero_value):

* `0` for all **integer** types,
* `0.0` for **floating point** numbers,
* `false` for **booleans**,
* `""` for **strings**,
* `nil` for **interfaces**, **slices**, **channels**, **maps**, **pointers** and **functions**.

The elements of an **array** or **struct** will have its fields zeroed if no value is specified. This initialization is done recursively:

```text
type T struct {
    n int
    f float64
    next *T
}
fmt.Println([2]T{}) // [{0 0 <nil>} {0 0 <nil>}]
```

## Operators: complete list

### Arithmetic <a id="arithmetic"></a>

| Operator | Name | Types |
| :--- | :--- | :--- |
| `+` | sum | integers, floats, complex values, strings |
| `-` | difference | integers, floats, complex values |
| `*` | product |  |
| `/` | quotient |  |
| `%` | remainder | integers |
| `&` | bitwise AND |  |
| `|` | bitwise OR |  |
| `^` | bitwise XOR |  |
| `&^` | bit clear \(AND NOT\) |  |
| `<<` | left shift | integer &lt;&lt; unsigned integer |
| `>>` | right shift | integer &gt;&gt; unsigned integer |

See [Arithmetic operators](https://golang.org/ref/spec#Arithmetic_operators) in the Go language specification for complete definitions of the shift, quotient and remainder operators, integer overflow, and floating point behavior.

See [Bitwise operators cheat sheet](https://yourbasic.org/golang/bitwise-operator-cheat-sheet/) for more about how to manipulate bits with operators and functions in package math/bits \(bitcount, rotate, reverse, leading and trailing zeros\).

### Comparison <a id="comparison"></a>

Comparison operators compare two operands and yield an untyped boolean value.

| Operator | Name | Types |
| :--- | :--- | :--- |
| `==` | equal | comparable |
| `!=` | not equal |  |
| `<` | less | integers, floats, strings |
| `<=` | less or equal |  |
| `>` | greater |  |
| `>=` | greater or equal |  |

* Boolean, integer, floats, complex values and strings are comparable.
* Strings are ordered lexically byte-wise.
* Two pointers are equal if they point to the same variable or if both are nil.
* Two channel values are equal if they were created by the same call to make or if both are nil.
* Two interface values are equal if they have identical dynamic types and equal concrete values or if both are nil.
* A value `x` of non-interface type `X` and a value t of interface type `T` are equal if `t`’s dynamic type is identical to `X` and `t`’s concrete value is equal to `x`.
* Two struct values are equal if their corresponding non-blank fields are equal.
* Two array values are equal if their corresponding elements are equal.

### Logical <a id="logical"></a>

Logical operators apply to boolean values. The right operand is evaluated conditionally.

| Operator | Name | Description |
| :--- | :--- | :--- |
| `&&` | conditional AND | `p && q`   means "if p then q else false" |
| `||` | conditional OR | `p || q`   means "if p then true else q" |
| `!` | NOT | `!p`   means "not p" |

### Pointers and channels <a id="pointers-and-channels"></a>

| Operator | Name | Description |
| :--- | :--- | :--- |
| `&` | address of | `&x`   generates a pointer to `x` |
| `*` | pointer indirection | `*x`   denotes the variable pointed to by `x` |
| `<-` | receive | `<-ch`   is the value received from channel `ch` |

### Operator precedence <a id="operator-precedence"></a>

Go operator precedence spells MACAO

#### Unary operators <a id="unary-operators"></a>

Unary operators have the highest priority and bind the strongest.

#### Binary operators \(MACAO\) <a id="binary-operators-macao"></a>

| Priority | Operators | Note |
| :--- | :--- | :--- |
| 1 | `*`  `/`  `%`  `<<`  `>>`  `&`  `&^` | **M**ultiplicative |
| 2 | `+`  `-`  `|`  `^` | **A**dditive |
| 3 | `==`  `!=`  `<`  `<=`  `>`  `>=` | **C**omparison |
| 4 | `&&` | **A**nd |
| 5 | `||` | **O**r |

Binary operators of the same priority associate from **left to right**.

#### Statement operators <a id="statement-operators"></a>

The `++` and `- -` operators **form statements** and fall outside the operator hierarchy.

#### Examples <a id="examples"></a>

| Expression | Evaluation order |
| :--- | :--- |
| `x / y * z` | `(x / y) * z` |
| `*p++` | `(*p)++` |
| `^a >> b` | `(^a) >> b` |
| `1 + 2*a[i]` | `1 + (2*a[i])` |
| `m == n+1 && <-ch > 0` | `(m == (n+1)) && ((<-ch) > 0)` |

## Conversions \[complete list\]



### Basics <a id="basics"></a>

The expression `T(x)` converts the value `x` to the type `T`.

```text
x := 5.1
n := int(x) // convert float to int
```

The conversion rules are extensive but predictable:

* all conversions between typed expressions must be explicitly stated,
* illegal conversions are caught by the compiler.

Conversions to and from numbers and strings may **change the representation** and have a **run-time cost**. All other conversions only change the type but not the representation of `x`.

### Interfaces <a id="interfaces"></a>

> To “convert” an [interface](https://yourbasic.org/golang/interfaces-explained/) to a string, struct or map you should use a **type assertion** or a **type switch**. A type assertion doesn’t really convert an interface to another data type, but it provides access to an interface’s concrete value, which is typically what you want.[Type assertions and type switches](https://yourbasic.org/golang/type-assertion-switch/)

### Integers <a id="integers"></a>

* When converting to a shorter integer type, the value is **truncated** to fit in the result type’s size.
* When converting to a longer integer type,
  * if the value is a signed integer, it is [**sign extended**](https://en.wikipedia.org/wiki/Sign_extension);
  * otherwise it is **zero extended**.

```text
a := uint16(0x10fe) // 0001 0000 1111 1110
b := int8(a)        //           1111 1110 (truncated to -2)
c := uint16(b)      // 1111 1111 1111 1110 (sign extended to 0xfffe)
```

### Floats <a id="floats"></a>

* When converting a floating-point number to an integer, the **fraction is discarded** \(truncation towards zero\).
* When converting an integer or floating-point number to a floating-point type, the result value is **rounded to the precision** specified by the destination type.

```text
var x float64 = 1.9
n := int64(x) // 1
n = int64(-x) // -1

n = 1234567890
y := float32(n) // 1.234568e+09
```

> **Warning:** In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.[The Go Programming Language Specification: Conversions](https://golang.org/ref/spec#Conversions)

### Integer to string <a id="integer-to-string"></a>

* When converting an integer to a string, the value is interpreted as a Unicode code point, and the resulting string will contain the character represented by that code point, encoded in UTF-8.
* If the value does not represent a valid code point \(for instance if it’s negative\), the result will be `"\ufffd"`, the Unicode replacement character �.

```text
string(97) // "a"
string(-1) // "\ufffd" == "\xef\xbf\xbd"
```

> Use [`strconv.Itoa`](https://golang.org/pkg/strconv/#Itoa) to get the decimal string representation of an integer.
>
> ```text
> strconv.Itoa(97) // "97"
> ```

### Strings and byte slices <a id="strings-and-byte-slices"></a>

* Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.
* Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string.

```text
string([]byte{97, 230, 151, 165}) // "a日"
[]byte("a日")                     // []byte{97, 230, 151, 165}
```

### Strings and rune slices <a id="strings-and-rune-slices"></a>

* Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.
* Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.

```text
string([]rune{97, 26085}) // "a日"
[]rune("a日")             // []rune{97, 26085}
```

### Underlying type <a id="underlying-type"></a>

A non-constant value can be converted to type `T` if it has the same underlying type as `T`.

In this example, the underlying type of `int64`, `T1`, and `T2` is `int64`.

```text
type (
	T1 int64
	T2 T1
)
```

It’s idiomatic in Go to convert the type of an expression to access a specific method.

```text
var n int64 = 12345
fmt.Println(n)                // 12345
fmt.Println(time.Duration(n)) // 12.345µs
```

\(The underlying type of [`time.Duration`](https://golang.org/pkg/time/#Duration) is `int64`, and the `time.Duration` type has a [`String`](https://golang.org/pkg/time/#Duration.String) method that returns the duration formatted as a time.\)

### Implicit conversions <a id="implicit-conversions"></a>

The only implicit conversion in Go is when an untyped constant is used in a situation where a type is required.

In this example the untyped literals `1` and `2` are implicitly converted.

```text
var x float64
x = 1 // Same as x = float64(1)

t := 2 * time.Second // Same as t := time.Duration(2) * time.Second
```

\(The implicit conversions are necessary since there is no mixing of numeric types in Go. You can only multiply a `time.Duration` with another `time.Duration`.\)

When the type can’t be inferred from the context, an untyped constant is converted to a `bool`, `int`, `float64`, `complex128`, `string` or `rune` depending on the syntactical format of the constant.

```text
n := 1   // Same as n := int(1)
x := 1.0 // Same as x := float64(1.0)
s := "A" // Same as s := string("A")
c := 'A' // Same as c := rune('A')
```

Illegal implicit conversions are caught by the compiler.

```text
var b byte = 256 // Same as var b byte = byte(256)
```

```text
../main.go:2:6: constant 256 overflows byte
```

### Pointers <a id="pointers"></a>

The Go compiler does not allow conversions between pointers and integers. Package [`unsafe`](https://golang.org/pkg/unsafe/) implements this functionality under restricted circumstances.

> **Warning:** The built-in package unsafe, known to the compiler, provides facilities for low-level programming including operations that violate the type system. A package using unsafe must be vetted manually for type safety and may not be portable.[The Go Programming Language Specification: Package unsafe](https://golang.org/ref/spec#Package_unsafe)

## Methods and Interfaces

## Methods explained

You can define methods on any type declared in a type definition.

* A method is a function with an extra **receiver** argument.
* The receiver sits between the `func` keyword and the method name.

In this example, the `HasGarage` method is associated with the `House` type. The method receiver is called `p`.

```text
type House struct {
    garage bool
}

func (p *House) HasGarage() bool { return p.garage }

func main() {
    house := new(House)
    fmt.Println(house.HasGarage()) // Prints "false" (zero value)
}
```

#### Conversions and methods <a id="conversions-and-methods"></a>

If you [convert](https://yourbasic.org/golang/conversions/) a value to a different type, the new value will have the methods of the new type, but not the old.

```text
type MyInt int

func (m MyInt) Positive() bool { return m > 0 }

func main() {
    var m MyInt = 2
    m = m * m // The operators of the underlying type still apply.

    fmt.Println(m.Positive())        // Prints "true"
    fmt.Println(MyInt(3).Positive()) // Prints "true"

    var n int
    n = int(m) // The conversion is required.
    n = m      // ILLEGAL
}
```

```text
../main.go:14:4: cannot use m (type MyInt) as type int in assignment
```

It’s idiomatic in Go to convert the type of an expression to access a specific method.

```text
var n int64 = 12345
fmt.Println(n)                // 12345
fmt.Println(time.Duration(n)) // 12.345µs
```

\(The underlying type of [`time.Duration`](https://golang.org/pkg/time/#Duration) is `int64`, and the `time.Duration` type has a [`String`](https://golang.org/pkg/time/#Duration.String) method that returns the duration formatted as a time.\)

## Type, value and equality of interfaces

### Interface type <a id="interface-type"></a>

> An interface type consists of a set of method signatures. A variable of interface type can hold any value that implements these methods.

In this example both `Temp` and `*Point` implement the `MyStringer` interface.

```text
type MyStringer interface {
	String() string
}
```

```text
type Temp int

func (t Temp) String() string {
	return strconv.Itoa(int(t)) + " °C"
}

type Point struct {
	x, y int
}

func (p *Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}
```

Actually, `*Temp` also implements `MyStringer`, since the method set of a pointer type `*T` is the set of all methods with receiver `*T` or `T`.

When you call a method on an interface value, the method of its underlying type is executed.

```text
var x MyStringer

x = Temp(24)
fmt.Println(x.String()) // 24 °C

x = &Point{1, 2}
fmt.Println(x.String()) // (1,2)
```

### Structural typing <a id="structural-typing"></a>

> A type implements an interface by implementing its methods. No explicit declaration is required.

In fact, the `Temp`, `*Temp` and `*Point` types also implement the standard library [`fmt.Stringer`](https://golang.org/pkg/fmt/#Stringer) interface. The `String` method in this interface is used to print values passed as an operand to functions such as [`fmt.Println`](https://golang.org/pkg/fmt/#Println).

```text
var x MyStringer

x = Temp(24)
fmt.Println(x) // 24 °C

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

### The empty interface <a id="the-empty-interface"></a>



The interface type that specifies no methods is known as the empty interface.

```text
interface{}
```

An empty interface can hold values of any type since every type implements at least zero methods.

```text
var x interface{}

x = 2.4
fmt.Println(x) // 2.4

x = &Point{1, 2}
fmt.Println(x) // (1,2)
```

The [`fmt.Println`](https://golang.org/pkg/fmt/#Println) function is a chief example. It takes any number of arguments of any type.

```text
func Println(a ...interface{}) (n int, err error)
```

### Interface values <a id="interface-values"></a>

> An **interface value** consists of a **concrete value** and a **dynamic type**: `[Value, Type]`

In a call to [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf), you can use `%v` to print the concrete value and `%T` to print the dynamic type.

```text
var x MyStringer
fmt.Printf("%v %T\n", x, x) // <nil> <nil>

x = Temp(24)
fmt.Printf("%v %T\n", x, x) // 24 °C main.Temp

x = &Point{1, 2}
fmt.Printf("%v %T\n", x, x) // (1,2) *main.Point

x = (*Point)(nil)
fmt.Printf("%v %T\n", x, x) // <nil> *main.Point
```

The **zero value** of an interface type is nil, which is represented as `[nil, nil]`.

Calling a method on a nil interface is a run-time error. However, it’s quite common to write methods that can handle a receiver value `[nil, Type]`, where `Type` isn’t nil.

You can use [**type assertions**](https://yourbasic.org/golang/type-assertion-switch/) or [**type switches**](https://yourbasic.org/golang/type-assertion-switch/) to access the dynamic type of an interface value. See [Find the type of an object](https://yourbasic.org/golang/find-type-of-object/) for more details.

### Equality <a id="equality"></a>

Two interface values are equal

* if they have equal concrete values **and** identical dynamic types,
* or if both are nil.

A value `t` of interface type `T` and a value `x` of non-interface type `X` are equal if

* `t`’s concrete value is equal to `x`
* **and** `t`’s dynamic type is identical to `X`.

```text
var x MyStringer
fmt.Println(x == nil) // true

x = (*Point)(nil)
fmt.Println(x == nil) // false
```

In the second print statement, the concrete value of `x` equals `nil`, but its dynamic type is `*Point`, which is not `nil`.

> **Warning:** See [Nil is not nil](https://yourbasic.org/golang/gotcha-why-nil-error-not-equal-nil/) for a real-world example where this definition of equality leads to puzzling results.

#### Further reading <a id="further-reading"></a>

[Generics \(alternatives and workarounds\)](https://yourbasic.org/golang/generics/) discusses how interfaces, multiple functions, type assertions, reflection and code generation can be use in place of parametric polymorphism in Go

## Named return values \[best practice\]

In Go return parameters may be named and used as regular variables. When the function returns, they are used as return values.

```text
func f() (i int, s string) {
    i = 17
    s = "abc"
    return // same as return i, s
}
```

Named return parameters are initialized to their [zero values](https://yourbasic.org/golang/default-zero-value/).

The names are not mandatory but can make for good documentation. Correctly used, named return parameters can also help clarify and clean up the code.

### Example <a id="example"></a>

This version of [`io.ReadFull`](https://golang.org/pkg/io/#ReadFull), taken from [Effective Go](https://golang.org/doc/effective_go.html#named-results), uses them to good effect.

```text
// ReadFull reads exactly len(buf) bytes from r into buf. It returns
// the number of bytes copied and an error if fewer bytes were read.
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

The code is both simpler and clearer, with named return values that are properly initialized and tied to a plain return.

## Optional parameters, default parameter values and method overloading

By design, Go does **not** support

* **optional parameters**,
* **default parameter values**,
* or **method overloading**.

> Method dispatch is simplified if it doesn’t need to do type matching as well. Experience with other languages told us that having a variety of methods with the same name but different signatures was occasionally useful but that it could also be confusing and fragile in practice. Matching only by name and requiring consistency in the types was a major simplifying decision in Go’s type system.[Go FAQ: Why does Go not support overloading of methods and operators?](https://golang.org/doc/faq#overloading)

However, there are

* [**variadic functions**](https://yourbasic.org/golang/variadic-function/) \(functions that accept a variable number of arguments\),
* and **dynamic method dispatch** is supported through [interfaces](https://yourbasic.org/golang/interfaces-explained/).

For more on this, see [Interfaces in 5 easy steps](https://yourbasic.org/golang/interfaces-explained/).

The idiomatic way to emulate optional parameters and method overloading in Go is to write several methods with different names. For example, the [`sort`](https://golang.org/pkg/sort/) package has five different functions for sorting a slice:

* the generic [`sort.Slice`](https://golang.org/pkg/sort/#Slice) and [`sort.SliceStable`](https://golang.org/pkg/sort/#SliceStable),
* and the three more specific [`sort.Float64s`](https://golang.org/pkg/sort/#Float64s), [`sort.Ints`](https://golang.org/pkg/sort/#Ints), and [`sort.Strings`](https://golang.org/pkg/sort/#Strings).

## Variadic functions \(...T\)

### Basics <a id="basics"></a>

If the **last parameter** of a function has type `...T` it can be called with **any number** of trailing arguments of type `T`.

```text
func Sum(nums ...int) int {
    res := 0
    for _, n := range nums {
        res += n
    }
    return res
}

func main()
    fmt.Println(Sum())        // 0
    fmt.Println(Sum(1, 2, 3)) // 6
}
```

The actual type of `...T` inside the function is `[]T`.

### Pass slice elements to a variadic function <a id="pass-slice-elements-to-a-variadic-function"></a>

You can pass the elements of a slice `s` directly to a variadic function using the `s...` notation. In this case no new slice is created.

```text
primes := []int{2, 3, 5, 7}
fmt.Println(Sum(primes...)) // 17
```

### Append is variadic <a id="append-is-variadic"></a>

The built-in [append function](https://yourbasic.org/golang/append-explained/) is variadic and can be used to append any number of elements to a slice.

As a special case, you can append a string to a byte slice:

```text
var buf []byte
buf = append(buf, 'a', 'b')
buf = append(buf, "cd"...)
fmt.Println(buf) // [97 98 99 100]
```

## Error Handling

## Error handling best practice

Go has two different error-handling mechanisms:

* most functions return [**errors**](https://yourbasic.org/golang/create-error/);
* only a truly unrecoverable condition, such as an out-of-range index, produces a run-time exception, known as a [**panic**](https://yourbasic.org/golang/recover-from-panic/).

Go’s multivalued return makes it easy to return a detailed error message alongside the normal return value. By convention, such messages have type `error`, a simple built-in [interface](https://yourbasic.org/golang/interfaces-explained/):

```text
type error interface {
    Error() string
}
```

#### Error handling example <a id="error-handling-example"></a>

The `os.Open` function returns a non-nil `error` value when it fails to open a file.

```text
func Open(name string) (file *File, err error)
```

The following code uses `os.Open` to open a file. If an `error` occurs it calls `log.Fatal` to print the error message and stop.

```text
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

### Custom errors <a id="custom-errors"></a>

To create a simple string-only `error` you can use [`errors.New`](https://golang.org/pkg/errors/#New):

```text
err := errors.New("Houston, we have a problem")
```

The `error` interface requires only an `Error` method, but specific `error` implementations often have additional methods, allowing callers to inspect the details of the error.

#### Learn more <a id="learn-more"></a>

See [3 simple ways to create an error](https://yourbasic.org/golang/create-error/) for more examples.

### Panic <a id="panic"></a>

Panics are similar to C++ and Java exceptions, but are only intended for run-time errors, such as following a nil pointer or attempting to index an array out of bounds.

#### Learn more <a id="learn-more-1"></a>

See [Recover from a panic](https://yourbasic.org/golang/recover-from-panic/) for a tutorial on how to recover from and test panics.

## 3 simple ways to create an error

### String-based errors <a id="string-based-errors"></a>

The standard library offers two out-of-the-box options.

```text
// simple string-based error
err1 := errors.New("math: square root of negative number")

// with formatting
err2 := fmt.Errorf("math: square root of negative number %g", x)
```

### Custom errors with data <a id="custom-errors-with-data"></a>

To define a custom error type, you must satisfy the predeclared `error` [interface](https://yourbasic.org/golang/interfaces-explained/).

```text
type error interface {
    Error() string
}
```

Here are two examples.

```text
type SyntaxError struct {
    Line int
    Col  int
}

func (e *SyntaxError) Error() string {
    return fmt.Sprintf("%d:%d: syntax error", e.Line, e.Col)
}
```

```text
type InternalError struct {
    Path string
}

func (e *InternalError) Error() string {
    return fmt.Sprintf("parse %v: internal error", e.Path)
}
```

If `Foo` is a function that can return a `SyntaxError` or an `InternalError`, you may handle the two cases like this.

```text
if err := Foo(); err != nil {
    switch e := err.(type) {
    case *SyntaxError:
        // Do something interesting with e.Line and e.Col.
    case *InternalError:
        // Abort and file an issue.
    default:
        log.Println(e)
    }
}
```

#### More code examples

[Go blueprints: code for com­mon tasks](https://yourbasic.org/golang/blueprint/) is a collection of handy code examples.

## Panics, stack traces and how to recover \[best practice\]

### A panic is an exception in Go <a id="a-panic-is-an-exception-in-go"></a>

Panics are similar to C++ and Java exceptions, but are only intended for run-time errors, such as following a nil pointer or attempting to index an array out of bounds. To signify events such as end-of-file, Go programs use the built-in `error` type. See [Error handling best practice](https://yourbasic.org/golang/errors-explained/) and [3 simple ways to create an error](https://yourbasic.org/golang/create-error/) for more on errors.

A panic stops the normal execution of a goroutine:

* When a program panics, it immediately starts to unwind the call stack.
* This continues until the program crashes and prints a stack trace,
* or until the built-in `recover` function is called.

A panic is caused either by a runtime error, or an explicit call to the built-in `panic` function.

### Stack traces <a id="stack-traces"></a>

A **stack trace** – a report of all active stack frames – is typically printed to the console when a panic occurs. Stack traces can be very useful for debugging:

* not only do you see **where** the error happened,
* but also **how** the program arrived in this place.

#### Interpret a stack trace <a id="interpret-a-stack-trace"></a>

Here’s an example of a stack trace:

```text
goroutine 11 [running]:
testing.tRunner.func1(0xc420092690)
    /usr/local/go/src/testing/testing.go:711 +0x2d2
panic(0x53f820, 0x594da0)
    /usr/local/go/src/runtime/panic.go:491 +0x283
github.com/yourbasic/bit.(*Set).Max(0xc42000a940, 0x0)
    ../src/github.com/bit/set_math_bits.go:137 +0x89
github.com/yourbasic/bit.TestMax(0xc420092690)
    ../src/github.com/bit/set_test.go:165 +0x337
testing.tRunner(0xc420092690, 0x57f5e8)
    /usr/local/go/src/testing/testing.go:746 +0xd0
created by testing.(*T).Run
    /usr/local/go/src/testing/testing.go:789 +0x2de
```

It can be read from the bottom up:

* `testing.(*T).Run` has called `testing.tRunner`,
* which has called `bit.TestMax`,
* which has called `bit.(*Set).Max`,
* which has called `panic`,
* which has called `testing.tRunner.func1`.

The indented lines show the source file and line number at which the function was called. The hexadecimal numbers refer to parameter values, including values of pointers and internal data structures. [Stack Traces in Go](https://www.goinggo.net/2015/01/stack-traces-in-go.html) has more details.

#### Print and log a stack trace <a id="print-and-log-a-stack-trace"></a>

To print the stack trace for the current goroutine, use [`debug.PrintStack`](https://golang.org/pkg/runtime/debug/#PrintStack) from package [`runtime/debug`](https://golang.org/pkg/runtime/debug/).

You can also examine the current stack trace programmatically by calling [`runtime.Stack`](https://golang.org/pkg/runtime/#Stack).

#### Level of detail <a id="level-of-detail"></a>

The [`GOTRACEBACK`](https://golang.org/pkg/runtime/#hdr-Environment_Variables) variable controls the amount of output generated when a Go program fails.

* `GOTRACEBACK=none` omits the goroutine stack traces entirely.
* `GOTRACEBACK=single` \(the default\) prints a stack trace for the current goroutine, eliding functions internal to the run-time system. The failure prints stack traces for all goroutines if there is no current goroutine or the failure is internal to the run-time.
* `GOTRACEBACK=all` adds stack traces for all user-created goroutines.
* `GOTRACEBACK=system` is like `all` but adds stack frames for run-time functions and shows goroutines created internally by the run-time.

### Recover and catch a panic <a id="recover-and-catch-a-panic"></a>

The built-in `recover` function can be used to regain control of a panicking goroutine and resume normal execution.

* A call to `recover` stops the unwinding and returns the argument passed to `panic`.
* If the goroutine is not panicking, `recover` returns `nil`.

Because the only code that runs while unwinding is inside [deferred functions](https://yourbasic.org/golang/defer/), `recover` is only useful inside such functions.

#### Panic handler example <a id="panic-handler-example"></a>

```text
func main() {
	n := foo()
	fmt.Println("main received", n)
}

func foo() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	m := 1
	panic("foo: fail")
	m = 2
	return m
}
```

```text
foo: fail
main received 0
```

Since the panic occurred before `foo` returned a value, `n` still has its initial zero value.

#### Return a value <a id="return-a-value"></a>

To return a value during a panic, you must use a [named return value](https://yourbasic.org/golang/named-return-values-parameters/).

```text
func main() {
	n := foo()
	fmt.Println("main received", n)
}

func foo() (m int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			m = 2
		}
	}()
	m = 1
	panic("foo: fail")
	m = 3
	return m
}
```

```text
foo: fail
main received 2
```

### Test a panic \(utility function\) <a id="test-a-panic-utility-function"></a>

In this example, we use reflection to check if a list of interface variables have types corre­sponding to the para­meters of a given function. If so, we call the function with those para­meters to check if there is a panic.

```text
// Panics tells if function f panics with parameters p.
func Panics(f interface{}, p ...interface{}) bool {
	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)
	if ft.NumIn() != len(p) {
		panic("wrong argument count")
	}
	pv := make([]reflect.Value, len(p))
	for i, v := range p {
		if reflect.TypeOf(v) != ft.In(i) {
			panic("wrong argument type")
		}
		pv[i] = reflect.ValueOf(v)
	}
	return call(fv, pv)
}

func call(fv reflect.Value, pv []reflect.Value) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = true
		}
	}()
	fv.Call(pv)
	return
}
```

## Bits and Pieces



## Blank identifier \(underscore\)

The blank identifier `_` is an anonymous placeholder. It may be used like any other identifier in a declaration, but it does not introduce a binding.

### Ignore values <a id="ignore-values"></a>

The blank identifier provides a way to ignore left-hand side values in an assignment.

```text
_, present := timeZone["CET"]

sum := 0
for _, n := range a {
	sum += n
}
```

### Import for side effects <a id="import-for-side-effects"></a>

It can also be used to import a package solely for its side effects.

```text
import _ "image/png" // init png decoder function
```

### Silence the compiler <a id="silence-the-compiler"></a>

It can be used to during development to avoid compiler errors about unused imports and variables in a half-written program.

```text
package main

import (
    "fmt"
    "log"
    "os"
)

var _ = fmt.Printf // DEBUG: delete when done

func main() {
    f, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    _ = f // TODO: read file
}
```

For an automatic solution, use the [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) tool, which rewrites a Go source file to have the correct imports. Many Go editors and IDEs run this tool automatically.

## Find the type of an object

### Use fmt for a string type description <a id="use-fmt-for-a-string-type-description"></a>

You can use the `%T` flag in the [fmt](https://golang.org/pkg/fmt/) package to get a Go-syntax representation of the type.

```text
var x interface{} = []int{1, 2, 3}
xType := fmt.Sprintf("%T", x)
fmt.Println(xType) // "[]int"
```

\(The [empty interface](https://yourbasic.org/golang/interfaces-explained/#the-empty-interface) denoted by `interface{}` can hold values of any type.\)

### A type switch lets you choose between types <a id="a-type-switch-lets-you-choose-between-types"></a>

Use a [type switch](https://yourbasic.org/golang/type-assertion-switch/) to do several [type assertions](https://yourbasic.org/golang/type-assertion-switch/) in series.

```text
var x interface{} = 2.3
switch v := x.(type) {
case int:
    fmt.Println("int:", v)
case float64:
    fmt.Println("float64:", v)
default:
    fmt.Println("unknown")
}
// Output: float64: 2.3
```

### Reflection gives full type information <a id="reflection-gives-full-type-information"></a>

Use the [reflect](https://golang.org/pkg/reflect/) package if the options above don’t suffice.

```text
var x interface{} = []int{1, 2, 3}
xType := reflect.TypeOf(x)
xValue := reflect.ValueOf(x)
fmt.Println(xType, xValue) // "[]int [1 2 3]"
```

## Generics \(alternatives and workarounds\)

Go has some built-in generic data types, such as slices and maps, and some generic functions, such as `append` and `copy`. However, there is no mechanism for writing your own.

Here are some techniques that can be used in place of parametric polymorphism in Go.

### Find a well-fitting interface <a id="find-a-well-fitting-interface"></a>

Describe the generic behaviour of your data with an [interface](https://yourbasic.org/golang/interfaces-explained/).

The [`io.Reader`](https://golang.org/pkg/io/#Reader) interface, which represents the read end of a stream of data, is a good example:

* many functions take an [`io.Reader`](https://golang.org/pkg/io/#Reader) as input,
* and many data types, including files, network connections, and ciphers, implement this interface.

### Use multiple functions <a id="use-multiple-functions"></a>

If you only need to support a few data types, consider offering a separate function for each type.

As an example, the two packages [`strings`](https://golang.org/pkg/strings/) and [`bytes`](https://golang.org/pkg/bytes/) come with pretty much the same set of functions.

If this leads to an unmanageable amount of copy and paste, consider using a code generation tool.

### Use the empty interface <a id="use-the-empty-interface"></a>

If little is known about the data, consider using the empty interface `interface{}` in combination with type assertions, and possibly also reflection. Libraries such as [`fmt`](https://golang.org/pkg/fmt/) and [`encoding/json`](https://golang.org/pkg/encoding/json/) couldn’t have been written in any other way.

### Write an experience report <a id="write-an-experience-report"></a>

If none of these solutions are effective, consider submitting an experience report:

> This page collects experience reports about problems with Go that might inform our design of solutions to those problems. These reports should focus on the problems: they should not focus on and need not propose solutions.
>
> We hope to use these experience reports to understand where people are having trouble writing Go, to help us prioritize future changes to the Go ecosystem.
>
> [The Go Wiki: Experience Reports](https://github.com/golang/go/wiki/ExperienceReports)

## Pick the right one: int vs. int64

### Use int for indexing <a id="use-int-for-indexing"></a>

An **index**, **length** or **capacity** should normally be an `int`. The `int` type is either 32 or 64 bits, and always big enough to hold the maximum possible length of an array.

See [Maximum value of an int](https://yourbasic.org/golang/max-min-int-uint/) for code to compute the maximum value of an `int`.

### Use int64 and friends for data <a id="use-int64-and-friends-for-data"></a>

The types `int8`, `int16`, `int32`, and `int64` \(and their unsigned counterparts\) are best suited for **data**. An `int64` is the typical choice when memory isn’t an issue.

In particular, you can use a **`byte`**, which is an alias for `uint8`, to be extra clear about your intent. Similarly, you can use a **`rune`**, which is an alias for `int32`, to emphasize than an integer represents a [code point](https://yourbasic.org/golang/rune/).

> Sometimes it makes little or no difference if you use 32 or 64 bits for data, and then it’s quite common to simply use an int. However, I prefer to be explicit. It forces you to spend a moment thinking about the choice and also makes the code a bit clearer.

### Examples <a id="examples"></a>

In this code, the slice elements and the `max` variable have type `int64`, while the index and the length of the slice have type `int`.

```text
func Max(a []int64) int64 {
	max := a[0]
	for i := 1; i < len(a); i++ {
		if max < a[i] {
			max = a[i]
		}
	}
	return max
}
```

The implementation of [`time.Duration`](https://golang.org/pkg/time/#Duration) is a typical example from the standard library where an `int64` is used to store data:

```text
type Duration int64
```

A `Duration` represents the time between two instants as a nanosecond count. This limits the largest possible duration to about 290 years.

#### Further reading <a id="further-reading"></a>

The [Maximum value of an int](https://yourbasic.org/golang/max-min-int-uint/) article shows how to compute the size and limit values of an `int` as untyped constants.

## 3 dots in 4 places

### Variadic function parameters <a id="variadic-function-parameters"></a>

If the **last parameter** of a function has type `...T`, it can be called with any number of trailing arguments of type `T`. The actual type of `...T` inside the function is `[]T`.

This example function can be called with, for instance, `Sum(1, 2, 3)` or `Sum()`.

```text
func Sum(nums ...int) int {
    res := 0
    for _, n := range nums {
        res += n
    }
    return res
}
```

### Arguments to variadic functions <a id="arguments-to-variadic-functions"></a>

You can pass a slice `s` directly to a variadic function if you unpack it with the `s...` notation. In this case no new slice is created.

In this example, we pass a slice to the `Sum` function.

```text
primes := []int{2, 3, 5, 7}
fmt.Println(Sum(primes...)) // 17
```

### Array literals <a id="array-literals"></a>

In an array literal, the `...` notation specifies a length equal to the number of elements in the literal.

```text
stooges := [...]string{"Moe", "Larry", "Curly"} // len(stooges) == 3
```

### The go command <a id="the-go-command"></a>

Three dots are used by the [`go`](https://golang.org/cmd/go/) command as a wildcard when describing package lists.

This command tests all packages in the current directory and its subdirectories.

```text
$ go test ./...
```

## Redeclaring variables

You can’t redeclare a variable which has already been declared in the same block.

```text
func main() {
	m := 0
	m := 1
	fmt.Println(m)
}
```

```text
../main.go:3:4: no new variables on left side of :=
```

However, variables can be redeclared in short multi-variable declarations where at least one new variable is introduced.

```text
func main() {
	m := 0
	m, n := 1, 2
	fmt.Println(m, n)
}
```

> Unlike regular variable declarations, a short variable declaration may redeclare variables provided they were originally declared earlier in the same block \(or the parameter lists if the block is the function body\) with the same type, and at least one of the non-blank variables is new. \[…\] Redeclaration does not introduce a new variable; it just assigns a new value to the original.
>
> – [_The Go Programming Language Specification: Short variable declarations_](https://golang.org/ref/spec#Short_variable_declarations)

## Standard Library



## How to use JSON with Go \[best practices\]

### Default types <a id="default-types"></a>

The default Go types for decoding and encoding JSON are

* `bool` for JSON booleans,
* `float64` for JSON numbers,
* `string` for JSON strings, and
* `nil` for JSON null.

Additionally, [`time.Time`](https://golang.org/pkg/time/#Time) and the numeric types in the [`math/big`](https://golang.org/pkg/math/big/) package can be automatically encoded as JSON strings.

Note that JSON doesn’t support basic integer types. They can often be approximated by floating-point numbers.

> Since software that implements IEEE 754-2008 binary64 \(double precision\) numbers is generally available and widely used, good interoperability can be achieved by implementations that expect no more precision or range than these provide \[…\]
>
> Note that when such software is used, numbers that are integers and are in the range \[-253 + 1, 253 - 1\] are interoperable in the sense that implementations will agree exactly on their numeric values.[RFC 7159: The JSON Data Interchange Format](https://tools.ietf.org/html/rfc7159#section-6)

### Encode \(marshal\) struct to JSON <a id="encode-marshal-struct-to-json"></a>

The [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) function in package [`encoding/json`](https://golang.org/pkg/encoding/json/) generates JSON data.

```text
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

basket := FruitBasket{
    Name:    "Standard",
    Fruit:   []string{"Apple", "Banana", "Orange"},
    Id:      999,
    private: "Second-rate",
    Created: time.Now(),
}

var jsonData []byte
jsonData, err := json.Marshal(basket)
if err != nil {
    log.Println(err)
}
fmt.Println(string(jsonData))
```

Output:

```text
{"Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999,"Created":"2018-04-09T23:00:00Z"}
```

Only data that can be represented as JSON will be encoded; see [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) for the complete rules.

* Only the exported \(public\) fields of a struct will be present in the JSON output. **Other fields are ignored**.
* A field with a `json:` **struct tag** is stored with its tag name instead of its variable name.
* Pointers will be encoded as the values they point to, or `null` if the pointer is `nil`.

### Pretty print <a id="pretty-print"></a>

Replace `json.Marshal` with [`json.MarshalIndent`](https://golang.org/pkg/encoding/json/#MarshalIndent) in the example above to indent the JSON output.

```text
jsonData, err := json.MarshalIndent(basket, "", "    ")
```

Output:

```text
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}
```

### Decode \(unmarshal\) JSON to struct <a id="decode-unmarshal-json-to-struct"></a>

The [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) function in package [`encoding/json`](https://golang.org/pkg/encoding/json/) parses JSON data.

```text
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64 `json:"ref"`
    Created time.Time
}

jsonData := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

var basket FruitBasket
err := json.Unmarshal(jsonData, &basket)
if err != nil {
    log.Println(err)
}
fmt.Println(basket.Name, basket.Fruit, basket.Id)
fmt.Println(basket.Created)
```

Output:

```text
Standard [Apple Banana Orange] 999
2018-04-09 23:00:00 +0000 UTC
```

Note that `Unmarshal` allocated a new slice all by itself. This is how unmarshaling works for slices, maps and pointers.

For a given JSON key `Foo`, `Unmarshal` will attempt to match the struct fields in this order:

1. an exported \(public\) field with a struct tag `json:"Foo"`,
2. an exported field named `Foo`, or
3. an exported field named `FOO`, `FoO`, or some other case-insensitive match.

Only fields thar are found in the destination type will be decoded:

* This is useful when you wish to pick only a few specific fields.
* In particular, any unexported fields in the destination struct will be unaffected.

### Arbitrary objects and arrays <a id="arbitrary-objects-and-arrays"></a>

The [encoding/json](https://golang.org/pkg/encoding/json/) package uses

* `map[string]interface{}` to store arbitrary JSON objects, and
* `[]interface{}` to store arbitrary JSON arrays.

It will unmarshal any valid JSON data into a plain `interface{}` value.

Consider this JSON data:

```text
{
    "Name": "Eve",
    "Age": 6,
    "Parents": [
        "Alice",
        "Bob"
    ]
}
```

The [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) function will parse it into a map whose keys are strings, and whose values are themselves stored as empty interface values:

```text
map[string]interface{}{
    "Name": "Eve",
    "Age":  6.0,
    "Parents": []interface{}{
        "Alice",
        "Bob",
    },
}
```

We can iterate through the map with a range statement and use a type switch to access its values.

```text
jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

var v interface{}
json.Unmarshal(jsonData, &v)
data := v.(map[string]interface{})

for k, v := range data {
    switch v := v.(type) {
    case string:
        fmt.Println(k, v, "(string)")
    case float64:
        fmt.Println(k, v, "(float64)")
    case []interface{}:
        fmt.Println(k, "(array):")
        for i, u := range v {
            fmt.Println("    ", i, u)
        }
    default:
        fmt.Println(k, v, "(unknown)")
    }
}
```

Output:

```text
Name Eve (string)
Age 6 (float64)
Parents (array):
     0 Alice
     1 Bob
```

### JSON file example <a id="json-file-example"></a>

The [`json.Decoder`](https://golang.org/pkg/encoding/json/#Decoder) and [`json.Encoder`](https://golang.org/pkg/encoding/json/#Encoder) types in package [`encoding/json`](https://golang.org/pkg/encoding/json/) offer support for reading and writing streams, e.g. files, of JSON data.

The code in this example

* reads a stream of JSON objects from a [Reader](https://yourbasic.org/golang/io-reader-interface-explained/) \([`strings.Reader`](https://golang.org/pkg/strings/#Reader)\),
* removes the `Age` field from each object,
* and then writes the objects to a [Writer](https://yourbasic.org/golang/io-writer-interface-explained/) \([`os.Stdout`](https://golang.org/pkg/os/#pkg-variables)\).

```text
const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
`
reader := strings.NewReader(jsonData)
writer := os.Stdout

dec := json.NewDecoder(reader)
enc := json.NewEncoder(writer)

for {
    // Read one JSON object and store it in a map.
    var m map[string]interface{}
    if err := dec.Decode(&m); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }

    // Remove all key-value pairs with key == "Age" from the map.
    for k := range m {
        if k == "Age" {
            delete(m, k)
        }
    }

    // Write the map as a JSON object.
    if err := enc.Encode(&m); err != nil {
        log.Println(err)
    }
}
```

Output:

```text
{"Name":"Alice"}
{"Name":"Bob"}
```

## fmt.Printf formatting tutorial and cheat sheet

### Basics <a id="basics"></a>

With the Go [`fmt`](https://golang.org/pkg/fmt) package you can format numbers and strings padded with spaces or zeroes, in different bases, and with optional quotes.

You submit a **template string** that contains the text you want to format plus some **annotation verbs** that tell the `fmt` functions how to format the trailing arguments.

#### Printf <a id="printf"></a>

In this example, [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf) formats and writes to standard output:

```text
fmt.Printf("Binary: %b\\%b", 4, 5) // Prints `Binary: 100\101`
```

* the **template string** is `"Binary: %b\\%b"`,
* the **annotation verb** `%b` formats a number in binary, and
* the **special value** `\\` is a backslash.

As a special case, the verb `%%`, which consumes no argument, produces a percent sign:

```text
fmt.Printf("%d %%", 50) // Prints `50 %`
```

#### Sprintf \(format without printing\) <a id="sprintf-format-without-printing"></a>

Use [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) to format a string without printing it:

```text
s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
```

#### Find fmt errors with vet <a id="find-fmt-errors-with-vet"></a>

If you try to compile and run this incorrect line of code

```text
fmt.Printf("Binary: %b\\%b", 4) // An argument to Printf is missing.
```

you’ll find that the program will compile, and then print

```text
Binary: 100\%!b(MISSING)
```

To catch this type of errors early, you can use the [vet command](https://golang.org/cmd/vet/) – it can find calls whose arguments do not align with the format string.

```text
$ go vet example.go
example.go:8: missing argument for Printf("%b"): format reads arg 2, have only 1 args
```

### Cheat sheet <a id="cheat-sheet"></a>

#### Default formats and type <a id="default"></a>

* **Value:** `[]int64{0, 1}`

| Format | Verb | Description |
| :--- | :--- | :--- |
| \[0 1\] | `%v` | Default format |
| \[\]int64{0, 1} | `%#v` | Go-syntax format |
| \[\]int64 | `%T` | The type of the value |

#### Integer \(indent, base, sign\) <a id="integer"></a>

* **Value:** `15`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 15 | `%d` | Base 10 |
| +15 | `%+d` | Always show sign |
| ␣␣15 | `%4d` | Pad with spaces \(width 4, right justified\) |
| 15␣␣ | `%-4d` | Pad with spaces \(width 4, left justified\) |
| 0015 | `%04d` | Pad with zeroes \(width 4\) |
| 1111 | `%b` | Base 2 |
| 17 | `%o` | Base 8 |
| f | `%x` | Base 16, lowercase |
| F | `%X` | Base 16, uppercase |
| 0xf | `%#x` | Base 16, with leading 0x |

#### Character \(quoted, Unicode\) <a id="character"></a>

* **Value:** `65`   \(Unicode letter A\)

| Format | Verb | Description |
| :--- | :--- | :--- |
| A | `%c` | Character |
| 'A' | `%q` | Quoted character |
| U+0041 | `%U` | Unicode |
| U+0041 'A' | `%#U` | Unicode with character |

#### Boolean \(true/false\) <a id="boolean"></a>

Use `%t` to format a boolean as `true` or `false`.

#### Pointer \(hex\) <a id="pointer"></a>

Use `%p` to format a pointer in base 16 notation with leading `0x`.

#### Float \(indent, precision, scientific notation\) <a id="float"></a>

* **Value:** `123.456`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 1.234560e+02 | `%e` | Scientific notation |
| 123.456000 | `%f` | Decimal point, no exponent |
| 123.46 | `%.2f` | Default width, precision 2 |
| ␣␣123.46 | `%8.2f` | Width 8, precision 2 |
| 123.456 | `%g` | Exponent as needed, necessary digits only |

#### String or byte slice \(quote, indent, hex\) <a id="string-or-byte-slice"></a>

* **Value:** `"café"`

| Format | Verb | Description |
| :--- | :--- | :--- |
| café | `%s` | Plain string |
| ␣␣café | `%6s` | Width 6, right justify |
| café␣␣ | `%-6s` | Width 6, left justify |
| "café" | `%q` | Quoted string |
| 636166c3a9 | `%x` | Hex dump of byte values |
| 63 61 66 c3 a9 | `% x` | Hex dump with spaces |

#### Special values <a id="special-values"></a>

| Value | Description |
| :--- | :--- |
| `\a` | U+0007 alert or bell |
| `\b` | U+0008 backspace |
| `\\` | U+005c backslash |
| `\t` | U+0009 horizontal tab |
| `\n` | U+000A line feed or newline |
| `\f` | U+000C form feed |
| `\r` | U+000D carriage return |
| `\v` | U+000b vertical tab |

Arbitrary values can be encoded with backslash escapes and can be used in any `""` string literal.

There are four different formats:

* `\x` followed by exactly two hexadecimal digits,
* `\` followed by exactly three octal digits,
* `\u` followed by exactly four hexadecimal digits,
* `\U` followed by exactly eight hexadecimal digits.

The escapes `\u` and `\U` represent Unicode code points.

```text
fmt.Println("\\caf\u00e9") // Prints \café
```

## How to use the io.Reader interface

### Basics <a id="basics"></a>

The [`io.Reader`](https://golang.org/pkg/io/#Reader) [interface](https://yourbasic.org/golang/interfaces-explained/) represents an entity from which you can read a stream of bytes.

```text
type Reader interface {
        Read(buf []byte) (n int, err error)
}
```

`Read` reads up to `len(buf)` bytes into `buf` and returns the number of bytes read – it returns an [`io.EOF`](https://golang.org/pkg/io/#pkg-variables) error when the stream ends.

The standard library provides numerous Reader [implementations](https://golang.org/search?q=Read#Global) \(including in-memory byte buffers, files and network connections\), and Readers are accepted as input by many utilities \(including the HTTP client and server implementations\).

### Use a built-in reader <a id="use-a-built-in-reader"></a>

As an example, you can create a Reader from a string using the [`strings.Reader`](https://golang.org/pkg/strings/#Reader) function and then pass the Reader directly to the [`http.Post`](https://golang.org/pkg/net/http/#Post) function in package [`net/http`](https://golang.org/pkg/net/http/). The Reader is then used as the source for the data to be posted.

```text
r := strings.NewReader("my request")
resp, err := http.Post("http://foo.bar",
	"application/x-www-form-urlencoded", r)
```

Since `http.Post` uses a Reader instead of a `[]byte` it’s trivial to, for instance, use the contents of a file instead.

### Read directly from a byte stream <a id="read-directly-from-a-byte-stream"></a>

You can use the `Read` function directly \(this is the least common use case\).

```text
r := strings.NewReader("abcde")

buf := make([]byte, 4)
for {
	n, err := r.Read(buf)
	fmt.Println(n, err, buf[:n])
	if err == io.EOF {
		break
	}
}
```

```text
4 <nil> [97 98 99 100]
1 <nil> [101]
0 EOF []
```

Use [`io.ReadFull`](https://golang.org/pkg/io/#ReadFull) to read exactly `len(buf)` bytes into `buf`:

```text
r := strings.NewReader("abcde")

buf := make([]byte, 4)
if _, err := io.ReadFull(r, buf); err != nil {
	log.Fatal(err)
}
fmt.Println(buf)

if _, err := io.ReadFull(r, buf); err != nil {
	fmt.Println(err)
}
```

```text
[97 98 99 100]
unexpected EOF
```

Use [`ioutil.ReadAll`](https://golang.org/pkg/io/ioutil/#ReadAll) to read everything:

```text
r := strings.NewReader("abcde")

buf, err := ioutil.ReadAll(r)
if err != nil {
	log.Fatal(err)
}
fmt.Println(buf)
```

```text
[97 98 99 100 101]
```

### Buffered reading and scanning <a id="buffered-reading-and-scanning"></a>

The [`bufio.Reader`](https://golang.org/pkg/bufio/#Reader) and [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) types wrap a Reader creating another Reader that also implements the interface but provides buffering and some help for textual input.

In this example we use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to count the number of words in a text.

```text
const input = `Beware of bugs in the above code;
I have only proved it correct, not tried it.`

scanner := bufio.NewScanner(strings.NewReader(input))
scanner.Split(bufio.ScanWords) // Set up the split function.

count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Println(err)
}
fmt.Println(count)
```

```text
16
```

## How to use the io.Writer interface

### Basics <a id="basics"></a>

The [`io.Writer`](https://golang.org/pkg/io/#Writer) [interface](https://yourbasic.org/golang/interfaces-explained/) represents an entity to which you can write a stream of bytes.

```text
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

`Write` writes up to `len(p)` bytes from `p` to the underlying data stream – it returns the number of bytes written and any error encountered that caused the write to stop early.

The standard library provides numerous Writer [implementations](https://golang.org/search?q=Write#Global), and Writers are accepted as input by many utilities.

### How to use a built-in writer \(3 examples\) <a id="how-to-use-a-built-in-writer-3-nbsp-examples"></a>

As a first example, you can write directly into a [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) using the [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) function. This works since

* `bytes.Buffer` has a `Write` method, and
* `fmt.Fprintf` takes a `Writer` as its first argument.

```text
var buf bytes.Buffer
fmt.Fprintf(&buf, "Size: %d MB.", 85)
s := buf.String()) // s == "Size: 85 MB."
```

Similarly, you can write directly into files or other streams, such as http connections. See the [HTTP server example](https://yourbasic.org/golang/http-server-example/) article for a complete code example.

This is a very common pattern in Go. As yet another example, you can compute the hash value of a file by copying the file into the `io.Writer` function of a suitable [`hash.Hash`](https://golang.org/pkg/hash/#Hash) object. See [Hash checksums](https://yourbasic.org/golang/hash-md5-sha256-string-file/#file) for code.

### Optimize string writes <a id="optimize-string-writes"></a>

Some Writers in the standard library have an additional `WriteString` method. This method can be more efficient than the standard `Write` method since it writes a string directly without allocating a byte slice.

You can take direct advantage of this optimization by using the [`io.WriteString()`](https://golang.org/pkg/io/#WriteString) function.

```text
func WriteString(w Writer, s string) (n int, err error)
```

If `w` implements a `WriteString` method, it is invoked directly. Otherwise, `w.Write` is called exactly once.

## Create a new image

Use the [`image`](https://golang.org/pkg/image/), [`image/color`](https://golang.org/pkg/image/color/), and [`image/png`](https://golang.org/pkg/image/png/) packages to create a PNG image.

```text
width := 200
height := 100

upLeft := image.Point{0, 0}
lowRight := image.Point{width, height}

img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

// Colors are defined by Red, Green, Blue, Alpha uint8 values.
cyan := color.RGBA{100, 200, 200, 0xff}

// Set color for each pixel.
for x := 0; x < width; x++ {
    for y := 0; y < height; y++ {
        switch {
        case x < width/2 && y < height/2: // upper left quadrant
            img.Set(x, y, cyan)
        case x >= width/2 && y >= height/2: // lower right quadrant
            img.Set(x, y, color.White)
        default:
            // Use zero value.
        }
    }
}

// Encode as PNG.
f, _ := os.Create("image.png")
png.Encode(f, img)
```

Output \([image.png](https://yourbasic.org/golang/image.png)\):

```text

```

**Note:** The upper right and lower left quadrants of the image are transparent \(the alpha value is 0\) and will be the same color as the background.

### Go image support <a id="go-image-support"></a>

The [`image`](https://golang.org/pkg/image/) package implements a basic 2-D image library without painting or drawing functionality. The article [The Go image package](https://blog.golang.org/go-image-package) has a nice introduction to images, color models, and image formats in Go.

Additionally, the [`image/draw`](https://golang.org/pkg/image/draw) package provides image composition functions that can be used to perform a number of common image manipulation tasks. The article [The Go image/draw package](https://blog.golang.org/go-imagedraw-package) has plenty of examples.

## Write log to file \(or /dev/null\)

This code appends a log message to the file `text.log`. It creates the file if it doesn’t already exist.

```text
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()

logger := log.New(f, "prefix", log.LstdFlags)
logger.Println("text to append")
logger.Println("more text to append")
```

Contents of `text.log`:

```text
prefix: 2017/10/20 07:52:58 text to append
prefix: 2017/10/20 07:52:58 more text to append
```

* [`log.New`](https://golang.org/pkg/log/#New) creates a new [`log.Logger`](https://golang.org/pkg/log/#Logger) that writes to `f`.
* The prefix appears at the beginning of each generated log line.
* The [`flag`](https://golang.org/pkg/log/#pkg-constants) argument defines which text to prefix to each log entry.

### Disable logging <a id="disable-logging"></a>

To turn off all output from a [`log.Logger`](https://golang.org/pkg/log/#Logger), set the output destination to [`ioutil.Discard`](https://golang.org/pkg/io/ioutil/#pkg-variables), a writer on which all calls succeed without doing anything.

```text
log.SetOutput(ioutil.Discard)
```

## Hello world HTTP server example

### A basic web server <a id="a-basic-web-server"></a>

If you access the URL `http://localhost:8080/world` on a machine where the program below is running, you will be greeted by this page.

```text
package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
```

* The call to [`http.HandleFunc`](https://golang.org/pkg/net/http/#HandleFunc) tells the [`net.http`](https://golang.org/pkg/net/http/) package to handle all requests to the web root with the `HelloServer` function.
* The call to [`http.ListenAndServe`](https://golang.org/pkg/net/http/#ListenAndServe) tells the server to listen on the TCP network address `:8080`. This function blocks until the program is terminated.
* Writing to an [`http.ResponseWriter`](https://golang.org/pkg/net/http/#ResponseWriter) sends data to the HTTP client.
* An [`http.Request`](https://golang.org/pkg/net/http/#Request) is a data structure that represents a client HTTP request.
* `r.URL.Path` is the path component of the requested URL. In this case, `"/world"` is the path component of `"http://localhost:8080/world"`.

### Further reading: a complete wiki <a id="further-reading-a-complete-wiki"></a>

The [Writing Web Applications](https://golang.org/doc/articles/wiki/) tutorial shows how to extend this small example into a complete wiki.

The tutorial covers how to

* create a data structure with load and save methods,
* use the [`net/http`](https://golang.org/pkg/net/http/) package to build web applications,
* use the [`html/template`](https://golang.org/pkg/html/template/) package to process HTML templates,
* use the [`regexp`](https://yourbasic.org/golang/regexp-cheat-sheet/) package to validate user input.

