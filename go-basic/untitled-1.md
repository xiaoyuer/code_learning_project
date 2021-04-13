# Object-oriented programming without inheritance

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



