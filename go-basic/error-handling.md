# Error Handling

## Error handling best practice

Go has two different error-handling mechanisms:

* most functions return [**errors**](https://yourbasic.org/golang/create-error/);
* only a truly unrecoverable condition, such as an out-of-range index, produces a run-time exception, known as a [**panic**](https://yourbasic.org/golang/recover-from-panic/).

Go’s multivalued return makes it easy to return a detailed error message alongside the normal return value. By convention, such messages have type `error`, a simple built-in [interface](https://yourbasic.org/golang/interfaces-explained/):

```text
type error interface {
    Error() string
}
```

#### Error handling example <a id="error-handling-example"></a>

The `os.Open` function returns a non-nil `error` value when it fails to open a file.

```text
func Open(name string) (file *File, err error)
```

The following code uses `os.Open` to open a file. If an `error` occurs it calls `log.Fatal` to print the error message and stop.

```text
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
// do something with the open *File f
```

### Custom errors <a id="custom-errors"></a>

To create a simple string-only `error` you can use [`errors.New`](https://golang.org/pkg/errors/#New):

```text
err := errors.New("Houston, we have a problem")
```

The `error` interface requires only an `Error` method, but specific `error` implementations often have additional methods, allowing callers to inspect the details of the error.

#### Learn more <a id="learn-more"></a>

See [3 simple ways to create an error](https://yourbasic.org/golang/create-error/) for more examples.

### Panic <a id="panic"></a>

Panics are similar to C++ and Java exceptions, but are only intended for run-time errors, such as following a nil pointer or attempting to index an array out of bounds.

#### Learn more <a id="learn-more-1"></a>

See [Recover from a panic](https://yourbasic.org/golang/recover-from-panic/) for a tutorial on how to recover from and test panics.

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

#### More code examples

[Go blueprints: code for com­mon tasks](https://yourbasic.org/golang/blueprint/) is a collection of handy code examples.

## Panics, stack traces and how to recover \[best practice\]

### A panic is an exception in Go <a id="a-panic-is-an-exception-in-go"></a>

Panics are similar to C++ and Java exceptions, but are only intended for run-time errors, such as following a nil pointer or attempting to index an array out of bounds. To signify events such as end-of-file, Go programs use the built-in `error` type. See [Error handling best practice](https://yourbasic.org/golang/errors-explained/) and [3 simple ways to create an error](https://yourbasic.org/golang/create-error/) for more on errors.

A panic stops the normal execution of a goroutine:

* When a program panics, it immediately starts to unwind the call stack.
* This continues until the program crashes and prints a stack trace,
* or until the built-in `recover` function is called.

A panic is caused either by a runtime error, or an explicit call to the built-in `panic` function.

### Stack traces <a id="stack-traces"></a>

A **stack trace** – a report of all active stack frames – is typically printed to the console when a panic occurs. Stack traces can be very useful for debugging:

* not only do you see **where** the error happened,
* but also **how** the program arrived in this place.

#### Interpret a stack trace <a id="interpret-a-stack-trace"></a>

Here’s an example of a stack trace:

```text
goroutine 11 [running]:
testing.tRunner.func1(0xc420092690)
    /usr/local/go/src/testing/testing.go:711 +0x2d2
panic(0x53f820, 0x594da0)
    /usr/local/go/src/runtime/panic.go:491 +0x283
github.com/yourbasic/bit.(*Set).Max(0xc42000a940, 0x0)
    ../src/github.com/bit/set_math_bits.go:137 +0x89
github.com/yourbasic/bit.TestMax(0xc420092690)
    ../src/github.com/bit/set_test.go:165 +0x337
testing.tRunner(0xc420092690, 0x57f5e8)
    /usr/local/go/src/testing/testing.go:746 +0xd0
created by testing.(*T).Run
    /usr/local/go/src/testing/testing.go:789 +0x2de
```

It can be read from the bottom up:

* `testing.(*T).Run` has called `testing.tRunner`,
* which has called `bit.TestMax`,
* which has called `bit.(*Set).Max`,
* which has called `panic`,
* which has called `testing.tRunner.func1`.

The indented lines show the source file and line number at which the function was called. The hexadecimal numbers refer to parameter values, including values of pointers and internal data structures. [Stack Traces in Go](https://www.goinggo.net/2015/01/stack-traces-in-go.html) has more details.

#### Print and log a stack trace <a id="print-and-log-a-stack-trace"></a>

To print the stack trace for the current goroutine, use [`debug.PrintStack`](https://golang.org/pkg/runtime/debug/#PrintStack) from package [`runtime/debug`](https://golang.org/pkg/runtime/debug/).

You can also examine the current stack trace programmatically by calling [`runtime.Stack`](https://golang.org/pkg/runtime/#Stack).

#### Level of detail <a id="level-of-detail"></a>

The [`GOTRACEBACK`](https://golang.org/pkg/runtime/#hdr-Environment_Variables) variable controls the amount of output generated when a Go program fails.

* `GOTRACEBACK=none` omits the goroutine stack traces entirely.
* `GOTRACEBACK=single` \(the default\) prints a stack trace for the current goroutine, eliding functions internal to the run-time system. The failure prints stack traces for all goroutines if there is no current goroutine or the failure is internal to the run-time.
* `GOTRACEBACK=all` adds stack traces for all user-created goroutines.
* `GOTRACEBACK=system` is like `all` but adds stack frames for run-time functions and shows goroutines created internally by the run-time.

### Recover and catch a panic <a id="recover-and-catch-a-panic"></a>

The built-in `recover` function can be used to regain control of a panicking goroutine and resume normal execution.

* A call to `recover` stops the unwinding and returns the argument passed to `panic`.
* If the goroutine is not panicking, `recover` returns `nil`.

Because the only code that runs while unwinding is inside [deferred functions](https://yourbasic.org/golang/defer/), `recover` is only useful inside such functions.

#### Panic handler example <a id="panic-handler-example"></a>

```text
func main() {
	n := foo()
	fmt.Println("main received", n)
}

func foo() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	m := 1
	panic("foo: fail")
	m = 2
	return m
}
```

```text
foo: fail
main received 0
```

Since the panic occurred before `foo` returned a value, `n` still has its initial zero value.

#### Return a value <a id="return-a-value"></a>

To return a value during a panic, you must use a [named return value](https://yourbasic.org/golang/named-return-values-parameters/).

```text
func main() {
	n := foo()
	fmt.Println("main received", n)
}

func foo() (m int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			m = 2
		}
	}()
	m = 1
	panic("foo: fail")
	m = 3
	return m
}
```

```text
foo: fail
main received 2
```

### Test a panic \(utility function\) <a id="test-a-panic-utility-function"></a>

In this example, we use reflection to check if a list of interface variables have types corre­sponding to the para­meters of a given function. If so, we call the function with those para­meters to check if there is a panic.

```text
// Panics tells if function f panics with parameters p.
func Panics(f interface{}, p ...interface{}) bool {
	fv := reflect.ValueOf(f)
	ft := reflect.TypeOf(f)
	if ft.NumIn() != len(p) {
		panic("wrong argument count")
	}
	pv := make([]reflect.Value, len(p))
	for i, v := range p {
		if reflect.TypeOf(v) != ft.In(i) {
			panic("wrong argument type")
		}
		pv[i] = reflect.ValueOf(v)
	}
	return call(fv, pv)
}

func call(fv reflect.Value, pv []reflect.Value) (b bool) {
	defer func() {
		if err := recover(); err != nil {
			b = true
		}
	}()
	fv.Call(pv)
	return
}
```



