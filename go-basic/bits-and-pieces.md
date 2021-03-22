# Bits and Pieces

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



