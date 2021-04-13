# Functional Programming

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



