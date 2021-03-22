# Statements

## 4 basic if-else statement patterns

### Basic syntax <a id="basic-syntax"></a>

```text
if x > max {
    x = max
}
```

```text
if x <= y {
    min = x
} else {
    min = y
}
```

An **if statement** executes one of two branches according to a boolean expression.

* If the expression evaluates to true, the **if** branch is executed,
* otherwise, if present, the **else** branch is executed.

### With init statement <a id="with-init-statement"></a>

```text
if x := f(); x <= y {
    return x
}
```

The expression may be preceded by a **simple statement**, which executes before the expression is evaluated. The **scope** of `x` is limited to the if statement.

### Nested if statements <a id="nested-if-statements"></a>

```text
if x := f(); x < y {
    return x
} else if x > z {
    return z
} else {
    return y
}
```

Complicated conditionals are often best expressed in Go with a **switch statement**. See [5 switch statement patterns](https://yourbasic.org/golang/switch-statement/) for details.

### Ternary ? operator alternatives <a id="ternary-operator-alternatives"></a>



You can’t write a short one-line conditional in Go; there is no ternary conditional operator. Instead of

```text
res = expr ? x : y
```

you write

```text
if expr {
    res = x
} else {
    res = y
}
```

In some cases, you may want to create a dedicated function.

```text
func Min(x, y int) int {
    if x <= y {
        return x
    }
    return y
}
```

## 5 switch statement patterns

### Basic switch with default <a id="basic-switch-with-default"></a>

* A switch statement runs the first case equal to the condition expression.
* The cases are evaluated from top to bottom, stopping when a case succeeds.
* If no case matches and there is a default case, its statements are executed.

```text
switch time.Now().Weekday() {
case time.Saturday:
    fmt.Println("Today is Saturday.")
case time.Sunday:
    fmt.Println("Today is Sunday.")
default:
    fmt.Println("Today is a weekday.")
}
```

> Unlike C and Java, the case expressions do not need to be constants.

### No condition <a id="no-condition"></a>

A switch without a condition is the same as switch true.

```text
switch hour := time.Now().Hour(); { // missing expression means "true"
case hour < 12:
    fmt.Println("Good morning!")
case hour < 17:
    fmt.Println("Good afternoon!")
default:
    fmt.Println("Good evening!")
}
```

### Case list <a id="case-list"></a>

```text
func WhiteSpace(c rune) bool {
    switch c {
    case ' ', '\t', '\n', '\f', '\r':
        return true
    }
    return false
}
```

### Fallthrough <a id="fallthrough"></a>

* A `fallthrough` statement transfers control to the next case.
* It may be used only as the final statement in a clause.

```text
switch 2 {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")
    fallthrough
case 3:
    fmt.Println("3")
}
```

```text
2
3
```

### Exit with break <a id="exit-with-break"></a>

A `break` statement terminates execution of the **innermost** `for`, `switch`, or `select` statement.

If you need to break out of a surrounding loop, not the switch, you can put a **label** on the loop and break to that label. This example shows both uses.

```text
Loop:
    for _, ch := range "a b\nc" {
        switch ch {
        case ' ': // skip space
            break
        case '\n': // break at newline
            break Loop
        default:
            fmt.Printf("%c\n", ch)
        }
    }
```

```text
a
b
```

### Execution order <a id="execution-order"></a>

* First the switch expression is evaluated once.
* Then case expressions are evaluated left-to-right and top-to-bottom:
  * the first one that equals the switch expression triggers execution of the statements of the associated case,
  * the other cases are skipped.

```text
// Foo prints and returns n.
func Foo(n int) int {
    fmt.Println(n)
    return n
}

func main() {
    switch Foo(2) {
    case Foo(1), Foo(2), Foo(3):
        fmt.Println("First case")
        fallthrough
    case Foo(4):
        fmt.Println("Second case")
    }
}
```

```text
2
1
2
First case
Second case
```

## 5 basic for loop patterns



### Three-component loop <a id="three-component-loop"></a>

This version of the Go for loop works just as in C or Java.

```text
sum := 0
for i := 1; i < 5; i++ {
    sum += i
}
fmt.Println(sum) // 10 (1+2+3+4)
```

1. The init statement, `i := 1`, runs.
2. The condition, `i < 5`, is computed.
   * If true, the loop body runs,
   * otherwise the loop is done.
3. The post statement, `i++`, runs.
4. Back to step 2.

The scope of `i` is limited to the loop.

### While loop <a id="while-loop"></a>

If you skip the init and post statements, you get a while loop.

```text
n := 1
for n < 5 {
    n *= 2
}
fmt.Println(n) // 8 (1*2*2*2)
```

1. The condition, `n < 5`, is computed.
   * If true, the loop body runs,
   * otherwise the loop is done.
2. Back to step 1.

### Infinite loop <a id="infinite-loop"></a>

If you skip the condition as well, you get an infinite loop.

```text
sum := 0
for {
    sum++ // repeated forever
}
fmt.Println(sum) // never reached
```

### For-each range loop <a id="for-each-range-loop"></a>

Looping over elements in _slices_, _arrays_, _maps_, _channels_ or _strings_ is often better done with a range loop.

```text
strings := []string{"hello", "world"}
for i, s := range strings {
    fmt.Println(i, s)
}
```

```text
0 hello
1 world
```

See [4 basic range loop patterns](https://yourbasic.org/golang/for-loop-range-array-slice-map-channel/) for a complete set of examples.

### Exit a loop <a id="exit-a-loop"></a>

The break and continue keywords work just as they do in C and Java.

```text
sum := 0
for i := 1; i < 5; i++ {
    if i%2 != 0 { // skip odd numbers
        continue
    }
    sum += i
}
fmt.Println(sum) // 6 (2+4)
```

* A **continue** statement begins the next iteration of the innermost **for** loop at its post statement \(`i++`\).
* A **break** statement leaves the innermost **for**, [**switch**](https://yourbasic.org/golang/switch-statement/) or [**select**](https://yourbasic.org/golang/select-explained/) statement.

## 2 patterns for a do-while loop in Go

There is no **do-while loop** in Go. To emulate the C/Java code

```text
do {
    work();
} while (condition);
```

you may use a [for loop](https://yourbasic.org/golang/for-loop/) in one of these two ways:

```text
for ok := true; ok; ok = condition {
    work()
}
```

```text
for {
    work()
    if !condition {
        break
    }
}
```

#### Repeat-until loop <a id="repeat-until-loop"></a>

To write a **repeat-until loop**

```text
repeat
    work();
until condition;
```

simply change the condition in the code above to its complement:

```text
for ok := true; ok; ok = !condition {
    work()
}
```

```text
for {
    work()
    if condition {
        break
    }
}
```

## 4 basic range loop \(for-each\) patterns



### Basic for-each loop \(slice or array\) <a id="basic-for-each-loop-slice-or-array"></a>

```text
a := []string{"Foo", "Bar"}
for i, s := range a {
    fmt.Println(i, s)
}
```

```text
0 Foo
1 Bar
```

* The range expression, `a`, is **evaluated once** before beginning the loop.
* The iteration values are assigned to the respective iteration variables, `i` and `s`, **as in an assignment statement**.
* The second iteration variable is optional.
* For a nil slice, the number of iterations is 0.

### String iteration: runes or bytes <a id="string-iteration-runes-or-bytes"></a>

For strings, the range loop iterates over [Unicode code points](https://yourbasic.org/golang/rune/).

```text
for i, ch := range "日本語" {
    fmt.Printf("%#U starts at byte position %d\n", ch, i)
}
```

```text
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```

* The index is the first byte of a UTF-8-encoded code point; the second value, of type `rune`, is the value of the code point.
* For an invalid UTF-8 sequence, the second value will be 0xFFFD, and the iteration will advance a single byte.

> To loop over individual bytes, simply use a [normal for loop](https://yourbasic.org/golang/for-loop/) and string indexing:
>
> ```text
> const s = "日本語"
> for i := 0; i < len(s); i++ {
>     fmt.Printf("%x ", s[i])
> }
> ```
>
> ```text
> e6 97 a5 e6 9c ac e8 aa 9e
> ```

### Map iteration: keys and values <a id="map-iteration-keys-and-values"></a>

The iteration order over [maps](https://yourbasic.org/golang/maps-explained/) is not specified and is not guaranteed to be the same from one iteration to the next.

```text
m := map[string]int{
    "one":   1,
    "two":   2,
    "three": 3,
}
for k, v := range m {
    fmt.Println(k, v)
}
```

```text
two 2
three 3
one 1
```

* If a map entry that has not yet been reached is removed during iteration, this value will not be produced.
* If a map entry is created during iteration, that entry may or may not be produced.
* For a nil map, the number of iterations is 0.

### Channel iteration <a id="channel-iteration"></a>

For [channels](https://yourbasic.org/golang/channels-explained/), the iteration values are the successive values sent on the channel until closed.

```text
ch := make(chan int)
go func() {
    ch <- 1
    ch <- 2
    ch <- 3
    close(ch)
}()
for n := range ch {
    fmt.Println(n)
}
```

```text
1
2
3
```

* For a nil channel, the range loop blocks forever.

### Gotchas <a id="gotchas"></a>

Here are two traps that you want to avoid when using range loops:

* [Unexpected values in range loop](https://yourbasic.org/golang/gotcha-unexpected-values-range/)
* [Can’t change entries in range loop](https://yourbasic.org/golang/gotcha-change-value-range/)

## Defer a function call \(with return value\)

### Defer statement basics <a id="defer-statement-basics"></a>

A `defer` statement postpones the execution of a function until the surrounding function returns, either normally or through a panic.

```text
func main() {
    defer fmt.Println("World")
    fmt.Println("Hello")
}
```

```text
Hello
World
```

Deferred calls are executed even when the function panics:

```text
func main() {
    defer fmt.Println("World")
    panic("Stop")
    fmt.Println("Hello")
}
```

```text
World
panic: Stop

goroutine 1 [running]:
main.main()
    ../main.go:3 +0xa0
```

#### Order of execution <a id="order-of-execution"></a>

The deferred call’s **arguments are evaluated immediately**, even though the function call is not executed until the surrounding function returns.

If there are several deferred function calls, they are executed in last-in-first-out order.

```text
func main() {
    fmt.Println("Hello")
    for i := 1; i <= 3; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("World")
}
```

```text
Hello
World
3
2
1
```

#### Use func to return a value <a id="use-func-to-return-a-value"></a>

Deferred anonymous functions may access and modify the surrounding function’s named return parameters.

In this example, the `foo` function returns “Change World”.

```text
func foo() (result string) {
    defer func() {
        result = "Change World" // change value at the very last moment
    }()
    return "Hello World"
}
```

### Common applications <a id="common-applications"></a>

Defer is often used to perform clean-up actions, such as closing a file or unlocking a mutex. Such actions should be performed both when the function returns normally and when it panics.

#### Close a file <a id="close-a-file"></a>

In this example, defer statements are used to ensure that all files are closed before leaving the `CopyFile` function, whichever way that happens.

```text
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}
```

#### Error handling: catch a panic <a id="error-handling-catch-a-panic"></a>

The [Recover from a panic](https://yourbasic.org/golang/recover-from-panic/#recover-and-catch-a-panic) code example shows how to use a defer statement to recover from a panic and update the return value.

## Type assertions and type switches

### Type assertions <a id="type-assertions"></a>

A **type assertion** doesn’t really convert an [interface](https://yourbasic.org/golang/interfaces-explained/) to another data type, but it provides access to an interface’s concrete value, which is typically what you want.

The type assertion `x.(T)` asserts that the concrete value stored in `x` is of type `T`, and that `x` is not nil.

* If `T` is not an interface, it asserts that the dynamic type of `x` is identical to `T`.
* If `T` is an interface, it asserts that the dynamic type of `x` implements `T`.

```text
var x interface{} = "foo"

var s string = x.(string)
fmt.Println(s)     // "foo"

s, ok := x.(string)
fmt.Println(s, ok) // "foo true"

n, ok := x.(int)
fmt.Println(n, ok) // "0 false"

n = x.(int)        // ILLEGAL
```

```text
panic: interface conversion: interface {} is string, not int
```

### Type switches <a id="type-switches"></a>

A **type switch** performs several type assertions in series and runs the first case with a matching type.

```text
var x interface{} = "foo"

switch v := x.(type) {
case nil:
    fmt.Println("x is nil")            // here v has type interface{}
case int: 
    fmt.Println("x is", v)             // here v has type int
case bool, string:
    fmt.Println("x is bool or string") // here v has type interface{}
default:
    fmt.Println("type unknown")        // here v has type interface{}
}
```

```text
x is bool or string
```

## Type alias explained

An **alias declaration** has the form

```text
type T1 = T2
```

as opposed to a standard **type definition**

```text
type T1 T2
```

An alias declaration doesn’t create a new distinct type different from the type it’s created from. It just introduces an alias name `T1`, an alternate spelling, for the type denoted by `T2`.

Type aliases are not meant for everyday use. They were introduced to support gradual code repair while moving a type between packages during large-scale refactoring. [Codebase Refactoring \(with help from Go\)](https://talks.golang.org/2016/refactor.article) covers this in detail.



