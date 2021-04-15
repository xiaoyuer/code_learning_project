# Pprof 火焰图

[https://slcjordan.github.io/posts/pprof/](https://slcjordan.github.io/posts/pprof/)

```text
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"time"
)

// func main() {
// 	go func() {
// 		for {
// 			LocalTz()

// 			doSomething([]byte(`{"a": 1, "b": 2, "c": 3}`))
// 		}
// 	}()

// 	fmt.Println("start api server...")
// 	panic(http.ListenAndServe(":8080", nil))
// }

// func doSomething(s []byte) {
// 	var m map[string]interface{}
// 	err := json.Unmarshal(s, &m)
// 	if err != nil {
// 		panic(err)
// 	}

// 	s1 := make([]string, 0)
// 	s2 := ""
// 	for i := 0; i < 100; i++ {
// 		s1 = append(s1, string(s))
// 		s2 += string(s)
// 	}
// }

// func LocalTz() *time.Location {
// 	tz, _ := time.LoadLocation("Asia/Shanghai")
// 	return tz
// }

var tz *time.Location

func main() {
	go func() {
		for {
			LocalTz()

			doSomething([]byte(`{"a": 1, "b": 2, "c": 3}`))
		}
	}()

	fmt.Println("start api server...")
	panic(http.ListenAndServe(":8080", nil))
}

// func doSomething(s []byte) {
// 	var m map[string]interface{}
// 	err := json.Unmarshal(s, &m)
// 	if err != nil {
// 		panic(err)
// 	}

// 	s1 := make([]string, 0)
// 	s2 := ""
// 	for i := 0; i < 100; i++ {
// 		s1 = append(s1, string(s))
// 		s2 += string(s)
// 	}
// }

func LocalTz() *time.Location {
	if tz == nil {
		tz, _ = time.LoadLocation("Asia/Shanghai")
	}
	return tz
}

// func doSomething(s []byte) {
// 	var m map[string]interface{}
// 	err := json.Unmarshal(s, &m)
// 	if err != nil {
// 		panic(err)
// 	}

// 	s1 := make([]string, 0)
// 	var buff bytes.Buffer
// 	for i := 0; i < 100; i++ {
// 		s1 = append(s1, string(s))
// 		buff.Write(s)
// 	}
// }

func doSomething(s []byte) {
	var m map[string]interface{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		panic(err)
	}

	s1 := make([]string, 0, 100)
	var buff bytes.Buffer
	for i := 0; i < 100; i++ {
		s1 = append(s1, string(s))
		buff.Write(s)
	}
}
```

```text
go tool pprof http://127.0.0.1:8080/debug/pprof/profile
```

```text
go tool pprof -http=:8081 ~/pprof/pprof.samples.cpu.001.pb.gz
```

图中,从上往下是方法的调用栈,长度代表cpu时长。

使用一个bytes.Buffer类型代替原有的字符串拼接,之后要使用只要buff.String\(\)则可,这里就不在列出。当然buffer并不是线程安全的,如果要考虑并发问题则需做另行打算。

以json.Unmarshal项做参考,可以看到concatstring项已经被bytes.\(\*Buffer\).Write代替,而且仅仅是json.Unmarshal的1/2左右,而原来的concatstring是json.Unmarshal的3倍左右

由于s1这个slice初始化容量为0,在append时,会频繁扩容,带来很大的开销,而此处容量其实是已知项。所以我们可以给他一个初始化容量

可以看到runtime.growslice项已经不存在了。

