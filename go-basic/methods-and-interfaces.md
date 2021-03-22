# Methods and Interfaces

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



