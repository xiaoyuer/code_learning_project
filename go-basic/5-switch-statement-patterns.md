# 5 switch statement patterns

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

