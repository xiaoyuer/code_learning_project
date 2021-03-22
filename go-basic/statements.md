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



