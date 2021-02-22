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

