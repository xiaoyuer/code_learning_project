# Go Code For Common Tasks

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

