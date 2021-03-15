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

