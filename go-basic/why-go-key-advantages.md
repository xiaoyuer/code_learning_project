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




