# Packages explained: declare, import, download, document

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
$ godoc -http=:6060 
```

For more on how to access and create documentation, see the [Package documentation](https://yourbasic.org/golang/package-documentation/) article.

