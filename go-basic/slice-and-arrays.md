# Slice and Arrays

## Slices/arrays explained: create, index, slice, iterate



### Basics <a id="basics"></a>

A slice doesn’t store any data, it just describes a section of an underlying [array](https://yourbasic.org/algorithms/time-complexity-arrays/).

* When you change an element of a slice, you modify the corresponding element of its underlying array, and other slices that share the same underlying array will see the change.
* A slice can grow and shrink within the bounds of the underlying array.
* Slices are indexed in the usual way: `s[i]` accesses the `i`th element, starting from zero. 

### Construction <a id="construction"></a>

```text
var s []int                   // a nil slice
s1 := []string{"foo", "bar"}
s2 := make([]int, 2)          // same as []int{0, 0}
s3 := make([]int, 2, 4)       // same as new([4]int)[:2]
fmt.Println(len(s3), cap(s3)) // 2 4
```

* The default **zero value** of a slice is `nil`. The functions `len`, `cap` and `append` all regard `nil` as an empty slice with 0 capacity.
* You create a slice either by a **slice literal** or a call to the [`make`](https://golang.org/pkg/builtin/#make) function, which takes the **length** and an optional **capacity** as arguments.
* The built-in [`len`](https://golang.org/pkg/builtin/#len) and [`cap`](https://golang.org/pkg/builtin/#cap) functions retrieve the length and capacity.

### Slicing <a id="slicing"></a>

```text
a := [...]int{0, 1, 2, 3} // an array
s := a[1:3]               // s == []int{1, 2}        cap(s) == 3
s = a[:2]                 // s == []int{0, 1}        cap(s) == 4
s = a[2:]                 // s == []int{2, 3}        cap(s) == 2
s = a[:]                  // s == []int{0, 1, 2, 3}  cap(s) == 4
```

You can also create a slice by slicing an existing array or slice.

* A slice is formed by specifying a low bound and a high bound: `a[low:high]`. This selects a half-open range which includes the first element, but excludes the last.
* You may omit the high or low bounds to use their defaults instead. The default is zero for the low bound and the length of the slice for the high bound.

```text
s := []int{0, 1, 2, 3, 4} // a slice
s = s[1:4]                // s == []int{1, 2, 3}
s = s[1:2]                // s == []int{2} (index relative to slice)
s = s[:3]                 // s == []int{2, 3, 4} (extend length)
```

When you slice a slice, the indexes are relative to the slice itself, not to the backing array.

* The high bound is not bound by the slice’s length, but by it’s capacity, which means you can extend the length of the slice.
* Trying to extend beyond the capacity causes a panic.

### Iteration <a id="iteration"></a>

```text
s := []string{"Foo", "Bar"}
for i, v := range s {
    fmt.Println(i, v)
}
```

```text
0 Foo
1 Bar
```

* The range expression, `s`, is **evaluated once** before beginning the loop.
* The iteration values are assigned to the respective iteration variables, `i` and `v`, **as in an assignment statement**.
* The second iteration variable is optional.
* If the slice is `nil`, the number of iterations is 0.

### Append and copy <a id="append-and-copy"></a>

* The `append` function appends elements to a slice. It will **automatically allocate** a larger backing array if the capacity is exceeded. See [Append function](https://yourbasic.org/golang/append-explained/).
* The `copy` function copies elements into a destination slice `dst` from a source slice `src`. The number of elements copied is the **minimum** of `len(dst)` and `len(src)`. See [Copy function](https://yourbasic.org/golang/copy-explained/).

### Stacks and queues <a id="stacks-and-queues"></a>

The idiomatic way to implement a stack or queue in Go is to use a slice directly. For code examples, see

* [Implement a stack \(LIFO\)](https://yourbasic.org/golang/implement-stack/)
* [Implement a FIFO queue](https://yourbasic.org/golang/implement-fifo-queue/)

## 3 ways to compare slices \(arrays\)

### Basic case <a id="basic-case"></a>

In most cases, you will want to write your own code to compare the elements of two [**slices**](https://yourbasic.org/golang/slices-explained/).

```text
// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
```

For [**arrays**](https://yourbasic.org/golang/slices-explained/), however, you can use the comparison operators `==` and `!=`.

```text
a := [2]int{1, 2}
b := [2]int{1, 3}
fmt.Println(a == b) // false
```

> Array values are comparable if values of the array element type are comparable. Two array values are equal if their corresponding elements are equal.[The Go Programming Language Specification: Comparison operators](https://golang.org/ref/spec#Comparison_operators)

### Optimized code for byte slices <a id="optimized-code-for-byte-slices"></a>

To compare byte slices, use the optimized [`bytes.Equal`](https://golang.org/pkg/bytes/#Equal). This function also treats nil arguments as equivalent to empty slices.

### General-purpose code for recursive comparison <a id="general-purpose-code-for-recursive-comparison"></a>

For testing purposes, you may want to use [`reflect.DeepEqual`](https://golang.org/pkg/reflect/#DeepEqual). It compares two elements of any type recursively.

```text
var a []int = nil
var b []int = make([]int, 0)
fmt.Println(reflect.DeepEqual(a, b)) // false
```

The performance of this function is much worse than for the code above, but it’s useful in test cases where simplicity and correctness are crucial. The semantics, however, are quite complicated.

## How to best clear a slice: empty vs. nil

### Remove all elements <a id="remove-all-elements"></a>

To remove all elements, simply set the slice to `nil`.

```text
a := []string{"A", "B", "C", "D", "E"}
a = nil
fmt.Println(a, len(a), cap(a)) // [] 0 0
```

This will release the underlying array to the garbage collector \(assuming there are no other references\).

### Keep allocated memory <a id="keep-allocated-memory"></a>

To keep the underlying array, slice the slice to zero length.

```text
a := []string{"A", "B", "C", "D", "E"}
a = a[:0]
fmt.Println(a, len(a), cap(a)) // [] 0 5
```

If the slice is extended again, the original data reappears.

```text
fmt.Println(a[:2]) // [A B]
```

### Empty slice vs. nil slice <a id="empty-slice-vs-nil-slice"></a>

In practice, **nil slices** and **empty slices** can often be treated in the same way:

* they have zero length and capacity,
* they can be used with the same effect in [for loops](https://yourbasic.org/golang/for-loop/) and [append functions](https://yourbasic.org/golang/append-explained/),
* and they even look the same when printed.

```text
var a []int = nil
fmt.Println(len(a)) // 0
fmt.Println(cap(a)) // 0
fmt.Println(a)      // []
```

However, if needed, you can tell the difference.

```text
var a []int = nil
var a0 []int = make([]int, 0)

fmt.Println(a == nil)  // true
fmt.Println(a0 == nil) // false

fmt.Printf("%#v\n", a)  // []int(nil)
fmt.Printf("%#v\n", a0) // []int{}
```

The official Go wiki recommends using nil slices over empty slices.

> \[…\] the nil slice is the preferred style.
>
> Note that there are limited circumstances where a non-nil but zero-length slice is preferred, such as when encoding JSON objects \(a nil slice encodes to null, while \[\]string{} encodes to the JSON array \[\]\).
>
> When designing interfaces, avoid making a distinction between a nil slice and a non-nil, zero-length slice, as this can lead to subtle programming errors.[The Go wiki: Declaring empty slices](https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices)

#### Further reading <a id="further-reading"></a>

[Slices and arrays in 6 easy steps](https://yourbasic.org/golang/slices-explained/)

## 2 ways to delete an element from a slice

### Fast version \(changes order\) <a id="fast-version-changes-order"></a>

```text
a := []string{"A", "B", "C", "D", "E"}
i := 2

// Remove the element at index i from a.
a[i] = a[len(a)-1] // Copy last element to index i.
a[len(a)-1] = ""   // Erase last element (write zero value).
a = a[:len(a)-1]   // Truncate slice.

fmt.Println(a) // [A B E D]
```

The code copies a single element and runs in **constant time**.

### Slow version \(maintains order\) <a id="slow-version-maintains-order"></a>

```text
a := []string{"A", "B", "C", "D", "E"}
i := 2

// Remove the element at index i from a.
copy(a[i:], a[i+1:]) // Shift a[i+1:] left one index.
a[len(a)-1] = ""     // Erase last element (write zero value).
a = a[:len(a)-1]     // Truncate slice.

fmt.Println(a) // [A B D E]
```

The code copies len\(a\) - i - 1 elements and runs in **linear time**.

[Slices and arrays in 6 easy steps](https://yourbasic.org/golang/slices-explained/)

## Find element in slice/array with linear or binary search

### Linear search <a id="linear-search"></a>

Go doesn’t have an out-of-the-box linear search function for [slices and arrays](https://yourbasic.org/golang/slices-explained/). Here are two example **linear search** implemen­tations, which you can use as templates.

```text
// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func Find(a []string, x string) int {
    for i, n := range a {
        if x == n {
            return i
        }
    }
    return len(a)
}

// Contains tells whether a contains x.
func Contains(a []string, x string) bool {
    for _, n := range a {
        if x == n {
            return true
        }
    }
    return false
}
```

### Binary search <a id="binary-search"></a>



