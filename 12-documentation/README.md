#### 12.1 go doc命令

go语言内置的doc子命令可以通过命令行来查询文档，我们可以使用go help命令来查看如何使用go doc

```
$ go help doc

        go doc
                Show documentation for current package.
        go doc Foo
                Show documentation for Foo in the current package.
                (Foo starts with a capital letter so it cannot match
                a package path.)
        go doc encoding/json
                Show documentation for the encoding/json package.
        go doc json
                Shorthand for encoding/json.
        go doc json.Number (or go doc json.number)
                Show documentation and method summary for json.Number.
        go doc json.Number.Int64 (or go doc json.number.int64)
                Show documentation for json.Number's Int64 method.
        go doc cmd/doc
                Show package docs for the doc command.
        go doc -cmd cmd/doc
                Show package docs and exported symbols within the doc command.
        go doc template.new
                Show documentation for html/template's New function.
                (html/template is lexically before text/template)
        go doc text/template.new # One argument
                Show documentation for text/template's New function.
        go doc text/template new # Two arguments
                Show documentation for text/template's New function.

```
根据上面的描述信息，假如我们想要查看fmt.Println的文档可以使用下面的方式来查询。
```
$ go doc fmt.Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

```
另外默认情况下直接使用go doc将列举当前位置的库的文档信息，后面我们可以看到如何去编写自己写的库的文档信息。
 
#### 12.2 godoc命令

go语言安装后会自动的同时安装godoc命令，这个命令有一个内置的web服务，可以在本地启动一个标准库的文档服务器，用于查询当前标准库的相关内容。使用下面的命令启动该web服务：

```sh
godoc.exe -http=127.0.0.1:8080
```
可以通过web服务中的搜索框或者Packages来查询，同时godoc本身支持类似于go doc查询库包的命令

```go
$ godoc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline
    is appended. It returns the number of bytes written and any write error
    encountered.

$ go doc fmt Println
func Println(a ...interface{}) (n int, err error)
    Println formats using the default formats for its operands and writes to
    standard output. Spaces are always added between operands and a newline is
    appended. It returns the number of bytes written and any write error
    encountered.

```

#### 12.3 使用godoc索引

假如自己写了一个简单的go语言的库，想要别人能够通过godoc检索到，该如何操作？

其实很简单，这里我们以当前文件夹中的mylib库为例子，我们推送到git上后，会获得该库的一个访问路径，以下是实例的连接：
```
https://github.com/zhangmingkai4315/golang-essentials/tree/master/12-documentation/mylib
```

我们将该链接放置到godoc中的搜索框后点击Go按钮即可，系统会自动的重定向到一个新的URL:
```
https://godoc.org/
```
而且所有的文档信息会自动的显示出来。这时候任何人如果通过搜索框都可以查询到该库的相关信息。

同时我们可以使用godoc命令行来查询相关链接下的文档信息：

```sh
 godoc github.com/zhangmingkai4315/golang-essentials/12-documentation/mylib

```

#### 12.4 编写文档

如何编写文档，可以通过学习一些标准库的编写规范，比如errors库, 代码一般包含以下几个部分：
- 头部版权信息
- 包通用描述信息
- 针对具体函数或者类型的定义信息（第一个单词必须以类型名称或者函数名开始）

如下面的实例所示：

```
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors implements functions to manipulate errors.
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

```

对于一些包的描述信息如果比较长，包含一些示例的话，可以创建一个单独的文件用于保存这些信息，比如fmt包中存在一个独立的文件doc.go文件来保存这些信息，具体实现可访问链接[doc.go](https://golang.org/src/fmt/doc.go)

如果我们查询标准库的话，会看到很多库都带有example示例，这些示例其实也是通过解析代码库中的相关代码提取出来的，而且这些示例提取后，可以直接形成可执行的go代码。

比如golang/example示例中的stringutil包，为了提供示例代码我们可以编写example_test.go文件：
```
package stringutil_test

import (
    "fmt"

    "github.com/golang/example/stringutil"
)

func ExampleReverse() {
    fmt.Println(stringutil.Reverse("hello"))
    // Output: olleh
}
```

文件的package名称为包名称加test，同时示例函数必须以Examle开头，对任何函数或者类型的示例后面跟函数名称和结构体名称。
示例代码中的```Output: olleh```会自动被转换输出

![](https://blog.golang.org/examples/reverse.png)

同时当我们使用go test的时候，example代码也会被自动执行测试
```
$ go test -v
=== RUN TestReverse
--- PASS: TestReverse (0.00s)
=== RUN: ExampleReverse
--- PASS: ExampleReverse (0.00s)
PASS
ok      github.com/golang/example/stringutil    0.009s
```

这里的example的测试比较实际的输出和Output中定义的输出，**如果不包含Output也就不会执行相关测试**

多个输出的时候，按行完成output的编写：

```go

    fmt.Println(people)
    sort.Sort(ByAge(people))
    fmt.Println(people)

    // Output:
    // [Bob: 31 John: 42 Michael: 17 Jenny: 26]
    // [Michael: 17 Jenny: 26 Bob: 31 John: 42]
```

#### 12.5 附录

[1.golang官方文档](https://golang.org/doc/)

[2.godoc.org索引](https://godoc.org/)

[3.example](https://blog.golang.org/examples)
