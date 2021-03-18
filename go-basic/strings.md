# Strings

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

### Further reading <a id="further-reading"></a>

[40+ practical string tips \[cheat sheet\]](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)

## Runes and character encoding

### Characters, ASCII and Unicode <a id="characters-ascii-and-unicode"></a>

> The `rune` type is an alias for `int32`, and is used to emphasize than an integer represents a code point.

**ASCII** defines 128 characters, identified by the **code points** 0–127. It covers English letters, Latin numbers, and a few other characters.

**Unicode**, which is a superset of ASCII, defines a codespace of 1,114,112 code points. Unicode version 10.0 covers 139 modern and historic scripts \(including the runic alphabet, but not Klingon\) as well as multiple symbol sets.

### Strings and UTF-8 encoding <a id="strings-and-utf-8-encoding"></a>

> A `string` is a sequence of bytes, not runes.

However, strings often contain Unicode text encoded in [UTF-8](https://research.swtch.com/utf8), which encodes all Unicode code points using one to four bytes. \(ASCII characters are encoded with one byte, while other code points use more.\)

Since Go source code itself is encoded as UTF-8, string literals will automatically get this encoding.

For example, in the string `"café"` the character `é` \(code point 233\) is encoded using two bytes, while the ASCII characters `c`, `a` and `f` \(code points 99, 97 and 102\) only use one:

```text
fmt.Println([]byte("café")) // [99 97 102 195 169]
fmt.Println([]rune("café")) // [99 97 102 233]
```

See [Convert between byte array/slice and string](https://yourbasic.org/golang/convert-string-to-byte-slice/) and [Convert between rune array/slice and string](https://yourbasic.org/golang/convert-string-to-rune-slice/).

#### Further reading <a id="further-reading"></a>

[Escapes and multiline strings](https://yourbasic.org/golang/multiline-string/)

### Clean and simple string building <a id="clean-and-simple-string-building"></a>

For simple cases where performance is a non-issue, [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend. It’s clean, simple and fairly efficient.

```text
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

The [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) lists the most common formatting verbs and flags.

## Efficient string concatenation \[full guide\]



### Clean and simple string building <a id="clean-and-simple-string-building"></a>

For simple cases where performance is a non-issue, [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) is your friend. It’s clean, simple and fairly efficient.

```text
s := fmt.Sprintf("Size: %d MB.", 85) // s == "Size: 85 MB."
```

The [fmt cheat sheet](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) lists the most common formatting verbs and flags.

### High-performance string concatenation[Go 1.10](https://golang.org/doc/go1.10) <a id="high-performance-string-concatenation"></a>

A [`strings.Builder`](https://golang.org/pkg/strings/#Builder) is used to efficiently append strings using write methods.

* It offers a subset of the [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) methods that allows it to safely avoid extra copying when converting a builder to a string.
* You can use the [`fmt`](https://golang.org/pkg/fmt/) package for formatting since the builder implements the [`io.Writer`](https://yourbasic.org/golang/io-writer-interface-explained/) interface.
* The [`Grow`](https://golang.org/pkg/strings/#Builder.Grow) method can be used to preallocate memory when the maximum size of the string is known.

```text
var b strings.Builder
b.Grow(32)
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&b, "%d:%d, ", i+1, p)
}
s := b.String()   // no copying
s = s[:b.Len()-2] // no copying (removes trailing ", ")
fmt.Println(s)
```

```text
1:2, 2:3, 3:5, 4:7, 5:11, 6:13
```

### Before Go 1.10 <a id="before-go-1-10"></a>

Use [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) to print into a [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer).

```text
var buf bytes.Buffer
for i, p := range []int{2, 3, 5, 7, 11, 13} {
    fmt.Fprintf(&buf, "%d:%d, ", i+1, p)
}
buf.Truncate(buf.Len() - 2) // Remove trailing ", "
s := buf.String()           // Copy into a new string
fmt.Println(s)
```

```text
1:2, 2:3, 3:5, 4:7, 5:11, 6:13
```

This solution is pretty efficient but may generate some excess garbage. For higher performance, you can try to use the append functions in package [`strconv`](https://golang.org/pkg/strconv/).

```text
buf := []byte("Size: ")
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```

If the expected maximum length of the string is known, you may want to preallocate the slice.

```text
buf := make([]byte, 0, 16)
buf = append(buf, "Size: "...)
buf = strconv.AppendInt(buf, 85, 10)
buf = append(buf, " MB."...)
s := string(buf)
```

## Escapes and multiline strings

### Raw string literals <a id="raw-string-literals"></a>

Raw [string literals](https://golang.org/ref/spec#String_literals), delimited by **backticks** \(back quotes\), are interpreted literally. They can contain line breaks, and backslashes have no special meaning.

```text
const s = `First line
Second line`
fmt.Println(s)
```

```text
First line
Second line
```

#### Backtick escape <a id="backtick-escape"></a>

It’s [not possible](https://github.com/golang/go/issues/24475) to include a backtick in a raw string literal, but you can do

```text
fmt.Println("`" + "foo" + "`") // Prints: `foo`
```

### Interpreted string literals <a id="interpreted-string-literals"></a>

To insert escape characters, use interpreted string literals delimited by **double quotes**.

```text
const s = "\tFirst line\n" +
"Second line"
fmt.Println(s)
```

```text
   First line
Second line
```

The escape character `\t` denotes a horizontal tab and `\n` is a line feed or newline.

#### Double quote escape <a id="double-quote-escape"></a>

Use `\"` to insert a double quote in an interpreted string literal:

```text
fmt.Println("\"foo\"") // Prints: "foo"
```

### Escape HTML <a id="escape-html"></a>

Use [`html.EscpapeString`](https://golang.org/pkg/html/#EscapeString) to encode a string so that it can be safely placed inside HTML text. The function escapes the five characters `<`, `>`, `&`, `'` and `"`.

```text
const s = `"Foo's Bar" <foobar@example.com>`
fmt.Println(html.EscapeString(s))
```

```text
&#34;Foo&#39;s Bar&#34; &lt;foobar@example.com&gt;
```

[`html.UnescapeString`](https://golang.org/pkg/html/#UnescapeString) does the inverse transformation.

### Escape URL <a id="escape-url"></a>

Use [`url.PathEscape`](https://golang.org/pkg/net/url/#PathEscape) in package [`net/url`](https://golang.org/pkg/net/url/) to encode a string so it can be safely placed inside a URL. The function uses [percent-encoding](https://en.wikipedia.org/wiki/Percent-encoding).

```text
const s = `Foo's Bar?`
fmt.Println(url.PathEscape(s))
```

```text
Foo%27s%20Bar%3F
```

[`url.PathUnescape`](https://golang.org/pkg/net/url/#PathUnescape) does the inverse transformation.

### All escape characters <a id="all-escape-characters"></a>

Arbitrary character values can be encoded with backslash escapes and used in string or [rune literals](https://golang.org/ref/spec#Rune_literals). There are four different formats:

* `\x` followed by exactly two hexadecimal digits,
* `\` followed by exactly three octal digits,
* `\u` followed by exactly four hexadecimal digits,
* `\U` followed by exactly eight hexadecimal digits,

where the escapes `\u` and `\U` represent Unicode code points.

The following special escape values are also available.

| Value | Description |
| :--- | :--- |
| `\a` | Alert or bell |
| `\b` | Backspace |
| `\\` | Backslash |
| `\t` | Horizontal tab |
| `\n` | Line feed or newline |
| `\f` | Form feed |
| `\r` | Carriage return |
| `\v` | Vertical tab |
| `\'` | Single quote \(only in rune literals\) |
| `\"` | Double quote \(only in string literals\) |

```text
fmt.Println("\\caf\u00e9") // Prints string: \café
fmt.Printf("%c", '\u00e9') // Prints character: é
```

## 3 ways to split a string into a slice

### Split on comma or other substring <a id="split-on-comma-or-other-substring"></a>

Use the [`strings.Split`](https://golang.org/pkg/strings/#Split) function to split a string into its comma separated values.

```text
s := strings.Split("a,b,c", ",")
fmt.Println(s)
// Output: [a b c]
```

To include the separators, use [`strings.SplitAfter`](https://golang.org/pkg/strings/#SplitAfter). To split only the first n values, use [`strings.SplitN`](https://golang.org/pkg/strings/#SplitN) and [`strings.SplitAfterN`](https://golang.org/pkg/strings/#SplitAfterN).

You can use [`strings.TrimSpace`](https://yourbasic.org/golang/trim-whitespace-from-string/) to strip leading and trailing whitespace from the resulting strings.

### Split by whitespace and newline <a id="split-by-whitespace-and-newline"></a>

Use the [`strings.Fields`](https://golang.org/pkg/strings/#Fields) function to split a string into substrings removing any space characters, including newlines.

```text
s := strings.Fields(" a \t b \n")
fmt.Println(s)
// Output: [a b]
```

### Split on regular expression <a id="split-on-regular-expression"></a>

In more complicated situations, the regexp [`Split`](https://yourbasic.org/golang/regexp-cheat-sheet/#split) method might do the trick.

It splits a string into substrings separated by a regular expression. The method takes an integer argument `n`; if `n >= 0`, it returns at most `n` substrings.

```text
a := regexp.MustCompile(`a`)              // a single `a`
fmt.Printf("%q\n", a.Split("banana", -1)) // ["b" "n" "n" ""]
fmt.Printf("%q\n", a.Split("banana", 0))  // [] (nil slice)
fmt.Printf("%q\n", a.Split("banana", 1))  // ["banana"]
fmt.Printf("%q\n", a.Split("banana", 2))  // ["b" "nana"]

zp := regexp.MustCompile(` *, *`)             // spaces and one comma
fmt.Printf("%q\n", zp.Split("a,b ,  c ", -1)) // ["a" "b" "c "]
```

See this [Regexp tutorial and cheat sheet](https://yourbasic.org/golang/regexp-cheat-sheet/) for a gentle introduction to the Go regexp package with plenty of examples.

## Convert between rune array/slice and string

### Convert string to runes <a id="convert-string-to-runes"></a>

* When you convert a string to a rune slice, you get a new slice that contains the [Unicode code points](https://yourbasic.org/golang/rune/) \(runes\) of the string.
* For an invalid UTF-8 sequence, the rune value will be `0xFFFD` for each invalid byte.

```text
r := []rune("ABC€")
fmt.Println(r)        // [65 66 67 8364]
fmt.Printf("%U\n", r) // [U+0041 U+0042 U+0043 U+20AC]
```

> You can also use a [range loop](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) to access the code points of a string.

### Convert runes to string <a id="convert-runes-to-string"></a>

* When you convert a slice of runes to a string, you get a new string that is the concatenation of the runes converted to UTF-8 encoded strings.
* Values outside the range of valid Unicode code points are converted to `\uFFFD`, the Unicode replacement character `�`.

```text
s := string([]rune{'\u0041', '\u0042', '\u0043', '\u20AC', -1})
fmt.Println(s) // ABC€�
```

### Performance <a id="performance"></a>

These conversions create a new slice or string, and therefore have [time complexity](https://yourbasic.org/algorithms/time-complexity-explained/) proportional to the number of bytes that are processed.

#### More efficient alternative <a id="more-efficient-alternative"></a>

In some cases, you might be able to use a [string builder](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/), which can concatenate strings without redundant copying:

[Efficient string concatenation \[full guide\]](https://yourbasic.org/golang/build-append-concatenate-strings-efficiently/)

## Convert between float and string

### String to float <a id="string-to-float"></a>

Use the [`strconv.ParseFloat`](https://golang.org/pkg/strconv/#ParseFloat) function to parse a string as a floating-point number with the precision specified by `bitSize`: 32 for `float32`, or 64 for `float64`.

```text
func ParseFloat(s string, bitSize int) (float64, error)
```

When `bitSize` is 32, the result still has type `float64`, but it will be convertible to `float32` without changing its value.

```text
f := "3.14159265"
if s, err := strconv.ParseFloat(f, 32); err == nil {
    fmt.Println(s) // 3.1415927410125732
}
if s, err := strconv.ParseFloat(f, 64); err == nil {
    fmt.Println(s) // 3.14159265
}
```

### Float to string <a id="float-to-string"></a>

Use the [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) method to format a floating-point number as a string.

```text
s := fmt.Sprintf("%f", 123.456) // s == "123.456000"
```

| Formatting | Description | Verb |
| :--- | :--- | :--- |
| 1.234560e+02 | Scientific notation | `%e` |
| 123.456000 | Decimal point, no exponent | `%f` |
| 123.46 | Default width, precision 2 | `%.2f` |
| ␣␣123.46 | Width 8, precision 2 | `%8.2f` |
| 123.456 | Exponent as needed, necessary digits only | `%g` |

## Convert between int, int64 and string

### int/int64 to string <a id="int-int64-to-string"></a>

Use [`strconv.Itoa`](https://golang.org/pkg/strconv/#Itoa) to convert an int to a decimal string.

```text
s := strconv.Itoa(97) // s == "97"
```

> **Warning:** In a plain [conversion](https://yourbasic.org/golang/conversions/) the value is interpreted as a Unicode code point, and the resulting string will contain the character represented by that code point, encoded in UTF-8.
>
> ```text
> s := string(97) // s == "a"
> ```

Use [`strconv.FormatInt`](https://golang.org/pkg/strconv/#FormatInt) to format an int64 in a given base.

```text
var n int64 = 97
s := strconv.FormatInt(n, 10) // s == "97" (decimal)
```

```text
var n int64 = 97
s := strconv.FormatInt(n, 16) // s == "61" (hexadecimal)
```

### string to int/int64 <a id="string-to-int-int64"></a>

Use [`strconv.Atoi`](https://golang.org/pkg/strconv/#Atoi) to parse a decimal string to an int.

```text
s := "97"
if n, err := strconv.Atoi(s); err == nil {
    fmt.Println(n+1)
} else {
    fmt.Println(s, "is not an integer.")
}
// Output: 98
```

Use [`strconv.ParseInt`](https://golang.org/pkg/strconv/#ParseInt) to parse a decimal string \(base `10`\) and check if it fits into an int64.

```text
s := "97"
n, err := strconv.ParseInt(s, 10, 64)
if err == nil {
    fmt.Printf("%d of type %T", n, n)
}
// Output: 97 of type int64
```

The two numeric arguments represent a base \(0, 2 to 36\) and a bit size \(0 to 64\).

If the first argument is 0, the base is implied by the string’s prefix: base 16 for `"0x"`, base 8 for `"0"`, and base 10 otherwise.

The second argument specifies the integer type that the result must fit into. Bit sizes 0, 8, 16, 32, and 64 correspond to `int`, `int8`, `int16`, `int32`, and `int64`.

### int to int64 \(and back\) <a id="int-to-int64-and-back"></a>

The size of an `int` is implementation-specific, it’s either 32 or 64 bits, and hence you won’t lose any information when converting from int to  int64.

```text
var n int = 97
m := int64(n) // safe
```

> However, when converting to a shorter integer type, the value is **truncated** to fit in the result type's size.

```text
var m int64 = 2 << 32
n := int(m)    // truncated on machines with 32-bit ints
fmt.Println(n) // either 0 or 4,294,967,296
```

* See [Maximum value of an int](https://yourbasic.org/golang/max-min-int-uint/) for code to compute the size of an `int`.
* See [Pick the right one: int vs. int64](https://yourbasic.org/golang/int-vs-int64/) for best practices.

### General formatting \(width, indent, sign\) <a id="general-formatting-width-indent-sign"></a>

The [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) function is a useful general tool for converting data to string:

```text
s := fmt.Sprintf("%+8d", 97)
// s == "     +97" (width 8, right justify, always show sign)
```

## Convert interface to string

Use [`fmt.Sprintf`](https://golang.org/pkg/fmt/#Sprintf) to convert an [interface value](https://yourbasic.org/golang/interfaces-explained/) to a string.

```text
var x interface{} = "abc"
str := fmt.Sprintf("%v", x)
```

In fact, the same technique can be used to get a string representation of any data structure.

```text
var x interface{} = []int{1, 2, 3}
str := fmt.Sprintf("%v", x)
fmt.Println(str) // "[1 2 3]"

```

## Remove all duplicate whitespace



```text
space := regexp.MustCompile(`\s+`)
s := space.ReplaceAllString("Hello  \t \n world!", " ")
fmt.Printf("%q", s) // "Hello world!"
```

`\s+` is a [regular expression](https://yourbasic.org/golang/regexp-cheat-sheet/):

* the character class `\s` matches a space, tab, new line, carriage return or form feed,
* and `+` says “one or more of those”.

In other words, the code will replace each whitespace substring with a single space character.

#### Trim leading and trailing space <a id="trim-leading-and-trailing-space"></a>

To [trim leading and trailing whitespace](https://yourbasic.org/golang/trim-whitespace-from-string/), use the `strings.TrimSpace` function.

## 3 ways to trim whitespace \(or other characters\) from a string

Use the [`strings.TrimSpace`](https://golang.org/pkg/strings/#TrimSpace) function to remove leading and trailing whitespace as defined by Unicode.

```text
s := strings.TrimSpace("\t Goodbye hair!\n ")
fmt.Printf("%q", s) // "Goodbye hair!"
```

* To remove other leading and trailing characters, use [`strings.Trim`](https://golang.org/pkg/strings/#Trim).
* To remove only the leading or the trailing characters, use [`strings.TrimLeft`](https://golang.org/pkg/strings/#TrimLeft) or [`strings.TrimRight`](https://golang.org/pkg/strings/#TrimRight).

## How to reverse a string by byte or rune

### Byte by byte <a id="byte-by-byte"></a>

It’s pretty straightforward to reverse a string one byte at a time.

```text
// Reverse returns a string with the bytes of s in reverse order.
func Reverse(s string) string {
    var b strings.Builder
    b.Grow(len(s))
    for i := len(s) - 1; i >= 0; i-- {
        b.WriteByte(s[i])
    }
    return b.String()
}
```

### Rune by rune <a id="rune-by-rune"></a>

To reverse a string by UTF-8 encoded characters is a bit trickier.

```text
// ReverseRune returns a string with the runes of s in reverse order.
// Invalid UTF-8 sequences, if any, will be reversed byte by byte.
func ReverseRune(s string) string {
    res := make([]byte, len(s))
    prevPos, resPos := 0, len(s)
    for pos := range s {
        resPos -= pos - prevPos
        copy(res[resPos:], s[prevPos:pos])
        prevPos = pos
    }
    copy(res[0:], s[prevPos:])
    return string(res)
}
```

#### Example usage <a id="example-usage"></a>

```text
for _, s := range []string{
	"Ångström",
	"Hello, 世界",
	"\xff\xfe\xfd", // invalid UTF-8
} {
	fmt.Printf("%q\n", ReverseRune(s))
}
```

```text
"mörtsgnÅ"
"界世 ,olleH"
"\xfd\xfe\xff"
```



