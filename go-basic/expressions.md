# Expressions

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



