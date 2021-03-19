# Maps

## Maps explained: create, add, get, delete

Go maps are implemented by hash tables and have efficient add, get and delete operations.



### Create a new map <a id="create-a-new-map"></a>

```text
var m map[string]int                // nil map of string-int pairs

m1 := make(map[string]float64)      // Empty map of string-float64 pairs
m2 := make(map[string]float64, 100) // Preallocate room for 100 entries

m3 := map[string]float64{           // Map literal
    "e":  2.71828,
    "pi": 3.1416,
}
fmt.Println(len(m3))                // Size of map: 2
```

* A map \(or dictionary\) is an **unordered** collection of **key-value** pairs, where each key is **unique**.
* You create a new map with a [**make**](https://golang.org/pkg/builtin/#make) statement or a **map literal**.
* The default **zero value** of a map is `nil`. A nil map is equivalent to an empty map except that **elements can’t be added**.
* The [**`len`**](https://golang.org/pkg/builtin/#len) function returns the **size** of a map, which is the number of key-value pairs.

> **Warning:** If you try to add an element to an uninitialized map you get the mysterious run-time error [_Assignment to entry in nil map_](https://yourbasic.org/golang/gotcha-assignment-entry-nil-map/).

### Add, update, get and delete keys/values <a id="add-update-get-and-delete-keys-values"></a>

```text
m := make(map[string]float64)

m["pi"] = 3.14             // Add a new key-value pair
m["pi"] = 3.1416           // Update value
fmt.Println(m)             // Print map: "map[pi:3.1416]"

v := m["pi"]               // Get value: v == 3.1416
v = m["pie"]               // Not found: v == 0 (zero value)

_, found := m["pi"]        // found == true
_, found = m["pie"]        // found == false

if x, found := m["pi"]; found {
    fmt.Println(x)
}                           // Prints "3.1416"

delete(m, "pi")             // Delete a key-value pair
fmt.Println(m)              // Print map: "map[]"
```

* When you index a map you get two return values; the second one \(which is optional\) is a boolean that indicates if the key exists.
* If the key doesn’t exist, the first value will be the default [zero value](https://yourbasic.org/golang/default-zero-value/).

### For-each range loop <a id="for-each-range-loop"></a>

```text
m := map[string]float64{
    "pi": 3.1416,
    "e":  2.71828,
}
fmt.Println(m) // "map[e:2.71828 pi:3.1416]"

for key, value := range m { // Order not specified 
    fmt.Println(key, value)
}
```

* Iteration order is not specified and may vary from iteration to iteration.
* If an entry that has not yet been reached is removed during iteration, the corresponding iteration value will not be produced.
* If an entry is created during iteration, that entry may or may not be produced during the iteration.

> Starting with [Go 1.12](https://tip.golang.org/doc/go1.12), the [fmt package](https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/) prints maps in key-sorted order to ease testing.

### Performance and implementation <a id="performance-and-implementation"></a>

* Maps are backed by [hash tables](https://yourbasic.org/algorithms/hash-tables-explained/).
* Add, get and delete operations run in **constant** expected time. The time complexity for the add operation is [amortized](https://yourbasic.org/algorithms/amortized-time-complexity-analysis/).
* The comparison operators `==` and `!=` must be defined for the key type.



