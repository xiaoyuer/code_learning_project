# 27 Go Gotcha Ninja Pitfalls

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

