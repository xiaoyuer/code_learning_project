# Type, value and equality of interfaces

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


