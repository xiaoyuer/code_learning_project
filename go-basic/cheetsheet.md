# CheetSheet

### String literals \(escape characters\) <a id="string-literals-escape-characters"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `""` |  | [Default zero value](https://yourbasic.org/golang/default-zero-value/) for type `string` |
| `"Japan 日本"` | Japan 日本 | Go code is [Unicode text encoded in UTF‑8](https://yourbasic.org/golang/rune/) |
| `"\xe6\x97\xa5"` | 日 | `\xNN` specifies a byte |
| `"\u65E5"` | 日 | `\uNNNN` specifies a Unicode value |
| `"\\"` | \ | Backslash |
| `"\""` | " | Double quote |
| `"\n"` |  | Newline |
| `"\t"` |  | Tab |
| ```\xe6``` | \xe6 | Raw string literal\* |
| `html.EscapeString("<>")` | &lt;&gt; | HTML escape for &lt;, &gt;, &, ' and " |
| `url.PathEscape("A B")` | A%20B | URL percent-encoding net/url |

\* In ```````` string literals, text is interpreted literally and backslashes have no special meaning. See [Escapes and multiline strings](https://yourbasic.org/golang/multiline-string/) for more on raw strings, escape characters and string encodings.

### Concatenate <a id="concatenate"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `"Ja" + "pan"` | Japan | Concatenation |

> **Performance tips**  
> See [3 tips for efficient string concatenation](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/) for how to best use a string builder to concatenate strings without redundant copying.

### Equal and compare \(ignore case\) <a id="equal-and-compare-ignore-case"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `"Japan" == "Japan"` | true | Equality |
| `strings.EqualFold("Japan", "JAPAN")` | true | Unicode case folding |
| `"Japan" < "japan"` | true | Lexicographic order |

### Length in bytes or runes <a id="length-in-bytes-or-runes"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `len("日")` | 3 | Length in bytes |
| `utf8.RuneCountInString("日")` | 1 | in runes unicode/utf8 |
| `utf8.ValidString("日")` | true | UTF-8? unicode/utf8 |

### Index, substring, iterate <a id="index-substring-iterate"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `"Japan"[2]` | 'p' | Byte at position 2 |
| `"Japan"[1:3]` | ap | Byte indexing |
| `"Japan"[:2]` | Ja |  |
| `"Japan"[2:]` | pan |  |

A Go [range loop](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) iterates over UTF-8 encoded characters \([runes](https://yourbasic.org/golang/rune/)\):

```text
for i, ch := range "Japan 日本" {
    fmt.Printf("%d:%q ", i, ch)
}
// Output: 0:'J' 1:'a' 2:'p' 3:'a' 4:'n' 5:' ' 6:'日' 9:'本'
```

Iterating over bytes produces nonsense characters for non-ASCII text:

```text
s := "Japan 日本"
for i := 0; i < len(s); i++ {
    fmt.Printf("%q ", s[i])
}
// Output: 'J' 'a' 'p' 'a' 'n' ' ' 'æ' '\u0097' '¥' 'æ' '\u009c' '¬'
```

### Search \(contains, prefix/suffix, index\) <a id="search-contains-prefix-suffix-index"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Contains("Japan", "abc")` | false | Is abc in Japan? |
| `strings.ContainsAny("Japan", "abc")` | true | Is a, b or c in Japan? |
| `strings.Count("Banana", "ana")` | 1 | Non-overlapping instances of ana |
| `strings.HasPrefix("Japan", "Ja")` | true | Does Japan start with Ja? |
| `strings.HasSuffix("Japan", "pan")` | true | Does Japan end with pan? |
| `strings.Index("Japan", "abc")` | -1 | Index of first abc |
| `strings.IndexAny("Japan", "abc")` | 1 | a, b or c |
| `strings.LastIndex("Japan", "abc")` | -1 | Index of last abc |
| `strings.LastIndexAny("Japan", "abc")` | 3 | a, b or c |

### Replace \(uppercase/lowercase, trim\) <a id="replace-uppercase-lowercase-trim"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Replace("foo", "o", ".", 2)` | f.. | Replace first two “o” with “.” Use -1 to replace all |
| `f := func(r rune) rune {     return r + 1 } strings.Map(f, "ab")` | bc | Apply function to each character |
| `strings.ToUpper("Japan")` | JAPAN | Uppercase |
| `strings.ToLower("Japan")` | japan | Lowercase |
| `strings.Title("ja pan")` | Ja Pan | Initial letters to uppercase |
| `strings.TrimSpace(" foo\n")` | foo | Strip leading and trailing white space |
| `strings.Trim("foo", "fo")` |  | Strip _leading and trailing_ f:s and o:s |
| `strings.TrimLeft("foo", "f")` | oo | _only leading_ |
| `strings.TrimRight("foo", "o")` | f | _only trailing_ |
| `strings.TrimPrefix("foo", "fo")` | o |  |
| `strings.TrimSuffix("foo", "o")` | fo |  |

### Split by space or comma <a id="split-by-space-or-comma"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Fields(" a\t b\n")` | `["a" "b"]` | Remove white space |
| `strings.Split("a,b", ",")` | `["a" "b"]` | Remove separator |
| `strings.SplitAfter("a,b", ",")` | `["a," "b"]` | Keep separator |

### Join strings with separator <a id="join-strings-with-separator"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strings.Join([]string{"a", "b"}, ":")` | a:b | Add separator |
| `strings.Repeat("da", 2)` | dada | 2 copies of “da” |

### Format and convert <a id="format-and-convert"></a>

| Expression | Result | Note |
| :--- | :--- | :--- |
| `strconv.Itoa(-42)` | `"-42"` | Int to string |
| `strconv.FormatInt(255, 16)` | `"ff"` | Base 16 |

#### Sprintf <a id="sprintf"></a>

The [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function is often your best friend when formatting data:

```text
s := fmt.Sprintf("%.4f", math.Pi) // s == "3.1416
```

This [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) covers the most common formatting flags.

### Regular expression <a id="regular-expressions"></a>

For more advanced string handling, see this [Regular expressions tutorial](https://yourbasic.org/golang/regexp-cheat-sheet/), a gentle introduction to the `regexp` package with cheat sheet and plenty of examples.

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
>
> ### Basics <a id="basics"></a>
>
> The expression `T(x)` converts the value `x` to the type `T`.
>
> ```text
> x := 5.1
> n := int(x) // convert float to int
> ```
>
> The conversion rules are extensive but predictable:
>
> * all conversions between typed expressions must be explicitly stated,
> * illegal conversions are caught by the compiler.
>
> Conversions to and from numbers and strings may **change the representation** and have a **run-time cost**. All other conversions only change the type but not the representation of `x`.
>
> ### Interfaces <a id="interfaces"></a>
>
> > To “convert” an [interface](https://yourbasic.org/golang/interfaces-explained/) to a string, struct or map you should use a **type assertion** or a **type switch**. A type assertion doesn’t really convert an interface to another data type, but it provides access to an interface’s concrete value, which is typically what you want.[Type assertions and type switches](https://yourbasic.org/golang/type-assertion-switch/)
>
> ### Integers <a id="integers"></a>
>
> * When converting to a shorter integer type, the value is **truncated** to fit in the result type’s size.
> * When converting to a longer integer type,
>   * if the value is a signed integer, it is [**sign extended**](https://en.wikipedia.org/wiki/Sign_extension);
>   * otherwise it is **zero extended**.
>
> ```text
> a := uint16(0x10fe) // 0001 0000 1111 1110
> b := int8(a)        //           1111 1110 (truncated to -2)
> c := uint16(b)      // 1111 1111 1111 1110 (sign extended to 0xfffe)
> ```
>
> ### Floats <a id="floats"></a>
>
> * When converting a floating-point number to an integer, the **fraction is discarded** \(truncation towards zero\).
> * When converting an integer or floating-point number to a floating-point type, the result value is **rounded to the precision** specified by the destination type.
>
> ```text
> var x float64 = 1.9
> n := int64(x) // 1
> n = int64(-x) // -1
>
> n = 1234567890
> y := float32(n) // 1.234568e+09
> ```
>
> > **Warning:** In all non-constant conversions involving floating-point or complex values, if the result type cannot represent the value the conversion succeeds but the result value is implementation-dependent.[The Go Programming Language Specification: Conversions](https://golang.org/ref/spec#Conversions)
>
> ### Integer to string <a id="integer-to-string"></a>
>
> * When converting an integer to a string, the value is interpreted as a Unicode code point, and the resulting string will contain the character represented by that code point, encoded in UTF-8.
> * If the value does not represent a valid code point \(for instance if it’s negative\), the result will be `"\ufffd"`, the Unicode replacement character �.
>
> ```text
> string(97) // "a"
> string(-1) // "\ufffd" == "\xef\xbf\xbd"
> ```
>
> > Use [`strconv.Itoa`](https://golang.org/pkg/strconv/#Itoa) to get the decimal string representation of an integer.
> >
> > ```text
> > strconv.Itoa(97) // "97"
> > ```
>
> ### Strings and byte slices <a id="strings-and-byte-slices"></a>
>
> * Converting a slice of bytes to a string type yields a string whose successive bytes are the elements of the slice.
> * Converting a value of a string type to a slice of bytes type yields a slice whose successive elements are the bytes of the string.
>
> ```text
> string([]byte{97, 230, 151, 165}) // "a日"
> []byte("a日")                     // []byte{97, 230, 151, 165}
> ```
>
> ### Strings and rune slices <a id="strings-and-rune-slices"></a>
>
> * Converting a slice of runes to a string type yields a string that is the concatenation of the individual rune values converted to strings.
> * Converting a value of a string type to a slice of runes type yields a slice containing the individual Unicode code points of the string.
>
> ```text
> string([]rune{97, 26085}) // "a日"
> []rune("a日")             // []rune{97, 26085}
> ```
>
> ### Underlying type <a id="underlying-type"></a>
>
> A non-constant value can be converted to type `T` if it has the same underlying type as `T`.
>
> In this example, the underlying type of `int64`, `T1`, and `T2` is `int64`.
>
> ```text
> type (
> 	T1 int64
> 	T2 T1
> )
> ```
>
> It’s idiomatic in Go to convert the type of an expression to access a specific method.
>
> ```text
> var n int64 = 12345
> fmt.Println(n)                // 12345
> fmt.Println(time.Duration(n)) // 12.345µs
> ```
>
> \(The underlying type of [`time.Duration`](https://golang.org/pkg/time/#Duration) is `int64`, and the `time.Duration` type has a [`String`](https://golang.org/pkg/time/#Duration.String) method that returns the duration formatted as a time.\)
>
> ### Implicit conversions <a id="implicit-conversions"></a>
>
> The only implicit conversion in Go is when an untyped constant is used in a situation where a type is required.
>
> In this example the untyped literals `1` and `2` are implicitly converted.
>
> ```text
> var x float64
> x = 1 // Same as x = float64(1)
>
> t := 2 * time.Second // Same as t := time.Duration(2) * time.Second
> ```
>
> \(The implicit conversions are necessary since there is no mixing of numeric types in Go. You can only multiply a `time.Duration` with another `time.Duration`.\)
>
> When the type can’t be inferred from the context, an untyped constant is converted to a `bool`, `int`, `float64`, `complex128`, `string` or `rune` depending on the syntactical format of the constant.
>
> ```text
> n := 1   // Same as n := int(1)
> x := 1.0 // Same as x := float64(1.0)
> s := "A" // Same as s := string("A")
> c := 'A' // Same as c := rune('A')
> ```
>
> Illegal implicit conversions are caught by the compiler.
>
> ```text
> var b byte = 256 // Same as var b byte = byte(256)
> ```
>
> ```text
> ../main.go:2:6: constant 256 overflows byte
> ```
>
> ### Pointers <a id="pointers"></a>
>
> The Go compiler does not allow conversions between pointers and integers. Package [`unsafe`](https://golang.org/pkg/unsafe/) implements this functionality under restricted circumstances.
>
> > **Warning:** The built-in package unsafe, known to the compiler, provides facilities for low-level programming including operations that violate the type system. A package using unsafe must be vetted manually for type safety and may not be portable.[The Go Programming Language Specification: Package unsafe](https://golang.org/ref/spec#Package_unsafe)

