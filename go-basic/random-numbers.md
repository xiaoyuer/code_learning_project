# Random Numbers

## Generate random numbers, characters and slice elements

### Go pseudo-random number basics <a id="go-pseudo-random-number-basics"></a>

Use the [`rand.Seed`](https://golang.org/pkg/math/rand/#Seed) and [`rand.Int63`](https://golang.org/pkg/math/rand/#Int63) functions in package [`math/rand`](https://golang.org/pkg/math/rand/) to generate a non-negative pseudo-random number of type `int64`:

```text
rand.Seed(time.Now().UnixNano())
n := rand.Int63() // for example 4601851300195147788
```

Similarly, [`rand.Float64`](https://golang.org/pkg/math/rand/#Float64) generates a pseudo-random float x, where 0 ≤ x &lt; 1:

```text
x := rand.Float64() // for example 0.49893371771268225
```

> **Warning:** Without an initial call to `rand.Seed`, you will get the same sequence of numbers each time you run the program.

See [What’s a seed in a random number generator?](https://yourbasic.org/algorithms/random-number-generator-seed/) for an explanation of pseuodo-random number generators.

#### Several random sources <a id="several-random-sources"></a>

The functions in the [`math/rand`](https://golang.org/pkg/math/rand/) package all use a single random source.

If needed, you can create a new random generator of type [`Rand`](https://golang.org/pkg/math/rand/#Rand) with its own source, and then use its methods to generate random numbers:

```text
generator := rand.New(rand.NewSource(time.Now().UnixNano()))
n := generator.Int63()
x := generator.Float64()
```

### Integers and characters in a given range <a id="integers-and-characters-in-a-given-range"></a>

#### Number between a and b <a id="number-between-a-and-b"></a>

Use [`rand.Intn(m)`](https://golang.org/pkg/math/rand/#Intn), which returns a pseudo-random number n, where 0 ≤ n &lt; m.

```text
n := a + rand.Intn(b-a+1) // a ≤ n ≤ b
```

#### Character between 'a' and 'z' <a id="character-between-39-a-39-and-39-z-39"></a>

```text
c := 'a' + rune(rand.Intn('z'-'a'+1)) // 'a' ≤ c ≤ 'z'
```

### Random element from slice <a id="random-element-from-slice"></a>

To generate a character from an arbitrary set, choose a random index from a slice of characters:

```text
chars := []rune("AB⌘")
c := chars[rand.Intn(len(chars))] // for example '⌘'

```

[Runes and character encoding](https://yourbasic.org/golang/rune/)

## Generate a random string \(password\)

### Random string <a id="random-string"></a>

This code generates a random string of numbers and characters from the Swedish alphabet \(which includes the non-ASCII characters å, ä and ö\).

```text
rand.Seed(time.Now().UnixNano())
chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
    "abcdefghijklmnopqrstuvwxyzåäö" +
    "0123456789")
length := 8
var b strings.Builder
for i := 0; i < length; i++ {
    b.WriteRune(chars[rand.Intn(len(chars))])
}
str := b.String() // E.g. "ExcbsVQs"
```

> **Warning:** To generate a password, you should use cryptographically secure pseudorandom numbers. See [User-friendly access to crypto/rand](https://yourbasic.org/golang/crypto-rand-int/).

### Random string with restrictions <a id="random-string-with-restrictions"></a>

This code generates a random ASCII string with at least one digit and one special character.

```text
rand.Seed(time.Now().UnixNano())
digits := "0123456789"
specials := "~=+%^*/()[]{}/!@#$?|"
all := "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
    "abcdefghijklmnopqrstuvwxyz" +
    digits + specials
length := 8
buf := make([]byte, length)
buf[0] = digits[rand.Intn(len(digits))]
buf[1] = specials[rand.Intn(len(specials))]
for i := 2; i < length; i++ {
    buf[i] = all[rand.Intn(len(all))]
}
rand.Shuffle(len(buf), func(i, j int) {
    buf[i], buf[j] = buf[j], buf[i]
})
str := string(buf) // E.g. "3i[g0|)z"
```

#### Before Go 1.10 <a id="before-go-1-10"></a>

In code before Go 1.10, replace the call to [rand.Shuffle](https://golang.org/pkg/math/rand/#Shuffle) with this code:

```text
for i := len(buf) - 1; i > 0; i-- { // Fisher–Yates shuffle
    j := rand.Intn(i + 1)
    buf[i], buf[j] = buf[j], buf[i]
}
```

## Generate a unique string \(UUID, GUID\)

A [universally unique identifier](https://en.wikipedia.org/wiki/Universally_unique_identifier) \(UUID\), or globally unique identifier \(GUID\), is a 128-bit number used to identify information.

* A UUID is for practical purposes unique: the probability that it will be duplicated is very close to zero.
* UUIDs don’t depend on a central authority or on coordination between those generating them.

The string representation of a UUID consists of 32 hexadecimal digits displayed in 5 groups separated by hyphens. For example:

```text
123e4567-e89b-12d3-a456-426655440000
```

### UUID generator example <a id="uuid-generator-example"></a>

You can use the [`rand.Read`](https://golang.org/pkg/crypto/rand/#Read) function from package [`crypto/rand`](https://golang.org/pkg/crypto/rand/) to generate a basic UUID.

```text
b := make([]byte, 16)
_, err := rand.Read(b)
if err != nil {
    log.Fatal(err)
}
uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
    b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
fmt.Println(uuid)
```

```text
9438167c-9493-4993-fd48-950b27aad7c9
```

#### Limitations <a id="limitations"></a>

This UUID doesn’t conform to [RFC 4122](https://tools.ietf.org/html/rfc4122). In particular, it doesn’t contain any version or variant numbers.

> **Warning:** The `rand.Read` call returns an error if the underlying system call fails. For instance if it can't read `/dev/urandom` on a Unix system, or if [`CryptAcquireContext`](https://msdn.microsoft.com/en-us/library/windows/desktop/aa379886%28v=vs.85%29.aspx) fails on a Windows system.

## Shuffle a slice or array

The [`rand.Shuffle`](https://golang.org/pkg/math/rand/#Shuffle) function in package [`math/rand`](https://golang.org/pkg/math/rand/) shuffles an input sequence using a given swap function.

```text
a := []int{1, 2, 3, 4, 5, 6, 7, 8}
rand.Seed(time.Now().UnixNano())
rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
```

```text
[5 8 6 4 3 7 2 1]
```

> **Warning:** Without the call to `rand.Seed` you will get the same sequence of pseudo­random numbers each time you run the program.

#### Before Go 1.10 <a id="before-go-1-10"></a>

Use the [`rand.Seed`](https://golang.org/pkg/math/rand/#Seed) and [`rand.Intn`](https://golang.org/pkg/math/rand/#Intn) functions in package [`math/rand`](https://golang.org/pkg/math/rand/).

```text
a := []int{1, 2, 3, 4, 5, 6, 7, 8}
rand.Seed(time.Now().UnixNano())
for i := len(a) - 1; i > 0; i-- { // Fisher–Yates shuffle
    j := rand.Intn(i + 1)
    a[i], a[j] = a[j], a[i]
}
```

## User-friendly access to crypto/rand

Go has two packages for random numbers:

* [`math/rand`](https://golang.org/pkg/math/rand/) implements a large selection of pseudo-random number generators.
* [`crypto/rand`](https://golang.org/pkg/crypto/rand/) implements a cryptographically secure pseudo-random number generator with a limited interface.

The two packages can be combined by calling [`rand.New`](https://golang.org/pkg/math/rand/#New) in package `math/rand` with a source that gets its data from `crypto/rand`.

```text
import (
    crand "crypto/rand"
    rand "math/rand"

    "encoding/binary"
    "fmt"
    "log"
)

func main() {
    var src cryptoSource
    rnd := rand.New(src)
    fmt.Println(rnd.Intn(1000)) // a truly random number 0 to 999
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
    return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
    err := binary.Read(crand.Reader, binary.BigEndian, &v)
    if err != nil {
        log.Fatal(err)
    }
    return v
}
```

> **Warning:** The `crand.Reader` returns an error if the underlying system call fails. For instance if it can't read `/dev/urandom` on a Unix system, or if [`CryptAcquireContext`](https://msdn.microsoft.com/en-us/library/windows/desktop/aa379886%28v=vs.85%29.aspx) fails on a Windows system.



