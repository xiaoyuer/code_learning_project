# Scripting

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

## Read a file \(stdin\) line by line

Use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to read a file line by line.

```text
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

### Read from stdin <a id="read-from-stdin"></a>

Use [`os.Stdin`](https://golang.org/pkg/os/#pkg-variables) to read from the standard input stream.

```text
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	log.Println(err)
}
```

### Read from any stream <a id="read-from-any-stream"></a>

A bufio.Scanner can read from any stream of bytes, as long as it implements the [`io.Reader`](https://golang.org/pkg/io/#Reader) interface. See [How to use the io.Reader interface](https://yourbasic.org/golang/io-reader-interface-explained/).

#### Further reading <a id="further-reading"></a>

For more advanced scanning, see the examples in the [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) documentation.  


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

## Command-line arguments and flags

## The [`os.Args`](https://golang.org/pkg/os/#pkg-variables) variable holds the command-line arguments – starting with the program name – which are passed to a Go program.

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

## Access environment variables

```text
key:"SHELL" value:"/bin/bash"
key:"SESSION" value:"ubuntu"
key:"TERM" value:"xterm-256color"
key:"LANG" value:"en_US.UTF-8"
key:"XMODIFIERS" value:"@im=ibus"
…
```

```text
for _, s := range os.Environ() {
    kv := strings.SplitN(s, "=", 2) // unpacks "key=value"
    fmt.Printf("key:%q value:%q\n", kv[0], kv[1])
}
```

The [`os.Environ`](https://golang.org/pkg/os/#Environ) function returns a slice of `"key=value"` strings listing all environment variables.

```text
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/bash"

os.Unsetenv("SHELL")
fmt.Printf("%q\n", os.Getenv("SHELL")) // ""

os.Setenv("SHELL", "/bin/dash")
fmt.Printf("%q\n", os.Getenv("SHELL")) // "/bin/dash"
```

Use the [`os.Setenv`](https://golang.org/pkg/os/#Setenv), [`os.Getenv`](https://golang.org/pkg/os/#Getenv) and [`os.Unsetenv`](https://golang.org/pkg/os/#Unsetenv) functions to access environment variables.

##  

