# Go Tutorial

## Go beginner’s guide: top 4 resources to get you started\[already known\]

## How to use JSON with Go \[best practices\]

### Default types <a id="default-types"></a>

The default Go types for decoding and encoding JSON are

* `bool` for JSON booleans,
* `float64` for JSON numbers,
* `string` for JSON strings, and
* `nil` for JSON null.

Additionally, [`time.Time`](https://golang.org/pkg/time/#Time) and the numeric types in the [`math/big`](https://golang.org/pkg/math/big/) package can be automatically encoded as JSON strings.

Note that JSON doesn’t support basic integer types. They can often be approximated by floating-point numbers.

> Since software that implements IEEE 754-2008 binary64 \(double precision\) numbers is generally available and widely used, good interoperability can be achieved by implementations that expect no more precision or range than these provide \[…\]
>
> Note that when such software is used, numbers that are integers and are in the range \[-253 + 1, 253 - 1\] are interoperable in the sense that implementations will agree exactly on their numeric values.[RFC 7159: The JSON Data Interchange Format](https://tools.ietf.org/html/rfc7159#section-6)

### Encode \(marshal\) struct to JSON <a id="encode-marshal-struct-to-json"></a>

The [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) function in package [`encoding/json`](https://golang.org/pkg/encoding/json/) generates JSON data.

```text
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64  `json:"ref"`
    private string // An unexported field is not encoded.
    Created time.Time
}

basket := FruitBasket{
    Name:    "Standard",
    Fruit:   []string{"Apple", "Banana", "Orange"},
    Id:      999,
    private: "Second-rate",
    Created: time.Now(),
}

var jsonData []byte
jsonData, err := json.Marshal(basket)
if err != nil {
    log.Println(err)
}
fmt.Println(string(jsonData))
```

Output:

```text
{"Name":"Standard","Fruit":["Apple","Banana","Orange"],"ref":999,"Created":"2018-04-09T23:00:00Z"}
```

Only data that can be represented as JSON will be encoded; see [`json.Marshal`](https://golang.org/pkg/encoding/json/#Marshal) for the complete rules.

* Only the exported \(public\) fields of a struct will be present in the JSON output. **Other fields are ignored**.
* A field with a `json:` **struct tag** is stored with its tag name instead of its variable name.
* Pointers will be encoded as the values they point to, or `null` if the pointer is `nil`.

### Pretty print <a id="pretty-print"></a>

Replace `json.Marshal` with [`json.MarshalIndent`](https://golang.org/pkg/encoding/json/#MarshalIndent) in the example above to indent the JSON output.

```text
jsonData, err := json.MarshalIndent(basket, "", "    ")
```

Output:

```text
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}
```

### Decode \(unmarshal\) JSON to struct <a id="decode-unmarshal-json-to-struct"></a>

The [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) function in package [`encoding/json`](https://golang.org/pkg/encoding/json/) parses JSON data.

```text
type FruitBasket struct {
    Name    string
    Fruit   []string
    Id      int64 `json:"ref"`
    Created time.Time
}

jsonData := []byte(`
{
    "Name": "Standard",
    "Fruit": [
        "Apple",
        "Banana",
        "Orange"
    ],
    "ref": 999,
    "Created": "2018-04-09T23:00:00Z"
}`)

var basket FruitBasket
err := json.Unmarshal(jsonData, &basket)
if err != nil {
    log.Println(err)
}
fmt.Println(basket.Name, basket.Fruit, basket.Id)
fmt.Println(basket.Created)
```

Output:

```text
Standard [Apple Banana Orange] 999
2018-04-09 23:00:00 +0000 UTC
```

Note that `Unmarshal` allocated a new slice all by itself. This is how unmarshaling works for slices, maps and pointers.

For a given JSON key `Foo`, `Unmarshal` will attempt to match the struct fields in this order:

1. an exported \(public\) field with a struct tag `json:"Foo"`,
2. an exported field named `Foo`, or
3. an exported field named `FOO`, `FoO`, or some other case-insensitive match.

Only fields thar are found in the destination type will be decoded:

* This is useful when you wish to pick only a few specific fields.
* In particular, any unexported fields in the destination struct will be unaffected.

### Arbitrary objects and arrays <a id="arbitrary-objects-and-arrays"></a>

The [encoding/json](https://golang.org/pkg/encoding/json/) package uses

* `map[string]interface{}` to store arbitrary JSON objects, and
* `[]interface{}` to store arbitrary JSON arrays.

It will unmarshal any valid JSON data into a plain `interface{}` value.

Consider this JSON data:

```text
{
    "Name": "Eve",
    "Age": 6,
    "Parents": [
        "Alice",
        "Bob"
    ]
}
```

The [`json.Unmarshal`](https://golang.org/pkg/encoding/json/#Unmarshal) function will parse it into a map whose keys are strings, and whose values are themselves stored as empty interface values:

```text
map[string]interface{}{
    "Name": "Eve",
    "Age":  6.0,
    "Parents": []interface{}{
        "Alice",
        "Bob",
    },
}
```

We can iterate through the map with a range statement and use a type switch to access its values.

```text
jsonData := []byte(`{"Name":"Eve","Age":6,"Parents":["Alice","Bob"]}`)

var v interface{}
json.Unmarshal(jsonData, &v)
data := v.(map[string]interface{})

for k, v := range data {
    switch v := v.(type) {
    case string:
        fmt.Println(k, v, "(string)")
    case float64:
        fmt.Println(k, v, "(float64)")
    case []interface{}:
        fmt.Println(k, "(array):")
        for i, u := range v {
            fmt.Println("    ", i, u)
        }
    default:
        fmt.Println(k, v, "(unknown)")
    }
}
```

Output:

```text
Name Eve (string)
Age 6 (float64)
Parents (array):
     0 Alice
     1 Bob
```

### JSON file example <a id="json-file-example"></a>

The [`json.Decoder`](https://golang.org/pkg/encoding/json/#Decoder) and [`json.Encoder`](https://golang.org/pkg/encoding/json/#Encoder) types in package [`encoding/json`](https://golang.org/pkg/encoding/json/) offer support for reading and writing streams, e.g. files, of JSON data.

The code in this example

* reads a stream of JSON objects from a [Reader](https://yourbasic.org/golang/io-reader-interface-explained/) \([`strings.Reader`](https://golang.org/pkg/strings/#Reader)\),
* removes the `Age` field from each object,
* and then writes the objects to a [Writer](https://yourbasic.org/golang/io-writer-interface-explained/) \([`os.Stdout`](https://golang.org/pkg/os/#pkg-variables)\).

```text
const jsonData = `
    {"Name": "Alice", "Age": 25}
    {"Name": "Bob", "Age": 22}
`
reader := strings.NewReader(jsonData)
writer := os.Stdout

dec := json.NewDecoder(reader)
enc := json.NewEncoder(writer)

for {
    // Read one JSON object and store it in a map.
    var m map[string]interface{}
    if err := dec.Decode(&m); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }

    // Remove all key-value pairs with key == "Age" from the map.
    for k := range m {
        if k == "Age" {
            delete(m, k)
        }
    }

    // Write the map as a JSON object.
    if err := enc.Encode(&m); err != nil {
        log.Println(err)
    }
}
```

Output:

```text
{"Name":"Alice"}
{"Name":"Bob"}
```

#### Further reading

[Tutorials](https://yourbasic.org/golang/tutorials/) for beginners and experienced developers alike: best practices and production-quality code examples.

## Regexp tutorial and cheat sheet

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

