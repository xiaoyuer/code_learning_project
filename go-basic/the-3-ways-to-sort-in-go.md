# The 3 ways to sort in Go



### Sort a slice of ints, float64s or strings <a id="sort-a-slice-of-ints-float64s-or-strings"></a>

Use one of the functions

* [`sort.Ints`](https://golang.org/pkg/sort/#Ints)
* [`sort.Float64s`](https://golang.org/pkg/sort/#Float64s)
* [`sort.Strings`](https://golang.org/pkg/sort/#Strings)

```text
s := []int{4, 2, 3, 1}
sort.Ints(s)
fmt.Println(s) // [1 2 3 4]
```



### Sort with custom comparator <a id="sort-with-custom-comparator"></a>

* Use the function [`sort.Slice`](https://golang.org/pkg/sort/#Slice). It sorts a slice using a provided function `less(i, j int) bool`.
* To sort the slice while keeping the original order of equal elements, use [`sort.SliceStable`](https://golang.org/pkg/sort/#SliceStable) instead.

```text
family := []struct {
    Name string
    Age  int
}{
    {"Alice", 23},
    {"David", 2},
    {"Eve", 2},
    {"Bob", 25},
}

// Sort by age, keeping original order or equal elements.
sort.SliceStable(family, func(i, j int) bool {
    return family[i].Age < family[j].Age
})
fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
```



### Sort custom data structures <a id="sort-custom-data-structures"></a>

* Use the generic [`sort.Sort`](https://golang.org/pkg/sort/#Sort) and [`sort.Stable`](https://golang.org/pkg/sort/#Stable) functions.
* They sort any collection that implements the [`sort.Interface`](https://golang.org/pkg/sort/#Interface) [interface](https://yourbasic.org/golang/interfaces-explained/).

```text
type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less reports whether the element with
        // index i should sort before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
}
```

Here’s an example.

```text
type Person struct {
    Name string
    Age  int
}

// ByAge implements sort.Interface based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
    family := []Person{
        {"Alice", 23},
        {"Eve", 2},
        {"Bob", 25},
    }
    sort.Sort(ByAge(family))
    fmt.Println(family) // [{Eve 2} {Alice 23} {Bob 25}]
}
```



### Bonus: Sort a map by key or value <a id="bonus-sort-a-map-by-key-or-value"></a>

A [map](https://yourbasic.org/golang/maps-explained/) is an **unordered** collection of key-value pairs. If you need a stable iteration order, you must maintain a separate data structure.

This code example uses a slice of keys to sort a map in key order.

```text
m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}

keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)

for _, k := range keys {
    fmt.Println(k, m[k])
}
// Output:
// Alice 2
// Bob 3
// Cecil 1
```

//sorted but not change order

further reading

[https://yourbasic.org/golang/slices-explained/](https://yourbasic.org/golang/slices-explained/)



