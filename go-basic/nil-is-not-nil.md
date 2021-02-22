# Nil is not nil

Why is nil not equal to nil in this example?

```text
func Foo() error {
    var err *os.PathError = nil
    // …
    return err
}

func main() {
    err := Foo()
    fmt.Println(err)        // <nil>
    fmt.Println(err == nil) // false
}
```

### Answer <a id="answer"></a>

An interface value is equal to `nil` only if both its value and dynamic type are `nil`. In the example above, `Foo()` returns `[nil, *os.PathError]` and we compare it with `[nil, nil]`.  


You can think of the interface value `nil` as typed, and `nil` _without type_ doesn’t equal `nil` _with type_. If we convert `nil` to the correct type, the values are indeed equal.

```text
…
fmt.Println(err == (*os.PathError)(nil)) // true
…
```

#### A better approach <a id="a-better-approach"></a>



To avoid this problem use a variable of type `error` instead, for example a [named return value](https://yourbasic.org/golang/named-return-values-parameters/).

```text
func Foo() (err error) {
    // …
    return // err is unassigned and has zero value [nil, nil]
}

func main() {
    err := Foo()
    fmt.Println(err)        // <nil>
    fmt.Println(err == nil) // true
}
```

See [Interfaces in 5 easy steps](https://yourbasic.org/golang/interfaces-explained/) for an extensive guide to interfaces in Go.

