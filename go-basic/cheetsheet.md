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



## fmt.Printf formatting tutorial and cheat sheet

### Basics <a id="basics"></a>

With the Go [`fmt`](https://golang.org/pkg/fmt) package you can format numbers and strings padded with spaces or zeroes, in different bases, and with optional quotes.

You submit a **template string** that contains the text you want to format plus some **annotation verbs** that tell the `fmt` functions how to format the trailing arguments.

#### Printf <a id="printf"></a>

In this example, [`fmt.Printf`](https://golang.org/pkg/fmt/#Printf) formats and writes to standard output:

```text
fmt.Printf("Binary: %b\\%b", 4, 5) // Prints `Binary: 100\101`
```

* the **template string** is `"Binary: %b\\%b"`,
* the **annotation verb** `%b` formats a number in binary, and
* the **special value** `\\` is a backslash.

As a special case, the verb `%%`, which consumes no argument, produces a percent sign:

```text
fmt.Printf("%d %%", 50) // Prints `50 %`
```

#### Sprintf \(format without printing\) <a id="sprintf-format-without-printing"></a>

Use [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) to format a string without printing it:

```text
s := fmt.Sprintf("Binary: %b\\%b", 4, 5) // s == `Binary: 100\101`
```

#### Find fmt errors with vet <a id="find-fmt-errors-with-vet"></a>

If you try to compile and run this incorrect line of code

```text
fmt.Printf("Binary: %b\\%b", 4) // An argument to Printf is missing.
```

you’ll find that the program will compile, and then print

```text
Binary: 100\%!b(MISSING)
```

To catch this type of errors early, you can use the [vet command](https://golang.org/cmd/vet/) – it can find calls whose arguments do not align with the format string.

```text
$ go vet example.go
example.go:8: missing argument for Printf("%b"): format reads arg 2, have only 1 args
```

### Cheat sheet <a id="cheat-sheet"></a>

#### Default formats and type <a id="default"></a>

* **Value:** `[]int64{0, 1}`

| Format | Verb | Description |
| :--- | :--- | :--- |
| \[0 1\] | `%v` | Default format |
| \[\]int64{0, 1} | `%#v` | Go-syntax format |
| \[\]int64 | `%T` | The type of the value |

#### Integer \(indent, base, sign\) <a id="integer"></a>

* **Value:** `15`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 15 | `%d` | Base 10 |
| +15 | `%+d` | Always show sign |
| ␣␣15 | `%4d` | Pad with spaces \(width 4, right justified\) |
| 15␣␣ | `%-4d` | Pad with spaces \(width 4, left justified\) |
| 0015 | `%04d` | Pad with zeroes \(width 4\) |
| 1111 | `%b` | Base 2 |
| 17 | `%o` | Base 8 |
| f | `%x` | Base 16, lowercase |
| F | `%X` | Base 16, uppercase |
| 0xf | `%#x` | Base 16, with leading 0x |

#### Character \(quoted, Unicode\) <a id="character"></a>

* **Value:** `65`   \(Unicode letter A\)

| Format | Verb | Description |
| :--- | :--- | :--- |
| A | `%c` | Character |
| 'A' | `%q` | Quoted character |
| U+0041 | `%U` | Unicode |
| U+0041 'A' | `%#U` | Unicode with character |

#### Boolean \(true/false\) <a id="boolean"></a>

Use `%t` to format a boolean as `true` or `false`.

#### Pointer \(hex\) <a id="pointer"></a>

Use `%p` to format a pointer in base 16 notation with leading `0x`.

#### Float \(indent, precision, scientific notation\) <a id="float"></a>

* **Value:** `123.456`

| Format | Verb | Description |
| :--- | :--- | :--- |
| 1.234560e+02 | `%e` | Scientific notation |
| 123.456000 | `%f` | Decimal point, no exponent |
| 123.46 | `%.2f` | Default width, precision 2 |
| ␣␣123.46 | `%8.2f` | Width 8, precision 2 |
| 123.456 | `%g` | Exponent as needed, necessary digits only |

#### String or byte slice \(quote, indent, hex\) <a id="string-or-byte-slice"></a>

* **Value:** `"café"`

| Format | Verb | Description |
| :--- | :--- | :--- |
| café | `%s` | Plain string |
| ␣␣café | `%6s` | Width 6, right justify |
| café␣␣ | `%-6s` | Width 6, left justify |
| "café" | `%q` | Quoted string |
| 636166c3a9 | `%x` | Hex dump of byte values |
| 63 61 66 c3 a9 | `% x` | Hex dump with spaces |

#### Special values <a id="special-values"></a>

| Value | Description |
| :--- | :--- |
| `\a` | U+0007 alert or bell |
| `\b` | U+0008 backspace |
| `\\` | U+005c backslash |
| `\t` | U+0009 horizontal tab |
| `\n` | U+000A line feed or newline |
| `\f` | U+000C form feed |
| `\r` | U+000D carriage return |
| `\v` | U+000b vertical tab |

Arbitrary values can be encoded with backslash escapes and can be used in any `""` string literal.

There are four different formats:

* `\x` followed by exactly two hexadecimal digits,
* `\` followed by exactly three octal digits,
* `\u` followed by exactly four hexadecimal digits,
* `\U` followed by exactly eight hexadecimal digits.

The escapes `\u` and `\U` represent Unicode code points.

```text
fmt.Println("\\caf\u00e9") // Prints \café
```

### Further readin <a id="further-reading"></a>

[40+ practical string tips \[cheat sheet\]](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)

## Format a time or date \[complete guide\]

### Basic example <a id="basic-example"></a>

Go doesn’t use yyyy-mm-dd layout to format a time. Instead, you format a special **layout parameter**

`Mon Jan 2 15:04:05 MST 2006`

the same way as the time or date should be formatted. \(This date is easier to remember when written as `01/02 03:04:05PM ‘06 -0700`.\)

```text
const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
)
date := "1999-12-31"
t, _ := time.Parse(layoutISO, date)
fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
fmt.Println(t.Format(layoutUS)) // December 31, 1999
```

The function

* [`time.Parse`](https://golang.org/pkg/time/#Parse) parses a date string, and
* [`Format`](https://golang.org/pkg/time/#Time.Format) formats a [`time.Time`](https://golang.org/pkg/time/#Time).

They have the following signatures:

```text
func Parse(layout, value string) (Time, error)
func (t Time) Format(layout string) string
```

### Standard time and date formats <a id="standard-time-and-date-formats"></a>

| Go layout | Note |
| :--- | :--- |
| `January 2, 2006` | Date |
| `01/02/06` |  |
| `Jan-02-06` |  |
| `15:04:05` | Time |
| `3:04:05 PM` |  |
| `Jan _2 15:04:05` | Timestamp |
| `Jan _2 15:04:05.000000` | with microseconds |
| `2006-01-02T15:04:05-0700` | [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) \([RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)\) |
| `2006-01-02` |  |
| `15:04:05` |  |
| `02 Jan 06 15:04 MST` | [RFC 822](https://www.ietf.org/rfc/rfc822.txt) |
| `02 Jan 06 15:04 -0700` | with numeric zone |
| `Mon, 02 Jan 2006 15:04:05 MST` | [RFC 1123](https://www.ietf.org/rfc/rfc1123.txt) |
| `Mon, 02 Jan 2006 15:04:05 -0700` | with numeric zone |

The following predefined date and timestamp [format constants](https://golang.org/pkg/time/#pkg-constants) are also available.

```text
ANSIC       = "Mon Jan _2 15:04:05 2006"
UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
RFC822      = "02 Jan 06 15:04 MST"
RFC822Z     = "02 Jan 06 15:04 -0700"
RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700"
RFC3339     = "2006-01-02T15:04:05Z07:00"
RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
Kitchen     = "3:04PM"
// Handy time stamps.
Stamp      = "Jan _2 15:04:05"
StampMilli = "Jan _2 15:04:05.000"
StampMicro = "Jan _2 15:04:05.000000"
StampNano  = "Jan _2 15:04:05.000000000"
```

### Layout options <a id="layout-options"></a>

| Type | Options |
| :--- | :--- |
| Year | `06`   `2006` |
| Month | `01`   `1`   `Jan`   `January` |
| Day | `02`   `2`   `_2`   \(width two, right justified\) |
| Weekday | `Mon`   `Monday` |
| Hours | `03`   `3`   `15` |
| Minutes | `04`   `4` |
| Seconds | `05`   `5` |
| ms μs ns | `.000`   `.000000`   `.000000000` |
| ms μs ns | `.999`   `.999999`   `.999999999`   \(trailing zeros removed\) |
| am/pm | `PM`   `pm` |
| Timezone | `MST` |
| Offset | `-0700`   `-07`   `-07:00`   `Z0700`   `Z07:00` |

### Corner cases <a id="corner-cases"></a>

It’s not possible to specify that an hour should be rendered without a leading zero in a 24-hour time format.

It’s not possible to specify midnight as `24:00` instead of `00:00`. A typical usage for this would be giving opening hours ending at midnight, such as `07:00-24:00`.

It’s not possible to specify a time containing a leap second: `23:59:60`. In fact, the time package assumes a Gregorian calendar without leap seconds.

## Regexp tutorial and cheat sheet

A regular expression is a sequence of characters that define a search pattern.

### Basics <a id="basics"></a>

The regular expression `a.b` matches any string that starts with an `a`, ends with a `b`, and has a single character in between \(the period matches any character\).

To check if there is a **substring** matching `a.b`, use the [regexp.MatchString](https://golang.org/pkg/regexp/#MatchString) function.

```text
matched, err := regexp.MatchString(`a.b`, "aaxbb")
fmt.Println(matched) // true
fmt.Println(err)     // nil (regexp is valid)
```

To check if a **full string** matches `a.b`, anchor the start and the end of the regexp:

* the caret `^` matches the beginning of a text or line,
* the dollar sign `$` matches the end of a text.

```text
matched, _ := regexp.MatchString(`^a.b$`, "aaxbb")
fmt.Println(matched) // false
```

Similarly, we can check if a string **starts with** or **ends with** a pattern by using only the start or end anchor.

#### Compile <a id="compile"></a>

For more complicated queries, you should compile a regular expression to create a [`Regexp`](https://golang.org/pkg/regexp/#Regexp) object. There are two options:

```text
re1, err := regexp.Compile(`regexp`) // error if regexp invalid
re2 := regexp.MustCompile(`regexp`)  // panic if regexp invalid
```

#### Raw strings <a id="raw-strings"></a>

It’s convenient to use ```raw strings``` when writing regular expressions, since both ordinary string literals and regular expressions use backslashes for special characters.

A [raw string](https://yourbasic.org/golang/multiline-string/#raw-string-literals), delimited by backticks, is interpreted literally and backslashes have no special meaning.

### Cheat sheet <a id="cheat-sheet"></a>

#### Choice and grouping <a id="choice-and-grouping"></a>

| Regexp | Meaning |
| :--- | :--- |
| `xy` | `x` followed by `y` |
| `x|y` | `x` or `y`, prefer `x` |
| `xy|z` | same as `(xy)|z` |
| `xy*` | same as `x(y*)` |

#### Repetition \(greedy and non-greedy\) <a id="repetition-greedy-and-non-greedy"></a>

| Regexp | Meaning |
| :--- | :--- |
| `x*` | zero or more x, prefer more |
| `x*?` | prefer fewer \(non-greedy\) |
| `x+` | one or more x, prefer more |
| `x+?` | prefer fewer \(non-greedy\) |
| `x?` | zero or one x, prefer one |
| `x??` | prefer zero |
| `x{n}` | exactly n x |

#### Character classes <a id="character-classes"></a>

| Expression | Meaning |
| :--- | :--- |
| `.` | any character |
| `[ab]` | the character a or b |
| `[^ab]` | any character except a or b |
| `[a-z]` | any character from a to z |
| `[a-z0-9]` | any character from a to z or 0 to 9 |
| `\d` | a digit: `[0-9]` |
| `\D` | a non-digit: `[^0-9]` |
| `\s` | a whitespace character: `[\t\n\f\r ]` |
| `\S` | a non-whitespace character: `[^\t\n\f\r ]` |
| `\w` | a word character: `[0-9A-Za-z_]` |
| `\W` | a non-word character: `[^0-9A-Za-z_]` |
| `\p{Greek}` | Unicode character class\* |
| `\pN` | one-letter name |
| `\P{Greek}` | negated Unicode character class\* |
| `\PN` | one-letter name |

\* [RE2: Unicode character class names](https://github.com/google/re2/wiki/Syntax)

#### Special characters <a id="special-characters"></a>

To match a **special character** `\^$.|?*+-[]{}()` literally, escape it with a backslash. For example `\{` matches an opening brace symbol.

Other escape sequences are:

| Symbol | Meaning |
| :--- | :--- |
| `\t` | horizontal tab = `\011` |
| `\n` | newline = `\012` |
| `\f` | form feed = `\014` |
| `\r` | carriage return = `\015` |
| `\v` | vertical tab = `\013` |
| `\123` | octal character code \(up to three digits\) |
| `\x7F` | hex character code \(exactly two digits\) |

#### Text boundary anchors <a id="text-boundary-anchors"></a>

| Symbol | Matches |
| :--- | :--- |
| `\A` | at beginning of text |
| `^` | at beginning of text or line |
| `$` | at end of text |
| `\z` |  |
| `\b` | at ASCII word boundary |
| `\B` | not at ASCII word boundary |

#### Case-insensitive and multiline matches <a id="case-insensitive-and-multiline-matches"></a>

To change the default matching behavior, you can add a set of flags to the beginning of a regular expression.

For example, the prefix `"(?is)"` makes the matching case-insensitive and lets `.` match `\n`. \(The default matching is case-sensitive and `.` doesn’t match `\n`.\)

| Flag | Meaning |
| :--- | :--- |
| `i` | case-insensitive |
| `m` | let `^` and `$` match begin/end line in addition to begin/end text \(multi-line mode\) |
| `s` | let `.` match `\n` \(single-line mode\) |

### Code examples <a id="code-examples"></a>

#### First match <a id="first-match"></a>

Use the [`FindString`](https://golang.org/pkg/regexp/#Regexp.FindString) method to find the **text of the first match**. If there is no match, the return value is an empty string.

```text
re := regexp.MustCompile(`foo.?`)
fmt.Printf("%q\n", re.FindString("seafood fool")) // "food"
fmt.Printf("%q\n", re.FindString("meat"))         // ""
```

#### Location <a id="location"></a>

Use the [`FindStringIndex`](https://golang.org/pkg/regexp/#Regexp.FindStringIndex) method to find `loc`, the **location of the first match**, in a string `s`. The match is at `s[loc[0]:loc[1]]`. A return value of nil indicates no match.

```text
re := regexp.MustCompile(`ab?`)
fmt.Println(re.FindStringIndex("tablett"))    // [1 3]
fmt.Println(re.FindStringIndex("foo") == nil) // true
```

#### All matches <a id="all-matches"></a>

Use the [`FindAllString`](https://golang.org/pkg/regexp/#Regexp.FindAllString) method to find the **text of all matches**. A return value of nil indicates no match.

The method takes an integer argument `n`; if `n >= 0`, the function returns at most `n` matches.

```text
re := regexp.MustCompile(`a.`)
fmt.Printf("%q\n", re.FindAllString("paranormal", -1)) // ["ar" "an" "al"]
fmt.Printf("%q\n", re.FindAllString("paranormal", 2))  // ["ar" "an"]
fmt.Printf("%q\n", re.FindAllString("graal", -1))      // ["aa"]
fmt.Printf("%q\n", re.FindAllString("none", -1))       // [] (nil slice)
```

#### Replace <a id="replace"></a>

Use the [`ReplaceAllString`](https://golang.org/pkg/regexp/#Regexp.ReplaceAllString) method to **replace the text of all matches**. It returns a copy, replacing all matches of the regexp with a replacement string.

```text
re := regexp.MustCompile(`ab*`)
fmt.Printf("%q\n", re.ReplaceAllString("-a-abb-", "T")) // "-T-T-"
```

#### Split <a id="split"></a>

Use the [`Split`](https://golang.org/pkg/regexp/#Regexp.Split) method to **slice a string into substrings** separated by the regexp. It returns a slice of the substrings between those expression matches. A return value of nil indicates no match.

The method takes an integer argument `n`; if `n >= 0`, the function returns at most `n` matches.

```text
a := regexp.MustCompile(`a`)
fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

zp := regexp.MustCompile(`z+`)
fmt.Printf("%q\n", zp.Split("pizza", -1)) // ["pi" "a"]
fmt.Printf("%q\n", zp.Split("pizza", 0))  // [] (nil slice)
fmt.Printf("%q\n", zp.Split("pizza", 1))  // ["pizza"]
fmt.Printf("%q\n", zp.Split("pizza", 2))  // ["pi" "a"]
```

**More functions**

There are 16 functions following the naming pattern

```text
Find(All)?(String)?(Submatch)?(Index)?
```

For example: `Find`, `FindAllString`, `FindStringIndex`, …

* If `All` is present, the function matches successive non-overlapping matches.
* `String` indicates that the argument is a string; otherwise it’s a byte slice.
* If `Submatch` is present, the return value is a slice of successive submatches. Submatches are matches of parenthesized subexpressions within the regular expression. See [`FindSubmatch`](https://golang.org/pkg/regexp/#Regexp.FindSubmatch) for an example.
* If `Index` is present, matches and submatches are identified by byte index pairs.

### Implementation <a id="implementation"></a>

* The [`regexp`](https://golang.org/pkg/regexp/) package implements regular expressions with [RE2](https://golang.org/s/re2syntax) syntax.
* It supports UTF-8 encoded strings and Unicode character classes.
* The implementation is very efficient: the running time is linear in the size of the input.
* Backreferences are not supported since they cannot be efficiently implemented.

#### Further reading <a id="further-reading"></a>

[Regular expression matching can be simple and fast \(but is slow in Java, Perl, PHP, Python, Ruby, …\)](https://swtch.com/~rsc/regexp/regexp1.html).

##  Bitwise operators \[cheat sheet\]

### Number literals <a id="number-literals"></a>

The binary number 100002 can be written as `020`, `16` or `0x10` in Go.

| Literal | Base | Note |
| :--- | :--- | :--- |
| `020` | 8 | Starts with `0` |
| `16` | 10 | Never starts with `0` |
| `0x10` | 16 | Starts with `0x` |

### Built-in operators <a id="built-in-operators"></a>

| Operation | Result | Description |
| :--- | :--- | :--- |
| `0011 & 0101` | 0001 | Bitwise AND |
| `0011 | 0101` | 0111 | Bitwise OR |
| `0011 ^ 0101` | 0110 | Bitwise XOR |
| `^0101` | 1010 | Bitwise NOT \(same as `1111 ^ 0101`\) |
| `0011 &^ 0101` | 0010 | Bitclear \(AND NOT\) |
| `00110101<<2` | 11010100 | Left shift |
| `00110101<<100` | 00000000 | No upper limit on shift count |
| `00110101>>2` | 00001101 | Right shift |

* The binary numbers in the examples are for explanation only. Integer literals in Go must be specified in octal, decimal or hexadecimal.
* The bitwise operators take both signed and unsigned integers as input. The right-hand side of a shift operator, however, must be an unsigned integer.
* Shift operators implement arithmetic shifts if the left operand is a signed integer and logical shifts if it is an unsigned integer.

### Package [`math/bits`](https://golang.org/pkg/math/bits/) <a id="package-math-bits"></a>

| Operation | Result | Description |
| :--- | :--- | :--- |
| `bits.UintSize` | 32 or 64 | Size of a uint in bits |
| `bits.OnesCount8(00101110)` | 4 | Number of one bits \(population count\) |
| `bits.Len8(00101110)` | 6 | Bits required to represent number |
| `bits.Len8(00000000)` | 0 |  |
| `bits.LeadingZeros8(00101110)` | 2 | Number of leading zero bits |
| `bits.LeadingZeros8(00000000)` | 8 |  |
| `bits.TrailingZeros8(00101110)` | 1 | Number of trailing zero bits |
| `bits.TrailingZeros8(00000000)` | 8 |  |
| `bits.RotateLeft8(00101110, 3)` | 01110001 | The value rotated left by 3 bits |
| `bits.RotateLeft8(00101110, -3)` | 11000101 | The value rotated **right** by 3 bits |
| `bits.Reverse8(00101110)` | 01110100 | Bits in reversed order |
| `bits.ReverseBytes16(0x00ff)` | `0xff00` | Bytes in reversed order |

* The functions operate on **unsigned integers**.
* They come in different forms that take arguments of different sizes. For example, `Len`, `Len8`, `Len16`, `Len32`, and `Len64` apply to the types `uint`, `uint8`, `uint16`, `uint32`, and `uint64`, respectively.
* The functions are recognized by the compiler and on most architectures they are treated as [intrinsics](https://dave.cheney.net/2019/08/20/go-compiler-intrinsics) for additional performance.

### Bit manipulation code exampl <a id="bit-manipulation-code-example"></a>

[Bitmasks, bitsets and flags](https://yourbasic.org/golang/bitmask-flag-set-clear/) shows how to implement a bitmask, a small set of booleans, often called flags, represented by the bits in a single number.

## Start a new Go project \[standard layout\]

The repository at [github.com/yourbasic/fenwick](https://github.com/yourbasic/fenwick) is a small but complete Go library [package](https://yourbasic.org/golang/packages-explained/). It shows the structure of a basic project and can be used as a template.

In addition to source code and resources, it includes

* a [README](https://github.com/yourbasic/fenwick/blob/master/README.md) file with sections on
  * installation,
  * documentation and
  * compatibility policy,
* unit tests,
* benchmarks,
* godoc links,
* a testable doc example and
* a licence.

#### Further readin <a id="further-reading"></a>

[Your basic API](https://yourbasic.org/algorithms/your-basic-api/) is an introduction to API design with examples in Go and Java.  


## Learn to love your compiler

The Go compiler sometimes confuses and annoys developers who are new to the language.

This is a list of short articles with strategies and workarounds for common compiler error messages that tend to confuse fresh Go programmers.

* [`imported and not used`](https://yourbasic.org/golang/unused-imports/)Programs with unused imports won't compile.
* [`declared and not used`](https://yourbasic.org/golang/unused-local-variables/)You must use all local variables.
* [`multiple-value in single-value context`](https://yourbasic.org/golang/gotcha-multiple-value-single-value-context/)When a function returns multiple values, you must use all of them.
* [`syntax error: unexpected newline, expecting comma or }`](https://yourbasic.org/golang/gotcha-missing-comma-slice-array-map-literal/)In a multi-line slice, array or map literal, every line must end with a comma.
* [`cannot assign to …`](https://yourbasic.org/golang/gotcha-strings-are-immutable/)Go strings are immutable and behave like read-only byte slices.
* [`constant overflows int`](https://yourbasic.org/golang/gotcha-constant-overflows-int/)An untyped constant is converted before it is assigned to a variable.
* [`syntax error: unexpected ++, expecting expression comma or )`](https://yourbasic.org/golang/gotcha-increment-decrement-statement/)Increment and decrement operations can’t be used as expressions, only as statements.
* [`syntax error: non-declaration statement outside function body`](https://yourbasic.org/golang/short-variable-declaration-outside-function/)Short variable declarations can only be used inside functions.
* [`missing function body for …`](https://yourbasic.org/golang/opening-brace-separate-line/)An opening brace cannot appear on a line by itself.

#### Further reading

[Tutorials](https://yourbasic.org/golang/tutorials/) for beginners and experienced developers alike: best practices and production-quality code examples.



