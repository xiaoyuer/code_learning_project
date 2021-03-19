# Files

## Read a file \(stdin\) line by line

### Read from file <a id="read-from-file"></a>

Use a [`bufio.Scanner`](https://golang.org/pkg/bufio/#Scanner) to read a file line by line.

```text
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

### Read from stdin <a id="read-from-stdin"></a>

Use [`os.Stdin`](https://golang.org/pkg/os/#pkg-variables) to read from the standard input stream.

```text
scanner := bufio.NewScanner(os.Stdin)
for scanner.Scan() {
	fmt.Println(scanner.Text())
}

if err := scanner.Err(); err != nil {
	log.Println(err)
}
```

### Read from any stream <a id="read-from-any-stream"></a>

A bufio.Scanner can read from any stream of bytes, as long as it implements the [`io.Reader`](https://golang.org/pkg/io/#Reader) interface. See [How to use the io.Reader interface](https://yourbasic.org/golang/io-reader-interface-explained/).

## Append text to a file

This code appends a line of text to the file `text.log`. It creates the file if it doesn’t already exist.

```text
f, err := os.OpenFile("text.log",
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
	log.Println(err)
}
defer f.Close()
if _, err := f.WriteString("text to append\n"); err != nil {
	log.Println(err)
}

```

## Find current working directory



### Current directory <a id="current-directory"></a>

Use [`os.Getwd`](https://golang.org/pkg/os/#Getwd) to find the path name for the current directory.

```text
path, err := os.Getwd()
if err != nil {
    log.Println(err)
}
fmt.Println(path)  // for example /home/user
```

> **Warning:** If the current directory can be reached via multiple paths \(due to symbolic links\), Getwd may return any one of them.

### Current executable <a id="current-executable"></a>

Use [`os.Executable`](https://golang.org/pkg/os/#Executable) to find the path name for the executable that started the current process.

```text
path, err := os.Executable()
if err != nil {
    log.Println(err)
}
fmt.Println(path) // for example /tmp/go-build872132473/b001/exe/main
```

> **Warning:** There is no guarantee that the path is still pointing to the correct executable. If a symlink was used to start the process, depending on the operating system, the result might be the symlink or the path it pointed to. If a stable result is needed, [`path/filepath.EvalSymlinks`](https://golang.org/pkg/path/filepath/#EvalSymlinks) might help.

## List all files \(recursively\) in a directory

### Directory listing <a id="directory-listing"></a>

Use the [`ioutil.ReadDir`](https://golang.org/pkg/io/ioutil/#ReadDir) function in package [`io/ioutil`](https://golang.org/pkg/io/ioutil/). It returns a sorted slice containing elements of type [`os.FileInfo`](https://golang.org/pkg/os/#FileInfo).

The code in this example prints a sorted list of all file names in the current directory.

```text
files, err := ioutil.ReadDir(".")
if err != nil {
    log.Fatal(err)
}
for _, f := range files {
    fmt.Println(f.Name())
}
```

Example output:

```text
dev
etc
tmp
usr
```

### Visit all files and folders in a directory tree <a id="visit-all-files-and-folders-in-a-directory-tree"></a>

Use the [`filepath.Walk`](https://golang.org/pkg/path/filepath/#Walk) function in package [`path/filepath`](https://golang.org/pkg/path/filepath/).

* It walks a file tree calling a function of type [`filepath.WalkFunc`](https://golang.org/pkg/path/filepath/#WalkFunc) for each file or directory in the tree, including the root.
* The files are walked in lexical order.
* Symbolic links are not followed.

The code in this example lists the paths and sizes of all files and directories in the file tree rooted at the current directory.

```text
err := filepath.Walk(".",
    func(path string, info os.FileInfo, err error) error {
    if err != nil {
        return err
    }
    fmt.Println(path, info.Size())
    return nil
})
if err != nil {
    log.Println(err)
}
```

Example output:

```text
. 1644
dev 1644
dev/null 0
dev/random 0
dev/urandom 0
dev/zero 0
etc 1644
etc/group 116
etc/hosts 20
etc/passwd 0
etc/resolv.conf 0
tmp 548
usr 822
usr/local 822
usr/local/go 822
usr/local/go/lib 822
usr/local/go/lib/time 822
usr/local/go/lib/time/zoneinfo.zip 366776
```

## Create a temporary file or directory

### File <a id="file"></a>

Use [`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile) in package [`io/ioutil`](https://golang.org/pkg/io/ioutil/) to create a **globally unique temporary file**. It’s your own job to remove the file when it’s no longer needed.

```text
file, err := ioutil.TempFile("dir", "prefix")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(file.Name())

fmt.Println(file.Name()) // For example "dir/prefix054003078"
```

The call to [`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile)

* creates a new file with a name starting with `"prefix"` in the directory `"dir"`,
* opens the file for reading and writing,
* and returns the new [`*os.File`](https://golang.org/pkg/os/#File).

To put the new file in [`os.TempDir()`](https://golang.org/pkg/os/#TempDir), the default directory for temporary files, call [`ioutil.TempFile`](https://golang.org/pkg/io/ioutil/#TempFile) with an empty directory string.

#### Add a suffix to the temporary file name[Go 1.11](https://tip.golang.org/doc/go1.11) <a id="add-suffix"></a>

Starting with Go 1.11, if the second string given to `TempFile` includes a `"*"`, the random string replaces this `"*"`.

```text
file, err := ioutil.TempFile("dir", "myname.*.bat")
if err != nil {
    log.Fatal(err)
}
defer os.Remove(file.Name())

fmt.Println(file.Name()) // For example "dir/myname.054003078.bat"
```

If no `"*"` is included the old behavior is retained, and the random digits are appended to the end.

### Directory <a id="directory"></a>

Use [`ioutil.TempDir`](https://golang.org/pkg/io/ioutil/#TempDir) in package [`io/ioutil`](https://golang.org/pkg/io/ioutil/) to create a **globally unique temporary directory**.

```text
dir, err := ioutil.TempDir("dir", "prefix")
if err != nil {
	log.Fatal(err)
}
defer os.RemoveAll(dir)
```

The call to [`ioutil.TempDir`](https://golang.org/pkg/io/ioutil/#TempDir)

* creates a new directory with a name starting with `"prefix"` in the directory `"dir"`
* and returns the path of the new directory.

To put the new directory in [`os.TempDir()`](https://golang.org/pkg/os/#TempDir), the default directory for temporary files, call [`ioutil.TempDir`](https://golang.org/pkg/io/ioutil/#TempDir) with an empty directory string.



