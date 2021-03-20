# Language Basics

## Packages explained: declare, import, download, document

### Basics <a id="basics"></a>

Every Go program is made up of packages and each package has an **import path**:

* `"fmt"`
* `"math/rand"`
* `"github.com/yourbasic/graph"`

Packages in the standard library have short import paths, such as `"fmt"` and `"math/rand"`. Third-party packages, such as `"github.com/yourbasic/graph"`, typically have an import path that includes a hosting service \(`github.com`\) and an organization name \(`yourbasic`\).

By convention, the **package name** is the same as the last element of the import path:

* `fmt`
* `rand`
* `graph`

References to other packages’ definitions must always be prefixed with their package names, and only the capitalized names from other packages are accessible.

```text
package main

import (
    "fmt"
    "math/rand"

    "github.com/yourbasic/graph"
)

func main() {
    n := rand.Intn(100)
    g := graph.New(n)
    fmt.Println(g)
}
```

### Declare a package <a id="declare-a-package"></a>

Every Go source file starts with a package declaration, which contains only the package name.

For example, the file [`src/math/rand/exp.go`](https://golang.org/src/math/rand/exp.go), which is part of the implementation of the [`math/rand`](https://golang.org/pkg/math/rand/) package, contains the following code.

```text
package rand
  
import "math"
  
const re = 7.69711747013104972
…
```

You don’t need to worry about package name collisions, only the import path of a package must be unique. [How to Write Go Code](https://golang.org/doc/code.html) shows how to organize your code and its packages in a file structure.

### Package name conflicts <a id="package-name-conflicts"></a>

You can customize the name under which you refer to an imported package.

```text
package main

import (
    csprng "crypto/rand"
    prng "math/rand"

    "fmt"
)

func main() {
    n := prng.Int() // pseudorandom number
    b := make([]byte, 8)
    csprng.Read(b) // cryptographically secure pseudorandom number
    fmt.Println(n, b)
}
```

### Dot imports <a id="dot-imports"></a>

If a period `.` appears instead of a name in an import statement, all the package’s exported identifiers can be accessed without a qualifier.

```text
package main

import (
    "fmt"
    . "math"
)

func main() {
    fmt.Println(Sin(Pi/2)*Sin(Pi/2) + Cos(Pi)/2) // 0.5
}
```

Dot imports can make programs hard to read and **generally should be avoided**.

### Package download <a id="package-download"></a>

The [`go get`](https://golang.org/cmd/go/#hdr-Download_and_install_packages_and_dependencies) command downloads packages named by import paths, along with their dependencies, and then installs the packages.

```text
$ go get github.com/yourbasic/graph
```

The import path corresponds to the repository hosting the code. This reduces the likelihood of future name collisions.

The [Go Wiki](https://github.com/golang/go/wiki/Projects) and [Awesome Go](https://github.com/avelino/awesome-go) provide lists of high-quality Go packages and resources.

For more information on using remote repositories with the go tool, see [Command go: Remote import paths](https://golang.org/cmd/go/#hdr-Remote_import_paths).

### Package documentation <a id="package-documentation"></a>

The [GoDoc](https://godoc.org/) web site hosts documentation for all public Go packages on Bitbucket, GitHub, Google Project Hosting and Launchpad:

* [`https://godoc.org/fmt`](https://godoc.org/fmt)
* [`https://godoc.org/math/rand`](https://godoc.org/math/rand)
* [`https://godoc.org/github.com/yourbasic/graph`](https://godoc.org/github.com/yourbasic/graph)

The [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command extracts and generates documentation for all locally installed Go programs. The following command starts a web server that presents the documentation at `http://localhost:6060/`.

```text
$ godoc -http=:6060 &
```

## Package documentation

### godoc.org website <a id="godoc-org-website"></a>

The [GoDoc](https://godoc.org/) website hosts docu­men­tation for all public Go [packages](https://yourbasic.org/golang/packages-explained/) on Bitbucket, GitHub, Google Project Hosting and Launchpad.

### Local godoc server <a id="local-godoc-server"></a>

The [godoc](https://godoc.org/golang.org/x/tools/cmd/godoc) command extracts and generates documentation for all locally installed Go programs, both your own code and the standard libraries.

The following command starts a web server that presents the documentation at `http://localhost:6060/`.

```text
$ godoc -http=:6060 &
```

![Web browser localhost:6060](https://yourbasic.org/golang/localhost-6060.png)

The documentation is tightly coupled with the code. For example, you can navigate from a function’s documentation to its implementation with a single click.

### go doc command-line tool <a id="go-doc-command-line-tool"></a>

The [go doc](https://golang.org/cmd/go/#hdr-Show_documentation_for_package_or_symbol) command prints plain text documentation to standard output:

```text
$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.
```

### Create documentation <a id="create-documentation"></a>

To document a function, type, constant, variable, or even a complete package, write a regular comment directly preceding its declaration, with no blank line in between. For example, this is the documentation for the [`fmt.Println`](https://golang.org/src/fmt/print.go?s=7388:7437#L246) function:

```text
// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
…
```

For best practices on how to document Go code, see [Effective Go: Commentary](https://golang.org/doc/effective_go.html#commentary).

### Runnable documentation examples <a id="runnable-documentation-examples"></a>

You can add example code snippets to the package documentation; this code is verified by running it as a test. For more information on how to create such testable examples, see [The Go Blog: Testable Examples in Go](https://blog.golang.org/examples).

## Package initialization and program execution order



### Basics <a id="basics"></a>

* First the `main` [package](https://yourbasic.org/golang/packages-explained/) is initialized.
  * Imported packages are initialized before the package itself.
  * Packages are initialized one at a time:
  * first package-level variables are initialized in declaration order,
  * then the `init` functions are run.
* Finally the `main` function is called.

### Program execution <a id="program-execution"></a>

Program execution begins by initializing the `main` package and then calling the function `main`. When `main` returns, the program exits. It **does not wait** for other goroutines to complete.

### Package initialization <a id="package-initialization"></a>

* Package-level variables are initialized in **declaration order**, but after any of the variables they **depend** on.
* Initialization of variables declared in multiple files is done in **lexical file name order**. Variables declared in the first file are declared before any of the variables declared in the second file.
* Initialization cycles are **not allowed**.
* Dependency analysis is performed **per package**; only references referring to variables, functions, and methods declared in the current package are considered.

#### Example <a id="example"></a>

In this example, taken directly from the [Go language specification](https://golang.org/ref/spec#Package_initialization), the initialization order is d, b, c, a.

```text
var (
    a = c + b
    b = f()
    c = f()
    d = 3
)

func f() int {
    d++
    return d
}
```

### Init function <a id="init-function"></a>

Variables may also be initialized using `init` functions.

```text
func init() { … }
```

Multiple such functions may be defined. They cannot be called from inside a program.

* A package with **no imports** is initialized
  * by assigning initial values to all its package-level variables,
  * followed by calling all `init` functions in the order they appear in the source.
* Imported packages are initialized before the package itself.
* Each package is initialized **once**, regardless if it’s imported by multiple other packages.

It follows that there can be **no cyclic dependencies**.

Package initialization happens in a single goroutine, sequentially, one package at a time.

### Warning <a id="warning"></a>

Lexical ordering according to file names is not part of the formal language specification.

> To ensure reproducible initialization behavior, build systems are encouraged to present multiple files belonging to the same package in lexical file name order to a compiler.  
> [The Go Programming Language Specification: Package initialization](https://golang.org/ref/spec#Package_initialization)

{% embed url="https://tutorialedge.net/golang/the-go-init-function/" %}



