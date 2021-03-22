# Standard Library

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

## How to use the io.Reader interface

### Basics <a id="basics"></a>

The [`io.Reader`](https://golang.org/pkg/io/#Reader) [interface](https://yourbasic.org/golang/interfaces-explained/) represents an entity from which you can read a stream of bytes.

```text
type Reader interface {
        Read(buf []byte) (n int, err error)
}
```

`Read` reads up to `len(buf)` bytes into `buf` and returns the number of bytes read – it returns an [`io.EOF`](https://golang.org/pkg/io/#pkg-variables) error when the stream ends.

The standard library provides numerous Reader [implementations](https://golang.org/search?q=Read#Global) \(including in-memory byte buffers, files and network connections\), and Readers are accepted as input by many utilities \(including the HTTP client and server implementations\).

### Use a built-in reader <a id="use-a-built-in-reader"></a>

As an example, you can create a Reader from a string using the [`strings.Reader`](https://golang.org/pkg/strings/#Reader) function and then pass the Reader directly to the [`http.Post`](https://golang.org/pkg/net/http/#Post) function in package [`net/http`](https://golang.org/pkg/net/http/). The Reader is then used as the source for the data to be posted.

```text
r := strings.NewReader("my request")
resp, err := http.Post("http://foo.bar",
	"application/x-www-form-urlencoded", r)
```

Since `http.Post` uses a Reader instead of a `[]byte` it’s trivial to, for instance, use the contents of a file instead.

### Read directly from a byte stream <a id="read-directly-from-a-byte-stream"></a>

You can use the `Read` function directly \(this is the least common use case\).

```text
r := strings.NewReader("abcde")

buf := make([]byte, 4)
for {
	n, err := r.Read(buf)
	fmt.Println(n, err, buf[:n])
	if err == io.EOF {
		break
	}
}
```

```text
4 <nil> [97 98 99 100]
1 <nil> [101]
0 EOF []
```

Use [`io.ReadFull`](https://golang.org/pkg/io/#ReadFull) to read exactly `len(buf)` bytes into `buf`:

```text
r := strings.NewReader("abcde")

buf := make([]byte, 4)
if _, err := io.ReadFull(r, buf); err != nil {
	log.Fatal(err)
}
fmt.Println(buf)

if _, err := io.ReadFull(r, buf); err != nil {
	fmt.Println(err)
}
```

```text
[97 98 99 100]
unexpected EOF
```

Use [`ioutil.ReadAll`](https://golang.org/pkg/io/ioutil/#ReadAll) to read everything:

```text
r := strings.NewReader("abcde")

buf, err := ioutil.ReadAll(r)
if err != nil {
	log.Fatal(err)
}
fmt.Println(buf)
```

```text
[97 98 99 100 101]
```

### Buffered reading and scanning <a id="buffered-reading-and-scanning"></a>

The [`bufio.Reader`](https://golang.org/pkg/bufio/#Reader) and [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) types wrap a Reader creating another Reader that also implements the interface but provides buffering and some help for textual input.

In this example we use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to count the number of words in a text.

```text
const input = `Beware of bugs in the above code;
I have only proved it correct, not tried it.`

scanner := bufio.NewScanner(strings.NewReader(input))
scanner.Split(bufio.ScanWords) // Set up the split function.

count := 0
for scanner.Scan() {
    count++
}
if err := scanner.Err(); err != nil {
    fmt.Println(err)
}
fmt.Println(count)
```

```text
16
```

## How to use the io.Writer interface

### Basics <a id="basics"></a>

The [`io.Writer`](https://golang.org/pkg/io/#Writer) [interface](https://yourbasic.org/golang/interfaces-explained/) represents an entity to which you can write a stream of bytes.

```text
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

`Write` writes up to `len(p)` bytes from `p` to the underlying data stream – it returns the number of bytes written and any error encountered that caused the write to stop early.

The standard library provides numerous Writer [implementations](https://golang.org/search?q=Write#Global), and Writers are accepted as input by many utilities.

### How to use a built-in writer \(3 examples\) <a id="how-to-use-a-built-in-writer-3-nbsp-examples"></a>

As a first example, you can write directly into a [`bytes.Buffer`](https://golang.org/pkg/bytes/#Buffer) using the [`fmt.Fprintf`](https://golang.org/pkg/fmt/#Fprintf) function. This works since

* `bytes.Buffer` has a `Write` method, and
* `fmt.Fprintf` takes a `Writer` as its first argument.

```text
var buf bytes.Buffer
fmt.Fprintf(&buf, "Size: %d MB.", 85)
s := buf.String()) // s == "Size: 85 MB."
```

Similarly, you can write directly into files or other streams, such as http connections. See the [HTTP server example](https://yourbasic.org/golang/http-server-example/) article for a complete code example.

This is a very common pattern in Go. As yet another example, you can compute the hash value of a file by copying the file into the `io.Writer` function of a suitable [`hash.Hash`](https://golang.org/pkg/hash/#Hash) object. See [Hash checksums](https://yourbasic.org/golang/hash-md5-sha256-string-file/#file) for code.

### Optimize string writes <a id="optimize-string-writes"></a>

Some Writers in the standard library have an additional `WriteString` method. This method can be more efficient than the standard `Write` method since it writes a string directly without allocating a byte slice.

You can take direct advantage of this optimization by using the [`io.WriteString()`](https://golang.org/pkg/io/#WriteString) function.

```text
func WriteString(w Writer, s string) (n int, err error)
```

If `w` implements a `WriteString` method, it is invoked directly. Otherwise, `w.Write` is called exactly once.

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

## Write log to file \(or /dev/null\)

This code appends a log message to the file `text.log`. It creates the file if it doesn’t already exist.

```text
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()

logger := log.New(f, "prefix", log.LstdFlags)
logger.Println("text to append")
logger.Println("more text to append")
```

Contents of `text.log`:

```text
prefix: 2017/10/20 07:52:58 text to append
prefix: 2017/10/20 07:52:58 more text to append
```

* [`log.New`](https://golang.org/pkg/log/#New) creates a new [`log.Logger`](https://golang.org/pkg/log/#Logger) that writes to `f`.
* The prefix appears at the beginning of each generated log line.
* The [`flag`](https://golang.org/pkg/log/#pkg-constants) argument defines which text to prefix to each log entry.

### Disable logging <a id="disable-logging"></a>

To turn off all output from a [`log.Logger`](https://golang.org/pkg/log/#Logger), set the output destination to [`ioutil.Discard`](https://golang.org/pkg/io/ioutil/#pkg-variables), a writer on which all calls succeed without doing anything.

```text
log.SetOutput(ioutil.Discard)
```

## Hello world HTTP server example

### A basic web server <a id="a-basic-web-server"></a>

If you access the URL `http://localhost:8080/world` on a machine where the program below is running, you will be greeted by this page.

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



